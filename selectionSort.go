package main

import "fmt"

func selectionSort(list *[]int) {
	sorted := make([]int, 0, len(*list))
	unSorted := *list

	for i := 0; i < len(*list); i++ {
		min := unSorted[0]
		index := 0

		for i, e := range unSorted {
			if e <= min {
				min = e
				index = i
			}
		}

		sorted = append(sorted, (unSorted)[index])
		if index == 0 {
			unSorted = unSorted[1:]
		} else if index == (len(*list) - 1) {
			unSorted = unSorted[:index]
		} else {
			unSorted = append(unSorted[0:index], unSorted[index+1:]...)
		}

	}

	*list = sorted
}

func main() {
	x := []int{5, 2, 3, 1, 4, 6, 7, 21, 8, 10, 9, 15, 13, 16}
	fmt.Println(x)

	selectionSort(&x)

	fmt.Println(x)

}
