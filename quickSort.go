package main

import "fmt"

func quickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := make([]int, 1)
	pivot[0] = list[0]
	leftList := []int{}
	rightList := []int{}

	for i := 1; i < len(list); i++ {
		if list[i] > pivot[0] {
			rightList = append(rightList, list[i])
		} else {
			leftList = append(leftList, list[i])
		}
	}

	leftSublist := quickSort(leftList)
	rightSublist := quickSort(rightList)

	answer := []int{}
	answer = append(leftSublist, pivot...)
	answer = append(answer, rightSublist...)
	return answer
}

func main() {
	x := []int{4, 6, 3, 2, 9, 7, 3, 5}
	fmt.Println(quickSort(x))
}
