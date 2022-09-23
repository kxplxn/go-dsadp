package main

import (
	"log"
	"math/rand"

	"github.com/kxplxn/learn_go-dsadp/algorithms/sort"
)

// https://www.geeksforgeeks.org/merge-sort

func MergeSort[T sort.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	l := MergeSort(arr[:mid])
	r := MergeSort(arr[mid:])
	return merge(l, r)
}

func merge[T sort.Ordered](l, r []T) []T {
	res := make([]T, len(l)+len(r))

	iL, iR := 0, 0
	for iL < len(l) && iR < len(r) {
		if l[iL] <= r[iR] {
			res[iL+iR] = l[iL]
			iL++
		} else {
			res[iL+iR] = r[iR]
			iR++
		}
	}

	for iL < len(l) {
		res[iL+iR] = l[iL]
		iL++
	}

	for iR < len(r) {
		res[iL+iR] = r[iR]
		iR++
	}

	return res
}

func main() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	log.Printf("before merge sort: %v", arr)
	arr = MergeSort(arr)
	log.Printf("after merge sort: %v", arr)
}
