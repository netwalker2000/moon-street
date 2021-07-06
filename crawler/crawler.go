package main

import (
	"encoding/json"
	"fmt"
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
	//real site url
	shopeeUrl := "https://shopee.sg/api/v4/search/search_items?by=relevancy&keyword=iphone&limit=60&newest=0&order=desc&page_type=search&scenario=PAGE_GLOBAL_SEARCH&version=2"

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

	//log.Printf("%v", productJsons)

	for _, product := range productJsons {
		fmt.Printf("-- %d\n", product.ID)
		fmt.Printf("insert into product_db.product_tab(id, name) values(%d, '%s');\n", product.ID, product.Name)
		for _, image := range product.Images {
			fmt.Printf("insert into product_db.photo_tab(url, product_id) values('%s', %d);\n", image, product.ID)
		}
	}
}
