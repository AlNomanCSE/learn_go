package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int16)
	go func() {
		fmt.Println("Sending 42 channel ...")
		ch <- 42
		fmt.Println("Sent!")
	}()
	fmt.Println("Waiting to receive ...")
	value := <-ch
	fmt.Println("Received: ", value)

	time.Sleep(time.Microsecond * 100)

}
