package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var once = sync.Once{}

type srv struct {
	callCount int

	sync.Mutex
}

func (s *srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	s.callCount += 1
	s.Unlock()

	fmt.Fprint(w, "Hello World")

	// s.callRepository()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill)
}

func main() {
	// Recommendation: don't use a buffered channel unless you reaaaally know you
	// need it.
	// sem := make(chan struct{}, 10)

	// receiveOnly(sem)

	// 100 orders

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i += 1 {
		wg.Add(1)
		go func() {
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Hello World")

	http.ListenAndServe(":8080", &srv{})
}

// func receiveOnly(c <-chan struct{}) {
// c <- struct{}{}
// }
