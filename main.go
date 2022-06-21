package main

import (
	"fmt"
	"log"
	"resty-testy/pcn_reqs"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	stock, err := pcn_reqs.FetchStockData()
	if err != nil {
		panic("WTF IS GOING ON")
	}

	fmt.Println(stock.Msg)
}
