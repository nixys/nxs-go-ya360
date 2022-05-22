package ya360

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// UsersRx contains users list
type UsersRx struct {
	Users   []UserRx `json:"users"`
	Page    int64    `json:"page"`
	Pages   int64    `json:"pages"`
	PerPage int64    `json:"perPage"`
	Total   int64    `json:"total"`
}

// UserRx contains user data
type UserRx struct {
	About        string          `json:"about"`
	Aliases      []string        `json:"aliases"`
	AvatarID     string          `json:"avatarId"`
	Birthday     string          `json:"birthday"`
	Contacts     []UserContactRx `json:"contacts"`
	CreatedAt    string          `json:"createdAt"`
	DepartmentID int64           `json:"departmentId"`
	Email        string          `json:"email"`
	ExternalID   string          `json:"externalId"`
	Gender       string          `json:"gender"`
	Groups       []int64         `json:"groups"`
	ID           string          `json:"id"`
	IsAdmin      bool            `json:"isAdmin"`
	IsDismissed  bool            `json:"isDismissed"`
	IsEnabled    bool            `json:"isEnabled"`
	IsRobot      bool            `json:"isRobot"`
	Language     string          `json:"language"`
	Name         UserName        `json:"name"`
	Nickname     string          `json:"nickname"`
	Position     string          `json:"position"`
	Timezone     string          `json:"timezone"`
	UpdatedAt    string          `json:"updatedAt"`
}

// UserName contains user name data
type UserName struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Middle string `json:"middle"`
}

// UserContactRx contains user contacts
type UserContactRx struct {
	Alias     bool            `json:"alias"`
	Main      bool            `json:"main"`
	Synthetic bool            `json:"synthetic"`
	Type      UserContactType `json:"type"`
	Value     string          `json:"value"`
}

// UserDeleteRx contains result of user delete operation
type UserDeleteRx struct {
	ID      string `json:"id"`
	Removed bool   `json:"removed"`
}

// UserAliasDeleteRx contains result of user alias delete operation
type UserAliasDeleteRx struct {
	Alias   string `json:"alias"`
	Removed bool   `json:"removed"`
}

// UserCreateTx contains data to create new user
type UserCreateTx struct {
	About        string          `json:"about,omitempty"`
	Birthday     string          `json:"birthday,omitempty"`
	Contacts     []UserContactTx `json:"contacts,omitempty"`
	DepartmentID int64           `json:"departmentId"`
	ExternalID   string          `json:"externalId,omitempty"`
	Gender       string          `json:"gender,omitempty"`
	IsAdmin      bool            `json:"isAdmin,omitempty"`
	Language     string          `json:"language,omitempty"`
	Name         UserName        `json:"name,omitempty"`
	Nickname     string          `json:"nickname"`
	Password     string          `json:"password"`
	Position     string          `json:"position,omitempty"`
	Timezone     string          `json:"timezone,omitempty"`
}

// UserUpdateTx contains data to update user
type UserUpdateTx struct {
	About                  string          `json:"about,omitempty"`
	Birthday               string          `json:"birthday,omitempty"`
	Contacts               []UserContactTx `json:"contacts,omitempty"`
	DepartmentID           int64           `json:"departmentId,omitempty"`
	ExternalID             string          `json:"externalId,omitempty"`
	Gender                 string          `json:"gender,omitempty"`
	IsAdmin                bool            `json:"isAdmin,omitempty"`
	IsEnabled              bool            `json:"isEnabled,omitempty"`
	Language               string          `json:"language,omitempty"`
	Name                   UserName        `json:"name,omitempty"`
	Nickname               string          `json:"nickname,omitempty"`
	Password               string          `json:"password,omitempty"`
	PasswordChangeRequired bool            `json:"passwordChangeRequired,omitempty"`
	Position               string          `json:"position,omitempty"`
	Timezone               string          `json:"timezone,omitempty"`
}

// UserContactTx contains user contacts for transmit operations
type UserContactTx struct {
	Type  UserContactType `json:"type"`
	Value string          `json:"value"`
}

// UserAliasAddTx contains data to add new alias for user
type UserAliasAddTx struct {
	Alias string `json:"alias"`
}

type UserContactType string

const (
	UserContactTypeEmail          UserContactType = "email"
	UserContactTypePhoneExtension UserContactType = "phone_extension"
	UserContactTypePhone          UserContactType = "phone"
	UserContactTypeSite           UserContactType = "site"
	UserContactTypeICQ            UserContactType = "icq"
	UserContactTypeTwitter        UserContactType = "twitter"
	UserContactTypeSkype          UserContactType = "skype"
)

func (t UserContactType) String() string {
	return string(t)
}

// UserCreate creates new user
// Link: https://yandex.ru/dev/api360/doc/ref/UserService/UserService_Create.html
func (ya *Ya360) UserCreate(user UserCreateTx) (UserRx, error) {

	var (
		resp UserRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users", ya.s.OrgID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPost, ur, user, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// UserGet gets specified user
// Link: https://yandex.ru/dev/api360/doc/ref/UserService/UserService_Get.html
func (ya *Ya360) UserGet(userID string) (UserRx, error) {

	var (
		resp UserRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users/%s", ya.s.OrgID, userID),
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

// UsersList gets users list
// Link: https://yandex.ru/dev/api360/doc/ref/UserService/UserService_List.html
func (ya *Ya360) UsersList(page, perPage int64) (UsersRx, error) {

	var (
		resp UsersRx
	)

	urlParams := url.Values{}

	urlParams.Add("page", strconv.FormatInt(page, 10))
	urlParams.Add("perPage", strconv.FormatInt(perPage, 10))

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users", ya.s.OrgID),
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

// UserUpdate updates specified user with new `user` data
// Link: https://yandex.ru/dev/api360/doc/ref/UserService/UserService_Update.html
func (ya *Ya360) UserUpdate(userID string, user UserUpdateTx) (UserRx, error) {

	var (
		resp UserRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users/%s", ya.s.OrgID, userID),
		RawQuery: urlParams.Encode(),
	}

	status, err := ya.alter(http.MethodPatch, ur, user, &resp)
	if err != nil {
		return resp, Error{
			Code: status,
			Text: err.Error(),
		}
	}

	return resp, nil
}

// UserAliasAdd adds new alias to specified user
// Link: https://yandex.ru/dev/api360/doc/ref/UserService/UserService_CreateUserAlias.html
func (ya *Ya360) UserAliasAdd(userID string, alias UserAliasAddTx) (UserRx, error) {

	var (
		resp UserRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users/%s/aliases", ya.s.OrgID, userID),
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

// UserAliasDelete deletes alias from specified user
// Link: https://yandex.ru/dev/api360/doc/ref/UserService/UserService_DeleteUserAlias.html
func (ya *Ya360) UserAliasDelete(userID, alias string) (UserAliasDeleteRx, error) {

	var (
		resp UserAliasDeleteRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users/%s/aliases/%s", ya.s.OrgID, userID, alias),
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

// UserDelete deletes user
// Link: not implemented yet in Yandex 360
func (ya *Ya360) UserDelete(userID string) (UserDeleteRx, error) {

	var (
		resp UserDeleteRx
	)

	urlParams := url.Values{}

	ur := url.URL{
		Path:     fmt.Sprintf("/directory/v1/org/%d/users/%s", ya.s.OrgID, userID),
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
