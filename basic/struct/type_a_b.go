package main

import "fmt"

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼，假装自己是一只鸭子\n")
}

//定义了一个新类型
//扩展这个fish的方法
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼\n")
}

//定义一个新类型
type StrongFakeFish Fish

func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼，嘎嘎嘎\n")
}

func main() {
	fake := FakeFish{}
	//这里无法直接调用Swim方法
	fake.FakeSwim()

	//转换为Fish就可以调用Swim了
	td := Fish(fake)
	//真的成了鱼
	td.Swim()

	sFake := StrongFakeFish{}
	sFake.Swim()

	td = Fish(sFake)
	td.Swim()
}
