package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	// This nil check is to avoid an expensive lock operation on every call to
	// GetInstance by making sure that the instance is nil to begin with.
	if instance == nil {
		// This block is to ensure that if multiple threads bypass the first
		// nil check, only one of them will be able to create the instance.
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			fmt.Println("Creating new instance now.")
			instance = &Singleton{}
		} else {
			fmt.Println("Instance is already created.")
		}
	} else {
		fmt.Println("Instance is already created.")
	}
	return instance
}

func main() {
	for i := 0; i < 5; i++ {
		go GetInstance()
	}
	time.Sleep(100 * time.Millisecond)
}
