package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

    now := time.Now()
    ch := make(chan int)
	defer close(ch)

    for j := 0; j < 7; j++ {
        wg.Add(1)
        go func() {
            time.Sleep(time.Second * 2)
            i := <-ch
            fmt.Println(i)
            wg.Done()
        }()
		fmt.Println("Récépteur", j, "prêt")
    }

    for j := 0; j < 7; j++ {
        wg.Add(1)
        go func() {
            time.Sleep(time.Second * 2)
            ch <- 50

            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println(time.Since(now))
}