package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
    fmt.Println(s1, s2)

    i1 := strings.Index("ABCDE", "A")
    i2 := strings.Index("ABCDE", "D")
    fmt.Println(i1, i2)
    i3 := strings.LastIndex("ABCDABCD", "BC")
    i4 := strings.IndexAny("ABC", "ABC")
    i5 := strings.LastIndexAny("ABC", "ABC")
    fmt.Println(i3, i4, i5)
    b1 := strings.HasPrefix("ABC", "A")
    b2 := strings.HasSuffix("ABC", "C")
    b3 := strings.HasSuffix("ABC", "A")
    fmt.Println(b1, b2, b3)
    b4 := strings.Contains("ABC", "B")
    b5 := strings.ContainsAny("ABCDE", "BD")
    fmt.Println(b4, b5)
    i6 := strings.Count("ABCABC", "B")
    i7 := strings.Count("ABCABC", "")
    fmt.Println(i6, i7)

    s3 := strings.Repeat("ABC", 4)
    s0 := strings.Repeat("ABC", 0)
    fmt.Println(s3, s0)

    s5 := strings.Replace("AAAAAAA", "A", "B", 1)
    s7 := strings.Replace("AAAAAAA", "A", "B", -1)
    fmt.Println(s5, s7)

    s8 := strings.Split("A,B,C,D,E", ",")
    fmt.Println(s8)
    s9 := strings.SplitAfter("A,B,C,D,E", ",")
    fmt.Println(s9)
    s10 := strings.SplitN("A,B,C,D,E", ",", 2)
    fmt.Println(s10)
    s11 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
    fmt.Println(s11)

    a := strings.ToLower("ABC")
    b := strings.ToUpper("abc")
    fmt.Println(a, b)

    c := strings.TrimSpace(" .  = ABV     ")
    d := strings.TrimSpace("　.　　= ABV     ")
    fmt.Println(c, d)

    e := strings.Fields("a b c")
    fmt.Println(e)
    fmt.Println(e[1])
}
