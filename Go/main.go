package main

import (
	"fmt"
    "os"
)

func main() {
	fmt.Println("show")
	fmt.Print("show")
	fmt.Println("Hello!")
    
    fmt.Println("Hello", "World!")

    fmt.Printf("%s\n", "Hello")
    fmt.Printf("%s#v\n", "Hello")

    s := fmt.Sprintf("Hello")
    s1 := fmt.Sprintf("%v\n", "Hello")
    s2 := fmt.Sprintln("Hello")
    fmt.Println(s)
    fmt.Println(s1)
    fmt.Println(s2)

    fmt.Fprint(os.Stdout, "Hello")
    fmt.Fprintf(os.Stdout, "Hello")
    fmt.Fprintln(os.Stdout, "Hello")

    f, _ := os.Create("test1.txt")
    defer f.Close()
    fmt.Fprintln(f, "Fprint")
}
