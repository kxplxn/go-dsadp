package main

import (
	"fmt"
	"github.com/kxplxn/learn_go-dsadp/algorithms/sort"
	"math/rand"
)

// https://www.geeksforgeeks.org/selection-sort

func SelectionSort[T sort.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		iMin := i
		for j := iMin; j < len(arr); j++ {
			if arr[j] < arr[iMin] {
				iMin = j
			}
		}
		arr[i], arr[iMin] = arr[iMin], arr[i]
	}
}

func main() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}

	fmt.Printf("before: %v\n", arr)
	SelectionSort(arr)
	fmt.Printf("after:  %v\n", arr)
}
