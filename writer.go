package gsheets

import (
	"context"
	"fmt"

	sheets "google.golang.org/api/sheets/v4"
)

// Update updates a single row.
// `rowNo` is started from zero.
func (c *Client) Update(ctx context.Context, spreadsheetID, sheetName string, rowNo int, values []interface{}) error {

	req := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
		Data: []*sheets.ValueRange{
			{
				MajorDimension: "ROWS",
				Range:          fmt.Sprintf("%[1]s!%[2]d:%[2]d", sheetName, rowNo+1),
				Values:         [][]interface{}{values},
			},
		},
	}

	_, err := c.srv.Spreadsheets.Values.BatchUpdate(spreadsheetID, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("unable to update sheet: %v", err)
	}

	return nil
}

// UpdateValue is data structure for BatchUpdate method.
// `RowNo` is started from zero.
type UpdateValue struct {
	SheetName string
	RowNo     int
	Values    []interface{}
}

// BatchUpdate updates multiple rows.
func (c *Client) BatchUpdate(ctx context.Context, spreadsheetID string, updateValues ...UpdateValue) error {

	data := []*sheets.ValueRange{}
	for _, updateValue := range updateValues {
		data = append(data, &sheets.ValueRange{
			MajorDimension: "ROWS",
			Range:          fmt.Sprintf("%[1]s!%[2]d:%[2]d", updateValue.SheetName, updateValue.RowNo+1),
			Values:         [][]interface{}{updateValue.Values},
		})
	}

	req := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
		Data:             data,
	}

	_, err := c.srv.Spreadsheets.Values.BatchUpdate(spreadsheetID, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("unable to update sheet: %v", err)
	}

	return nil
}
