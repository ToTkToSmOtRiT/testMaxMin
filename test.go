package main

import "fmt"

func maxMin(arr []int) {
	maxCount := [2]int{0, 0}
	minCount := [2]int{0, 0}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				if maxCount[1]-maxCount[0] < j-i {
					maxCount[1] = j
					maxCount[0] = i

				}
				break
			}
			if j+1 == len(arr) {
				if maxCount[1]-maxCount[0] < j+1-i {
					maxCount[1] = j + 1
					maxCount[0] = i
				}
			}
		}

		for j := i + 1; j < len(arr); j++ {
			if arr[j-1] < arr[j] {
				if minCount[1]-minCount[0] < j-i {
					minCount[1] = j
					minCount[0] = i

				}
				break
			}
			if j+1 == len(arr) {
				if minCount[1]-minCount[0] < j+1-i {
					minCount[1] = j + 1
					minCount[0] = i
				}
			}
		}

	}
	fmt.Print(arr[maxCount[0]:maxCount[1]], "\n")
	fmt.Print(arr[minCount[0]:minCount[1]])
}

func main() {
	x := []int{-1, -2, 0, 1, 2, 3, 2, 1}
	maxMin(x)
}
