package unmarshaller

import (
	"encoding/json"
	"fmt"
)

const js = `{
	"name": "",
	"country": "",
	"opening_times": {
		"twentyfourseven": false,
 		"regular_hours": [
			{
 				"weekday": 1,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			}, 
			{
 				"weekday": 2,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			},
			{
 				"weekday": 3,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			},
			{
 				"weekday": 4,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			},
			{
 				"weekday": 5,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			},
			{
 				"weekday": 6,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			},
			{
 				"weekday": 7,
 				"period_begin": "08:00",
 				"period_end": "20:00"
 			}],
		"exceptional_openings": [
			{
				"period_begin": "2014-06-21T09:00:00Z",
				"period_end": "2014-06-21T12:00:00Z"
			}],
		"exceptional_closings": [
			{
				"period_begin": "2014-06-24T00:00:00Z",
				"period_end": "2014-06-25T00:00:00Z"
			}]
	}
}`

type Location struct {
	Name         string       `json:"name,omitempty"`
	Country      string       `json:"country,omitempty"`
	OpeningTimes OpeningTimes `json:"opening_times"`
}

type OpeningTimes struct {
	TwentyForSeven bool `json:"twentyforseven,omitempty"`
	RegularHours   []struct {
		Weekday     int    `json:"weekday"`
		PeriodBegin string `json:"period_begin"`
		PeriodEnd   string `json:"period_end"`
	} `json:"regular_hours,omitempty"`
	ExceptionalOpenings []struct {
		PeriodBegin string `json:"period_begin"`
		PeriodEnd   string `json:"period_end"`
	} `json:"exceptional_openings,omitempty"`
	ExceptionalClosings []struct {
		PeriodBegin string `json:"period_begin"`
		PeriodEnd   string `json:"period_end"`
	} `json:"exceptional_closings,omitempty"`
	Valid bool `json:"-"`
}

func unmarshalOT(str string) Location {
	loc := Location{}
	if err := json.Unmarshal([]byte(js), &loc); err != nil {
		fmt.Println("error unmarshalling json: ", err)
	}
	fmt.Printf("%+v", loc) // prints readable one-line result
	// fmt.Printf("%#v", ot) // prints detailed result (type info)
	return loc
}
