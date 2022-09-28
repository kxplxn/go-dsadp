package main

import (
	"fmt"
)

// https://refactoring.guru/design-patterns/singleton

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		fmt.Println("Creating new instance now.")
		instance = &singleton{}
	} else {
		fmt.Println("Instance is already created.")
	}
	return instance
}

func main() {
	for i := 0; i < 5; i++ {
		GetInstance()
	}
}
