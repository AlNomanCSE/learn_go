package main

import (
	"fmt"
	"time"
)

func performTask(data int, callback func(int) string) string {
	fmt.Println("Processing data:", data)
	time.Sleep(1 * time.Second)
	return callback(data)
}
func main() {
	myCallBack := func(result int) string {
		return fmt.Sprintf("Processed Result: %d ", result*2)
	}
	output := performTask(5, myCallBack)

	fmt.Println(output)
}
