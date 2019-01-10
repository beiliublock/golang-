package main

import (
	"fmt"
	"unsafe"
)

type pig struct {
	name string
	age int
	age2 uint
	name2 string
}

func (p *pig)SetName(name string) {
	p.name = name
}

func (p pig)Name() string {
	return p.name
}

func main() {
	//示例1：
	p := pig{"little pig",1,1,"2"}
	pp := &p  //获取变量p的内存地址
	pPtr := uintptr(unsafe.Pointer(pp)) //通过两次转换将指针值转换为uintptr类型的值
	fmt.Println(pPtr)
	namePtr := pPtr + unsafe.Offsetof(pp.name)
	fmt.Println(namePtr)	//我们会发现结构体的地址就是首字段的地址
	//fmt.Println(unsafe.Sizeof(p.name2))	//我们通过Sizeof方法可以得知字段所占的字节大小
	//fmt.Println(unsafe.Alignof(p.age))
	nameP := (*string)(unsafe.Pointer(namePtr)) //指针值与unsafe.Pointer之间的转换
	fmt.Printf("nameP == &(pp.name)?%v\n",nameP == &(pp.name))

	*nameP = "monster" //修改p的name
	fmt.Printf("The name of pig is %q\n",p.name)
	//Output:
	//nameP == &(pp.name)?true
	//The name of pig is "monster"

	//示例2：
	//下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number:%d\n",num)
	//Output:
	//The name of pig is "monster"
	//This is an unexpected number:4979112
}
