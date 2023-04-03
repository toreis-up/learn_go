package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibN := 0
	fibM := 1
	return func() int {
		res := fibN
		fibN = fibM
		fibM = res + fibM
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
