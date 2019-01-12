package main

import "fmt"

func main() {
	//示例1
	value3 := [...]int8{0,1,2,3,4,5,6}
	/*switch value3[4] { 无法通过编译，因为子句中有重复项
	case 0,1,2:
		fmt.Println("0 or 1 or 2")
	case 2,3,4:
		fmt.Println("2 or 3 or 4")
	case 4,5,6:
		fmt.Println("4 or 5 or 6")
	}*/

	//示例2：
	switch value3[3] {
	case value3[0],value3[1],value3[2]:
		//虽然这里通过索引表达式成功骗过了编译器
		//但是是不可取的
		fmt.Println("0 or 1 or 2")
	case value3[2],value3[3],value3[4]:
		fmt.Println("2 or 3 or 4")
	case value3[4],value3[5],value3[6]:
		fmt.Println("4 or 5 or 6")
	}
	//Output:
	//2 or 3 or 4

	//示例3：
	/*value4 := interface{}(byte(127))
	switch t := value4.(type) {   //uint8和byte重复
	case uint8,uint16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Printf("unsupported type:%T",t)
	}*/
}
