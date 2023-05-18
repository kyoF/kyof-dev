package main

import (
	"fmt"
)

// type Stringfy interface {
// 	ToString() string
// }

// type Person struct {
// 	Name string
// 	Age int
// }
// func (p *Person) ToString() string {
// 	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
// }

// type Car struct {
// 	Number string
// 	Model string
// }
// func (p *Car) ToString() string {
// 	return fmt.Sprintf("Number=%v, Model=%v", p.Number, p.Model)
// }

type MyError struct {
	Msg string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Msg
}

// type Point struct {
// 	A int
// 	B string
// }

// func (p *Point) String() string {
// 	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
// }

func main() {
	// p := Point{100, "ABC"}
	// fmt.Println(p)
	// fmt.Println(p.String())
	err := MyError{Msg: "error death", ErrCode: 1243}
	fmt.Println(err.Error())
	fmt.Println(err.ErrCode)

	// fmt.Println(err.Msg)
	// if ok {
	// 	fmt.Println(e.ErrCode)
	// }
	// vs := []Stringfy{
	// 	&Person{Name: "fujiki", Age: 21},
	// 	&Car{Number: "1234", Model: "Ver3"},
	// }

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }
}
