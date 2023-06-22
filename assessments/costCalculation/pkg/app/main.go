package main

import (
	"flag"
	"log"
	"reflect"

	"costCalculation/internal/path"
	"costCalculation/pkg/calculator"
	"costCalculation/pkg/models"
	"costCalculation/pkg/util"
)

// main initialises the program to calculate the total cost for every metering point in the provided data set
func main() {
	var filepath string
	flag.StringVar(&filepath, "filepath", "", "user specified input file")
	flag.Parse()

	file := path.Root + "/input.csv"
	if filepath != "" {
		file = filepath
	}

	data, err := util.ReadDataFromFile(file)
	if err != nil {
		log.Fatal("data could not be read: ", err)
	}

	table := util.ConvDataToTable(data)
	if !reflect.DeepEqual(table, []models.Table{}) {
		log.Println("resulting table parsed for cost calculation")
	}

	calculator.Calculate(table)
}
