package main

import (
	"fmt"
)

func Plus(x, y int) int {
	return x + y
}

// func Div(x, y int) (int, int) {
// 	q := x / y
// 	r := x % y
// 	return q, r
// }

// func Double(price int) (result int) {
// 	result = price * 2
// 	return
// }

// func Noreturn() {
// 	fmt.Println("No Return")
// 	return
// }

// func ReturnFunc() func() {
// 	return func() {
// 		fmt.Println("I'm a function")
// 	}
// }

// func CallFunc(f func()) {
// 	f()
// }

// func Later() func(string) string {
// 	var store string
// 	return func(next string) string {
// 		s := store
// 		store = next
// 		return s
// 	}
// }

func CountUp() func() int {
	i := 0
	return func() (int) {
		i++
		return i
	}
}

func main() {
	a := CountUp()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	
	b := CountUp()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	// f := Later()
	// fmt.Println(f("Hello"))
	// fmt.Println(f("My"))
	// fmt.Println(f("name"))
	// fmt.Println(f("is"))
	// fmt.Println(f("GoLang"))
	// i := Plus(1, 2)
	// fmt.Println(i)
	// i2, i3 := Div(9, 4)
	// fmt.Println(i2, i3)
	// i4, _ := Div(9, 4)
	// fmt.Println(i4)
	// i5 := Double(2)
	// fmt.Println(i5)
	// Noreturn()
	// f := func(x, y int) int {
	// 	return x*y
	// }
	// i := f(1, 3)
	// fmt.Println(i)
	
	// i2 := func(x, y int) int {
	// 	return x*y
	// }(2, 3)
	// fmt.Println(i2)
	// f := ReturnFunc()
	// f()
	// CallFunc(func() {
	// 	fmt.Println("Call Function!")
	// })
}	
