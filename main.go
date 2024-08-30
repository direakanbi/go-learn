package main

import "fmt"

func main() {
	user := make(map[string]string)
	user["name"] = "John"
	user["age"] = "30"
}

fmt.Println(user)