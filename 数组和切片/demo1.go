package main

import "fmt"

func ExampleDemo11() {
	s1 := make([]int,5)
	//切片由三部分组成：指向底层数组的索引，长度，容量
	//这里我们只指定length，那么capacity会和length一致
	fmt.Printf("The length of s1:%d\n",len(s1))
	fmt.Printf("The capacity of s1:%d\n",cap(s1))
	fmt.Printf("The value of s1:%d\n",s1)
	s2 := make([]int,5,8)
	//这里我们同时指定length和capacity
	fmt.Printf("The length of s2:%d\n",len(s2))
	fmt.Printf("The capacity of s2:%d\n",cap(s2))
	fmt.Printf("The value of s2:%d\n",s2)
	//Output:
	//The length of s1:5
	//The capacity of s1:5
	//The value of s1:[0 0 0 0 0]
	//The length of s2:5
	//The capacity of s2:8
	//The value of s2:[0 0 0 0 0]
}

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	//Output:
	//The length of s4: 3
	//The capacity of s4: 5
	//The value of s4: [4 5 6]
}
