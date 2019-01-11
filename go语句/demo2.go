package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{},10)
	for i:=0;i<10;i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- struct{}{}
		}(i)
	}
	//方法1：小睡一会
	//time.Sleep(1*time.Second)
	//方法2：利用通道传递信息
	for i:=0;i<10;i++ {
		<- ch
	}
}
