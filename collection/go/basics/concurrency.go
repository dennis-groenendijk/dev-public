package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numberOfProducers := 12
	numberOfConsumers := 3
	numberOfIntegers := 100

	integer := make(chan int)
	var producer sync.WaitGroup
	producer.Add(numberOfProducers)
	var consumer sync.WaitGroup
	consumer.Add(numberOfConsumers)

	for i := 0; i < numberOfProducers; i++ {
		go fillChannel(integer, &producer, numberOfIntegers)
	}

	for i := 0; i < numberOfConsumers; i++ {
		go consumeChannel(integer, &consumer)
	}

	producer.Wait()
	close(integer)

	consumer.Wait()
	fmt.Println("Done!")
}

func fillChannel(input chan int, s *sync.WaitGroup, numberOfIntegers int) {
	for i := 0; i < numberOfIntegers; i++ {
		input <- rand.Int()
	}
	s.Done()
	close(input)
}

func consumeChannel(input chan int, s *sync.WaitGroup) {
	for i := range input {
		fmt.Println(i)
	}
	s.Done()
}
