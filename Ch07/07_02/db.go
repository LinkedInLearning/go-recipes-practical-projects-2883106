// Dummy bid model
package main

import "time"

var state = 0

func bestBid(url string) Bid {
	state = 1 - state // toggle state
	if state == 1 {
		time.Sleep(2 * time.Millisecond)
		return Bid{
			Price: 0.035,
			URL:   "https://j.mp/3f3Dpkb",
		}
	}

	time.Sleep(bidTimeout + 20*time.Millisecond)
	return Bid{
		Price: 0.018,
		URL:   "https://j.mp/39oEJe7",
	}

}
