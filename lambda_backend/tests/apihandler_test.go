package tests

import (
	"backend/apihandler"
	"testing"
)

func TestFetchCountries(t *testing.T) {

	fd := apihandler.APIHandler("india")
	t.Log(fd)
}
