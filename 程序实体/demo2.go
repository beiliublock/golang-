package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name","everyone","The greeting object")
	name = getTheFlag()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","proEntity")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello,%v",*name)
}

func getTheFlag() *string {
	return flag.String("name","everyone","The greeting object")
}
