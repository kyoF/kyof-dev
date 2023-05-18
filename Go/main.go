package main

import (
	"fmt"
)

// type User struct {
// 	Name string
// 	Age int
// 	// X, Y int
// }

// type Users []*User

// type Users struct {
// 	Users []*Users
// }

// type T struct {
// 	User
// }

// func (u *User) SetName() {
// 	u.Name = "A"
// }

// func UpdateUser(user User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// func UpdateUser2(user *User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// func (u User) SayName() {
// 	fmt.Println(u.Name)
// }

// func (u User) SetNameTmp(name string) {
// 	u.Name = name
// }

// func (u *User) SetName(name string) {
// 	u.Name = name
// }

// func NewUser(name string, age int) *User {
// 	return &User{Name: name, Age: age}
// }

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	mi.Print()

	// i := 100 as MyInt
	// fmt.Println(mi + i)
	// m := map[int]User{
	// 	1: {Name: "user1", Age: 10},
	// 	2: {Name: "user2", Age: 20},
	// }
	// fmt.Println(m)

	// m2 := map[User]string{
	// 	{Name: "user1", Age: 10}: "tokyo",
	// 	{Name: "user2", Age: 20}: "LA",
	// }
	// fmt.Println(m2)

	// m3 := make(map[int]User)
	// fmt.Println(m3)
	// m3[1] = User{Name: "user3"}
	// fmt.Println(m3)

	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// user1 := User{Name: "user1", Age: 10}
	// user2 := User{Name: "user2", Age: 20}
	// user3 := User{Name: "user3", Age: 30}
	// user4 := User{Name: "user4", Age: 40}
	// users := Users{}
	// users = append(users, &user1)
	// users = append(users, &user2)
	// users = append(users, &user3)
	// users = append(users, &user4)

	// fmt.Println(users)
	// for _, v := range users {
	// 	fmt.Println(*v)
	// }

	// users2 := make([]*User, 0)
	// users2 = append(users2, &user1, &user2)
	// for _, v := range users2 {
	// 	fmt.Println(*v)
	// }
	// user := NewUser("fujiki", 10)
	// fmt.Println(*user)
	// t := T{User: User{Name: "fujiki", Age: 10}}
	// fmt.Println(t)
	// fmt.Println(t.User)
	// fmt.Println(t.User.Name)
	// fmt.Println(t.Name)

	// t.User.SetName()
	// fmt.Println(t)
	// user1 := User{Name: "fujiki"}
	// user1.SayName()

	// user1.SetName("A")
	// user1.SayName()

	// user2 := &User{Name: "kyosuke"}
	// user2.SayName()
	// user2.SetNameTmp("B")
	// user2.SayName()

	// var user1 User
	// fmt.Println(user1)
	// user1.Name = "user1"
	// user1.Age = 10
	// fmt.Println(user1)

	// user2 := User{}
	// fmt.Println(user2)
	// user2.Name = "user2"
	// fmt.Println(user2)

	// user3 := User{Name: "user3", Age: 30}
	// fmt.Println(user3)

	// user4 := User{"user4", 40}
	// fmt.Println(user4)

	// // user5 := User{30, "user5"}
	// // fmt.Println(user5)

	// user6 := User{Name: "user6"}
	// fmt.Println(user6)

	// user7 := new(User)
	// fmt.Println(user7)

	// user8 := &User{}
	// fmt.Println(user8)

	// UpdateUser(user1)
	// UpdateUser2(user8)
	// fmt.Println(user1)
	// fmt.Println(user8)
}
