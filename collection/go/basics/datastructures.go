package main

import (
	"fmt"
)

// Basics of Data Structures (Array, Stack, Queue, Map, Set)

func main() {
	fmt.Println("Array is an sequential collection of items")

	array := []int{2, 6, 4, 3, 45, 7, 8, 8, 76, 6}
	fmt.Println(array)
	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Println("Stack is Last In First Out / First In Last Out")

	stack := []int{}

	stack = pushToStack(5, stack) // [5]
	fmt.Println(stack)
	stack = pushToStack(10, stack) // [5, 10]
	fmt.Println(stack)
	stack = pushToStack(3, stack) // [5, 10, 3]
	fmt.Println(stack)
	stack = pushToStack(7, stack) // [5, 10, 3, 7]
	fmt.Println(stack)

	element := 0
	element, stack = popFromStack(stack)
	fmt.Println("popped ", element, "from stack, stack is now", stack)
	element, stack = popFromStack(stack)
	fmt.Println("popped ", element, "from stack, stack is now", stack)
	element, stack = popFromStack(stack)
	fmt.Println("popped ", element, "from stack, stack is now", stack)
	element, stack = popFromStack(stack)
	fmt.Println("popped ", element, "from stack, stack is now", stack)

	fmt.Println("Queue is First In First Out")

	// empty array again for reuse in queue example
	stack = []int{}

	stack = pushToQueue(5, stack) // [5]
	fmt.Println(stack)
	stack = pushToQueue(10, stack) // [5, 10]
	fmt.Println(stack)
	stack = pushToQueue(3, stack) // [5, 10, 3]
	fmt.Println(stack)
	stack = pushToQueue(7, stack) // [5, 10, 3, 7]
	fmt.Println(stack)

	element, stack = popFromQueue(stack)
	fmt.Println("popped ", element, "from Queue, Queue is now", stack)
	element, stack = popFromQueue(stack)
	fmt.Println("popped ", element, "from Queue, Queue is now", stack)
	element, stack = popFromQueue(stack)
	fmt.Println("popped ", element, "from Queue, Queue is now", stack)
	element, stack = popFromQueue(stack)
	fmt.Println("popped ", element, "from Queue, Queue is now", stack)

	fmt.Println("Map")
	m := map[string]int{"Batman": 2, "Flash": 2, "Superman": 2, "Aquaman": 2}
	fmt.Println(m)
	m["Flash"] = 21
	fmt.Println(m)
	m["Flash"] = 24
	fmt.Println(m)
	delete(m, "Aquaman")
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}

	fmt.Println("map key Batman has value", m["Batman"])
	v, ok := m["Cyborg"]
	fmt.Println("map key Cyborg has entry:", ok, "value retrieved was", v)
	v, ok = m["Batman"]
	fmt.Println("map key Batman has entry:", ok, "value retrieved was", v)

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
