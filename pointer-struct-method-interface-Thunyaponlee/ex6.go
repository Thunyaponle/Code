package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Breathe() {
	fmt.Printf("%s กำลังหายใจ\n", a.Name)
}

type Cat struct {
	Animal
	Breed string
}

func (c Cat) Meow() {
	fmt.Printf("%s กำลังร้อง: เมี้ยว! เมี้ยว!\n", c.Name)
}

/*func main() {
    myCat := Cat{Animal: Animal{Name: "ตุนตัง"}, Breed: "มันช์กิ้น"}
    myCat.Breathe() // "ตุนตัง กำลังหายใจ"
    myCat.Meow()    // "ตุนตัง กำลังร้อง: เมี้ยว! เมี้ยว!"
}*/
