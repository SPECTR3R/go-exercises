package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

func GetComic(comicID int) (comic Comic, err error) {
	url := fmt.Sprintf(baseXkcdURL, comicID)
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		log.Println(err)
	}

	return comic, nil
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start).Seconds())
	}()

	comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan Comic)

	for _, id := range comicsNeeded {
		go func(id int) {
			comic, err := GetComic(id)
			if err != nil {
				log.Println("error getting comic", err)
			}
			ch <- comic
		}(id)
	}

	var comics []Comic
	for i := 0; i < len(comicsNeeded); i++ {
		comm := <-ch
		comics = append(comics, comm)
		fmt.Println("Fetched comic", comicsNeeded[i], "of title:", comics[i].Title)
	}
}
