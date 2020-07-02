package main

import (
	"fmt"
	"time"
)

type consumer struct {
	id             int
	widgetCapacity []string
	name           string
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
				//	fmt.Println(p[i].name, "doesnt have anymore widgets,", len(p[i].widgetChan), c.name, "switching to", p[0].name, len(p[0].widgetChan))
				i = 0
				widgetChan = p[0].widgetChan
			} else {
				i++
				//	fmt.Println(p[i-1].name, "doesnt have anymore widgets,", len(p[i-1].widgetChan), c.name, "switching to", p[i].name, len(p[i-1].widgetChan))
				widgetChan = p[i].widgetChan
			}
		}
	}
}
