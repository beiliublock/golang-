package main

import "fmt"

type animal interface {
	//ScientificName  用于获取动物的学名
	ScientificName()  string
	//Category  用于获取动物的基本分类
	Category() string
}

type named interface {
	//Name  用于获取名字
	Name() string
}

type pet interface {
	animal		//嵌入animal接口
	named		//嵌入named接口
}

type petTag struct {
	name string
	owner string
}

func (pt petTag) Name() string {
	return pt.name
}

func (pt petTag)Owner() string {
	return pt.owner
}

type dog struct {
	petTag		//嵌入petTag结构体
	scientificName string
}

func (d dog) ScientificName() string {
	return d.scientificName
}

func (d dog) Category() string {
	return "dog"
}

func main() {
	pt := petTag{name:"little pig"}
	_,ok := interface{}(pt).(named)
	fmt.Printf("petTag implements interface named:%v\n",ok)
	d := dog{
		petTag:pt,
		scientificName:"Labrador Retriever",
	}
	_,ok = interface{}(d).(animal)
	fmt.Printf("dog implements interface animal:%v\n",ok)
	_,ok = interface{}(d).(named)
	fmt.Printf("dog implements interface named:%v\n",ok)
	_,ok = interface{}(d).(pet)
	fmt.Printf("dog implements interface pet:%v\n",ok)
	//Output:
	//petTag implements interface named:true
	//dog implements interface animal:true
	//dog implements interface named:true
	//dog implements interface pet:true
}