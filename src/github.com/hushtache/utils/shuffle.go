package utils

import (
    "math/rand"
)

func ShuffleStrings(a []string) {
  for i := range a {
      j := rand.Intn(i + 1)
      a[i], a[j] = a[j], a[i]
  }
}

func ShuffleObjects(a []interface{}) {
  for i := range a {
      j := rand.Intn(i + 1)
      a[i], a[j] = a[j], a[i]
  }
}


func Slice(s []int, i int) []int {
  s[len(s)-1], s[i] = s[i], s[len(s)-1]
  return s[:len(s)-1]
}