package tests

import (
	"backend/config"
	"backend/services"
	"fmt"
	"testing"
)

func TestFetchCountries(t *testing.T) {
	var configs config.Conf
	configs.ReadConf()

	cl := services.FetchCountries(configs)
	for i, s := range cl {
		fmt.Println(i, s)
	}
}
