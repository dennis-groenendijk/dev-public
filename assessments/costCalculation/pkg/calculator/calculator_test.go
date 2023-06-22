package calculator

import (
	"log"
	"testing"

	"costCalculation/internal/path"
	"costCalculation/pkg/util"
	"github.com/stretchr/testify/assert" // requirement for assessment
)

func TestCalculate(t *testing.T) {
	var (
		header   []string
		values   []string
		expected [][]string
	)

	header = append(header, "metering_point_id", "total_price")
	values = append(values, "1", "0.02")
	expected = append(expected, header, values)

	file := path.Root + "/input_test.csv"
	data, _ := util.ReadDataFromFile(file)
	table := util.ConvDataToTable(data)

	result := Calculate(table)
	transposed := util.Transpose(result)

	assert.Equal(t, expected, transposed, "metering_point_id should be 1, total_price should be 0.02")
}

func TestCalculateUsage(t *testing.T) {
	var tests = []struct {
		name     string
		readings []float64
		expected []float64
	}{
		{
			"has negative reading",
			[]float64{166.606, 166.714, 166.694, 166.822}, // readings[2] is negative reading
			[]float64{0, 0.10800000000000409, 0.04399999999999693, 0.12800000000001432},
		}, // todo: add more tests
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateUsage(tt.readings)
			assert.Equal(t, tt.expected, got, "the negative reading should be skipped and replaced by linear consumption")
		})
	}
}

func TestCalculateCosts(t *testing.T) {
	readings := []float64{166.606, 166.694, 166.714, 166.822}
	expected := "0.02"

	file := path.Root + "/input_test.csv"
	data, _ := util.ReadDataFromFile(file)
	table := util.ConvDataToTable(data)

	var ids []string
	for _, v := range table {
		ids = append(ids, v.MeteringPointId)
	}

	usages := calculateUsage(readings)
	costs := calculateCosts(table, usages)
	_, totalCost := calculateTotalCosts(ids, costs)

	assert.Equal(t, expected, totalCost[1], "total cost is not what was expected")
}

func TestSaveResult(t *testing.T) {
	file := path.Root + "/expected.csv"
	expected, err := util.ReadDataFromFile(file)
	if err != nil {
		log.Fatal("file 'expected.csv' could not be read:\n", err)
	}

	file2 := path.Root + "/result.csv"
	result, err := util.ReadDataFromFile(file2)
	if err != nil {
		log.Fatal("file 'result.csv' could not be read:\n", err)
	}

	assert.Equal(t, expected, result, "result is not what was expected")
}
