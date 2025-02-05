package main

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r *Rectangle) Resize(w, h float64) {
	r.Width = w
	r.Height = h
}

/*func main() {
    r := Rectangle{Width: 10, Height: 5}
    fmt.Println(r.Area()) // 50
    r.Resize(2, 2)
    fmt.Println(r.Area()) // 4
}*/
