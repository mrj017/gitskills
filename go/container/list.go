package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l)
	l.PushBack("er")
	l.PushBack("tyty")

	l.PushFront("ooppoop")
	e := l.Front()
	fmt.Println(e.Value)
	fmt.Println(e.Prev())
	l.Init()
	fmt.Println(l)

	ring.New()
}
