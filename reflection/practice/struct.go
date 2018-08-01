package main

import (
	"fmt"
)

type test interface {
	hello() string
}

type hotdog int

func main() {
	var d hotdog
	fmt.Printf("%T\n", d)
	fmt.Println(d.hello())
	bye(d)
}

func (d hotdog) hello() string {
	return "hello"
}

func bye(_ hotdog) {
	fmt.Println("Bye")
}
