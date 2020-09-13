package models

//SubregionData ->
type SubregionData struct {
	Country    string   `json:"country"`
	Subregions []string `json:"subregions"`
}

//FinalData ->
type FinalData struct {
	Region []string
}
