package gsheets

func NewSheet(values [][]interface{}) *Sheet {
	return &Sheet{values: values}
}
