package main

import (
	"fmt"
)

func mergeSort(list []int) []int {
	// If list only have 1 number or none, return it.
	if len(list) < 2 {
		return list
	}

	// Split recursion
	left := mergeSort(list[:len(list)/2])
	right := mergeSort(list[len(list)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	answer := []int{}
	leftCount := 0
	rightCount := 0

	for leftCount < len(left) && rightCount < len(right) {
		if left[leftCount] < right[rightCount] {
			answer = append(answer, left[leftCount])
			leftCount++
		} else {
			answer = append(answer, right[rightCount])
			rightCount++
		}
	}

	for ; leftCount < len(left); leftCount++ {
		answer = append(answer, left[leftCount])
	}

	for ; rightCount < len(right); rightCount++ {
		answer = append(answer, right[rightCount])
	}

	return answer
}

func main() {
	unsortedArray := []int{8, 2, 5, 3, 4, 7, 6, 1}
	sortedArray := mergeSort(unsortedArray)

	fmt.Println(sortedArray)
}
