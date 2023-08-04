package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	err := OpenDatabase()
	if err != nil {
		log.Printf("error opening database connection %v", err)
	}
	var wg sync.WaitGroup

	wg.Add(3)
	go getAll(&wg)
	go generateNumbers(10000, &wg)
	go printNumbers(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

func getAll(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Beginning Query")
	rows, err := DB.Query("SELECT * FROM album;")
	if err != nil {
		log.Printf("error querying books table %v", err)
		return
	}

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			log.Printf("scanning books table rows into struct Book %v", err)
			return
		}
		albums = append(albums, album)
		fmt.Println(album.Artist)
	}

	j, err := json.Marshal(albums)
	if err != nil {
		log.Printf("error marshalling books into json %v", err)
		return
	}

	fmt.Printf("json data: %s\n", j)
}

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= 150; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}
