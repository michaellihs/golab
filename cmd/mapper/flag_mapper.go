// Copyright © 2017 Michael Lihs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mapper

import (
	"errors"
	"reflect"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"time"
)

type FlagMapper struct {
	cmd *cobra.Command
}

func New(cmd *cobra.Command) FlagMapper {
	return FlagMapper{cmd: cmd}
}

func (m FlagMapper) SetFlags(flags interface{}) {
	v := reflect.ValueOf(flags).Elem()
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag
		f := v.Field(i)
		flagName := tag.Get("flag_name")
		shortHand := tag.Get("short")

		switch f.Type().String() {
		case "*int":
			m.cmd.PersistentFlags().IntP(flagName, shortHand, 0, flagUsage(tag))
		case "*string":
			m.cmd.PersistentFlags().StringP(flagName, shortHand, "", flagUsage(tag))
		case "*bool":
			m.cmd.PersistentFlags().BoolP(flagName, shortHand, false, flagUsage(tag))
		case "*[]string":
			m.cmd.PersistentFlags().StringArrayP(flagName, shortHand, nil, flagUsage(tag))
		default:
			panic("Unknown type " + f.Type().String())
		}
	}
}
func flagUsage(tag reflect.StructTag) string {
	description := tag.Get("description")
	required := tag.Get("required")
	usage := ""
	if required == "yes" {
		usage = "(required) "
	} else {
		usage = "(optional) "
	}
	return usage + description
}

func (m FlagMapper) Map(flags interface{}, opts interface{}) {
	flagsReflected := reflect.ValueOf(flags).Elem()
	optsReflected := reflect.ValueOf(opts).Elem()

	for i := 0; i < flagsReflected.NumField(); i++ {
		flag := flagsReflected.Field(i)
		tag := flagsReflected.Type().Field(i).Tag

		flagName := tag.Get("flag_name")
		flagChanged := m.cmd.PersistentFlags().Changed(flagName) // flagChanged --> value for flag has been set on command line

		// see https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
		// see https://stackoverflow.com/questions/40060131/reflect-assign-a-pointer-struct-value
		if flagChanged {
			fieldName := flagsReflected.Type().Field(i).Name
			opt := optsReflected.FieldByName(fieldName)
			mapOpt(opt, tag, m, flagName, flag, fieldName)
			mapFlag(flag, m, flagName)
		} else {
			// TODO what do we want to do, if flag is not changed / given by command
			// TODO e.g. provide default value in annotation?
		}
	}
}

func mapFlag(value reflect.Value, mapper FlagMapper, tagName string) {
	mapValue(value, mapper, tagName, value)
}

func mapOpt(opt reflect.Value, tag reflect.StructTag, mapper FlagMapper, flagName string, value reflect.Value, fieldName string) {
	if opt.IsValid() {
		// A Value can be changed only if it is addressable and was not obtained by the use of unexported struct fields.
		if opt.CanSet() {
			if transform := tag.Get("transform"); transform != "" {
				value, err := mapper.cmd.PersistentFlags().GetString(flagName)
				if err != nil {
					panic(err.Error())
				}
				transformAndSet(transform, opt, value)
			} else {
				mapValue(value, mapper, flagName, opt)
			}
		} else {
			panic(fieldName + " can not be set")
		}
	} else {
		// for the moment, we want to ignore flags, that are not available in opts
		// panic(fieldName + " is not valid")
	}
}

func mapValue(value reflect.Value, mapper FlagMapper, flagName string, opt reflect.Value) {
	switch value.Type().String() {
	case "*int":
		mapInt(mapper, flagName, opt)
	case "*string":
		mapString(mapper, flagName, opt)
	case "*bool":
		mapBool(mapper, flagName, opt)
	case "*[]string":
		mapStringArray(mapper, flagName, opt)
	default:
		panic("Unknown type " + value.Type().String())
	}
}

func mapInt(m FlagMapper, flagName string, opt reflect.Value) {
	value, err := m.cmd.PersistentFlags().GetInt(flagName)
	if err != nil {
		panic(err.Error())
	}
	if typesMatch(opt, &value) {
		opt.Set(reflect.ValueOf(&value))
	}
}

func mapString(m FlagMapper, flagName string, opt reflect.Value) {
	value, err := m.cmd.PersistentFlags().GetString(flagName)
	if err != nil {
		panic(err.Error())
	}
	if typesMatch(opt, &value) {
		opt.Set(reflect.ValueOf(&value))
	}
}

func mapStringArray(m FlagMapper, flagName string, opt reflect.Value) {
	value, err := m.cmd.PersistentFlags().GetStringArray(flagName)
	if err != nil {
		panic(err.Error())
	}
	if typesMatch(opt, &value) {
		opt.Set(reflect.ValueOf(&value))
	}
}

func mapBool(m FlagMapper, flagName string, opt reflect.Value) {
	value, err := m.cmd.PersistentFlags().GetBool(flagName)
	if err != nil {
		panic(err.Error())
	}
	if typesMatch(opt, &value) {
		opt.Set(reflect.ValueOf(&value))
	}
}

func transformAndSet(transform string, opt reflect.Value, value string) {
	fieldType := opt.Type()

	transformedValue, err := call(funcs, transform, value)
	if err != nil {
		panic(err.Error())
	}

	opt.Set(transformedValue[0].Convert(fieldType))
}

func str2Visibility(s string) *gitlab.VisibilityValue {
	if s == "private" {
		return gitlab.Visibility(gitlab.PrivateVisibility)
	}
	if s == "internal" {
		return gitlab.Visibility(gitlab.InternalVisibility)
	}
	if s == "public" {
		return gitlab.Visibility(gitlab.PublicVisibility)
	}
	return nil
}

func string2IsoTime(s string) *gitlab.ISOTime {
	iso8601 := "2006-01-02"
	isotime, err := time.Parse(`"`+iso8601+`"`, string(s))
	if err != nil {
		panic(err.Error())
	}
	t := gitlab.ISOTime(isotime)
	return &t
}

func str2AccessLevel(s string) *gitlab.AccessLevelValue {
	if s == "10" {
		return gitlab.AccessLevel(gitlab.GuestPermissions)
	}
	if s == "20" {
		return gitlab.AccessLevel(gitlab.ReporterPermissions)
	}
	if s == "30" {
		return gitlab.AccessLevel(gitlab.DeveloperPermissions)
	}
	if s == "40" {
		return gitlab.AccessLevel(gitlab.MasterPermissions)
	}
	if s == "50" {
		return gitlab.AccessLevel(gitlab.OwnerPermission)
	}
	panic("Unknown access level: " + s)
}

var funcs = map[string]interface{}{
	"string2visibility": str2Visibility,
	"string2IsoTime":    string2IsoTime,
	"str2AccessLevel":   str2AccessLevel,
}

func call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is not adapted")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func typesMatch(target reflect.Value, source interface{}) bool {
	targetType := target.Type().String()
	sourceType := reflect.ValueOf(source).Type().String()

	return targetType == sourceType
}
