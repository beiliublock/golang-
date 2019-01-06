package main

import "fmt"

func main() {
	var bMap map[int]string
	defer func() {
		if err := recover();err != nil {
			fmt.Printf("The error %v",err)
		}
	}()
	key := 2
	elem,ok := bMap[key]
	fmt.Printf("The element paired with key %d in nil map[%d]%v,%v\n",key,key,elem,ok)
	fmt.Printf("The length of bMap is %d\n",len(bMap))
	fmt.Printf("Delete the key-element pair by key %d...\n", key)
	delete(bMap,key)
	elem = "two"
	fmt.Println("Add a key-element pair to a nil map...")
	bMap[key] = elem
	//Output:
	//The element paired with key 2 in nil map[2],false
	//The length of bMap is 0
	//Delete the key-element pair by key 2...
	//Add a key-element pair to a nil map...
	//The error assignment to entry in nil map
}
