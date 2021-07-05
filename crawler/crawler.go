package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ProductBasicJson struct {
	ID     uint     `json:"itemid"`
	Name   string   `json:"name"`
	Images []string `json:"images"`
}

type ProductResponseJson struct {
	Items []struct {
		ItemBasic ProductBasicJson `json:"item_basic"`
	} `json:"items"`
}

func main() {
	log.Printf("Begin Crawling...")

	shopeeUrl := "https://shopee.sg/api/v4/search/search_items?by=relevancy&order=desc&page_type=search&limit=10"

	resp, err := http.Get(shopeeUrl)

	if err != nil {
		log.Printf("error when http.Get %v", err)
		return
	}

	var respData ProductResponseJson
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		log.Printf("error when json parse %v", err)
		return
	}

	productJsons := make([]ProductBasicJson, 10)

	for _, item := range respData.Items {
		productJsons = append(productJsons, item.ItemBasic)
		//todo : print to insert sql statment for further usage
	}
}
