package utils

import "math/rand"

func Generate(length, lower, upper int) (slice []int) {
	slice = make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn((upper - lower) + lower)
	}
	return
}
