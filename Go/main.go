package main

import (
	"fmt"
    // "strconv"
    // "os"
    // "time"
)

// func anything(a interface{}) {
//     // fmt.Println(a)
//     switch v:= a.(type) {
//     case string:
//         fmt.Println(v+"kyosuke")
//     case int:
//         fmt.Println(v+100000)
//     }
// }

// func TestDefer() {
//     defer fmt.Println("END")
//     fmt.Println("START")
// }

// func RunDefer() {
//     defer fmt.Println("a")
//     defer fmt.Println("b")
//     defer fmt.Println("c")
// }

// func sub() {
//     for {
//         fmt.Println("sub loop")
//         time.Sleep(100*time.Millisecond)
//     }
// }

func init() {
    fmt.Println("init")
}
func init() {
    fmt.Println("init 2")
}

func main() {
    fmt.Println("main")
    // go sub()
    // go sub()
    // for {
    //     fmt.Println("main loop")
    //     time.Sleep(200 * time.Millisecond)
    // }
    // defer func() {
    //     if x := recover(); x != nil {
    //         fmt.Println(x)
    //     }
    // }()
    // panic("runtime error")
    // fmt.Println("start")
    // TestDefer()
    // RunDefer()
    // defer func() {
    //     fmt.Println("1")
    //     fmt.Println("2")
    //     fmt.Println("3")
    //     fmt.Println("4")
    // }()

    // file, err := os.Create("test.txt")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // defer file.Close()
    // file.Write([]byte("Hello"))

    // Loop: for {
    //     for {
    //         for {
    //             fmt.Println("START")
    //             break Loop
    //         }
    //         fmt.Println("no return")
    //     }
    //     fmt.Println("no return no")
    // }
    // fmt.Println("END")

    // CLoop: for i:=0; i<3; i++ {
    //     for j:=0; j<3; j++ {
    //         if j > 1 {
    //             continue CLoop
    //         }
    //         fmt.Println(i, j)
    //     }
    //     fmt.Println("No, no")
    // }

    // anything("aaa")
    // anything(1)
    
    // var x interface{} = 3
    // i, isInt := x.(int)
    // fmt.Println(i + 2, isInt)

    // f, isFloat64 := x.(float64)
    // fmt.Println(f, isFloat64)

    // if x == nil {
    //     fmt.Println("none")
    // } else if i, isInt := x.(int); isInt {
    //     fmt.Println(i, "x is Int")
    // } else if s, isString := x.(string); isString {
    //     fmt.Println(s, isString)
    // } else {
    //     fmt.Println("i don't know")
    // }

    // switch v := x.(type) {
    // case int:
    //     fmt.Println("int", v)
    // case string:
    //     fmt.Println("string", v)
    // default:
    //     fmt.Println("i dont't know", v)
    // }

    // n := 5
    // switch n {
    // case 1, 2:
    //     fmt.Println("1 or 2")
    // case 3, 4:
    //     fmt.Println("3 or 4")
    // default:
    //     fmt.Println("I don't know")
    // }

    // switch n:=3; n {
    // case 1, 2:
    //     fmt.Println("A")
    // case 4:
    //     fmt.Println("B")
    // default:
    //     fmt.Println("C")
    // }

    // n := 3
    // switch {
    // case n >= 0 && n <= 1:
    //     fmt.Println("0<=n<=1")
    // case n>1 && n<=2:
    //     fmt.Println("1<n<=2")
    // default:
    //     fmt.Print("2<n")
    // } 

    // i := 0
    // for {
    //     i++
    //     if i == 3 {
    //         break
    //     }
    //     fmt.Println("loop")
    // }

    // point := 0
    // for point < 10 {
    //     fmt.Println(point)
    //     point++
    // }

    // for i:=0; i<10; i++ {
    //     if i == 3 {
    //         continue
    //     }
    //     if i == 6 {
    //         break
    //     }
    //     fmt.Println(i)
    // }

    // arr := [3]int{1, 2, 3}
    // for i:=0; i<len(arr); i++ {
    //     fmt.Println(arr[i])
    // }
    // for idx, val := range arr {
    //     fmt.Println(idx, val)
    // }

    // sl := []string{"python", "go", "javascript"}
    // for i, v := range sl {
    //     fmt.Println(i, v)
    // }

    // m := map[string]int {"apple":100, "banana":200}
    // for i, v := range m {
    //     fmt.Println(i, v)
    // }

    // var s string = "A"

    // i, err := strconv.Atoi(s)
    // if err != nil {
    //     fmt.Println(err)
    //     fmt.Println("error")
    // } else {
    //     fmt.Printf("i = %T\n", i)
    // }

    // a := 1
    // if a == 2 {
    //     fmt.Println("two")
    // } else if a == 1 {
    //     fmt.Println("one")
    // } else {
    //     fmt.Println("I don't know")
    // }

    // // x := 1
    // if x := 0; true {
    //     fmt.Println(x)
    // }
    // fmt.Println(x)
}	
