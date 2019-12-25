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
	var mapInt = []int{1, 2, 4, 5, 6, 7, 7, 8}
	basic.CustomRange(mapInt)
	fmt.Println(vertx.Abs())
	basic.CustomRead()
	go basic.Say("aaa")
	basic.Say("zzz")
	//web.WebStore()
	//web.DymicStaticStore()
	//web.MuxWeb()
	//web.CustomJson()
	basic.GetLocalIp()
}
