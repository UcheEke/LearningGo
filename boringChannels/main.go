package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(name string) <-chan string {
	ch := make(chan string)
	boringWords := []string{"blah ", "wibble ", "blee-blah ", "hmmm... ","erm... ", "zoiing ", "ah..", "wobble "}

	go func() {
		for i:=0; ;i++ {
			var msg string
			rand.Seed(int64(time.Now().Nanosecond()))
			for i:=0; i < 3; i++ {
				msg += boringWords[rand.Intn(len(boringWords))]
			}
			ch <- fmt.Sprintf("%s: %s",name, msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)
	for _, ch := range channels {
		go func(channel <-chan string){
			out<- <-channel
		}(ch)
	}
	return out
}

func main(){
	for i := 0; i < 20; i++ {
		fmt.Printf("%s\n",<-fanIn(boring("Ann"),boring("Jon")))
		time.Sleep(time.Millisecond * 30)
	}
	fmt.Println("Wow, people! You're reeeeally boring! I'm leaving...")
}
