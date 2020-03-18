package main

import (
  "fmt"
  "os"
  "strconv"
  "github.com/lirossarvet/euler_solutions/pkg/math_helpers"
)

func main() {
  number, _ := strconv.Atoi(os.Args[1])

  if math_helpers.IsPrime(number) {
    fmt.Printf("%d is PRIME\n", number)
  } else {
    fmt.Printf("%d is COMPOSITE\n", number)
  }
}
