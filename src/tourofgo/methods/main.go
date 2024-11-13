package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	fmt.Println(v)

	Scale(&v, 10)
	fmt.Println(Abs(v))

	v1 := Vertex{6, 8}
	v1.Scale(2)
	Scale(&v1, 100)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v1, *p)

	p1 := &Vertex{18, 13}
	fmt.Println(p1.Abs())
	fmt.Println(Abs(*p1))

	v2 := &Vertex{1.7, 0.4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs())
	v2.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs())
}

/*
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)


2005 rue king west
lunetrie iris

21 juin 14h

819-8210064

*/
