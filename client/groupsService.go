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

package client

import (
	. "github.com/michaellihs/golab/model"
	"strconv"
)

type GroupListParams struct {
	skip_groups   []int  // (optional) Skip the group IDs passes
	all_available bool   // (optional) Show all the groups you have access to
	search        string // (optional) Return list of authorized groups matching the search criteria
	order_by      string // (optional) Order groups by name or path.Default is name
	sort          string // (optional) Order groups in asc or desc order.Default is asc
	statistics    bool   // (optional) Include group statistics (admins only)
	owned         bool   // (optional) Limit by groups owned by the current user
}

type GroupsService struct {
	Client *GitlabClient
}

func (groupService *GroupsService) GetGroup(id string) (*Group, error) {
	group := new(Group)
	req, err := groupService.Client.NewGetRequest("/groups/" + id)
	if err != nil {
		return nil, err
	}
	_, err2 := groupService.Client.Do(req, group)
	if err2 != nil {
		return nil, err2
	}
	return group, nil
}

func (groupsService *GroupsService) Namespacify(group string) (int, error) {
	if group == "" {
		return 0, nil
	}

	// if group is an int, it's already a namespace id
	if namespace_id, err := strconv.Atoi(group); err == nil {
		return namespace_id, nil
	}

	// if group is a string, we have to resolve group id
	groupInfo, err := groupsService.GetGroup(group)
	if err != nil {
		return 0, err
	}

	return groupInfo.ID, nil
}