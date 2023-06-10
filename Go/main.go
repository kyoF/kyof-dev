package main

import (
	"fmt"
    "regexp"
)

func main() {
    match, _ := regexp.MatchString("A", "ABC")
    fmt.Println(match)

    a, _ := regexp.Compile("A")
    match = a.MatchString("ABC")
    fmt.Println(match)

    b := regexp.MustCompile("A")
    match = b.MatchString("ABC")
    fmt.Println(match)

    regexp.MustCompile("\\d")
    regexp.MustCompile(`\d`)
}
