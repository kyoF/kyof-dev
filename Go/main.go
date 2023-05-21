package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	// f, err := os.Open("foo.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// bs := make([]byte, 128)
	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// bs2 := make([]byte, 128)
	// nn, err := f.ReadAt(bs2, 10)
	// fmt.Println(nn)
	// fmt.Println(string(bs2))

	// f, _ := os.Create("foo.txt")
	// f.Write([]byte("Hello\n"))
	// f.WriteAt([]byte("Golang"), 6)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("YaaH")
	// f, err := os.Open("test1.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Println("length=%d\n", len(os.Args))
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }
	// os.Exit(1)
	// fmt.Println("Start")

	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
