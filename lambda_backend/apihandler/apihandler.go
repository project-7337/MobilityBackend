package apihandler

import (
	"backend/config"
	"backend/services"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//APIHandler ->
func APIHandler(configs config.Conf) []string {
	countriesList := services.FetchCountries(configs)
	//var finalData models.FinalData
	//wg.Add(len(countriesList))
	for i, s := range countriesList {
		fmt.Println(i, s)
		//go services.FetchSubRegion(configs, s, finalData, &wg)
	}
	//wg.Wait()
	return countriesList
}
