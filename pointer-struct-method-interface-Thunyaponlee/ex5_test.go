package main

import (
	"testing"
)

func TestBankAccount(t *testing.T) {
	acc := BankAccount{owner: "Bob", balance: 1000.0}
	acc.Deposit(500)
	if acc.Balance() != 1500.0 {
		t.Errorf("Expected balance to be 1500.0, got %f", acc.Balance())
	}

	err := acc.Withdraw(200)
	if err != nil {
		t.Errorf("Unexpected error when withdrawing 200: %v", err)
	}
	if acc.Balance() != 1300.0 {
		t.Errorf("Expected balance to be 1300.0 after withdrawal, got %f", acc.Balance())
	}

	err = acc.Withdraw(2000)
	if err == nil {
		t.Errorf("Expected an error when withdrawing more than balance")
	}
	if acc.Balance() != 1300.0 {
		t.Errorf("Balance should remain unchanged at 1300.0 after failed withdrawal, got %f", acc.Balance())
	}
}
