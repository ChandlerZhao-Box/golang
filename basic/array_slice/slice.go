package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3)
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 3, 4)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := append(s3, 7)
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))

	s5 := append(s3, 7, 8)
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

	//修改底层元素，查看基于这个切片创建出来的切片数据共享情况
	s3[2] = 100
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))
	fmt.Printf("s5: %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

}
