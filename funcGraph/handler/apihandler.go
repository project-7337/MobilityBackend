package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graph/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	appleCountiresURL = "https://disease.sh/v3/covid-19/apple/countries"
)

// APIHandler ->
func APIHandler(country string, region string) models.RegionGraphData {
	graphData := fetchData(country, region)
	return graphData
}

func fetchData(country string, region string) models.RegionGraphData {
	url := appleCountiresURL + "/" + country + "/" + region
	url = strings.ReplaceAll(url, " ", "%20")
	var graphdata models.RegionGraphData
	urlresponse := makeRestCall(url)
	if urlresponse.StatusCode == 200 {
		data, _ := ioutil.ReadAll(urlresponse.Body)
		if err := json.Unmarshal(data, &graphdata); err != nil {
			fmt.Println(bytes.NewBuffer(data).String())
			log.Fatalf("FetchSubRegion Unmarshal Error for %s\n Err: %s\n", url, err)
		}
	}
	return graphdata
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
