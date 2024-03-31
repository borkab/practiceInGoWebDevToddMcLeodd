package main

import "fmt"

func main() {
	m := map[string]int{"Karajan": 23, "henrietta": 4, "rabbat": 37}

	fmt.Println("our map example: ", m)

	for k := range m {
		fmt.Println("this is the key: ", k)

	}

	for k, v := range m {
		fmt.Println("key: ", k, "\n", "value: ", v)
	}

}
