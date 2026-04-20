package main

import (
	"fmt"
)

func main() {
	A := "test"
	salt := "salt"
	var B int
	for i := 0; i < len(A); i++ {
		B += int(A[i])
	}
	
	for i := 0; i < len(salt); i++ {
		B += int(salt[i])
	}
	fmt.Println(B)
}
