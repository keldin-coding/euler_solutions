package main

import "fmt"

var goalSum int = 1000

func main() {
  var value int

  for a := 1; a < goalSum; a++ {
    for b := 1; b < goalSum; b++ {
      for c := 1; c < goalSum; c++ {
        if a+b+c == goalSum && validTriangle(a, b, c) && isPythagTriple(a, b, c) {
          value = a * b * c
        }
      }
    }
  }

  fmt.Printf("Value is %v", value)
}

func isPythagTriple(a, b, c int) bool {
  if a*a+b*b == c*c {
    return true
  }

  return false
}

func validTriangle(a, b, c int) bool {
  if a+b > c && a+c > b && b+c > a {
    return true
  }

  return false
}
