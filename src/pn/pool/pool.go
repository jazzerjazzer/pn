package pool

import (
	"strings"
	"sync"
)

type Pool struct {
	id          string
	concurrency int
	wg          *sync.WaitGroup
	job         chan []string
	result      chan string
}

func New(id string, concurrency int, wg *sync.WaitGroup, job chan []string, result chan string) *Pool {
	return &Pool{
		id:          id,
		concurrency: concurrency,
		wg:          wg,
		job:         job,
		result:      result,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.concurrency; i++ {
		p.wg.Add(1)
		go find(p.id, p.job, p.result, p.wg)
	}
}

func find(id string, jobs <-chan []string, results chan<- string, wg *sync.WaitGroup) {
	for job := range jobs {
		for _, row := range job {
			if strings.Contains(row, id) {
				results <- row
			}
		}
	}
	wg.Done()
}
