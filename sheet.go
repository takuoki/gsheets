package gsheets

import (
	"fmt"
	"testing"
)

// Sheet is a sheet of google spreadsheets.
type Sheet struct {
	values [][]interface{}
}

// NewSheet returns a new sheet instance with argument values.
// This function is for testing.
func NewSheet(t *testing.T, values [][]interface{}) *Sheet {
	t.Helper()
	return &Sheet{values: values}
}

// Rows returns rows which the sheet has.
func (s *Sheet) Rows() []Row {
	if s == nil {
		return nil
	}
	rs := make([]Row, len(s.values))
	for i, r := range s.values {
		rs[i] = Row{values: r}
	}
	return rs
}

// Value returns the string value.
func (s *Sheet) Value(row, clm int) string {
	if s == nil || row < 0 || len(s.values) <= row {
		return ""
	}
	if clm < 0 || len(s.values[row]) <= clm {
		return ""
	}
	return value(s.values[row][clm])
}

// Row is a row of google spreadsheets.
type Row struct {
	values []interface{}
}

// Value returns the string value.
func (r *Row) Value(clm int) string {
	if r == nil || clm < 0 || len(r.values) <= clm {
		return ""
	}
	return value(r.values[clm])
}

func value(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
