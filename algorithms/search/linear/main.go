package main

import (
	"errors"
	"log"
)

// https://www.geeksforgeeks.org/linear-search

func LinearSearch(items []string, query string) (int, error) {
	for i, v := range items {
		if v == query {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func main() {
	items := []string{"box", "red", "joke", "crooked", "delight"}

	query := "crooked"
	if i, err := LinearSearch(items, query); err != nil {
		log.Printf("err: %v\n", err)
	} else {
		log.Printf("%v found at index %v", query, i)
	}

	query = "wok"
	if i, err := LinearSearch(items, query); err != nil {
		log.Printf("err: %v\n", err)
	} else {
		log.Printf("%v found at index %v", query, i)
	}
}
