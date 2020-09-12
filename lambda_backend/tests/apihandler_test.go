package tests

import (
	"backend/apihandler"
	"testing"
)

func TestFetchCountries(t *testing.T) {

	cl := apihandler.APIHandler()
	for i, s := range cl.Region {
		t.Log(i, s)
	}
}
