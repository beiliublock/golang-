package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

//方案1
func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

//方案2
type calculateFunc func(x, y int) (int, error)

func genCalculate(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	//方案1
	x,y := 12,23
	op := func(x,y int) int{
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result:%d ,error:%v\n",result,err)
	result,err = calculate(x,y,nil)
	fmt.Printf("The result:%d ,error:%v\n",result,err)

	//方案2
	x,y = 56,78
	add := genCalculate(op)
	result,err = add(x,y)
	fmt.Printf("The result:%d ,error:%v\n",result,err)
	//Output:
	//The result:35 ,error:<nil>
	//The result:0 ,error:invalid operation
	//The result:134 ,error:<nil>
}
