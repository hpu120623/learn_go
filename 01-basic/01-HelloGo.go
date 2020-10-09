package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3 // 不是全局变量，包内部的
	bb = 4
)

//cc:= 5  // 错误不能这样搞

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)

}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)

}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	var s = "abc"
	fmt.Println(a, b, c, d, s)

}

func variableShorter() {
	a, b, c, d := 1, 2, 3, "hello"
	b = 6
	fmt.Println(a, b, c, d)

}

func euler() {
	a := 3 + 4i
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Abs(a))
}

func triangle() {
	var a, b = 1, 2
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}

// 常量
func consts()  {
	const (
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

// 普通值枚举
func enums()  {
	const (
		python = iota
		_
		java
		scala
		golang
	)
	fmt.Println(python, java, golang, scala)
}

// 自增值枚举
func enums2()  {
	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb)
	euler()
	triangle()
	consts()
	enums()
	enums2()
}
