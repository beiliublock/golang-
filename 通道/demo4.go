package main

import "fmt"

func ExampleDemo41() {
	ch := make(chan int,4)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("length:%d,capacity:%d\n",len(ch),cap(ch))
	//Output:
	//length:3,capacity:4
}

func main() {
	ch := make(chan int,4)
	a := 1
	ch <- a
	a = 2
	fmt.Printf("The element is %v",<-ch)
	//Output:
	//The element is 1
}

