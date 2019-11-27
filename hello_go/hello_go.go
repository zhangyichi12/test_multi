package main

import "fmt"
import "github.com/gofrs/uuid"

func main() {
	uuid, _ := uuid.NewV4()
	fmt.Println("hello world", uuid)
}
