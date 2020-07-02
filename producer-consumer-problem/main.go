package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	producers := make([]*producer, 0)
	for i := 1; i < 5; i++ {
		p := &producer{
			id:         i,
			name:       fmt.Sprintf("producer-%d", i),
			widgetChan: make(chan string, 10),
		}
		producers = append(producers, p)
	}

	for _, p := range producers {
		go p.produce()
	}

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 1; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			consumer := &consumer{
				id:   id,
				name: fmt.Sprintf("consumer-%d", id),
			}
			consumer.consume(producers)
		}(i)
	}
	wg.Wait()
}

func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
