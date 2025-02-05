package main

type Soundmaker interface {
	MakeSound() string
}
type Dog struct{}
type Cow struct{}

func (d Dog) MakeSound() string {
	return "Woof!"
}
func (c Cow) MakeSound() string {
	return "Moo!"
}
func identifyAnimal(s Soundmaker) string {
	switch s.(type) {
	case Dog:
		return "This is a Dog!"
	case Cow:
		return "This is a Cow!"
	default:
		return "Unknown Animal!"
	}

}

/*func main() {
    fmt.Println(identifyAnimal(Dog{})) // "This is a Dog!"
    fmt.Println(identifyAnimal(Cow{})) // "This is a Cow!"
}*/
