package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(c)
}

func StartInterval(interval int, count int) {
	PlayInterval(interval, count)
}

type Color int

const (
	a Color = iota
	b
	c
	d
	e
)

func StopInterval() {

}

func PlayInterval(interval int, count int) {

	for i := 0; i < count; i++ {
		ch := time.After(time.Duration(interval) * time.Second)
		<-ch
		fmt.Println("hi")
	}
}
