package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	//类型判断
	//方式1：类型断言
	_,ok1 := interface{}(container).([]string)
	_,ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n",container)
		return
	}
	fmt.Printf("The element is %q,container type: %T\n",container[1],container)

	//方式2：条件判断
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error:%s\n",err)
		return
	}
	fmt.Printf("The element is %q,container type: %T\n",elem,container)

}

func getElement(container interface{}) (elem string,err error) {
	switch t := container.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T",t)
		return
	}
	return
}
