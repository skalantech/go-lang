package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type Person struct {
	Name string
	Age  int
}

type IPAddr [4]byte

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t1 *T
	i = t1
	describe(i)
	i.M()

	var ei interface{}
	describeInterface(ei)

	ei = 42
	describeInterface(ei)

	ei = "hello"
	describeInterface(ei)

	// Type assertions
	var ia interface{} = "Holala"

	s := ia.(string)
	fmt.Println(s)

	s, ok := ia.(string)
	fmt.Println(s, ok)

	f, ok := ia.(float64)
	fmt.Println(f, ok)

	// f = ia.(float64) // panic
	// fmt.Println(f)

	do(21)
	do("hello")
	do(true)

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

//	type Stringer interface {
//	    String() string
//	}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
