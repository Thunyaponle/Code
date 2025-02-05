package main

import (
	"errors"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}
type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b BankAccount) Balance() float64 {
	return b.balance
}
func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.balance {
		return errors.New("ยอดเงินไม่พอสำหรับการถอน")
	} else {
		b.balance -= amount
		return nil
	}
}

/*func main() {
    acc := BankAccount{owner: "Bob", balance: 1000.0}
    acc.Deposit(500)
    fmt.Println(acc.Balance()) // 1500.0
    acc.Withdraw(1200)
    fmt.Println(acc.Balance()) // 300.0
}*/
