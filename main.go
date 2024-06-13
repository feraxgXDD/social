package main

import (
	"fmt"

	"just/structsSTR"
)

var userMap = make(map[string]*structsSTR.User)

func main() {
	structsSTR.NewUser(userMap, "John", "123456")

	fmt.Println(userMap["John"])
}
