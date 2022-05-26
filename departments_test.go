package ya360

import (
	"os"
	"strconv"
	"testing"
)

var (
	testDepartmentName        = "TestDepartmentName"
	testDepartmentUpdatedName = "TestDepartmentUpdatedName"

	testDepartmentLabel = "testdepartmentemail"
	testDepartmentAlias = "testdepartmentalias"
)

func TestDepartmentsCRUD(t *testing.T) {

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

	dCreated := testDepartmentCreate(t, y)
	defer testDepartmentDetele(t, y, dCreated.ID)

	testDepartmentGet(t, y, dCreated.ID)
	testDepartmentsList(t, y)

	testDepartmentUpdate(t, y, dCreated.ID)

	testDepartmentAliasAdd(t, y, dCreated.ID)
	testDepartmentAliasDelete(t, y, dCreated.ID)

}

func testDepartmentCreate(t *testing.T, y Ya360) DepartmentRx {

	d, err := y.DepartmentCreate(DepartmentCreateTx{
		Name:     testDepartmentName,
		ParentID: 1,
		Label:    testDepartmentLabel,
	})
	if err != nil {
		t.Fatal("Department create error:", err)
	}

	if d.Name != testDepartmentName {
		t.Fatal("Department create error: incorrect name")
	}

	t.Logf("Department create: success")

	return d
}

func testDepartmentGet(t *testing.T, y Ya360, departmentID int64) {

	d, err := y.DepartmentGet(departmentID)
	if err != nil {
		t.Fatal("Department get error:", err)
	}

	if d.ID != departmentID {
		t.Fatal("Department get error: incorrect ID")
	}

	if d.Name != testDepartmentName {
		t.Fatal("Department get error: incorrect name")
	}

	t.Logf("Department get: success")
}

func testDepartmentsList(t *testing.T, y Ya360) {

	d, err := y.DepartmentsList(1, 1000, 0, OrderByID)
	if err != nil {
		t.Fatal("Departments list error:", err)
	}

	for _, e := range d.Departments {
		if e.Name == testDepartmentName {
			t.Logf("Departments list: success")
			return
		}
	}

	t.Fatal("Departments list error: created department not found")
}

func testDepartmentUpdate(t *testing.T, y Ya360, departmentID int64) {

	d, err := y.DepartmentUpdate(departmentID, DepartmentUpdateTx{
		Name: testDepartmentUpdatedName,
	})
	if err != nil {
		t.Fatal("Department update error:", err)
	}

	if d.Name != testDepartmentUpdatedName {
		t.Fatalf("Department update error: incorrect new name (returned: %s)", d.Name)
	}

	t.Logf("Department update: success")
}

func testDepartmentAliasAdd(t *testing.T, y Ya360, departmentID int64) {

	d, err := y.DepartmentAliasAdd(departmentID, DepartmentAliasAddTx{
		Alias: testDepartmentAlias,
	})
	if err != nil {
		t.Fatal("Department add alias error:", err)
	}

	for _, a := range d.Aliases {
		if a == testDepartmentAlias {
			t.Logf("Department add alias: success")
			return
		}
	}

	t.Fatal("Department add alias error: created alias for department not found")
}

func testDepartmentAliasDelete(t *testing.T, y Ya360, departmentID int64) {

	d, err := y.DepartmentAliasDelete(departmentID, testDepartmentAlias)
	if err != nil {
		t.Fatal("Department delete alias error:", err)
	}

	if d.Alias != testDepartmentAlias {
		t.Fatal("Department delete alias error: alias not found for department")
	}

	t.Logf("Department delete alias: success")
}

func testDepartmentDetele(t *testing.T, y Ya360, departmentID int64) {

	d, err := y.DepartmentDelete(departmentID)
	if err != nil {
		t.Fatal("Department delete error:", err)
	}

	if d.ID != departmentID {
		t.Fatal("Department delete error: incorrect ID")
	}

	t.Logf("Department delete: success")
}
