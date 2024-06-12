package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func send(c chan byte) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)*100))
		toWriteData := letters[rand.Intn(52)]
		c <- toWriteData
	}
}

func main() {
	channel := make(chan byte, 100)

	ctx, stop := context.WithTimeout(context.Background(), time.Second*10)

	defer stop()

	go send(channel)

	go func(c chan byte) {
		for {
			time.Sleep(time.Millisecond)
			if <-c != byte(0) {
				fmt.Println("reader read:" + string(<-c))
			}
		}
	}(channel)
	<-ctx.Done()

}
