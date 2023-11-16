package concurrent_singleton_with_lock

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
}
