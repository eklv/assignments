package main

import (
	"fmt"
	"time"
)

type consumer struct {
	id   int
	name string
}

func (c *consumer) consume(p []*producer) {
	doneCounter := 0
	count := 0
	i := 0
	producer := p[i]
	widgetChan := producer.widgetChan
	for {
		select {
		case widget := <-widgetChan:
			fmt.Println(c.name, "consumed", widget)
			count++
			if count == 10 {
				count = 0
				duration := getRandomNumber(1, 5)
				fmt.Println(c.name, "Consumed 10 widgets... going to sleep for", duration, "seconds/seconds")
				doneCounter++
				if doneCounter == 10 {
					fmt.Println(c.name, "consumed 10 widgets 10 times, exiting...")
					return
				}
				time.Sleep(time.Duration(duration) * time.Second)
			}
		default:
			if i >= 3 {
				i = 0
				widgetChan = p[0].widgetChan
			} else {
				i++
				widgetChan = p[i].widgetChan
			}
		}
	}
}
