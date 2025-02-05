package interfaces

type LibraryItem interface {
	Info() string
}

func PrintItemInfo(item LibraryItem) string {
	return item.Info()
}

type Borrowable interface {
	Borrow() error
	Return() error
}
