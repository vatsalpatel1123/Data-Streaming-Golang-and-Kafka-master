package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	const numStreams = 1000
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numStreams; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			http.Post(fmt.Sprintf("http://localhost:8080/stream/start"), "application/json", nil)
			http.Post(fmt.Sprintf("http://localhost:8080/stream/%d/send", id), "application/json", nil)
			http.Get(fmt.Sprintf("http://localhost:8080/stream/%d/results", id))
		}(i)
	}
	wg.Wait()
	fmt.Printf("Processed %d streams in %s\n", numStreams, time.Since(start))
}
