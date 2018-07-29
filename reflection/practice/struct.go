package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	fmt.Printf("s: %v, type: %T, settability: %v, numField: %v\n", s, s.Type(), s.CanSet(), s.NumField())

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %v, type: %T\n", i, f, f)
	}
}
