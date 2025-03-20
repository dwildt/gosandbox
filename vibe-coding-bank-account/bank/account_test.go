package bank

import "testing"

func TestAccount(t *testing.T) {
	account := NewAccount("John Doe", 1000.0)

	// Test Deposit
	account.Deposit(500.0)
	if account.Balance != 1500.0 {
		t.Errorf("Expected balance 1500.0, got %f", account.Balance)
	}

	// Test Withdraw
	success := account.Withdraw(200.0)
	if !success || account.Balance != 1300.0 {
		t.Errorf("Expected balance 1300.0, got %f", account.Balance)
	}

	// Test Withdraw with insufficient funds
	success = account.Withdraw(2000.0)
	if success {
		t.Errorf("Withdrawal should have failed due to insufficient funds")
	}
}

func TestSpecialAccount(t *testing.T) {
	// Create a special account with a negative limit of 500
	account := NewSpecialAccount("Jane Doe", 1000.0, -500.0)

	// Test Withdraw within the limit
	success := account.Withdraw(1400.0)
	if !success || account.Balance != -400.0 {
		t.Errorf("Expected balance -400.0, got %f", account.Balance)
	}

	// Test Withdraw exceeding the limit
	success = account.Withdraw(200.0)
	if success {
		t.Errorf("Withdrawal should have failed due to exceeding the limit")
	}
}

func TestDefaultAccount(t *testing.T) {
	// Create a default account
	account := NewAccount("John Doe", 1000.0)

	// Test Withdraw exceeding the default limit (0)
	success := account.Withdraw(1100.0)
	if success {
		t.Errorf("Withdrawal should have failed due to insufficient funds")
	}
}
