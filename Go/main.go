package main

import (
	"fmt"
    "io"
    "crypto/md5"
)

func main() {
    h := md5.New()
    io.WriteString(h, "ABVDE")
    fmt.Println(h.Sum(nil))
    s := fmt.Sprintf("%x", h.Sum(nil))
    fmt.Println(s)
}
