package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step 1 :", url)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Step 2 :", url)
	defer response.Body.Close()
	
	fmt.Println("Step3: ", url)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second) // @NOTE: This is a hack, don't do this in production
}