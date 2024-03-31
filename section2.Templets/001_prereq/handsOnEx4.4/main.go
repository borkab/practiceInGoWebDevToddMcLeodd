package main

import "fmt"

type gator int
type flamingo bool

type swampCreatures interface {
	greeting()
}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x1 int
	fmt.Println(x1)
	fmt.Printf("%T\n", x1)

	x1 = int(g1)

	var g2 gator

	g2.greeting()

	var f1 flamingo
	f1.greeting()

	fmt.Println("-----------------")
	bayou(g1)
	bayou(f1)

}
func (g gator) greeting() {
	fmt.Println("Hello I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

func bayou(s swampCreatures) {
	s.greeting()
}
