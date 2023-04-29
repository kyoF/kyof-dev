package main

import (
	"fmt"
)

func main() {
	fmt.Println(1+3)
	fmt.Println(1-3)
	fmt.Println(1*3)
	fmt.Println(24/2)
	fmt.Println(4%2)
	fmt.Println("abc"+"def")

	n:=0
	n+=4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s:="abc"
	s+="def"
	fmt.Println(s)

	fmt.Println(1 == 1)
	fmt.Println(2 == 1)

	fmt.Println(4 <= 8)
	fmt.Println(4 >= 8)
	fmt.Println(4 > 8)

	fmt.Println(true == false)
	fmt.Println(true != false)

	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)
	fmt.Println(!false)
}	
