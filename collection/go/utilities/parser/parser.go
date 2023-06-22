package parser

import "strings"

type Status int

const (
	NULL        Status = iota // NULL = 0 and the default setting
	Available                 // Available = 1
	Blocked                   // Blocked = 2
	Charging                  // Charging = 3
	Inoperative               // Inoperative = 4
	OutOfOrder                // OutOfOrder = 5
	Planned                   // Planned = 6
	Removed                   // Removed = 7
	Reserved                  // Reserved = 8
	Unknown                   // Unknown = 9
)

var statusMap = map[string]Status{
	"":            NULL,
	"available":   Available,
	"blocked":     Blocked,
	"charging":    Charging,
	"inoperative": Inoperative,
	"outoforder":  OutOfOrder,
	"planned":     Planned,
	"removed":     Removed,
	"reserved":    Reserved,
	"unknown":     Unknown,
}

// ParseStatus converts a set enum value to string
func ParseStatus(str string) (Status, bool) {
	s, ok := statusMap[strings.ToLower(str)]
	return s, ok
}
