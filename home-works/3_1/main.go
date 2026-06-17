package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const baseUrl = "https://rest.coincap.io/v3"
const apiKey = "e0c10ccfb36b9c2cc85577afc85e691f0cce948ccb32cca21186fe897da4f249"

var targetNames = map[string]bool{
	"Bitcoin":     true,
	"Ethereum":    true,
	"Tether USDt": true,
	"BNB":         true,
	"USDC":        true,
}

type Asset struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

type APIResponse struct {
	Data      []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type OutputAsset struct {
	Name     string  `json:"name"`
	PriceUsd float64 `json:"priceUsd"`
}

func sendRequestAssets() (*http.Response, error) {
	action := baseUrl + "/assets"

	client := &http.Client{}
	req, err := http.NewRequest("GET", action, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return nil, err
	}

	return resp, nil
}

func formatResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка API: статус %d, тело: %s\n", resp.StatusCode, string(body))
		return nil, err
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return nil, err
	}

	var output []OutputAsset
	for _, asset := range apiResponse.Data {
		if targetNames[asset.Name] {
			price, err := strconv.ParseFloat(asset.PriceUsd, 64)
			if err != nil {
				fmt.Printf("Ошибка конвертации цены для %s: %v\n", asset.Name, err)
				continue
			}
			output = append(output, OutputAsset{
				Name:     asset.Name,
				PriceUsd: price,
			})
		}
	}

	outputJSON, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		fmt.Println("Ошибка создания JSON:", err)
		return nil, err
	}
	return outputJSON, nil
}

func main() {

	resp, err := sendRequestAssets()
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	outputJSON, err := formatResponse(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(outputJSON))
}
