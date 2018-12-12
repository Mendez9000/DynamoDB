package model

type Profile struct {
	GUI          string `json:"gui"`
	Vertical     string `json:"vertical"`
	Observations string `json:"observations"`
	Active       bool   `json:"active"`
	RawJsonData  string `json:"raw_json_data"`
}
