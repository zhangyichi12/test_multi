package hello_go

import "fmt"
import "github.com/gofrs/uuid"
// import "./helper/helper.go"
import "github.com/zhangyichi12/test_multi/hello_go/helper"
import "github.com/zhangyichi12/test_multi/hello"
// import "../hello"

func Hello_go() {
	uuid, _ := uuid.NewV4()
	fmt.Println("hello go", uuid)
	hello.Hello()
	hello.Hello2()
	helper.Helper()
	helper.Helper2()
}
