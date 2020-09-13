package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	appleCountiresURL = "https://disease.sh/v3/covid-19/apple/countries"
)

//APIHandler ->
func APIHandler() []string {
	countriesList := fetchCountries()
	return countriesList
}

func fetchCountries() []string {
	countriesResponse := makeRestCall(appleCountiresURL)
	var countriesList []string
	if countriesResponse.StatusCode == 200 {
		countries, _ := ioutil.ReadAll(countriesResponse.Body)
		if err := json.Unmarshal(countries, &countriesList); err != nil {
			log.Fatal(err)
		}
	}
	return countriesList
}

func makeRestCall(uri string) *http.Response {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatalf("MakeRestCall Error for %s\n Err: %s\n", uri, err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed HTTP Request %s\n", err)
	}
	return resp
}
