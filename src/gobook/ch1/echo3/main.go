// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(os.Args[1:])
	// Prints a slice with the arguments
}
