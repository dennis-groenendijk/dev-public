## Cost Calculation

### Assessment
by Dennis Groenendijk
```
```
### Goal

This program was written for making a cost calculation, based on the provided data table, for the usage of gas and/or electricity consumption.

The input data should consist of these 4 columns:
* **metering_point_id**
  
  The id of a metering point
* **type**

  The type of reading; '1' for electricity or '2' for gas
* **reading**

  The raw value from the metering point. The readings for electricity should be provided as Wh and the readings for gas should be provided as m3, with a 15 minutes interval
* **created_at**
  
  Contains a Unix timestamp for when the measurement was done

The program filters out incorrect reading values and interpolates these errors, based on a linear consumption. It then converts the reading values to kWh for calculation.

The costs are calculated differently for the usage of electricity and gas. A higher rate is used for electricity consumption during weekdays, between 07:00 and 23:00. 

The program outputs the file **result.csv** in the programs root directory. The file presents a new data table, consisting of 2 columns:
* **metering_point_id** 

  The id of the metering point
* **total_price**

  The total cost of consumption for this metering point

### Directory Index
```
.
└── costCalculation
    ├── README.md
    ├── breakdown.txt
    ├── input.csv
    ├── input_test.csv
    ├── expected.csv
    ├── result.csv
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── path
    │   │   └── root.go
    └── pkg
        ├── app
        │   └── main.go
        ├── calculator
        │   ├── calculator.go
        │   └── calculator_test.go
        ├── models
        │   └── table.go
        └── util
            ├── util.go
            └── util_test.go
```

### Architecture

To provide alternative input for this program, there are 2 options:
1. Before running the program, override the **input.csv** file located in the root directory.
2. Use the **-filepath** flag to redirect to another file location.

**root.go** Root returns the relative root dir of this project during runtime

**main.go** Initialises the program to calculate the total cost for every metering point in the provided data set.

**calculator.go** has these functions (made available for testing):
* **Calculate**

  Calculates the usage and costs of electricity/gas, per metering_point_id, based on the provided table.

* **ValidateTimestampInterval**

  Filters out the readings that don't have the correct timestamp interval.

* **CalculateUsage**

  Calculates the usage between each (validated) reading.

* **CalculateCosts**

  Calculates the cost of usage based on a set of time related conditions.

* **CalculateTotalCosts**

  Sums the calculated costs for each metering_point_id and returns all values for the result table.

* **SaveResult**

  Creates an output file, transposes the calculated data and writes the result to this file.

**util.go** has these utility functions, which can be called from any other function:
* **ReadDataFromFile**

  Reads and formats the .csv data, then returns its records.

* **ConvDataToTable**

  Converts the provided data to a table

* **ConvUnixToTime**

  Converts the provided timestamp from the Unix format to a human-readable time format.

* **SubtractTime**

  Calculates the duration in minutes between two time values.

* **Transpose**

  Rearranges the position of given values from rows to columns.