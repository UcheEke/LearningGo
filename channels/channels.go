package main

import (
	"fmt"
	_"github.com/UcheEke/LearningGo/channels/subtopics"
	_"sync"
	"github.com/UcheEke/LearningGo/channels/challenges"
)

//var wgA sync.WaitGroup

func main() {
	fmt.Println("Channels")
	/*
	subtopics.Unbuffered(20, wgA)
	wgA.Wait()

	subtopics.BlockingWithRange(20)

	subtopics.N2oneRace(wgA)

	subtopics.SemaOne()
	subtopics.SemaTwo()

	subtopics.One2Many()
	subtopics.PassingChannels()

	subtopics.DualAccumulator(10)
	*/

	//fmt.Printf("10! = %d\n",subtopics.Factorial(10))  // 120

	challenges.MultiFact()

}
