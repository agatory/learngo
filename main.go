package main

import (
	"fmt"

	"github.com/agatory/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
