package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	front := l.PushBack("1")
	element := list.Element{Value:"10"}
	l.PushFront("0")
	l.MoveAfter(&element,front)
	for i:=l.Front();i!=nil;i=i.Next(){
		fmt.Printf("%v  ",i.Value)
	}
}
