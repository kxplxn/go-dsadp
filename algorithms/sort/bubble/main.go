package main

import (
	"log"
	"math/rand"

	"github.com/kxplxn/learn_go-dsadp/algorithms/sort"
)

// https://www.geeksforgeeks.org/bubble-sort/

func BubbleSort[T sort.Ordered](arr []T) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
}

func main() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(101))
	}

	log.Printf("before bubble sort: %v", arr)
	BubbleSort(arr)
	log.Printf("after bubble sort: %v", arr)
}
