package parser

import "testing"

func Test_parser_status(t *testing.T) {
	tests := []struct {
		name    string
		entry   string
		want    Status
		wantErr bool
	}{
		{
			"parse status [Available]",
			"available",
			Available,
			false,
		},
		{
			"parse status [Unknown]",
			"unknown",
			Unknown,
			false,
		},
		{
			"parse capitalised status [Charging]",
			"CHARGING",
			Charging,
			false,
		},
		{
			"parsing undefined status results in status [NULL]",
			"damaged",
			NULL,
			false,
		},
		{
			"parsing empty string results in status [NULL]",
			"",
			NULL,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result, ok := ParseStatus(tt.entry); ok {
				if result == tt.want {
					t.Logf("parsed string '%s' to enum %v as expected", tt.entry, result)
				}
			} else {
				t.Logf("could not parse string '%s'", tt.entry)
			}
		})
	}
}
