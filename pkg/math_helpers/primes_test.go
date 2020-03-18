package math_helpers

import "testing"

func TestIsPrime(t *testing.T) {
  t.Run("2 is Prime", func(t *testing.T) {
    t.Parallel()

    if IsPrime(2) == false {
      t.Fail()
    }
  })

  t.Run("4 is Composite", func(t *testing.T) {
    t.Parallel()

    if IsPrime(4) == true {
      t.Fail()
    }
  })

  t.Run("100 is Composite", func(t *testing.T) {
    t.Parallel()

    if IsPrime(100) == true {
      t.Fail()
    }
  })

  t.Run("3 is Prime", func(t *testing.T) {
    t.Parallel()

    if IsPrime(3) == false {
      t.Fail()
    }
  })

  t.Run("17 is Prime", func(t *testing.T) {
    t.Parallel()

    if IsPrime(17) == false {
      t.Fail()
    }
  })
}
