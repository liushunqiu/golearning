package basic

import (
	"fmt"
	"io"
	"strings"
)

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

func CustomRange(mapInt []int) {
	for k, v := range mapInt {
		fmt.Println(k, v)
	}
}

func CustomRead() {
	read := strings.NewReader("ccccccccc")
	b := make([]byte, 8)
	for {
		n, err := read.Read(b)
		fmt.Println(n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
