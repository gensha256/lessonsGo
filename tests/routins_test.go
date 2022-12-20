package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

const (
	coinGeckoList  = "https://api.coingecko.com/api/v3/coins/list"
	coinGeckoPrice = "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=USD"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CoinPrice struct {
	Symbol string  `json:"symbol"`
	Usd    float64 `json:"usd"`
}

func TestRoutine(t *testing.T) {
	coins, err := FetchCoinList()
	if err != nil {
		t.Error(err)
		return
	}

	coins = coins[:20]

	stream := make(chan *CoinPrice)

	for _, value := range coins {
		go fetchPrice(value, stream)
	}

	prices := make([]*CoinPrice, 0)
	count := 0
	for p := range stream {
		count++
		prices = append(prices, p)
		if count >= len(coins) {
			close(stream)
		}
	}

	log.Println(prices)
}

func FetchCoinList() ([]Coin, error) {
	bodyBytes, err := call(coinGeckoList)
	if err != nil {
		log.Println("failed read body", err)
		return nil, err
	}

	coinList := make([]Coin, 0)

	err = json.Unmarshal(bodyBytes, &coinList)
	if err != nil {
		log.Println("failed unmarshal coin list", err)
		return nil, err
	}

	return coinList, nil
}

func FetchCoinPrice(coin Coin) (CoinPrice, error) {
	result := CoinPrice{
		Symbol: coin.Symbol,
	}

	url := fmt.Sprintf(coinGeckoPrice, coin.ID)
	bodyBytes, err := call(url)
	if err != nil {
		log.Println("failed read body", err)
		return result, err
	}

	priceResponse := make(map[string]CoinPrice, 0)

	err = json.Unmarshal(bodyBytes, &priceResponse)
	if err != nil {
		log.Println("failed unmarshal coin price", err)
		return result, err
	}

	result.Usd = priceResponse[coin.ID].Usd

	return result, nil
}

func fetchPrice(c Coin, stream chan *CoinPrice) {
	price, err := FetchCoinPrice(c)
	if err != nil {
		log.Println(err)
		stream <- nil
		return
	}

	stream <- &price
}

func call(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("failed fetch coin price", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("failed fetch coin list - bad status response", resp.StatusCode, resp.Status)
		return nil, errors.New("bad status")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed read body", err)
		return nil, err
	}

	return bodyBytes, nil
}
