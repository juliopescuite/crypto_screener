package coingecko

import (
	"encoding/json"
	"github.com/juliopescuite/crypto_screener/data_types/coingecko"
	"io"
	"net/http"
)

func FetchCoinGeckoMarketData(apiKey string) []coingecko.MarketData {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=cad&order=market_cap_asc"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data []coingecko.MarketData
	var err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data
}
