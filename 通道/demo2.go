package main

//对于缓冲通道的操作
func ExampleDemo21() {
	ch := make(chan int,1)
	ch <- 1
	ch <- 2	//通道已满，造成永久阻塞
	//Output:
	//fatal error: all goroutines are asleep - deadlock!
}

//对于非缓冲通道的操作
func ExampleDemo22() {
	ch := make(chan int)
	ch <- 1 //收发操作必须“同时”出现，不然就会造成永久阻塞
	<- ch
	//Output:
	//fatal error: all goroutines are asleep - deadlock!
}

//对于未初始化的通道的操作
func ExampleDemo23() {
	var ch chan int		//任何收发操作都会造成永久阻塞
	//ch <- 2
	<- ch
	//Output:
	//fatal error: all goroutines are asleep - deadlock!
}


