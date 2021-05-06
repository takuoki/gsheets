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

// WithCache returns a context with gsheets cache.
// If you want to use the cache, initialize the context.
func WithCache(ctx context.Context) context.Context {
	return context.WithValue(ctx, cacheKey, &gsCache{
		cache: map[string]interface{}{},
	})
}

func setCache(ctx context.Context, resp interface{}, srvName serviceName, keys ...string) error {
	v := ctx.Value(cacheKey)
	c, ok := v.(*gsCache)
	if !ok {
		return errors.New("the context has not been initialized for gsheets package")
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
		return nil, errors.New("the context has not been initialized for gsheets package")
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
