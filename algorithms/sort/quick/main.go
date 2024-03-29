package main

import (
	"log"
	"math/rand"

	"github.com/kxplxn/learn_go-dsadp/algorithms/sort"
)

// https://www.geeksforgeeks.org/quick-sort/

func QuickSort[T sort.Ordered](arr []T) {
	quickSortRange(arr, 0, len(arr)-1)
}

func quickSortRange[T sort.Ordered](arr []T, low, high int) {
	if len(arr) < 2 {
		return
	}

	if low < high {
		iPivot := partition(arr, low, high)
		quickSortRange(arr, low, iPivot-1)
		quickSortRange(arr, iPivot+1, high)
	}
}

func partition[T sort.Ordered](arr []T, low, high int) int {
	iSwap := low
	pivotItem := arr[high]

	for i := low; i < high; i++ {
		if arr[i] <= pivotItem {
			arr[iSwap], arr[i] = arr[i], arr[iSwap]
			iSwap++
		}
	}

	arr[iSwap], arr[high] = pivotItem, arr[iSwap]
	return iSwap
}

func main() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(101))
	}

	log.Printf("before quick sort: %v", arr)
	QuickSort(arr)
	log.Printf("after quick sort: %v", arr)
}
