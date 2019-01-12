package main

import "fmt"

func main() {
	//示例1
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i //将numbers1[i]和i按位或运算然后重新赋给numbers[i]
			//numbers[3] = 4 | 3
		}
	}
	fmt.Println(numbers1)
	fmt.Println()
	//Output:
	//[1 2 3 7 5 6]

	//示例2
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1 //=5
	for i, e := range numbers2 {
		if i == maxIndex2 {  //当i=5时，将第一个元素与第五个元素相加然后赋给第一个元素
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()
	//Output:
	//[7 3 5 7 9 11]

	//示例3：
	numbers3 := []int{1,2,3,4,5,6}
	maxIndex3 := len(numbers3) - 1
	for i,e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		}else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
	fmt.Println()
	//Output:
	//[22 3 6 10 15 21]

	//说明，数组是值类型的，切片是引用类型，而range表达式迭代的都是对象的副本
}
