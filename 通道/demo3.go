package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	//发送方，让其关闭通道
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v\n", i)
			ch <- i
		}
		fmt.Println("Sender:close the channel")
		close(ch)
	}()
	//接收方
	for {
		elem,ok := <- ch
		//先判断通道是否关闭
		if !ok {
			fmt.Println("Receiver:closed channel")
			break
		}
		fmt.Printf("Receiver:received an element:%v\n",elem)
	}
	fmt.Println("End")
}
