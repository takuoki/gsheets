package gsheets

import (
	"context"
	"errors"
	"strings"
	"sync"

	sheets "google.golang.org/api/sheets/v4"
)

type contextKey string

const cacheKey contextKey = "cache-key"

type gsCache struct {
	cache map[string]interface{}
	m     sync.RWMutex
}

// NewContext returns a new context of gsheets.
func NewContext() context.Context {
	return NewContextWithParent(context.Background())
}

// NewContextWithParent returns a new context of gsheets with parent context.
func NewContextWithParent(parent context.Context) context.Context {
	return context.WithValue(parent, cacheKey, &gsCache{
		cache: map[string]interface{}{},
	})
}

func setCache(ctx context.Context, resp interface{}, srvName serviceName, keys ...string) error {
	v := ctx.Value(cacheKey)
	c, ok := v.(*gsCache)
	if !ok {
		return errors.New("The context has not been initialized for gsheets package")
	}
	k := strings.Join(append([]string{string(srvName)}, keys...), "-")
	c.m.Lock()
	defer c.m.Unlock()
	c.cache[k] = resp
	return nil
}

func getCache(ctx context.Context, srvName serviceName, keys ...string) (interface{}, error) {
	v := ctx.Value(cacheKey)
	c, ok := v.(*gsCache)
	if !ok {
		return nil, errors.New("The context has not been initialized for gsheets package")
	}
	k := strings.Join(append([]string{string(srvName)}, keys...), "-")
	c.m.RLock()
	defer c.m.RUnlock()
	return c.cache[k], nil
}

type serviceName string

const (
	srvNameSpreadsheets       = "Spreadsheets"
	srvNameSpreadsheetsValues = "SpreadsheetsValues"
)

func getSpreadsheetsCache(ctx context.Context, spreadsheetID string) (*sheets.Spreadsheet, error) {
	v, err := getCache(ctx, srvNameSpreadsheets, spreadsheetID)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	s, _ := v.(*sheets.Spreadsheet)
	return s, nil
}

func getSpreadsheetsValuesCache(ctx context.Context, spreadsheetID, sheetName string) (*sheets.ValueRange, error) {
	v, err := getCache(ctx, srvNameSpreadsheetsValues, spreadsheetID)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	s, _ := v.(*sheets.ValueRange)
	return s, nil
}
