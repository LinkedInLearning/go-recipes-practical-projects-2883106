package main

import (
	"fmt"
	"log"
	"net/http"
)

func authRequest(url, user, passwd string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(user, passwd)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	return nil
}

func main() {
	user, passwd := "joe", "baz00ka"
	url := fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", user, passwd)

	if err := authRequest(url, user, passwd); err != nil {
		log.Fatal(err)
	}
	fmt.Println("OK")
}
