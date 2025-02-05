package models

import (
	"errors"
	"fmt"
)

type DVD struct {
	Title    string
	Duration int
	Stock    int
}

func (d DVD) Info() string {
	return fmt.Sprintf("DVD: %s (%d mins)", d.Title, d.Duration)
}

func NewDVD(Title string, Duration int, Stock int) *DVD {
	return &DVD{Title: Title, Duration: Duration, Stock: Stock}
}

func (d *DVD) AddStock(amount int) error {
	if d.Stock+amount < 0 {
		return errors.New("stock cannot be negative")
	}
	d.Stock += amount
	return nil
}

func (d *DVD) Borrow() error {
	if d.Stock > 0 {
		return d.AddStock(-1)
	}
	return errors.New("out of stock")
}
func (d *DVD) Return() error {
	d.AddStock(1)
	return nil
}
