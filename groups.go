package ya360

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GroupsRx contains groups list
type GroupsRx struct {
	Groups  []GroupRx `json:"groups"`
	Page    int64     `json:"page"`
	Pages   int64     `json:"pages"`
	PerPage int64     `json:"perPage"`
	Total   int64     `json:"total"`
}

// GroupRx contains group data
type GroupRx struct {
	AdminIDs     []string       `json:"adminIds"`
	Aliases      []string       `json:"aliases"`
	AuthorID     string         `json:"authorId"`
	CreatedAt    string         `json:"createdAt"`
	Description  string         `json:"description"`
	Email        string         `json:"email"`
	ExternalID   string         `json:"externalId"`
	ID           int64          `json:"id"`
	Label        string         `json:"label"`
	MemberOf     []int64        `json:"memberOf"`
	Members      []MemberIDType `json:"members"`
	MembersCount int64          `json:"membersCount"`
	Name         string         `json:"name"`
	Removed      bool           `json:"removed"`
	Type         string         `json:"type"`
}

// GroupMemberAddRx contains result of group member add operation
type GroupMemberAddRx struct {
	Added bool       `json:"added"`
	ID    string     `json:"id"`
	Type  MemberType `json:"type"`
}

// GroupMembersListRx contains group member lists
type GroupMembersListRx struct {
	Departments GroupMemberDepartmentRx `json:"departments"`
	Groups      GroupMemberGroupRx      `json:"groups"`
	Users       GroupMemberUserRx       `json:"users"`
}

// GroupMemberDepartmentRx contains group member for type `department`
type GroupMemberDepartmentRx struct {
	ID           int64  `json:"id"`
	MembersCount int64  `json:"membersCount"`
	Name         string `json:"name"`
}

// GroupMemberGroupRx contains group member for type `group`
type GroupMemberGroupRx struct {
	ID           int64  `json:"id"`
	MembersCount int64  `json:"membersCount"`
	Name         string `json:"name"`
}

// GroupMemberGroupRx contains group member for type `user`
type GroupMemberUserRx struct {
	AvatarId     string   `json:"avatarId"`
	DepartmentId int64    `json:"departmentId"`
	Email        string   `json:"email"`
	Gender       string   `json:"gender"`
	ID           string   `json:"id"`
	Name         UserName `json:"name"`
	Nickname     string   `json:"nickname"`
	Position     string   `json:"position"`
}

// GroupDeleteRx contains result of group delete operation
type GroupDeleteRx struct {
	ID      int64 `json:"id"`
	Removed bool  `json:"removed"`
}

// GroupCreateTx contains data to create new group
type GroupCreateTx struct {
	AdminIDs    []string       `json:"adminIds,omitempty"`
	Description string         `json:"description,omitempty"`
	ExternalID  string         `json:"externalId,omitempty"`
	Label       string         `json:"label,omitempty"`
	Members     []MemberIDType `json:"members,omitempty"`
	Name        string         `json:"name"`
}

// GroupUpdateTx contains data to update group
type GroupUpdateTx struct {
	AdminIDs    []string       `json:"adminIds,omitempty"`
	Description string         `json:"description,omitempty"`
	ExternalID  string         `json:"externalId,omitempty"`
	Label       string         `json:"label,omitempty"`
	Members     []MemberIDType `json:"members,omitempty"`
	Name        string         `json:"name,omitempty"`
}

// GroupMemberAddTx contains data to add new member for group
type GroupMemberAddTx struct {
	ID   string     `json:"id"`
	Type MemberType `json:"type,omitempty"`
}

// GroupCreate creates new group
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_Create.html
func (ya *Ya360) GroupCreate(group GroupCreateTx) (GroupRx, error) {

	var (
		resp GroupRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups", ya.s.OrgID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPost, ur, group, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// GroupGet gets specified group
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_Get.html
func (ya *Ya360) GroupGet(groupID int64) (GroupRx, error) {

	var (
		resp GroupRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups/%d", ya.s.OrgID, groupID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.get(ur, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// GroupsList gets groups list
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_List.html
func (ya *Ya360) GroupsList(page, perPage int64) (GroupsRx, error) {

	var (
		resp GroupsRx
	)

	urlParams := url.Values{}

	urlParams.Add("page", strconv.FormatInt(page, 10))
	urlParams.Add("perPage", strconv.FormatInt(perPage, 10))

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups", ya.s.OrgID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.get(ur, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// GroupMembersList adds members list for specified group
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_ListMembers.html
func (ya *Ya360) GroupMembersList(groupID int64) (GroupMembersListRx, error) {

	var (
		resp GroupMembersListRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups/%d/members", ya.s.OrgID, groupID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.get(ur, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// GroupUpdate updates specified group with new `group` data
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_Update.html
func (ya *Ya360) GroupUpdate(groupID int64, group GroupUpdateTx) (GroupRx, error) {

	var (
		resp GroupRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups/%d", ya.s.OrgID, groupID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPatch, ur, group, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// GroupMemberAdd adds new member into specified group
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_AddMember.html
func (ya *Ya360) GroupMemberAdd(groupID int64, member GroupMemberAddTx) (GroupMemberAddRx, error) {

	var (
		resp GroupMemberAddRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups/%d/members", ya.s.OrgID, groupID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPost, ur, member, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// GroupDelete deletes group
// Link: https://yandex.ru/dev/api360/doc/ref/GroupService/GroupService_Delete.html
func (ya *Ya360) GroupDelete(groupID int64) (GroupDeleteRx, error) {

	var (
		resp GroupDeleteRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/groups/%d", ya.s.OrgID, groupID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodDelete, ur, nil, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}
