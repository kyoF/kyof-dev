package main

import (
	"fmt"
	"strconv"
)

func main() {
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", i, i)
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)

	i10, _ := strconv.Atoi("123")
	fmt.Println(i10)
}
