package model

type CountryInfo struct {
	Count   int      `json:"count"`
	Name    string   `json:"name"`
	Country []County `json:"country"`
}

type County struct {
	CountryId   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
