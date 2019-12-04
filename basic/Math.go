package basic

import (
	"fmt"
	"math"
	"math/rand"
)

func CustomRand() {
	fmt.Println("测试", rand.Intn(10))
}

func CustomSqrt() {
	fmt.Println("sqrt", math.Sqrt(7))
}

func CustomPi() {
	fmt.Println("pi:", math.Pi)
}
