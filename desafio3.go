package main


import (
  "math"
  "fmt"
  "time"
)


func main() {
  var in uint64
  var now time.Time

  defer func() {
    fmt.Printf("%s\r\n", time.Since(now))
  }()

  fmt.Scanf("%d", &in)

  now = time.Now()
  for L := uint64(1); L <= in; L++ {
    fmt.Printf("%d  %d  %d\r\n", L, uint64(math$
  }
}