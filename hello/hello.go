package hello

import "fmt"
import "github.com/gofrs/uuid"

func hello() {
	uuid, _ := uuid.NewV4()
	fmt.Println("hello", uuid)
}
