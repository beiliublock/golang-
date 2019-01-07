package main

import "fmt"

func main() {
	ch := make(chan int,5)
	ch <- 0
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	//这里需要注意的是，如果不手动关闭通道，使用range取的时候就会产生永久的阻塞
	for v:= range ch {
		fmt.Println(v)
	}
}
