package bank

type Account struct {
	Owner   string
	Balance float64
	Limit   float64 // Overdraft limit (default is 0 for regular accounts)
}

// NewAccount creates a new regular bank account
func NewAccount(owner string, initialBalance float64) *Account {
	return &Account{
		Owner:   owner,
		Balance: initialBalance,
		Limit:   0, // Default limit for regular accounts
	}
}

// NewSpecialAccount creates a new special bank account with a specified limit
func NewSpecialAccount(owner string, initialBalance, limit float64) *Account {
	return &Account{
		Owner:   owner,
		Balance: initialBalance,
		Limit:   limit,
	}
}

// Deposit adds money to the account
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// Withdraw removes money from the account, considering the overdraft limit
func (a *Account) Withdraw(amount float64) bool {
	if a.Balance-amount < a.Limit {
		return false // Withdrawal exceeds the limit
	}
	a.Balance -= amount
	return true
}
