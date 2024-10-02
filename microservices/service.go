package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	fetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 100000.0,
	"ETH": 200000.0,
	"GG":  300000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given tiker (%s) is not supported.....", ticker)
	}
	return price, nil
}
