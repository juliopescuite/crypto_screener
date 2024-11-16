package main

import (
	"fmt"
	"github.com/juliopescuite/crypto_screener/api_client/coingecko"
	"github.com/juliopescuite/crypto_screener/parser"
	"os"
)

func main() {
	// API key check
	if len(os.Args) == 0 {
		panic("missing the CoinGecko API key program argument")
	}

	config := parser.ParseScreenerConfig()
	data := coingecko.FetchCoinGeckoMarketData(os.Args[1])

	fmt.Printf("Screening with : MarketCap %f CurrentPrice: %f PriceChangePercentage24h: %f \n", config.MarketCap, config.CurrentPrice, config.PriceChangePercentage24h)

	for _, element := range data {

		// Filter by Market cap
		if element.MarketCap <= config.MarketCap {

			// Filter by Price
			if element.CurrentPrice <= config.CurrentPrice {

				// filter by price appreciation
				if element.PriceChangePercentage24h >= config.PriceChangePercentage24h {
					fmt.Printf("Found a new coin: id: %s name: %s symbol: %s \n", element.Id, element.Name, element.Symbol)
				}
			}
		}
	}
}
