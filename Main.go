package main

import (
	"fmt"
	"golearning/basic"
)

func main() {
	basic.HelloWorld()
	basic.CustomRand()
	basic.CustomSqrt()
	basic.CustomPi()
	fmt.Println(basic.Add(1, 2))
	a, b := basic.Swap("222", "123123123")
	fmt.Println(a, b)
}
