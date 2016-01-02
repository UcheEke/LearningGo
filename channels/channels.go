package main

import (
	"fmt"
	"github.com/UcheEke/LearningGo/channels/subtopics"
	"sync"
)

var wgA sync.WaitGroup

func main() {
	fmt.Println("Channels")
	subtopics.Unbuffered(20, wgA)
	wgA.Wait()

	subtopics.BlockingWithRange(20)

	subtopics.N2oneRace(wgA)

	subtopics.SemaOne()
	subtopics.SemaTwo()

	subtopics.One2Many()
	subtopics.PassingChannels()
}
