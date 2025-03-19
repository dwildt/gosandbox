package bank

type Account struct {
	Owner   string
	Balance float64
}

// NewAccount creates a new bank account
func NewAccount(owner string, initialBalance float64) *Account {
	return &Account{
		Owner:   owner,
		Balance: initialBalance,
	}
}

// Deposit adds money to the account
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// Withdraw removes money from the account
func (a *Account) Withdraw(amount float64) bool {
	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}
