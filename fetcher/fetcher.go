package fetcher

import (
	"IMOOC/crawler_distributed/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching url: %s.\n", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Mac; Safari/604.1")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: status code %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
