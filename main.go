package main

import (
	"fmt"
	"time"
)

func main() {
	StartInterval(1, 10)
	for {
	}
}

func StartInterval(interval int, count int) {
	PlayInterval(interval, count)
}

var Ⲁ int

func StopInterval() {

}

func PlayInterval(interval int, count int) {

	for i := 0; i < count; i++ {
		ch := time.After(time.Duration(interval) * time.Second)
		<-ch
		fmt.Println("hi")
	}
}
