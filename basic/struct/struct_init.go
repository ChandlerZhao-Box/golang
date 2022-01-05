package main

import "fmt"

type Duck struct {

}

func (d Duck) Swim()  {
	fmt.Println("我是鸭子，我会游泳...")
}

func main() {
	//新建的是一个指针, duck1是*Duck
	duck1 := &Duck{}
	duck1.Swim()

	//是结构体的实例
	duck2 := Duck{}
	duck2.Swim()

	//是指针, 是*Duck
	duck3 := new(Duck)
	duck3.Swim()

	//当这样声名变量的时候，go就帮你分配好了内存
	//不用担心空指针的问题，因为他压根就不空
	var duck4 Duck
	duck4.Swim()

	var duck5 *Duck
	//这个会报错的，只分配了指针的内存
	//duck5.Swim()
	duck5 = &Duck{}
	duck5.Swim()
}
