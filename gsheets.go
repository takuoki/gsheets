package gsheets

import (
	"context"
	"fmt"
)

// GetTitle returns sheet title.
func (c *Client) GetTitle(ctx context.Context, spreadsheetID string) (string, error) {

	resp, err := c.srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return "", fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	return resp.Properties.Title, nil
}

// GetSheetNames returns sheet name list.
func (c *Client) GetSheetNames(ctx context.Context, spreadsheetID string) ([]string, error) {

	resp, err := c.srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	ss := []string{}
	for _, s := range resp.Sheets {
		ss = append(ss, s.Properties.Title)
	}

	return ss, nil
}

// GetSheet returns a Sheet.
func (c *Client) GetSheet(ctx context.Context, spreadsheetID, sheetName string) (*Sheet, error) {

	resp, err := c.srv.Spreadsheets.Values.Get(spreadsheetID, sheetName).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	return &Sheet{values: resp.Values}, nil
}
