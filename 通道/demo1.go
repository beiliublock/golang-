package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 2
	ch <- 1
	for i:=0;i<cap(ch);i++ {
		fmt.Printf("The %d element is %v\n",i,<-ch)
	}
	//Output:
	//The 0 element is 3
	//The 1 element is 2
	//The 2 element is 1
}
