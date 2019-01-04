package main

import "fmt"

func main() {
	a1 := [7]int{1,2,3,4,5,6,7}
	s1 := a1[1:4]
	fmt.Printf("s1:%d,len(s1):%d,cap(s1):%d\n",s1,len(s1),cap(s1))

	for i := 1;i<=5;i++ {
		//打印a1的地址
		fmt.Print("a1[6]:")
		for i:=1;i<7;i++ {
			fmt.Printf("%p ",&a1[i])
		}
		fmt.Println()
		//打印a1的值
		fmt.Print("a1[6]:")
		for i:=1;i<7;i++ {
			fmt.Printf("%d ",a1[i])
		}
		fmt.Println()
		s1 = append(s1,i)
		//打印s1的地址
		fmt.Printf("s1[%d]:",len(s1))
		for i := range s1 {
			fmt.Printf("%p ",&s1[i])
		}
		fmt.Printf("%d",cap(s1))
		fmt.Println()
		//打印s1的值
		fmt.Printf("s1[%d]:",len(s1))
		for _,v := range s1 {
			fmt.Printf("%d ",v)
		}
		fmt.Println()
	}
}
