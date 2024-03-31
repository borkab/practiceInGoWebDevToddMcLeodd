package main

import "fmt"

func main() {

	fmt.Println("our slice example: ")

	s := []int{2, 4, 6, 8, 0, 2, 4, 7}

	fmt.Println(s)

	for i := range s {
		fmt.Println("just the indexes ", i)
	}

	for _, ss := range s {
		fmt.Println("just the elements ", ss)
	}

	for i, ss := range s {
		fmt.Println("indexes and elements: ", i, ss)
	}
}
