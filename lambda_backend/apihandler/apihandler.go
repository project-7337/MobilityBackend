package apihandler

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var finalData models.FinalData

const (
	appleCountiresURL = "https://disease.sh/v3/covid-19/apple/countries"
)

//APIHandler ->
func APIHandler() models.FinalData {
	countriesList := fetchCountries()
	finalData.Region = make([]string, 0, 10)
	wg.Add(len(countriesList))
	for _, s := range countriesList {
		go fetchSubRegion(s, &wg)
	}
	wg.Wait()
	return finalData
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

func fetchSubRegion(country string, wg *sync.WaitGroup) {
	url := appleCountiresURL + "/" + country
	url = strings.ReplaceAll(url, " ", "%20")
	var subregionData models.SubregionData
	subregionResponse := makeRestCall(url)
	if subregionResponse.StatusCode == 200 {
		data, _ := ioutil.ReadAll(subregionResponse.Body)
		if err := json.Unmarshal(data, &subregionData); err != nil {
			fmt.Println(bytes.NewBuffer(data).String())
			log.Fatalf("FetchSubRegion Unmarshal Error for %s\n Err: %s\n", url, err)
		}
		for _, s := range subregionData.Subregions {
			finalData.Region = append(finalData.Region, s+" ("+country+")")
		}
	}
	defer wg.Done()
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
