package main

import "fmt"

// type person struct {
// 	fName   string
// 	lName   string
// 	favFood []string
// }

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

type gator int

var g1 gator

func main() {

	t := truck{vehicle: vehicle{doors: 2, color: "blue"}, fourWheel: true}
	s := sedan{vehicle: vehicle{doors: 5, color: "gold"}, luxury: true}

	fmt.Println(t, "\n", s)
	fmt.Println(t.fourWheel)
	fmt.Println(s.doors)

	tt := t.transportationDevice()
	ss := s.transportationDevice()

	fmt.Println(tt, "\n", ss)

	report(t)
	report(s)

	g1 = 42

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	// p1 := person{fName: "Lord", lName: "Voldemort", favFood: []string{"kiwi", "snakes", "elixirs"}}
	// fmt.Println(p1.fName)

	// fmt.Println(p1.favFood)
	// p1.favFood = append(p1.favFood, "souls")

	// for i, v := range p1.favFood {
	// 	fmt.Println(i, v)
	// }

	// fmt.Println(p1.walk())

}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I am a truck with", t.doors, "doors and I am rollin")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("I am a sedan in", s.color, "color and I am luxury shine")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

// func (p person) walk() string {
// 	return p.fName + " " + p.lName + " is walking"
// }
