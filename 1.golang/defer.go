package main

import "fmt"

func test01() {
	fmt.Println("我是测试01")
	a := 10
	defer fmt.Println("a=", a)
	a++
}

//未引用参数 跟测试一一样
func test02() {
	fmt.Println("我是测试02")
	a := 10
	defer func() {
		fmt.Println("a=", a)
	}()
	a++
}

//此处引用了参数
func test03() {
	fmt.Println("我是测试03")
	a := 10
	defer func(a int) {
		fmt.Println("a=", a)
	}(a)
	a++
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//
//	//fmt.Println(f1())
//	//fmt.Println(f2())
//	//fmt.Println(f3())
//	//fmt.Println(f4())
//
//	//test01()
//	//test02()
//	//test03()
//}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
