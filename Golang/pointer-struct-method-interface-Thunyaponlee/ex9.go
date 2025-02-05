package main

type Counter struct{ Count int }

func (c *Counter) Increment() {
	c.Count++
}
func (c Counter) GetCount() int {
	return c.Count
}

/*func main() {
    c := Counter{Count: 10}
    c.Increment()
    c.Increment()
    fmt.Println(c.GetCount()) // 12
}*/
