package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{
    	"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,  
    }
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    if u, ok := units[unit]; !ok {
        return false
    } else {
    	bill[item] += u
    	return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    b, okBill := bill[item]
    u, okUnits := units[unit]

    if !okBill || !okUnits || u > b {
        return false
    } else if u == b {
    	delete(bill, item)
    } else {
    	bill[item] -= u
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    u, ok := bill[item]
    return u, ok
}
