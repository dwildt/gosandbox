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
