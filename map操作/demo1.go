package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	if v, ok := aMap[k]; ok {
		fmt.Printf("The element of key %q:%d\n", k, v)
	} else {
		fmt.Println("Not found")
	}
	//Output:
	//The element of key "two":2
}
