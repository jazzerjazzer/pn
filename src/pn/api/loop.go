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

func Start(id string) string {
	start := time.Now()

	jobs := make(chan []string, 1000)
	results := make(chan string, 4000)

	var wg sync.WaitGroup
	pool := workerpool.New(id, goRoutines, &wg, jobs, results)
	pool.Start()

	r := reader.New("ids.csv", 1000)
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

	done := make(chan bool)
	found := ""
	go func() {
		read := 0
		for result := range results {
			if result != "" {
				fmt.Printf("Result: %s\n", result)
				found = result
			}
			read++
		}
		fmt.Printf("READ: %+v\n", read)
		done <- true
	}()
	wg.Wait()
	close(results)
	<-done
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s\n", elapsed)
	return found
}
