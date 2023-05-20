package main

import (
	"fmt"
	"foo/foo"
)

func appName() string {
	const AppName string = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "bbb"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}


func main() {
	fmt.Println(foo.Max)
	// fmt.Println(foo.min)
	fmt.Println(foo.ReturnMin())
	fmt.Println(appName())
	// fmt.Println(AppName, Version)
	fmt.Println(Do("AAA"))
}
