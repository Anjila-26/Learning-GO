package expenses

import "fmt"

// ---------------- Record & DaysPeriod ----------------
type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

// ---------------- Filter ----------------
// Filter returns records that satisfy the predicate function.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

// ---------------- ByDaysPeriod ----------------
// Returns a predicate function for filtering records by day period.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ---------------- ByCategory ----------------
// Returns a predicate function for filtering records by category.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// ---------------- TotalByPeriod ----------------
// Returns total amount of expenses for records inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	for _, r := range in {
		if r.Day >= p.From && r.Day <= p.To {
			total += r.Amount
		}
	}
	return total
}

// ---------------- CategoryExpenses ----------------
// Returns total expenses for a category inside a period.
// Returns error if the category does not exist in the records at all.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	categoryExists := false
	total := 0.0

	for _, r := range in {
		if r.Category == c {
			categoryExists = true
			if r.Day >= p.From && r.Day <= p.To {
				total += r.Amount
			}
		}
	}

	if !categoryExists {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return total, nil
}
