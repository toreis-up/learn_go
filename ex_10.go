package main

import "golang.org/x/tour/tree"
import "fmt"

// https://go.dev/play/p/kN20VmEg9bb
func walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	Walk(t, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go walk(t1, ch1)
	go walk(t2, ch2)

	for { // 無限ループ
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 && !ok2 {
			break // どっちも読み取れなくなった
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go walk(tree.New(2), ch)
	for ch := range ch {
		fmt.Println(ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
