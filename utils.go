package main

import (
	"log"
	"sync"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func deferDone(wg *sync.WaitGroup, start time.Time, id int, t *Times) {
	wg.Done()

	elapsed := time.Since(start)
	t.Elapsed = elapsed.Milliseconds()

	// log.Printf("Worker %v took %s", id, elapsed)
}
