package main

import (
	"fmt"
	"sort"
)

func main() {
	short := []int{0, 6, 11, 11, 18, 25, 28, 37, 40, 45, 47, 56, 59, 62, 74, 81, 81, 87, 89, 94}
	fmt.Println(binarySearch(short, 40, 0, len(short)))
}

func sortSlice(s []int, direction string) []int {
	if direction == "asc" {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	} else {
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
	}
	return s
}

func binarySearch(s []int, target, low, high int) int {
	val := 0
	for low != high {
		mid := (low + high) / 2
		if s[mid] == target {
			val = mid
			break
		}

		if s[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return val
}
