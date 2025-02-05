package models

import (
	"errors"
	"fmt"
)

func increaseStock(stockPtr *int) {
	*stockPtr++
}

// Book แทนข้อมูลหนังสือพื้นฐาน
type Book struct {
	title  string
	author string
	stock  int
}

// NewBook เป็นเหมือน Constructor Function
func NewBook(title, author string, stock int) *Book {
	return &Book{title: title, author: author, stock: stock}
}

// Getter Methods
func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetStock() int {
	return b.stock
}

func (b *Book) validateStock(amount int) error {

	if b.stock+amount < 0 {
		return errors.New("stock cannot be negative")
	}
	return nil
}

func (b Book) Details() string {
	return fmt.Sprintf("%s by %s (Stock: %d)", b.title, b.author, b.stock)
}

func (b *Book) AddStock(amount int) error {
	if b.stock+amount < 0 {
		return errors.New("stock cannot be negative")
	}
	b.stock += amount

	return nil
}

func (b Book) Info() string {
	return b.title + " by " + b.author
}

func (b *Book) Borrow() error {
	if b.stock > 0 {
		return b.AddStock(-1)
	}
	return errors.New("out of stock")
}

func (b *Book) Return() error {
	b.AddStock(1)
	return nil
}
