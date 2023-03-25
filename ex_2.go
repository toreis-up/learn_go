package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	answer := make([][]uint8, dy)
	for y := range answer {
		answer[y] = make([]uint8, dx)
		for x := range answer[y] {
			answer[y][x] = uint8((x + y) / 2)
		}
	}

	return answer
}

func main() {
	pic.Show(Pic)
}
