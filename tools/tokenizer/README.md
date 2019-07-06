# tokenizer

This tool generates a token for Google Sheets API.

## Installation

```bash
go get github.com/takuoki/gsheets/tools/tokenizer
```

## Requirement

This tool uses Google OAuth2.0. So before executing tool, you have to prepare credentials.json.
See [Go Quickstart](https://developers.google.com/sheets/api/quickstart/go), or [Blog (Japanese)](https://medium.com/veltra-engineering/how-to-use-google-sheets-api-with-golang-9e50ee9e0abc) for the details.

## Usage

### Create readonly token

```bash
$ tokenizer gen
Go to the following link in your browser then type the authorization code:
https://accounts.google.com/o/oauth2/auth?...
(input your authorization code here)
Saving credential file to: token.json
complete
```

### Create writable token

```bash
$ tokenizer gen -w
...
complete
```
