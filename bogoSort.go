package main

import (
	"fmt"
	"math/rand"
	"time"
)

// isSorted checks whether a list is sorted by comparing first and next value.
func isSorted(list []int) bool {
	for i, _ := range list {
		if i != (len(list) - 1) {
			if list[i] > list[i+1] {
				return false
			}
		}

	}

	return true

}

// bogoSort shuffles a list.
func bogoSort(list *[]int) {
	rand.Shuffle(len(*list), func(i, j int) {
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := []int{4, 1, 7, 5, 8}
	tries := 0

	for {
		tries++
		fmt.Println(tries)
		bogoSort(&x)
		fmt.Println(x)
		if isSorted(x) {
			break
		}

	}

	fmt.Println(x)
}
