package main

import (
	"fmt"
)

type ret func(x int)

func makeFunction(name ret) {
	fmt.Println("00000")
	name(1)
}

func makeFunction2(x int) {
	fmt.Println("from makeFn2 got: ", x)
}
