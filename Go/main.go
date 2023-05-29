package main

import (
	// "fmt"
    "os"
	"log"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")
	log.Println("message")	

	_, err := os.Open("afsdasdf")
	if err != nil {
		// log.Fatalln("Exit", err)
		logger.Fatalln("Exit", err)
	}

	// log.SetOutput(os.Stdout)

	// log.SetFlags(log.LstdFlags)
	// log.Println("A")
	// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	// log.Println("B")
	// log.SetFlags(log.Ltime | log.Lshortfile)
	// log.Println("C")
	// log.SetFlags(log.LstdFlags)
	// log.SetPrefix("[LOG]")
	// log.Println("E")
	
	// log.Print("log\n")
	// log.Println("Log2")
	// log.Printf("log%d\n", 3)

	// log.Fatal("log\n")
	// log.Fatalln("log2")
	// log.Fatalf("log%d\n", 3)

	// log.Panic("log\n")
	// log.Panicln("log2\n")
	// log.Panicf("log%d\n", 3)

	// f, err := os.Create("test.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(f)
	// log.Println("ファイルに書き込む")
}
