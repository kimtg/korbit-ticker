package main

// get cryptocurrency prices from Korbit
// https://apidocs.korbit.co.kr/#public

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func slurp(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return body, nil
}

type quote struct {
	Timestamp int64
	Last string
}

func getQuote(currencyPair string) (quote, error) {
	// {"timestamp":1499504423000,"last":"292000"}
	s, err := slurp("https://api.korbit.co.kr/v1/ticker?currency_pair=" + currencyPair)
	var m quote
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(s, &m)
	return m, nil
}

func main() {
	const interval = 5
	fmt.Println("Refreshes every", interval, "minutes.")
	for {
		fmt.Println(time.Now())
		for _, currencyPair := range []string{"btc_krw", "eth_krw", "etc_krw", "xrp_krw"} {
			if q, err := getQuote(currencyPair); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(currencyPair, ":", time.Unix(q.Timestamp / 1000, 0), q.Last)
			}
		}
		time.Sleep(interval * time.Minute)
	}
}