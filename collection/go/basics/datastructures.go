package main

import (
	"fmt"
)

// Data Structures (array, stack, queue, map, set)

func main() {
	fmt.Println("Array is an sequential collection of items")

	r := []int{2, 6, 4, 3, 45, 7, 8, 8, 76, 6}
	fmt.Println(r)
	for _, v := range r {
		fmt.Println(v)
	}

	fmt.Println("Stack is Last In First Out / First In Last Out")

	s := []int{}

	s = pushToStack(5, s) // [5]
	fmt.Println(s)
	s = pushToStack(10, s) // [5, 10]
	fmt.Println(s)
	s = pushToStack(3, s) // [5, 10, 3]
	fmt.Println(s)
	s = pushToStack(7, s) // [5, 10, 3, 7]
	fmt.Println(s)

	e := 0
	e, s = popFromStack(s)
	fmt.Println("popped ", e, "from stack, stack is now", s)
	e, s = popFromStack(s)
	fmt.Println("popped ", e, "from stack, stack is now", s)
	e, s = popFromStack(s)
	fmt.Println("popped ", e, "from stack, stack is now", s)
	e, s = popFromStack(s)
	fmt.Println("popped ", e, "from stack, stack is now", s)

	fmt.Println("Queue is First In First Out")

	// empty array again for reuse in queue fun
	s = []int{}

	s = pushToQueue(5, s) // [5]
	fmt.Println(s)
	s = pushToQueue(10, s) // [5, 10]
	fmt.Println(s)
	s = pushToQueue(3, s) // [5, 10, 3]
	fmt.Println(s)
	s = pushToQueue(7, s) // [5, 10, 3, 7]
	fmt.Println(s)

	e, s = popFromQueue(s)
	fmt.Println("popped ", e, "from Queue, Queue is now", s)
	e, s = popFromQueue(s)
	fmt.Println("popped ", e, "from Queue, Queue is now", s)
	e, s = popFromQueue(s)
	fmt.Println("popped ", e, "from Queue, Queue is now", s)
	e, s = popFromQueue(s)
	fmt.Println("popped ", e, "from Queue, Queue is now", s)

	fmt.Println("Map")
	m := map[string]int{"jemoeder": 2, "jevader": 2, "jezuster": 2, "jebroer": 2}
	fmt.Println(m)
	m["jetante"] = 21
	fmt.Println(m)
	m["jetante"] = 24
	fmt.Println(m)
	delete(m, "jetante")
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}

	fmt.Println("map key jevader has value", m["jevader"])
	v, ok := m["garble"]
	fmt.Println("map key garble has entry:", ok, "value retrieved was", v)
	v, ok = m["jevader"]
	fmt.Println("map key jevader has entry:", ok, "value retrieved was", v)

	fmt.Println("Set")
	mySet := map[string]bool{}
	mySet["key1"] = true
	mySet["key2"] = true
	mySet["key3"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	mySet["key1"] = true
	fmt.Println(mySet)

	numbers := []int{2, 5, 6, 3, 2, 2, 6, 75, 2, 35, 246, 246, 123, 3, 1, 5, 6, 3, 1, 3, 4}
	fmt.Println("numbers is", len(numbers), "long, with duplicates")
	fmt.Println(numbers)
	numbersSet := map[int]bool{}
	for _, v := range numbers {
		numbersSet[v] = true
	}
	fmt.Println("numbers as set:")
	fmt.Println(numbersSet)
	deduplicatedNumbers := []int{}
	for k, _ := range numbersSet {
		deduplicatedNumbers = append(deduplicatedNumbers, k)
	}
	fmt.Println("deduplicated:")
	fmt.Println(deduplicatedNumbers)


	fmt.Println("deduplication with sequence preservation")
	numbers2 := []int{2, 5, 6, 3, 2, 2, 6, 75, 2, 35, 246, 246, 123, 3, 1, 5, 6, 3, 1, 3, 4}
	numbers3 := []int{}
	alreadySeen := map[int]bool{}
	for _, v := range numbers2 {
		_, found := alreadySeen[v]
		if !found {
			alreadySeen[v] = true
			numbers3 = append(numbers3, v)
		}
	}
	fmt.Println(numbers2)
	fmt.Println(numbers3)

}

func pushToStack(elem int, stack []int) []int {
	return append(stack, elem)
}

func pushToQueue(elem int, queue []int) []int {
	return append(queue, elem)
}

func popFromStack(stack []int) (int, []int) {
	val := stack[len(stack)-1]
	newStack := stack[:len(stack)-1]
	return val, newStack
}

func popFromQueue(queue []int) (int, []int) {
	val := queue[0]
	newQueue := queue[1:len(queue)]
	return val, newQueue
}
