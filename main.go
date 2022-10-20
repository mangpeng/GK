package main

import (
	"fmt"
	"time"
)

func main() {
	fn := func(first int, second int) int {
		return first + second
	}(3, 4)

	println(fn)
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
