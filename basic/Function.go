package basic

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func CustomFor(end int) {
	for i := 0; i < end; i++ {
		fmt.Print(i)
	}
}
