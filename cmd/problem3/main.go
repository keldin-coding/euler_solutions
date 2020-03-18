package main

import (
  "fmt"
  "os"
  "strconv"

  "github.com/lirossarvet/euler_solutions/pkg/math_helpers"
)

func main() {
  number, _ := strconv.Atoi(os.Args[1])

  // Quick check to start: if our number is prime, immediately bail
  if math_helpers.IsPrime(number) {
    fmt.Printf("%d is the largest prime factor\n", number)
  }

  factors := make([]int, 0)

  for i := 2; i <= number; i++ {
    if math_helpers.IsPrime(i) && number%i == 0 {
      number = number / i
      factors = append(factors, i)
      // We set it to 1 because the ++ happens after, really resetting it to 2
      i = 1
    }
  }

  fmt.Printf("factors: %v", factors)
}
