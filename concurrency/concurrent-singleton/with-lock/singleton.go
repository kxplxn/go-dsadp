package concurrent_singleton_with_lock

import "sync"

// singleton uses a sync.RWMutex to differentiate between read and write locking
// and more efficiently handle cases where two different goroutines are reading
// from the concurrent-safe value as opposed to one reading and one writing.
//
// With sync.RWMutex, the lock can be held by an arbitrary number of readers OR
// a single writer.
type singleton struct {
	count int
	sync.RWMutex
}

var instance singleton

func GetInstance() *singleton { return &instance }

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count += 1
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
