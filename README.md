# gsheets

A golang wrapper package for `golang.org/x/oauth2` and `google.golang.org/api/sheets/v4`.

**!!! Only for personal use !!!**

## Installation

```bash
go get github.com/takuoki/gsheets
```

## Usage

### Create New Client

```go
c, err := gsheets.New(ctx, `{"credentials": "json"}`, `{"token": "json"}`)
```

```go
client, err := gsheets.NewForCLI(ctx, "credentials.json")
```

### Get Sheet Information

Currently, only the read function is supported.

```go
func (*Client) GetSheetNames(ctx context.Context, spreadsheetID string) ([]string, error)
```

```go
func (*Client) GetSheet(ctx context.Context, spreadsheetID, sheetName string) (Sheet, error)
```

### Manipulate Sheet Values

If the index is out of range, `Value` method returns empty string.

```go
s, err := c.GetSheet(ctx, "spreadsheetID", "sheetName")
if err != nil {
  return err
}

fmt.Println(s.Value(row, clm))

for _, r := range s.Rows() {
  fmt.Println(r.Value(clm))
}
```
