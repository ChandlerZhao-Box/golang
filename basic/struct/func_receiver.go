package main

import "fmt"

//方法接收器
//要改变内部数据，只有指针接收器才能改成功
func main() {

	u := User {
		Name: "Tom",
		Age: 10,
	}
	//结构体调结构体接收器
	u.ChangeName("Tom changed")
	//结构体调指针接收器
	u.ChangeAge(100)
	fmt.Printf("user: %v \n", u)

	u2 := &User{
		Name: "Chandler",
		Age: 12,
	}
	//指针调结构体接收器
	u2.ChangeName("Tom changed")
	//指针调结指针接收器
	u2.ChangeAge(100)
	fmt.Printf("user2: %v \n", u2)

}

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(newName string) {
	u.Name = newName
}

func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}
