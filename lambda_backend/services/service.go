package services

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"backend/config"
)

//FetchCountries ->
func FetchCountries(configs config.Conf) []string {
	countries := makeRestCall(configs.AppleCountries)
	var countriesList []string
	if err := json.Unmarshal(countries, &countriesList); err != nil {
		log.Fatal(err)
	}
	return countriesList
}

//FetchSubRegion ->
func FetchSubRegion(configs config.Conf, country string, finalData models.FinalData, wg *sync.WaitGroup) {
	url := configs.AppleCountries + "/" + country
	subregion := makeRestCall(url)
	var subregionData models.SubregionData
	if err := json.Unmarshal(subregion, &subregionData); err != nil {
		log.Fatal(err)
	}
	for i, s := range subregionData.Subregions {
		finalData.Region[i] = s + " (" + country + ")"
	}
	defer wg.Done()
}

func makeRestCall(uri string) []byte {
	fmt.Println("Making Rest Call")
	response, err := http.Get(uri)
	if err != nil {
		fmt.Printf("Failed HTTP Request %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data
}
