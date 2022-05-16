package main

import (
    "fmt"
)

func init() {
    // instructions
    fmt.Printf("Hi, let's play a game of FizzBuzz\nPlease input a number between 1 - 100\n")
}

func main() {
    // user input
    var input int
    fmt.Scan(&input)

    // loop
    for i := 1; i < input; i++ {
        fizzbuzzUserInput(i)
    }
}

func fizzbuzzUserInput(i int) {
    // substitute integer for string
    fizz := "Fizz"
    buzz := "Buzz"

    if i%15 == 0 { // if the number is divisible by 15 (both 3 & 5), print "FizzBuzz" instead of the number
        fmt.Println(fizz + buzz)
    } else if i%5 == 0 { // if the number is divisible by 5, print "Buzz" instead of the number
        fmt.Println(buzz)
    } else if i%3 == 0 { // if the number is divisible by 3, print "Fizz" instead of the number
        fmt.Println(fizz)
    } else { // print the numbers
        fmt.Println(i)
    }
}
