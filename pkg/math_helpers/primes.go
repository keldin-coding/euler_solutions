package math_helpers

import "math"

func IsPrime(n int) bool {
  root := int(math.Sqrt(float64(n)))

  for d := 2; d <= root+1; d++ {
    if n%int(d) == 0 {
      return false
    }
  }
  return true
}
