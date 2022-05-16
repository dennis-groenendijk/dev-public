/* Wordcounter

v1.0 - DONE
The goal is to build a program that can count the words provided by user input and store the result in a map.

v1.1 - DONE
Use the flag package to specify a particular file to be used as input.

v1.2 - DONE
Adjust the program to allow for input by flagging a local file or a POST request.
Gorilla/Mux should be used as a resource for setting up the server.

v1.3 - DONE
Instead of just being part of the program, the output should also be a response to the caller.

*/

package main

import (
	"bufio"
	"flag"
	. "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {

	// instructions
	Printf("Welcome! This program will count each word in a provided file and save the result as output.txt\n")
	Printf("You can either use the flag -filepath, followed by the file location to open and read that file")
	Printf(", or upload your file to localhost:8000 and receive the result as an immediate response\n\n")
}

func filepathHandler(filepath string) {
	input, errIn := os.Open(filepath)
	if errIn != nil {
		log.Fatalf("This file could not be opened: %v", errIn)
	}

	// an output file is created in the program's root folder, then the result of countWords is written to that file
	output, errOut := os.Create("output.txt")
	if errOut != nil {
		log.Fatalf("The output.txt file was not created: %v", errOut)
	}

	writer := bufio.NewWriter(output)

	defer func() {
		writer.Flush()
		output.Close()
	}()

	reader := bufio.NewReader(input)
	wordMap := countWords(reader)
	writeResult(wordMap, writer)

	Printf("[input received]\n")
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	writer := bufio.NewWriter(w)

	defer func() {
		writer.Flush()
	}()

	w.WriteHeader(http.StatusOK)                 // default, not necessary
	w.Header().Set("Content-Type", "text/plain") // added to be complete and proper

	reader := bufio.NewReader(r.Body)
	wordMap := countWords(reader)
	writeResult(wordMap, writer)
}

func main() {
	// flag -filepath
	var filepath string
	flag.StringVar(&filepath, "filepath", "", "user specified input file")
	flag.Parse()

	// the user input determines which handler is called
	if filepath != "" {
		filepathHandler(filepath)
	} else {
		// setting up a router and server for incoming traffic
		r := mux.NewRouter()

		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		// determines how to handle incoming POST requests on /
		//r.HandleFunc("/", httpHandler).Methods("POST") // TODO use the .Methods to restrict request types, for example "POST"
		r.HandleFunc("/", httpHandler)
		//r.HandleFunc("/endpoint1", httpHandlerAlt) // TODO use this to setup multiple handlers, when input needs to be handled in various ways

		// confirms the server has started and logs any errors
		Printf("Starting server at port 8000\n")
		log.Fatal(srv.ListenAndServe())
	}
}

func countWords(reader *bufio.Reader) map[string]int {
	// creates a map for indexing the words from the reader
	counter := make(map[string]int)

	for {
		// note: line is always just one line of text, hence the for loop to read all lines
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		// every line of the input is analysed and every tab (\t) is replaced for a space, to separate the words by
		text := string(line)
		text = strings.ReplaceAll(text, "\t", " ")
		words := strings.Split(text, " ")

		for _, word := range words {
			counter[word]++
		}
	}
	return counter
}

func writeResult(counter map[string]int, writer *bufio.Writer) {
	// the counter map is formatted and pushed to the writer
	result := Sprint(counter)

	_, err := writer.WriteString(result)
	if err != nil {
		log.Fatalf("The result could not be exported. Error: %s", err)
	}
}
