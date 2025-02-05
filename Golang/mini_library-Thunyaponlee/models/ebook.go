package models

import "fmt"

type EBook struct {
	Book
	fileSizeMB int
}

func NewEBook(title, author string, fileSizeMB int) *EBook {
	return &EBook{Book: Book{title: title, author: author}, fileSizeMB: fileSizeMB}
}

func (eb *EBook) GetFileSize() int {
	return eb.fileSizeMB
}

func (eb *EBook) Info() string {
	return fmt.Sprintf("%s by %s (E-Book, %dMB)", eb.title, eb.author, eb.fileSizeMB)
}
