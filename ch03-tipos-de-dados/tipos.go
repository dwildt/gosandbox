package main

import "fmt"

func soma(a, b int) int {
    return a+b
}

func par(numero rune) bool {
  return (numero%2 == 0)
}

func main() {
    fmt.Println(soma(4,5))
    fmt.Println(par(4))
}
