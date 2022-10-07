package utils

import "math/rand"

func GenerateRandomNumSlice(length, lower, upper int) (slice []int) {
	slice = make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn((upper - lower) + lower)
	}
	return
}

func GenerateSortedSlice(length int) (slice []int) {
	slice = make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i
	}
	return
}
