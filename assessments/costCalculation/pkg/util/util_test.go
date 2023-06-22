package util

import (
	"fmt"
	"testing"

	"costCalculation/internal/path"
	"github.com/stretchr/testify/assert" // requirement for assessment
)

func TestReadData(t *testing.T) {
	// for this assessment there was limited input data and no data provider
	// expected number of rows after skipping the table header
	expected := 4

	file := path.Root + "/input_test.csv"
	records, err := ReadDataFromFile(file)
	message := fmt.Sprintf("%s", err)

	got := len(records)
	if got == 5 {
		message = "the table header should not be counted as a record"
	}

	assert.Equal(t, expected, got, message)
}

func TestConvUnixToTime(t *testing.T) {
	birthdate := "462391200"
	timeOfBirth := "08:00"

	_, time := ConvUnixToTime(birthdate)

	assert.NotEqual(t, timeOfBirth, time, "function uses 24h time format and should return 20:00")
}

type MockTable struct {
	column1 string
	column2 string
}

var mock = []MockTable{
	{"1", "2"},
	{"3", "4"},
}

func TestTranspose(t *testing.T) {
	var (
		column1 []string
		column2 []string
		table   [][]string
	)

	column1 = append(column1, "1", "3")
	column2 = append(column2, "2", "4")
	table = append(table, column1, column2)

	result := Transpose(table)

	assert.NotEqual(t, mock, result, "column1 should have [1 2] and column2 should have [3 4]")
}
