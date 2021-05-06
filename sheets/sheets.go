// Package sheets is a package that defines a generic interface
// for working with data like google spreadsheets.
package sheets

// Sheet is a simple sheet interface.
type Sheet interface {
	Rows() []Row
	Value(row, clm int) string
}

// Row is a simple row interface.
type Row interface {
	Value(clm int) string
}
