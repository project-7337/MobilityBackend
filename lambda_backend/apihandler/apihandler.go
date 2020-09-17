package apihandler

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/appmanch/go-commons/logging"
)

var logger = logging.GetLogger()

var finalData models.FinalData

const (
	appleCountiresURL = "https://disease.sh/v3/covid-19/apple/countries"
)

//APIHandler ->
func APIHandler(country string) models.SubregionData {
	regionData := fetchSubRegion(country)
	return regionData
}

func fetchSubRegion(country string) models.SubregionData {
	url := appleCountiresURL + "/" + country
	url = strings.ReplaceAll(url, " ", "%20")
	var subregionData models.SubregionData
	subregionResponse := makeRestCall(url)
	if subregionResponse.StatusCode == 200 {
		data, _ := ioutil.ReadAll(subregionResponse.Body)
		if err := json.Unmarshal(data, &subregionData); err != nil {
			logger.InfoF(bytes.NewBuffer(data).String())
			logger.ErrorF("FetchSubRegion Unmarshal Error for %s\n Err: %s\n", url, err)
		}
	}
	return subregionData
}

func makeRestCall(uri string) *http.Response {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		logger.ErrorF("MakeRestCall Error for %s\n Err: %s\n", uri, err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorF("Failed HTTP Request %s\n", err)
	}
	return resp
}
