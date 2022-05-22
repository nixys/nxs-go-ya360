package ya360

import "fmt"

// Ya360 contains Yandex 360 parameters
type Ya360 struct {
	s Settings
}

// Settings contain settings for node connections
type Settings struct {
	URL   string
	OAuth string
	OrgID int64
}

type MemberIDType struct {
	ID   string     `json:"id"`
	Type MemberType `json:"type"`
}

type Error struct {
	Code int
	Text string
}

type errorRx struct {
	Code    int              `json:"code"`
	Details []errorRxDetails `json:"details"`
	Message string           `json:"message"`
}

type errorRxDetails struct {
	Type string `json:"@type"`
}

type MemberType string

const (
	MemberTypeUser       MemberType = "user"
	MemberTypeGroup      MemberType = "group"
	MemberTypeDepartment MemberType = "department"
)

func (t MemberType) String() string {
	return string(t)
}

const YaHostDefault = "https://api360.yandex.net"

// Init returns parametrized Node object
func Init(s Settings) Ya360 {

	if len(s.URL) == 0 {
		s.URL = YaHostDefault
	}

	return Ya360{
		s: s,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%d, %s", e.Code, e.Text)
}
