package mapper

import (
	"github.com/spf13/cobra"
	"reflect"
)

type FlagMapper struct {
	cmd *cobra.Command

}

func New(cmd *cobra.Command) FlagMapper {
	return FlagMapper{ cmd: cmd }
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
			m.cmd.PersistentFlags().IntP(flagName, shortHand, 0, tag.Get("description"))
		case "*string":
			m.cmd.PersistentFlags().StringP(flagName, shortHand,  "",tag.Get("description"))
		case "*bool":
			m.cmd.PersistentFlags().BoolP(flagName, shortHand,false, tag.Get("description"))
		case "*[]string":
			m.cmd.PersistentFlags().StringArrayP(flagName, shortHand,nil, tag.Get("description"))
		default:
			panic("Unknown type " + f.Type().String())
		}
	}
}



func (m FlagMapper) Map(flags interface{}, opts interface{}) {
	flagsReflected := reflect.ValueOf(flags).Elem()
	optsReflected := reflect.ValueOf(opts).Elem()

	for i := 0; i < flagsReflected.NumField(); i++ {
		f := flagsReflected.Field(i)
		tag := flagsReflected.Type().Field(i).Tag

		flagName := tag.Get("flag_name")
		flagChanged := m.cmd.PersistentFlags().Changed(flagName)   // flagChanged --> value for flag has been set on command line

		// see https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
		// see https://stackoverflow.com/questions/40060131/reflect-assign-a-pointer-struct-value
		if f.IsValid() {
			// A Value can be changed only if it is addressable and was not obtained by the use of unexported struct fields.
			if f.CanSet() {
				if flagChanged {
					fieldName := flagsReflected.Type().Field(i).Name
					opt := optsReflected.FieldByName(fieldName)
					switch f.Type().String() {
					case "*int":
						value, err := m.cmd.PersistentFlags().GetInt(flagName)
						if err != nil { panic(err.Error()) }
						opt.Set(reflect.ValueOf(&value))
					case "*string":
						value, err := m.cmd.PersistentFlags().GetString(flagName)
						if err != nil { panic(err.Error()) }
						opt.Set(reflect.ValueOf(&value))
					case "*bool":
						value, err := m.cmd.PersistentFlags().GetBool(flagName)
						if err != nil { panic(err.Error()) }
						opt.Set(reflect.ValueOf(&value))
					case "*[]string":
						value, err := m.cmd.PersistentFlags().GetStringArray(flagName)
						if err != nil { panic(err.Error()) }
						opt.Set(reflect.ValueOf(&value))
					default:
						panic("Unknown type " + f.Type().String())
					}

				} else {
					// TODO implement an additional tag that allows setting of "default" values
				}
			} else {
				panic("can not set " + flagName)
			}
		} else {
			panic(flagName + "is not valid")
		}
	}
}