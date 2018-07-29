package main

import (
	"fmt"
	"reflect"
)

func hello() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Printf("x : %f, Type: %T\n", x, x)
	fmt.Printf("p : %v, Type: %T, Settability: %v\n", p, p, p.CanSet())
	v := p.Elem()
	fmt.Printf("v : %v, Type: %T, Settability: %v\n", v, v, v.CanSet())
	v.SetFloat(7.1)
	fmt.Printf("x : %v, p: %v, v: %v\n", x, p, v)
}
