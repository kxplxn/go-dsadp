package main

import (
	"errors"
	"log"
)

// https://www.geeksforgeeks.org/binary-search

func BinarySearch(arr []string, query string) (int, error) {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := high + (low-high)/2
		if arr[mid] < query {
			low = mid + 1
		} else if arr[mid] > query {
			high = mid - 1
		} else {
			return mid, nil
		}
	}

	return -1, errors.New("not found")
}

func main() {
	items := []string{"box", "crooked", "delight", "joke", "red"}

	query := "crooked"
	if i, err := BinarySearch(items, query); err != nil {
		log.Printf("err: %v\n", err)
	} else {
		log.Printf("%v found at index %v", query, i)
	}

	query = "deli"
	if i, err := BinarySearch(items, query); err != nil {
		log.Printf("err on %v: %v\n", query, err)
	} else {
		log.Printf("%v found at index %v", query, i)
	}
}
