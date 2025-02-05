package main

type Person struct {
	Name string
	Age  int
}

func incrementAge(p *Person) {
	p.Age++
}

/*func main() {
    p := Person{Name: "Alice", Age: 30}
    incrementAge(&p)
    fmt.Println(p.Age) // 31
}*/
