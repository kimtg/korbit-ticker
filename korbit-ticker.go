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

func getQuote(currencyPair string) (int64, string, error) {
	// {"timestamp":1499504423000,"last":"292000"}
	s, err := slurp("https://api.korbit.co.kr/v1/ticker?currency_pair=" + currencyPair)
	if err != nil {
		return 0, "", err
	}
	var m map[string]interface{}
	err = json.Unmarshal(s, &m)
	return int64(m["timestamp"].(float64)), m["last"].(string), nil
}

func main() {
	const interval = 5
	fmt.Println("Refreshes every", interval, "minutes.")
	for {
		fmt.Println(time.Now())
		for _, v := range []string{"btc_krw", "eth_krw", "etc_krw", "xrp_krw"} {
			if t, p, err := getQuote(v); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(v, ":", time.Unix(t / 1000, 0), p)
			}
		}
		time.Sleep(interval * time.Minute)
	}
}