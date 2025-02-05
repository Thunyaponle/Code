package bank

import (
	"errors"
	"fmt"
)

// Custom error for insufficient funds
type InsufficientFundsError struct {
	message string
}

func (e *InsufficientFundsError) Error() string {
	return e.message
}

// SavingsAccount type
type SavingsAccount struct {
	balance float64
}

type interestAccount struct {
	*SavingsAccount
	interestRate float64
	//AddInterest  float64
}

func NewInterestAccount(initialBalance, rate float64) *interestAccount {
	if rate < 0 {
		rate = 0
	}
	return &interestAccount{
		SavingsAccount: NewSavingsAccount(initialBalance),
		interestRate:   rate,
	}
}

func NewSavingsAccount(initialBalance float64) *SavingsAccount {
	return &SavingsAccount{balance: initialBalance}
}

func (acc *SavingsAccount) GetBalance() float64 {
	return acc.balance
}

func (acc *SavingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc.balance += amount
	return nil
}

func (acc *SavingsAccount) Withdraw(amount float64) error {
	if amount > acc.balance {
		return &InsufficientFundsError{"insufficient funds"}
	}
	acc.balance -= amount
	return nil
}

// InterestAccount type (inherits SavingsAccount functionality)
/*type InterestAccount struct {
	SavingsAccount
	interestRate float64
	addInterest float64
}*/

/*func NewInterestAccount(initialBalance int, interestRate float64) *InterestAccount {
	return &InterestAccount{
		SavingsAccount: *NewSavingsAccount(initialBalance),
		interestRate:   interestRate,
	}
}*/

func (intAcc *interestAccount) AddInterest() {
	interest := float64(float64(intAcc.balance) * intAcc.interestRate)
	intAcc.balance += interest
}

func (intAcc *interestAccount) Withdraw(amount float64) error {
	fmt.Println("Withdrawing from interest account...")
	return intAcc.SavingsAccount.Withdraw(amount)
}

// Transfer function
func Transfer(from *SavingsAccount, to *SavingsAccount, amount float64) error {
	err := from.Withdraw(amount)
	if err != nil {
		return err
	}
	err = to.Deposit(amount)
	return err
}

// RiskyBankOperation with panic
func RiskyBankOperation(amount float64) {
	defer func() {
		if r := recover(); r != nil {
			// เมื่อเกิด Panic จะพิมพ์ข้อความนี้ออกมา
			fmt.Printf("เกิด Panic แล้วทำการ Recover สำเร็จ: %v\n", r)
		}
	}()

	if amount < 0 {
		// เกิด Panic เมื่อ amount เป็นลบ
		panic(fmt.Sprintf("negative amount: %f", amount))
	}
	fmt.Println("Operation successful!")
}
