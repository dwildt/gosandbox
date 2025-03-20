package main

import (
	"fmt"
	"vibe-coding-bank-account/bank"
)

func main() {
	// Test Default Account
	defaultAccount := bank.NewAccount("John Doe", 1000.0)
	fmt.Println("Default Account created:", defaultAccount)

	// Perform operations on the default account
	defaultAccount.Deposit(500.0)
	fmt.Println("After deposit of 500:", defaultAccount)
	withdrawSuccess := defaultAccount.Withdraw(1200.0)
	fmt.Printf("Withdraw 1200 success: %v, Balance: %f\n", withdrawSuccess, defaultAccount.Balance)
	withdrawSuccess = defaultAccount.Withdraw(500.0)
	fmt.Printf("Withdraw 500 success: %v, Balance: %f\n", withdrawSuccess, defaultAccount.Balance)

	// Test Special Account
	specialAccount := bank.NewSpecialAccount("Jane Doe", 1000.0, -500.0)
	fmt.Println("\nSpecial Account created:", specialAccount)

	// Perform operations on the special account
	specialAccount.Deposit(300.0)
	fmt.Println("After deposit of 300:", specialAccount)
	withdrawSuccess = specialAccount.Withdraw(1600.0)
	fmt.Printf("Withdraw 1600 success: %v, Balance: %f\n", withdrawSuccess, specialAccount.Balance)
	withdrawSuccess = specialAccount.Withdraw(300.0)
	fmt.Printf("Withdraw 300 success: %v, Balance: %f\n", withdrawSuccess, specialAccount.Balance)
}
