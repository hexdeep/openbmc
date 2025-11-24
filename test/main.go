package main

import (
	"fmt"
	"iter"
)

func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i >= 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	for v := range Countdown(5) {
		fmt.Println(v)
		if v == 2 {
			break // stops early
		}
	}
}
