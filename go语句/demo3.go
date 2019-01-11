package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//用原子变量的方式解决动态竞争
func main() {
	var count uint32 = 0
	trigger := func(i uint32,fn func()) {
		for {
			if n := atomic.LoadUint32(&count);n == i {
				fn()
				atomic.AddUint32(&count,1)
				break
			}
			time.Sleep(time.Millisecond)
		}
	}
	for i:=uint32(0);i<10;i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i,fn)
		}(i)
	}
	trigger(10,func(){})
}

//如果我们不用原子变量的方式操作会如何呢ExampleDemo3
func ExampleDemo3() {
	var count = 0
	trigger := func(i int, fn func()) {
		for {
			if count == i {
				fn()
				count++
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i:=0;i<10;i++ {
		go func(i int) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i,fn)
		}(i)
	}
	trigger(10,func(){})
	//我们发现并没有什么影响，但真的是这样吗，这里或许是单纯因为运气好而已，而这也往往会导致bug，所以一定
	//要注意使用原子操作保护竞态条件下的并发安全。
}
