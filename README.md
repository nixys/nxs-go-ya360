# nxs-go-ya360

This Go package provides access to Yandex 360 API.

Follows Yandex 360 resources are fully implemented at this moment:
- [DepartmentService](https://yandex.ru/dev/api360/doc/ref/DepartmentService.html)
- [GroupService](https://yandex.ru/dev/api360/doc/ref/GroupService.html)
- [UserService](https://yandex.ru/dev/api360/doc/ref/UserService.html)

## Install

```
go get github.com/nixys/nxs-go-ya360
```

## Example of usage

*You may find more examples in unit-tests in this repository*

**Get users from you organization in the Yandex 360:**

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	ya360 "github.com/nixys/nxs-go-ya360"
)

func main() {

	// Get variables from environment for connect to Yandex 360 server
	oAuth := os.Getenv("YA360_OAUTH")
	orgID, err := strconv.ParseInt(os.Getenv("YA360_ORG_ID"), 10, 64)
	if err != nil {
		fmt.Println("Init error: make sure environment variable `YA360_ORG_ID` correctly defined:", err)
		os.Exit(1)
	}
	if len(oAuth) == 0 {
		fmt.Println("Init error: make sure environment variable `YA360_OAUTH` correctly defined")
		os.Exit(1)
	}

	// Init Yandex 360 ctx
	y := ya360.Init(ya360.Settings{
		OAuth: oAuth,
		OrgID: orgID,
	})

	fmt.Println("Init: success")

	// Get users list
	u, err := y.UsersList(1, 1000)
	if err != nil {
		fmt.Println("Users list get error:", err)
		os.Exit(1)
	}

	fmt.Println("Users:")
	for _, e := range u.Users {
		fmt.Println("-", e.Nickname)
	}
}
```

Run:

```
YA360_OAUTH="YOUR_YA360_OAUTH" YA360_ORG_ID="YOUR_YA360_ORG_ID" go run main.go
```

## Feedback

For support and feedback please contact me:
- telegram: [@borisershov](https://t.me/borisershov)
- e-mail: b.ershov@nixys.ru

## License

nxs-go-ya360 is released under the [MIT License](LICENSE).
