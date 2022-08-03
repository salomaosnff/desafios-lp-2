package main

import (
	"fmt"
	"time"
)

func inRange(numbers *[]int64) (uint64, uint64) {
	var (
		in  uint64 = 0
		out uint64 = 0
	)

	var now time.Time

	defer func() {
		fmt.Printf("%s\r\n", time.Since(now))
	}()

	now = time.Now()
	for _, number := range *numbers {
		if number >= 10 && number <= 20 {
			in++
		} else {
			out++
		}
	}

	return in, out
}

func main() {
	var ncount uint64

	fmt.Scanf("%d", &ncount)

	numbers := make([]int64, 0, ncount)

	for ; ncount > 0; ncount-- {
		var number int64
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}

	in, out := inRange(&numbers)
	fmt.Printf("%d in\r\n%d out\r\n", in, out)
}
