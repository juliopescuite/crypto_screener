package data_types

type ScreeningConfig struct {
	CurrentPrice             float64 `yaml:"current_price"`
	MarketCap                float64 `yaml:"market_cap"`
	PriceChangePercentage24h float64 `yaml:"price_change_percentage_24_h"`
}
