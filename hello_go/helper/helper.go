package helper

import "fmt"
import "github.com/gofrs/uuid"

func Helper() {
	uuid, _ := uuid.NewV4()
	fmt.Println("helper", uuid)
}

func Helper2() {
	uuid, _ := uuid.NewV4()
	fmt.Println("helper 2", uuid)
}
