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
	basic.CustomAbs()
	fmt.Println(basic.Add(1, 2))
	a, b := basic.Swap("222", "123123123")
	fmt.Println(a, b)
	basic.CustomFor(100)
	var vertx = basic.Vertx{"java"}
	fmt.Println("\r\n", vertx.Source)
}
