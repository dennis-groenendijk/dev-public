package util

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"costCalculation/pkg/models"
)

// ReadDataFromFile reads and formats the .csv data, then returns its records
func ReadDataFromFile(file string) ([][]string, error) {
	input, err1 := os.Open(file)
	if err1 != nil {
		return [][]string{}, err1
	}
	log.Printf("%s was read\n", file)
	defer input.Close()

	data := csv.NewReader(input)
	// in case the .csv file contains a different delimiter or comment(s)
	if data.Comma == ';' || data.Comma == '|' || data.Comma == ' ' {
		data.Comma = ','
	}
	data.Comment = '#'

	// exclude the header
	if _, err2 := data.Read(); err2 != nil {
		return [][]string{}, err2
	}

	records, err3 := data.ReadAll()
	if err3 != nil {
		return [][]string{}, err3
	}

	return records, nil
}

// ConvDataToTable converts the provided data to a table
func ConvDataToTable(data [][]string) (table []models.Table) {
	if reflect.DeepEqual(data, [][]string{}) {
		log.Println("no data received to convert to table")
		return []models.Table{}
	}
	for _, column := range data {
		row := models.Table{
			MeteringPointId: column[0],
			ReadingType:     column[1],
			ReadingValue:    column[2],
			CreatedAt:       column[3],
		}
		table = append(table, row)
	}

	return table
}

// ConvUnixToTime converts the provided timestamp from the Unix format to a human-readable time format
func ConvUnixToTime(input string) (day string, timestamp string) {
	if input == "" {
		log.Println("no input received for conversion")
		return "", ""
	}

	conv, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Println("input could not be converted to time format")
		return "", ""
	}

	t := time.Unix(conv, 0)
	day = t.Format("Mon")
	timestamp = t.Format("15:04")

	return day, timestamp
}

// SubtractTime calculates the duration in minutes between two time values
func SubtractTime(time1, time2 time.Time) float64 {
	return time2.Sub(time1).Minutes()
}

// Transpose rearranges the position of given values from rows to columns
func Transpose(slice [][]string) [][]string {
	if reflect.DeepEqual(slice, [][]string{}) {
		log.Println("no data received to transpose")
		return [][]string{}
	}

	xAxis := len(slice[0])
	yAxis := len(slice)
	result := make([][]string, xAxis)

	for i := range result {
		result[i] = make([]string, yAxis)
	}
	for i := 0; i < xAxis; i++ {
		for j := 0; j < yAxis; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}
