package concepts

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg = &sync.WaitGroup{}

func PingWebsites() {
	websites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"http://www.netflix.com",
		"http://www.amazon.com",
	}

	resp := []string{}

	wg.Add(50 * len(websites))

	for i := 0; i < 50; i++ {
		for _, website := range websites {
			go pingWebsite(website, &resp)
		}
	}

	wg.Wait()
	fmt.Println(len(resp))
}

func pingWebsite(website string, responses *[]string) {
	resp, err := http.Get(website)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Status Code %d | %s \n", resp.StatusCode, website)
		respBytes, _ := io.ReadAll(resp.Body)
		*responses = append((*responses), string(respBytes))
	}
	wg.Done()
}
