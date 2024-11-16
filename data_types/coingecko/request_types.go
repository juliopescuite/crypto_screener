package coingecko

import "time"

type MarketDataRoi struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type MarketData struct {
	Id                           string        `json:"id"`                               // coin id
	Symbol                       string        `json:"symbol"`                           // coin symbol
	Name                         string        `json:"name"`                             // coin name
	Image                        string        `json:"image"`                            // coin image url
	CurrentPrice                 float64       `json:"current_price"`                    // coin current price in currency
	MarketCap                    float64       `json:"market_cap"`                       // coin market cap in currency
	MarketCapRank                float64       `json:"market_cap_rank"`                  // coin rank by market cap
	FullyDilutedValuation        float64       `json:"fully_diluted_valuation"`          // coin fully diluted valuation (fdv) in currency
	TotalVolume                  float64       `json:"total_volume"`                     // coin total trading volume in currency
	High24h                      float64       `json:"high_24h"`                         // coin 24h price high in currency
	Low24h                       float64       `json:"low_24h"`                          // coin 24h price low in currency
	PriceChange24h               float64       `json:"price_change_24h"`                 // coin 24h price change in currency
	PriceChangePercentage24h     float64       `json:"price_change_percentage_24h"`      // coin 24h price change in percentage
	MarketCapChange24h           float64       `json:"market_cap_change_24h"`            // coin 24h market cap change in currency
	MarketCapChangePercentage24h float64       `json:"market_cap_change_percentage_24h"` // coin 24h market cap change in percentage
	CirculatingSupply            float64       `json:"circulating_supply"`               // coin circulating supply
	TotalSupply                  float64       `json:"total_supply"`                     // coin total supply
	MaxSupply                    float64       `json:"max_supply"`                       // coin max supply
	Ath                          float64       `json:"ath"`                              // coin all-time high (ath) in currency
	AthChangePercentage          float64       `json:"ath_change_percentage"`            // coin all-time high (ath) change in percentage
	AthDate                      time.Time     `json:"ath_date"`                         // coin all-time high (ath) date
	Atl                          float64       `json:"atl"`                              // coin all-time low (atl) in currency
	AtlChangePercentage          float64       `json:"atl_change_percentage"`            // coin all-time low (atl) change in percentage
	AtlDate                      time.Time     `json:"atl_date"`                         // coin all-time low (atl) date
	Roi                          MarketDataRoi `json:"roi"`                              // roi
	LastUpdated                  time.Time     `json:"last_updated"`                     // coin last updated timestamp
	PriceChangePercentage1h      float64       `json:"price_change_percentage_1h"`       //  coin 1h price change in percentage
	SparklineIn7d                string        `json:"sparkline_in_7d"`                  // coin price sparkline in 7 days
	Price                        []float64     `json:"price"`                            // array of float64s
}
