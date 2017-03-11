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