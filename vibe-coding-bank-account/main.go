package main

import (
	"fmt"
	"vibe-coding-bank-account/bank"
)

func main() {
	account := bank.NewAccount("John Doe", 1000.0)
	fmt.Println("Account created:", account)
}
