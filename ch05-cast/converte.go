package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	a := 1         // int
	fmt.Println(a) // imprime 1

	s := strconv.Itoa(a)
	fmt.Println(reflect.TypeOf(s)) // imprime o tipo da variavel string
}
