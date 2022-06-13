package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// convTimeToUnix converts the provided timestamp from a human-readable time format to a Unix format
// this function was made to check just one particular timestamp at a time using the CLI
func main() {
	fmt.Println("Please input one timestamp in the format 01-02-2006 03:04 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	layout := "02-01-2006 03:04"
	t, _ := time.Parse(layout, scanner.Text())
	fmt.Printf("Time: %v", t.Unix())
}
