package calculator

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"costCalculation/internal/path"
	"costCalculation/pkg/models"
	"costCalculation/pkg/util"
)

// Calculate calculates the usage and costs of electricity/gas, per metering_point_id, based on the provided table
func Calculate(table []models.Table) [][]string {
	var (
		ids             []string
		timestamps      []string
		readings        []float64
		calculatedCosts [][]string
	)

	for _, v := range table {
		ids = append(ids, v.MeteringPointId)
		timestamps = append(timestamps, v.CreatedAt)

		value, _ := strconv.ParseFloat(v.ReadingValue, 64)
		if v.ReadingType == "1" {
			value = value / 1000.00
		}
		if v.ReadingType == "2" {
			value = value * 9.769
		}
		readings = append(readings, value)
	}

	// first filter bad readings by validating the correct timestamp interval
	// then calculate the total costs per metering_point_id by usage
	validReadings := validateTimestampInterval(timestamps, readings)
	usages := calculateUsage(validReadings)
	costs := calculateCosts(table, usages)
	mpIds, totalCost := calculateTotalCosts(ids, costs)

	calculatedCosts = append(calculatedCosts, mpIds, totalCost)
	saveResult(calculatedCosts)

	return calculatedCosts // only returns for tests
}

// validateTimestampInterval filters out the readings that don't have the correct timestamp interval
func validateTimestampInterval(timestamps []string, readings []float64) []float64 {
	// one of the requirements for this assessment was to drop any bad reading
	for i := 0; i < len(timestamps)-1; i++ {
		_, thisTime := util.ConvUnixToTime(timestamps[i])
		_, nextTime := util.ConvUnixToTime(timestamps[i+1])
		this, _ := time.Parse("15:04", thisTime)
		next, _ := time.Parse("15:04", nextTime)

		difference := util.SubtractTime(this, next)
		if difference != 15 {
			// trigger to more easily spot the error while testing, check log when needed
			readings = append(readings, -1)
			log.Fatalf("this timestamp interval is incorrect: %s - %s", thisTime, nextTime)
		}
	}

	return readings
}

// calculateUsage calculates the usage between each (validated) reading
func calculateUsage(readings []float64) []float64 {
	usages := make([]float64, 0)

	// add initial line for zero usage
	usages = append(usages, 0)

	for i := 0; i < len(readings)-1; i++ {
		usage := readings[i+1] - readings[i]
		// if reading is incorrect, replace it by a linear consumed interpolation (requirement for assessment)
		if usage < 0 || usage > 100 {
			usage = (readings[i+1] - readings[i-1]) / 2
		}
		usages = append(usages, usage)
	}

	return usages
}

// calculateCosts calculates the cost of usage based on a set of time related conditions
func calculateCosts(table []models.Table, usages []float64) []float64 {
	costs := make([]float64, 0)

	for i, v := range table {
		// set timeframe for higher electricity rate on a weekday
		start := "07:00"
		end := "23:00"

		// check for day and time by converting the Unix timestamp
		day, timestamp := util.ConvUnixToTime(v.CreatedAt)
		switch day {
		case "Mon", "Tue", "Wed", "Thu", "Fri":
			day = "Weekday"
		case "Sat", "Sun":
			day = "Weekend"
		}

		cost := usages[i] * 0.20
		if v.ReadingType == "1" && timestamp < start && timestamp > end && day == "Weekend" {
			cost = usages[i] * 0.18
		}
		if v.ReadingType == "2" {
			cost = usages[i] * 0.06
		}
		costs = append(costs, cost)
	}

	return costs
}

// calculateTotalCosts sums the calculated costs for each metering_point_id and returns all values for the result table
func calculateTotalCosts(ids []string, costs []float64) ([]string, []string) {
	// todo: check (and test) if sorting of mpIds is needed, when larger data set is available
	sumCosts := 0.00
	mpIds := make([]string, 0)
	totalCosts := make([]string, 0)

	// add table header values
	mpIds = append(mpIds, "metering_point_id")
	totalCosts = append(totalCosts, "total_price")

	ids = append(ids, "add extra element to prevent 'index out of range' error when checking ids[i+1]")
	for i := 0; i < len(ids)-1; i++ {
		// while the id is the same, sum costs
		if ids[i+1] == ids[i] {
			sumCosts += costs[i]
		}
		// when the id is no longer the same, sum costs for the next metering point
		if ids[i+1] != ids[i] {
			mpIds = append(mpIds, ids[i])
			totalCosts = append(totalCosts, fmt.Sprintf("%.2f", sumCosts))
			sumCosts = 0.00
		}
	}

	return mpIds, totalCosts
}

// saveResult creates an output file, transposes the calculated data and writes the result to this file
func saveResult(calculatedCosts [][]string) {
	file := path.Root + "/result.csv"
	output, err := os.Create(file)
	if err != nil {
		log.Fatal("failed to create and open file: ", err)
	}
	defer output.Close()
	writer := csv.NewWriter(output)

	// rearranges the position of given values, from rows to columns
	transposed := util.Transpose(calculatedCosts)

	err = writer.WriteAll(transposed) // calls Flush internally
	if err != nil {
		log.Fatal("the result could not be written to file", err)
	}
}
