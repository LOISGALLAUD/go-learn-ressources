package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

const (
	SITE = "https://scrapeme.live/shop/page"
	PRODUCT = "li.product"
	PRODUCT_URL = "a"
	PRODUCT_IMAGE = "img"
	PRODUCT_NAME = "h2"
	PRODUCT_PRICE = ".price"
	MAX_PAGE = 48
)

type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	log.Println("Program started.")

	// Initialize
	var products []PokemonProduct
	var wg sync.WaitGroup
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{"url", "image", "name", "price"})
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}

	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnHTML(PRODUCT, func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}
		pokemonProduct.url = e.ChildAttr(PRODUCT_URL, "href")
		pokemonProduct.image = e.ChildAttr(PRODUCT_IMAGE, "src")
		pokemonProduct.name = e.ChildText(PRODUCT_NAME)
		pokemonProduct.price = e.ChildText(PRODUCT_PRICE)
		products = append(products, pokemonProduct)
	})
	log.Println("Initialized successfully.")

	// Start scraping on the different pages
	wg.Add(MAX_PAGE)
	for i := 1; i <= MAX_PAGE; i++ {
		go func(i int) {
			defer wg.Done()
			c.Visit(fmt.Sprintf("%s/%d/", SITE, i))
		}(i)
	}
	log.Println("Scraping started successfully, please wait...")

	// Collect data from the sites
	wg.Wait()
	log.Println("Scraping completed successfully.")

	// Print the data
	for _, product := range products {
		err := writer.Write([]string{
		product.url,
		product.image,
		product.name,
		product.price})
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}

	log.Println("Program terminated.")
}
