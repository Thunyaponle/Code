package main

type Mover interface {
	Move() string
}
type Sounder interface {
	MakeSound() string
}
type Creature interface {
	Mover
	Sounder
}
type Horse struct {
}

func (h Horse) Move() string {
	return "Horse is galloping"
}
func (h Horse) MakeSound() string {
	return "Neigh!"
}
func describeCreature(c Creature) string {
	return c.Move() + " and " + c.MakeSound()
}

/*func main() {
	h := Horse{}
	fmt.Println(describeCreature(h)) // "Horse is galloping and Neigh!"
}*/
