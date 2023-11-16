package main

import (
	"fmt"
	"math/rand"

	"github.com/kxplxn/learn_go-dsadp/algorithms/sort"
)

func InsertionSort[T sort.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		iterator, temporary := currentIndex, arr[currentIndex]
		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}

func main() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Printf("before: %v\n", arr)
	InsertionSort(arr)
	fmt.Printf("after:  %v\n", arr)
}
