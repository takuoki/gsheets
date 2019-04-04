package gsheets

import (
	"context"
	"fmt"
)

// GetTitle returns sheet title.
func (c *Client) GetTitle(ctx context.Context, spreadsheetID string) (string, error) {
	var err error
	resp, _ := getSpreadsheetsCache(ctx, spreadsheetID)
	if resp == nil {
		resp, err = c.srv.Spreadsheets.Get(spreadsheetID).Do()
		if err != nil {
			return "", fmt.Errorf("Unable to retrieve data from sheet: %v", err)
		}
		setCache(ctx, resp, srvNameSpreadsheets, spreadsheetID)
	}

	return resp.Properties.Title, nil
}

// GetSheetNames returns sheet name list.
func (c *Client) GetSheetNames(ctx context.Context, spreadsheetID string) ([]string, error) {
	var err error
	resp, _ := getSpreadsheetsCache(ctx, spreadsheetID)
	if resp == nil {
		resp, err = c.srv.Spreadsheets.Get(spreadsheetID).Do()
		if err != nil {
			return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
		}
		setCache(ctx, resp, srvNameSpreadsheets, spreadsheetID)
	}

	ss := []string{}
	for _, s := range resp.Sheets {
		ss = append(ss, s.Properties.Title)
	}

	return ss, nil
}

// GetSheet returns a Sheet.
func (c *Client) GetSheet(ctx context.Context, spreadsheetID, sheetName string) (*Sheet, error) {
	var err error
	resp, _ := getSpreadsheetsValuesCache(ctx, spreadsheetID, sheetName)
	if resp == nil {
		resp, err = c.srv.Spreadsheets.Values.Get(spreadsheetID, sheetName).Do()
		if err != nil {
			return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
		}
		setCache(ctx, resp, srvNameSpreadsheetsValues, spreadsheetID, sheetName)
	}

	return &Sheet{values: resp.Values}, nil
}
