package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	value, _ := bcrypt.GenerateFromPassword([]byte("Engagement"), bcrypt.DefaultCost)
	fmt.Println(string(value))
}
