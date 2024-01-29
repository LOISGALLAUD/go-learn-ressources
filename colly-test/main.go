package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	url, image, name, price string
}


func main() {
	log.Println("Program started.")

	// Initialize
	var products []PokemonProduct
	var wg sync.WaitGroup
	const site string = "https://scrapeme.live/shop/"

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")
		products = append(products, pokemonProduct)
	})
	
	c.OnScraped(func(r *colly.Response) {
		log.Println("Finished", r.Request.URL)
	})

	log.Println("Initialized successfully.")


	// Start scraping on the sites
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Visit(site)
	}()
	log.Println("Scraping started successfully, please wait...")

	// Collect data from the sites
	wg.Wait()
	log.Println("Scraping completed successfully.")

	// Print the data
	for _, product := range products {
		fmt.Printf("URL: %s\nImage: %s\nName: %s\nPrice: %s\n\n",
		 product.url, product.image, product.name, product.price)
	}

	log.Println("Program terminated.")
}
