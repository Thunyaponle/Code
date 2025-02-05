package main

import (
	"errors"
)

func safeDereference(ptr *int) (int, error) {
	if ptr == nil {
		return 0, errors.New("err != nil")
	} else {
		return *ptr, nil
	}
}

/*func main() {
	var p *int
	val, err := safeDereference(p)
	fmt.Println("val =", val, ", err =", err)
	// val = 0, err != nil

	x := 42
	val, err = safeDereference(&x)
	fmt.Println("val =", val, ", err =", err)
	// val = 42, err = nil
}*/
