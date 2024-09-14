// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
		// fmt.Println( /*"%d: %v \n",*/ i, arg)
		fmt.Printf("%d: %v \n", i, arg)
	}
	fmt.Println(s)
}

/*
s := ""
var s string
var s = ""
var s string = ""
*/
