package utils

import (
	"fmt"
	"time"
)

func RunTimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("RunTimeCost  = %v\n", tc)
	}
}
