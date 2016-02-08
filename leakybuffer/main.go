package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var current_id int

type Buffer struct {
	id      int
	number  int
	message string
}

func (b *Buffer) get() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Int()
	fmt.Printf("Number assigned to buffer(%d) = %d\n", b.id, number)
	b.number = number
}

func (b *Buffer) process() error {
	switch {
	case b.message != "":
		fmt.Println("Buffer has already been processed. ")
		return errors.New("Previously Processed Buffer")
	default:
		b.message = fmt.Sprintf("[ID: %d] Internal number - %d", b.id, b.number)
		fmt.Println(b.message)
	}
	return nil
}

var available = make(chan Buffer, 10)
var toController = make(chan Buffer)

func worker() {
	for {
		var b Buffer
		select {
		case b = <-available: // If nothing is on the available channel, a new buffer is created
		default:
			b = Buffer{id: current_id}
			current_id++
		}
		b.get()
		fmt.Printf("Current Buffer: %v\n", b)
		toController <- b
	}
}

func controller() {
	for {
		b := <-toController
		fmt.Printf("Processing %v\n", b)
		if err := b.process(); err != nil {
			fmt.Println(err)
		} else {
			select {
			case available <- b:
			default: // Drop the buffer since nothing is available
			}
		}
	}
}

func main() {
	defer close(available)
	defer close(toController)
	go worker()
	go controller()
	time.Sleep(time.Millisecond * 2000)
}
