## About The Project

This is a straightforward cryptocurrency screener designed to learn GoLang principles and gain some practical experience.

Also, why not try to discover the next crypto gem :-).

It fetches data from the [CoinGecko API](https://www.coingecko.com/en/api) and applies basic filtering rules.

Currently, it is quite simple, but it is a work in progress.

### Built With

* [GoLang](https://go.dev)
* [yaml.v3](https://github.com/go-yaml/yaml/tree/v3)

## Usage
1. Obtain an API key from the [CoinGecko API](https://www.coingecko.com/en/api)
2. Build the program using the command: `go build github.com/juliopescuite/crypto_screener`.
3. Run the program by passing the API key as the first argument: `go run github.com/juliopescuite/crypto_screener api_key`.
4. The program will output a list of coins that match the screening criteria defined in the `crypto_screener/config/screener_config.yaml` file.

## License
Distributed under the MIT License. See `LICENSE.txt` for more information.

## Contact
#### Julio Pescuite - [linktr.ee](https://linktr.ee/juliopescuite)
#### Project Link: [https://github.com/juliopescuite/crypto_screener](https://github.com/juliopescuite/crypto_screener)