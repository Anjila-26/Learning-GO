package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {

    needs := map[string]int{
        "quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
    }
    return needs
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bills := make(map[string]int)
    return bills
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]
	if !ok {
		return false
	}

	bill[item] += quantity
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// 1. Check if item exists in bill
	current_quantity, item_ok := bill[item]
	if !item_ok {
		return false
	}

	// 2. Check if unit exists
	unit_quantity, unit_ok := units[unit]
	if !unit_ok {
		return false
	}

	new_quantity := current_quantity - unit_quantity

	// 3. Quantity cannot go below zero
	if new_quantity < 0 {
		return false
	}

	// 4. Remove item completely if quantity becomes zero
	if new_quantity == 0 {
		delete(bill, item)
		return true
	}

	// 5. Otherwise update quantity
	bill[item] = new_quantity
	return true
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    quantity, ok := bill[item]

    if !ok {
        return 0, false
    }

    return quantity, true
}
