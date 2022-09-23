package main

import (
	"errors"
	"log"
)

// https://www.geeksforgeeks.org/binary-search

func BinarySearchRange(arr []string, query string, low, high int) (int, error) {
	if high < low || len(arr) == 0 {
		return -1, errors.New("not found")
	}

	mid := low + (high-low)/2
	if arr[mid] > query {
		return BinarySearchRange(arr, query, low, mid-1)
	} else if arr[mid] < query {
		return BinarySearchRange(arr, query, mid+1, high)
	} else {
		return mid, nil
	}
}

func BinarySearch(arr []string, query string) (int, error) {
	return BinarySearchRange(arr, query, 0, len(arr)-1)
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
