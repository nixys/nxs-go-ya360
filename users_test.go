package ya360

import (
	"os"
	"strconv"
	"testing"
)

var (
	testUserFirstName = "TestUserFirstName"
	testUserLastName  = "TestUserLastName"
	testUserNickame   = "testusernickame"
	testUserPassword  = "TestUserPassword"

	testUserUpdatedFirstName = "TestUserUpdatedFirstName"
	testUserUpdatedLastName  = "TestUserUpdatedLastName"

	testUserAlias = "testusernickamealias"
)

func TestUsersCRUD(t *testing.T) {

	oAuth := os.Getenv("YA360_OAUTH")
	orgID, err := strconv.ParseInt(os.Getenv("YA360_ORG_ID"), 10, 64)
	if err != nil {
		t.Fatal("Init error: make sure environment variable `YA360_ORG_ID` correctly defined:", err)
	}
	if len(oAuth) == 0 {
		t.Fatal("Init error: make sure environment variable `YA360_OAUTH` correctly defined")
	}

	y := Init(Settings{
		OAuth: oAuth,
		OrgID: orgID,
	})

	uCreated := testUserCreate(t, y)
	defer testUserDetele(t, y, uCreated.ID)

	testUserGet(t, y, uCreated.ID)
	testUsersList(t, y)

	testUserUpdate(t, y, uCreated.ID)
	testUserAliasAdd(t, y, uCreated.ID)
	testUserAliasDelete(t, y, uCreated.ID)
}

func testUserCreate(t *testing.T, y Ya360) UserRx {

	u, err := y.UserCreate(UserCreateTx{
		Name: UserName{
			First: testUserFirstName,
			Last:  testUserLastName,
		},
		Nickname:     testUserNickame,
		Password:     testUserPassword,
		DepartmentID: 1,
	})
	if err != nil {
		t.Fatal("User create error:", err)
	}

	if u.Name.First != testUserFirstName || u.Name.Last != testUserLastName || u.Nickname != testUserNickame {
		t.Fatalf("User create error: incorrect first name, last name or nickname (returned: %s, %s, %s)", u.Name.First, u.Name.Last, u.Nickname)
	}

	t.Logf("User create: success")

	return u
}

func testUserGet(t *testing.T, y Ya360, userID string) {

	u, err := y.UserGet(userID)
	if err != nil {
		t.Fatal("User get error:", err)
	}

	if u.ID != userID {
		t.Fatal("User get error: incorrect ID")
	}

	t.Logf("User get: success")
}

func testUsersList(t *testing.T, y Ya360) {

	u, err := y.UsersList(1, 1000)
	if err != nil {
		t.Fatal("Users list error:", err)
	}

	for _, e := range u.Users {
		if e.Nickname == testUserNickame {
			t.Logf("Users list: success")
			return
		}
	}

	t.Fatal("Users list error: created user not found")
}

func testUserUpdate(t *testing.T, y Ya360, userID string) {

	u, err := y.UserUpdate(userID, UserUpdateTx{
		Name: UserName{
			First: testUserUpdatedFirstName,
			Last:  testUserUpdatedLastName,
		},
	})
	if err != nil {
		t.Fatal("User update error:", err)
	}

	if u.Name.First != testUserUpdatedFirstName || u.Name.Last != testUserUpdatedLastName {
		t.Fatalf("User update error: incorrect new first or last name (returned: %s, %s)", u.Name.First, u.Name.Last)
	}

	t.Logf("User update: success")
}

func testUserAliasAdd(t *testing.T, y Ya360, userID string) {

	u, err := y.UserAliasAdd(userID, UserAliasAddTx{
		Alias: testUserAlias,
	})
	if err != nil {
		t.Fatal("User add alias error:", err)
	}

	for _, a := range u.Aliases {
		if a == testUserAlias {
			t.Logf("User add alias: success")
			return
		}
	}

	t.Fatal("User add alias error: created alias for user not found")
}

func testUserAliasDelete(t *testing.T, y Ya360, userID string) {

	u, err := y.UserAliasDelete(userID, testUserAlias)
	if err != nil {
		t.Fatal("User delete alias error:", err)
	}

	if u.Alias != testUserAlias {
		t.Fatal("User delete alias error: alias not found for user")
	}

	t.Logf("User delete alias: success")
}

func testUserDetele(t *testing.T, y Ya360, userID string) {
	t.Logf("User must be deleted manually. Not implemented yet in Yandex 360")
}
