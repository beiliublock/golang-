package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	if value,ok := interface{}(container).([]string);ok {
		fmt.Println(value)
	}
	fmt.Printf("The element is %q\n", container[1])
}
