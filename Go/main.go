package main

import (
	"fmt"
    // "time"
)

// func Sum(s ...int) int {
//     n := 0
//     for _, v := range s {
//         n += v
//     }
//     return n
// }

// func reciever(c chan int) {
//     for {
//         i := <-c
//         fmt.Println(i)
//     }
// }

// func reciever(name string, c chan int) {
//     for {
//         i, ok := <-c
//         if !ok {
//             break
//         }
//         fmt.Println(name, i)
//     }
//     fmt.Println(name + "END")
// }

func main() {
    ch1 := make(chan int, 2)
    ch2 := make(chan string, 2)

    // ch2 <- "A"
    // ch1 <- 11
    select {
    case v1 := <-ch1:
        fmt.Println(v1 + 1000)
    case v2 := <-ch2:
        fmt.Println(v2 + "!?")
    default:
        fmt.Println("どちらでもない")
    }

    ch3 := make(chan int)
    ch4 := make(chan int)
    ch5 := make(chan int)

    go func() {
        for {
            i := <-ch3
            ch4 <- i*2
        }
    }()

    go func() {
        for {
            i2 := <-ch4
            ch5 <- i2 - 1
        }
    }()

    n := 0
    L: for {
        select {
        case ch3 <- n:
            n++
        case i3 := <-ch5:
            fmt.Println("recieved", i3)
        default:
            if n > 100 {
                break L
            }
        }
    }
    // ch2 <- "A"

    // v := <-ch1
    // v2 := <-ch2
    // fmt.Println(v)
    // fmt.Println(v2)
    // ch1 := make(chan int, 3)
    // ch1 <- 1
    // ch1 <- 2
    // ch1 <- 3
    // close(ch1)
    // for i := range ch1 {
    //     fmt.Println(i)
    // }
    // ch1 := make(chan int, 2)

    // go reciever("1.goroutin", ch1)
    // go reciever("2.goroutin", ch1)
    // go reciever("3.goroutin", ch1)

    // i := 0
    // for i<100 {
    //     ch1 <- i
    //     i++
    // }
    // close(ch1)
    // time.Sleep(3 * time.Second)
    // ch1 <- 1
    // close(ch1)
    // // ch1 <- 1
    // i, ok := <-ch1
    // fmt.Println(i, ok)
    // i2, ok := <-ch1
    // fmt.Println(i2, ok)
    // i3, ok := <-ch1
    // fmt.Println(i3, ok)

    // ch1 := make(chan int)
    // ch2 := make(chan int)

    // go reciever(ch1)
    // go reciever(ch2)

    // i := 0
    // for i<100 {
    //     ch1 <- i
    //     ch2 <- i
    //     time.Sleep(50 * time.Millisecond)
    //     i++
    // }
    // fmt.Println(<-ch)

    // var ch1 chan int
    // ch1 = make(chan int)
    // ch2 := make(chan int)
    // fmt.Println(cap(ch1))
    // fmt.Println(cap(ch2))

    // ch3 := make(chan int, 5)
    // fmt.Println(cap(ch3))

    // ch3 <- 1
    // fmt.Println(len(ch3))
    // ch3 <- 2
    // ch3 <-3 
    // fmt.Println(len(ch3))

    // i := <-ch3
    // fmt.Println(i)
    // fmt.Println("len",len(ch3))
    // i2 := <-ch3
    // fmt.Println(i2)
    // fmt.Println("len",len(ch3))
    // fmt.Println(<-ch3)
    // fmt.Println("len",len(ch3))

    // ch3 <-1
    // fmt.Println(ch3)
    // ch3 <-1
    // ch3 <-1
    // ch3 <-1
    // ch3 <-1
    // ch3 <-1
    // var ch2 <-chan int
    // var ch3 chan<- int

    // m := map[string]int{"A":100, "B":200, "C":300}
    // for i, v := range m {
    //     fmt.Println(i, v)
    // }

    // var m = map[string]int{"A":100, "B":200}
    // fmt.Println(m)
    // m2 := map[int]string{
    //     1: "A",
    //     2: "B",
    // }
    // fmt.Println(m2)

    // m3 := make(map[int]string)
    // fmt.Println(m3)
    // m3[1] = "japan"
    // m3[2] = "usa"
    // fmt.Println(m3)

    // fmt.Println(m["A"])
    // fmt.Println(m3[2])
    // fmt.Println(m3[4])

    // s, ok := m3[2]
    // if !ok {
    //     fmt.Println("error")
    // } else {
    //     fmt.Println(s, ok)
    // }

    // m3[3] = "china"
    // fmt.Println(m3)
    // delete(m3, 3)
    // fmt.Println(m3)
    // fmt.Println(len(m3))

    // fmt.Println(Sum(1,2,3))
    // fmt.Println(Sum(1,2,3,4,5,6,7,8,9))
    // fmt.Println(Sum())
    // sl := []int{1, 2,3}
    // fmt.Println(Sum(sl...))

    // sl := []string{"A", "B", "C"}
    // fmt.Println(sl)

    // for v := range sl {
    //     fmt.Println(v, sl[v])
    // }

    // for i:=0; i<len(sl); i++ {
    //     fmt.Println(sl[i])
    // }
    // sl := []int{100, 200}
    // sl2 := sl

    // sl2[0] = 1000
    // fmt.Println(sl)

    // var i int = 10
    // i2 := i
    // i2 = 200
    // fmt.Println(i, i2)
    // sl := []int{1,2,3,4,5}
    // sl2 := make([]int,5)
    // fmt.Println(sl2)
    // n := copy(sl2, sl)
    // fmt.Println(n, sl2)

    // sl := []int{100, 200}
    // fmt.Println(sl)

    // sl = append(sl, 300)
    // fmt.Println(sl)

    // sl = append(sl, 400, 500, 600)
    // fmt.Println(sl)

    // sl2 := make([]int, 5)
    // fmt.Println(sl2)
    // fmt.Println(len(sl2))
    // fmt.Println(cap(sl2))

    // sl3 := make([]int, 5, 10)
    // fmt.Println(len(sl3))
    // fmt.Println(cap(sl3))
    // sl3 = append(sl3, 1,2,3,4,5,6)
    // fmt.Println(len(sl3))
    // fmt.Println(cap(sl3))

    // var sl []int
    // fmt.Println(sl)

    // var sl2 []int = []int{100, 200}
    // fmt.Println(sl2)
    // var sl3 []string = []string{"a", "b"}
    // fmt.Println(sl3)

    // sl4 := make([]int, 5)
    // fmt.Println(sl4)

    // sl4[0] = 1000
    // fmt.Println(sl4)

    // sl5 := []int{1,2,3,4,5}
    // fmt.Println(sl5[0])
    // fmt.Println(sl5[2:4])
    // fmt.Println(sl5[:2])
    // fmt.Println(sl5[2:])
    // fmt.Println(sl5[:])
    // fmt.Println(sl5[len(sl5)-1])
    // fmt.Println(sl5[1:len(sl5)-1])
}
