package account

type Account struct {
	Name    string
	Balance int
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}
