package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1           // ch = [1]
	ch <- 2           // ch = [1, 2]
	fmt.Println(<-ch) // 1
	// ch = [2]
	ch <- 3           // ch = [2, 3]
	fmt.Println(<-ch) // 2
	// ch = [3]
	fmt.Println(<-ch) // 3
	// ch = []
	ch <- 4           // ch = [4]
	fmt.Println(<-ch) // 4
	// ch = []
}
