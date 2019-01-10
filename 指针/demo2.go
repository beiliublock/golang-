package main

type cat struct {
	name string
}

func myNew(name string) cat {
	return cat{name}
}

func (c *cat)SetName(name string) {
	c.name = name
}

func (c cat)Name() string {
	return c.name
}

func main() {
	//示例1：
	//myNew("little pig").SetName("monster")  //不能调用不可寻址的值的指针方法
	//# command-line-arguments
	//指针\demo2.go:21:21: cannot call pointer method on myNew("little pig")  果
	//指针\demo2.go:21:21: cannot take the address of myNew("little pig")  因

	//示例2：
	map[string]int{"one":1,"two":2,"three":3}["one"]++
	//注意，虽然自增和自减的左侧表达式必须是可取址的，但是map是一个例外
	map1 := map[string]int{"one":1,"two":2,"three":3}
	map1["one"]++
}