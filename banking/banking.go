package banking

// Account struct
type Account struct {
	owner   string
	Balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
