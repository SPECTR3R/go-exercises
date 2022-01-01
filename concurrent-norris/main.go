package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ChuckNorris struct {
	Quote      string   `json:"value"`
	Url        string   `json:"url"`
	Id         string   `json:"id"`
	Categories []string `json:"categories"`
}

const baseUrl = "https://api.chucknorris.io/jokes/random"

func getQuote() (quote *ChuckNorris, err error) {
	response, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func getQuotesConcurrently(numOfQuotes int) {
	var quotesMap sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < numOfQuotes; i++ {
		wg.Add(1)
		go func(idx int) {
			quote, err := getQuote()
			if err != nil {
				panic(err)
			}
			quotesMap.Store(idx, quote)
			fmt.Printf("New Chuck Norris quote: %v\n", quote.Quote)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("All quotes received")
}

func printExecutionTime(t time.Time) {
	fmt.Println("Execution time: ", time.Since(t))
}

func main() {
	startTime := time.Now()
	defer printExecutionTime(startTime)
	getQuotesConcurrently(100)
}
