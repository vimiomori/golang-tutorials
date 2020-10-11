package main

import "fmt"

func main() {
	is := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range is {
		if n % 2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}
}