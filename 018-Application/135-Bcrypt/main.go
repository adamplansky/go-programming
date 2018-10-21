package main

import (
	"fmt"

	"github.com/btcsuite/golangcrypto/bcrypt"
)

func main() {
	p := `password`
	fmt.Println(p)
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))

	err = bcrypt.CompareHashAndPassword(b, []byte(p))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("password match!")
	}

}
