package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// convUnixToTime converts the provided timestamp from a Unix format to a human-readable time format
// this function was made to check just one particular timestamp at a time using the CLI
func main() {
	fmt.Println("Please input one Unix formatted timestamp: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	timeConverted := time.Unix(input, 0)
	fmt.Printf("time.Time: %v", timeConverted)
}
