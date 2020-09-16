package models

// RegionGraphData ->
type RegionGraphData struct {
	Country   string       `json:"country"`
	Subregion string       `json:"subregion"`
	Data      []RegionData `json:"data"`
}

// RegionData ->
type RegionData struct {
	Date    string  `json:"date"`
	Driving float64 `json:"driving"`
	Transit float64 `json:"transit"`
	Walking float64 `json:"walking"`
}
