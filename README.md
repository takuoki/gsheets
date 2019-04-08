# gsheets

[![CircleCI](https://circleci.com/gh/takuoki/gsheets/tree/master.svg?style=shield&circle-token=30f6e95108024e7a0562f630c69209783e5086ec)](https://circleci.com/gh/takuoki/gsheets/tree/master)
[![codecov](https://codecov.io/gh/takuoki/gsheets/branch/master/graph/badge.svg)](https://codecov.io/gh/takuoki/gsheets)
[![GoDoc](https://godoc.org/github.com/takuoki/gsheets?status.svg)](https://godoc.org/github.com/takuoki/gsheets)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A golang wrapper package for `golang.org/x/oauth2` and `google.golang.org/api/sheets/v4`.
You can easily manipulate spreadsheets.

**!!! Only for personal use !!!**

## Installation

```bash
go get github.com/takuoki/gsheets
```

## Requirement

This package uses Google OAuth2.0. So before executing tool, you have to prepare credentials.json.
See [Go Quickstart](https://developers.google.com/sheets/api/quickstart/go), or [Blog (Japanese)](https://medium.com/veltra-engineering/how-to-use-google-sheets-api-with-golang-9e50ee9e0abc) for the details.

## Usage

### Create Cache

If you want to use the cache, initialize the context.

```go
ctx := gsheets.WithCache(ctx)
```

### Create New Client

```go
client, err := gsheets.New(ctx, `{"credentials": "json"}`, `{"token": "json"}`)
```

```go
client, err := gsheets.NewForCLI(ctx, "credentials.json")
```

### Get Sheet Information

Currently, only the read function is supported.

```go
func (*Client) GetTitle(ctx context.Context, spreadsheetID string) (string, error)
```

```go
func (*Client) GetSheetNames(ctx context.Context, spreadsheetID string) ([]string, error)
```

```go
func (*Client) GetSheet(ctx context.Context, spreadsheetID, sheetName string) (Sheet, error)
```

### Manipulate Sheet Values

If the index is out of range, `Value` method returns empty string.

```go
s, err := client.GetSheet(ctx, "spreadsheetID", "sheetName")
if err != nil {
  return err
}

fmt.Println(s.Value(row, clm))

for _, r := range s.Rows() {
  fmt.Println(r.Value(clm))
}
```
