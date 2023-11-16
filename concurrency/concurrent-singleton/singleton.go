package concurrent_singleton

type singleton struct{}

var instance singleton

func GetInstance() *singleton { return &instance }

func (s *singleton) AddOne() {
	addCh <- struct{}{}
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- struct{}{}
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

var addCh = make(chan struct{})
var getCountCh = make(chan chan int)
var quitCh = make(chan struct{})

func init() {
	var count int

	go func(
		addCh <-chan struct{},
		getCountCh <-chan chan int,
		quitCh <-chan struct{},
	) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				break
			}
		}
	}(addCh, getCountCh, quitCh)
}
