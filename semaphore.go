 package semaphore // import "https://github.com/mkyas/semaphore"

import (
	"log"
	"sync"
)

type Semaphore struct {
	count    uint32
	deferred uint32
	lock     sync.Mutex
	queue    *sync.Cond
}

func NewSemaphore(n uint32) (result *Semaphore) {
	result = &Semaphore{count: n, queue: sync.NewCond(&result.lock)}
	return
}

func (s *Semaphore) Acquire() {
	s.lock.Lock()
	if s.count < 1 {
		log.Println("Waiting for permit")
		s.deferred++
		s.queue.Wait()
	} else {
		s.count -= 1
	}
	s.lock.Unlock()
}

func (s *Semaphore) Release() {
	s.lock.Lock()
	if s.deferred > 0 {
		log.Println("Signalling deferred process")
		s.deferred--
		s.queue.Signal()
	} else {
		s.count += 1
	}
	s.lock.Unlock()
}
