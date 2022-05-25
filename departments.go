package ya360

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// DepartmentsRx contains departments list
type DepartmentsRx struct {
	Departments []DepartmentRx `json:"departments"`
	Page        int64          `json:"page"`
	Pages       int64          `json:"pages"`
	PerPage     int64          `json:"perPage"`
	Total       int64          `json:"total"`
}

// DepartmentRx contains department data
type DepartmentRx struct {
	Aliases      []string `json:"aliases"`
	CreatedAt    string   `json:"createdAt"`
	Description  string   `json:"description"`
	Email        string   `json:"email"`
	ExternalID   string   `json:"externalId"`
	HeadID       string   `json:"headId"`
	ID           int64    `json:"id"`
	Label        string   `json:"label"`
	MembersCount int64    `json:"membersCount"`
	Name         string   `json:"name"`
	ParentID     int64    `json:"parentId"`
}

// DepartmentDeleteRx contains result of department delete operation
type DepartmentDeleteRx struct {
	ID      int64 `json:"id"`
	Removed bool  `json:"removed"`
}

// DepartmentAliasDeleteRx contains result of department alias delete operation
type DepartmentAliasDeleteRx struct {
	Alias   string `json:"alias"`
	Removed bool   `json:"removed"`
}

// DepartmentCreateTx contains data to create new department
type DepartmentCreateTx struct {
	Description string `json:"description,omitempty"`
	ExternalID  string `json:"externalId,omitempty"`
	HeadID      string `json:"headId,omitempty"`
	Label       string `json:"label,omitempty"`
	Name        string `json:"name"`
	ParentID    int64  `json:"parentId"`
}

// DepartmentUpdateTx contains data to update department
type DepartmentUpdateTx struct {
	Description string `json:"description,omitempty"`
	ExternalID  string `json:"externalId,omitempty"`
	HeadID      string `json:"headId,omitempty"`
	Label       string `json:"label,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentID    int64  `json:"parentId,omitempty"`
}

// DepartmentAliasAddTx contains data to add new alias for department
type DepartmentAliasAddTx struct {
	Alias string `json:"alias"`
}

// DepartmentCreate creates new department
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_Create.html
func (ya *Ya360) DepartmentCreate(department DepartmentCreateTx) (DepartmentRx, error) {

	var (
		resp DepartmentRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments", ya.s.OrgID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPost, ur, department, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// DepartmentGet gets specified department
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_Get.html
func (ya *Ya360) DepartmentGet(departmentID int64) (DepartmentRx, error) {

	var (
		resp DepartmentRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments/%d", ya.s.OrgID, departmentID),
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

// DepartmentsList gets departments list
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_List.html
func (ya *Ya360) DepartmentsList(page, perPage, parentId int64, orderBy string) (DepartmentsRx, error) {

	var (
		resp DepartmentsRx
	)

	urlParams := url.Values{}

	urlParams.Add("page", strconv.FormatInt(page, 10))
	urlParams.Add("perPage", strconv.FormatInt(perPage, 10))
	urlParams.Add("parentId", strconv.FormatInt(parentId, 10))
	urlParams.Add("orderBy", orderBy)

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments", ya.s.OrgID),
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

// DepartmentUpdate updates specified department with new `department` data
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_Update.html
func (ya *Ya360) DepartmentUpdate(departmentID int64, department DepartmentUpdateTx) (DepartmentRx, error) {

	var (
		resp DepartmentRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments/%d", ya.s.OrgID, departmentID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPatch, ur, department, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// DepartmentAliasAdd adds new alias to specified department
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_CreateAlias.html
func (ya *Ya360) DepartmentAliasAdd(departmentID int64, alias DepartmentAliasAddTx) (DepartmentRx, error) {

	var (
		resp DepartmentRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments/%d/aliases", ya.s.OrgID, departmentID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPost, ur, alias, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// DepartmentAliasDelete deletes alias from specified department
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_DeleteAlias.html
func (ya *Ya360) DepartmentAliasDelete(departmentID int64, alias string) (DepartmentAliasDeleteRx, error) {

	var (
		resp DepartmentAliasDeleteRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments/%d/aliases/%s", ya.s.OrgID, departmentID, alias),
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

// DepartmentDelete deletes department
// Link: https://yandex.ru/dev/api360/doc/ref/DepartmentService/DepartmentService_Delete.html
func (ya *Ya360) DepartmentDelete(departmentID int64) (DepartmentDeleteRx, error) {

	var (
		resp DepartmentDeleteRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/departments/%d", ya.s.OrgID, departmentID),
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
