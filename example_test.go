package gsheets_test

import (
	"context"
	"fmt"
	"os"

	"github.com/takuoki/gsheets"
)

func Example() {

	// https://docs.google.com/spreadsheets/d/1-kaYsJxISA2MJMKtTIzWwOFsMm9Wp21U3Nbqnfru65w
	spreadsheetID := "1-kaYsJxISA2MJMKtTIzWwOFsMm9Wp21U3Nbqnfru65w"

	ctx := gsheets.WithCache(context.Background())
	client, err := gsheets.New(ctx,
		os.Getenv("GOOGLE_API_CREDENTIALS"), os.Getenv("GOOGLE_API_TOKEN"))

	title, err := client.GetTitle(ctx, spreadsheetID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("title: %s\n", title)

	sheetNames, err := client.GetSheetNames(ctx, spreadsheetID)
	if err != nil {
		panic(err)
	}

	for _, sheetName := range sheetNames {
		sheet, err := client.GetSheet(ctx, spreadsheetID, sheetName)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sheetName: %s, A1: %s\n", sheetName, sheet.Value(0, 0))
	}

	// Output:
	// title: test-sheet
	// sheetName: foo, A1: This value is written in A1 of the foo sheet.
	// sheetName: bar, A1: This value is written in A1 of the bar sheet.
}
