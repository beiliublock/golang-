package main

import (
	"flag"
	"fmt"
	"os"
)

//go run demo1.go -name BaoZhiHao
//output:Hello,BaoZhiHao
func main11() {
	var name string  //指定类型声明
	cmd := flag.NewFlagSet("proEntity",flag.ExitOnError)
	cmd.StringVar(&name,"name","everyone","The greeting object")
	cmd.Parse(os.Args[1:])//注意os.Args的第一个元素是程序名，所以我们忽略
	fmt.Printf("Hello,%v",name)
}

func main()  {
	cmd := flag.NewFlagSet("proEntity",flag.ExitOnError)
	name := cmd.String("name", "everyone", "The greeting object")
	//需要注意的是，这里不能直接使用*取出内容，不然无法解析，会一直输出默认值：everyone
	//短变量声明，注意这种声明方式只能用于函数内部
	cmd.Parse(os.Args[1:])
	fmt.Printf("Hello,%v",*name)//注意flag.String输出的是地址，所以我们用*取出内容
}
