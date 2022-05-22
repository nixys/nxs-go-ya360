package ya360

import (
	"os"
	"strconv"
	"testing"
)

var (
	testGroupName        = "TestGroupName"
	testGroupDescription = "Test group description"

	testGroupUpdatedName = "TestGroupUpdatedName"
)

func TestGroupsCRUD(t *testing.T) {

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

	gCreated := testGroupCreate(t, y)
	defer testGroupDetele(t, y, gCreated.ID)

	uCreated := testUserCreate(t, y)
	defer testUserDetele(t, y, uCreated.ID)

	testGroupGet(t, y, gCreated.ID)
	testGroupsList(t, y)

	testGroupUpdate(t, y, gCreated.ID)

	testGroupMemberAdd(t, y, gCreated.ID, uCreated.ID)
	testGroupMemberDel(t, y, gCreated.ID)
}

func testGroupCreate(t *testing.T, y Ya360) GroupRx {

	g, err := y.GroupCreate(GroupCreateTx{
		Name:        testGroupName,
		Description: testGroupDescription,
	})
	if err != nil {
		t.Fatal("Group create error:", err)
	}

	if g.Name != testGroupName {
		t.Fatal("Group create error: incorrect name")
	}

	t.Logf("Group create: success")

	return g
}

func testGroupGet(t *testing.T, y Ya360, groupID int64) {

	g, err := y.GroupGet(groupID)
	if err != nil {
		t.Fatal("Group get error:", err)
	}

	if g.ID != groupID {
		t.Fatal("Group get error: incorrect ID")
	}

	if g.Name != testGroupName {
		t.Fatal("Group get error: incorrect name")
	}

	t.Logf("Group get: success")
}

func testGroupsList(t *testing.T, y Ya360) {

	g, err := y.GroupsList(1, 1000)
	if err != nil {
		t.Fatal("Groups list error:", err)
	}

	for _, e := range g.Groups {
		if e.Name == testGroupName {
			t.Logf("Groups list: success")
			return
		}
	}

	t.Fatal("Groups list error: created group not found")
}

func testGroupUpdate(t *testing.T, y Ya360, groupID int64) {

	g, err := y.GroupUpdate(groupID, GroupUpdateTx{
		Name: testGroupUpdatedName,
	})
	if err != nil {
		t.Fatal("Group update error:", err)
	}

	if g.Name != testGroupUpdatedName {
		t.Fatalf("Group update error: incorrect new name (returned: %s)", g.Name)
	}

	t.Logf("Group update: success")
}

func testGroupMemberAdd(t *testing.T, y Ya360, groupID int64, userID string) {

	g, err := y.GroupMemberAdd(groupID, GroupMemberAddTx{
		ID:   userID,
		Type: MemberTypeUser,
	})
	if err != nil {
		t.Fatal("Group add member error:", err)
	}

	if g.ID != userID || g.Type != MemberTypeUser {
		t.Fatalf("Group add member error: incorrect added member (returned: %s, %s)", g.ID, g.Type)
	}

	t.Logf("Group add member: success")
}

func testGroupMemberDel(t *testing.T, y Ya360, groupID int64) {

	g, err := y.GroupUpdate(groupID, GroupUpdateTx{
		Members: []MemberIDType{},
	})
	if err != nil {
		t.Fatal("Group delete member error:", err)
	}

	if len(g.Members) > 0 {
		t.Fatalf("Group delete member error: incorrect members count (returned: %d)", len(g.Members))
	}

	t.Logf("Group delete member: success")
}

func testGroupDetele(t *testing.T, y Ya360, groupID int64) {

	g, err := y.GroupDelete(groupID)
	if err != nil {
		t.Fatal("Group delete error:", err)
	}

	if g.ID != groupID {
		t.Fatal("Group delete error: incorrect ID")
	}

	t.Logf("Group delete: success")
}
