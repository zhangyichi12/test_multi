package hello_go

import "fmt"
import "github.com/gofrs/uuid"

func Hello_go() {
	uuid, _ := uuid.NewV4()
	fmt.Println("hello go", uuid)
}
