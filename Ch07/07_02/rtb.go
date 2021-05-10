package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Bid struct {
	Price float64
	URL   string
}

var (
	defaultBid = Bid{
		Price: 0.02,
		URL:   "https://j.mp/3cbDsIY",
	}
	bidTimeout = 10 * time.Millisecond
)

func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		log.Printf("bid for %q timed out, returning default", url)
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), bidTimeout)
	defer cancel()
	bid := bidOn(ctx, "https://353solutions.com")
	fmt.Println(bid)
	// {0.035 https://j.mp/3f3Dpkb}

	ctx, cancel = context.WithTimeout(context.Background(), bidTimeout)
	defer cancel()
	bid = bidOn(ctx, "https://example.com")
	fmt.Println(bid)
	// 2021/05/02 19:10:00 bid for "https://example.com" timed out, returning default
	// {0.02 https://j.mp/3cbDsIY}
}
