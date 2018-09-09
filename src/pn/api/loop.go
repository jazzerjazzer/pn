package api

import (
	"fmt"
	"io"
	"log"
	workerpool "pn/pool"
	"pn/reader"
	"sync"
	"time"
)

const (
	goRoutines = 1000
)

func Search(id, filepath string) string {
	start := time.Now()

	jobs := make(chan []string, 1000)
	results := make(chan string, 4000)

	var wg sync.WaitGroup
	pool := workerpool.New(id, goRoutines, &wg, jobs, results)
	pool.Start()

	r := reader.New(filepath, 1000)
	for {
		rows, err := r.Read()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Error reading file: %v", err)
			}
			break
		}
		jobs <- rows
	}
	close(jobs)

	found := result(results, &wg)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)

	return found
}

func result(results chan string, wg *sync.WaitGroup) string {
	done := make(chan bool)
	found := ""
	go func() {
		for result := range results {
			if result != "" {
				found = result
			}
		}
		done <- true
	}()
	wg.Wait()
	close(results)
	<-done
	return found
}
