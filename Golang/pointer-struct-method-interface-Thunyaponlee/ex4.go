package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct{ Radius float64 }
type Square struct{ Side float64 }

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
func (s Square) Area() float64 {
	return s.Side * s.Side
}
func printArea(s Shape) {
	fmt.Println(s.Area())
}

/*func main() {
    c := Circle{Radius: 5}
    sq := Square{Side: 4}
    printArea(c)  // พิมพ์ค่าพื้นที่วงกลมรัศมี 5
    printArea(sq) // พิมพ์ค่าพื้นที่สี่เหลี่ยมจัตุรัสด้าน 4
}*/
