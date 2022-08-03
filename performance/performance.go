package performance

import (
	"fmt"
	"time"
)

func Measure(fn func()) {
	var now = time.Now()
	fn()
	var elapsed = time.Since(now)

	fmt.Println("Executado em ", elapsed)
}
