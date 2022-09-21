package main

import (
	"log"
	"math/rand"
)

// https://www.geeksforgeeks.org/bubble-sort/

func BubbleSort(arr []int) {
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
