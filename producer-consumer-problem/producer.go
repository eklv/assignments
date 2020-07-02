package main

import (
	"fmt"
	"math/rand"
)

type producer struct {
	id         int
	widgetChan chan string
	name       string
}

func (p *producer) produce() {
	for {
		widgetLimit := getRandomNumber(1, 5)
		for i := 0; i < widgetLimit; i++ {
			go func(index int) {
				widget := RandStringBytes(5)
				p.widgetChan <- widget
				fmt.Println(p.name, "produced-", widget)
			}(i)
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
