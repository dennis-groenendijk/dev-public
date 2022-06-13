package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// SimpleReader was written as a basic starting point for reading .csv data
func SimpleReader(input string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("this file could not be opened", err)
	}

	data := csv.NewReader(file)
	// in case the .csv file contains a different delimiter or comment(s)
	if data.Comma == ';' || data.Comma == '|' || data.Comma == ' ' {
		data.Comma = ','
	}
	data.Comment = '#'

	// all fields are read and printed to the cli
	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		for i := range record {
			fmt.Printf("%v\n", record[i])
		}
	}
	return
}
