package hello

import "fmt"
import "github.com/gofrs/uuid"

func Hello() {
	uuid, _ := uuid.NewV4()
	fmt.Println("hello", uuid)
}

func Hello2() {
	uuid, _ := uuid.NewV4()
	fmt.Println("hello 2", uuid)
}
