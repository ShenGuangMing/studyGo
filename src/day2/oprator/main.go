package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func arc() {
	var a float32 = 8
	var b float32 = 3
	var c int = 10
	var d int = 5
	fmt.Printf("a+b=%f\n", a+b)
	fmt.Printf("c-d=%d\n", c-d)
	fmt.Printf("a*b=%f\n", a*b)
	fmt.Printf("c/d=%d\n", c/d)
	fmt.Printf("d mod c=%d\n", d%c)
	fmt.Printf("5=%b\n", 5)
	fmt.Printf("~5=%b\n", ^5)
	fmt.Println("a > b? ", a > b)
	fmt.Println("c > d? ", d > c)
	fmt.Println("!(c > d)? ", !(d > c))
}

func bitGo() {
	fmt.Printf("os arch %s\nint size %d\n", runtime.GOARCH, strconv.IntSize)
	fmt.Printf("65 binery %b\n", 65)
}

func main() {
	arc()
	bitGo()
	//println(util.BinaryFormat(20))
}
