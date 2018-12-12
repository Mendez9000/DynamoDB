package model

type Profile struct {
	GUID         string `json:"guid"`
	Vertical     string `json:"vertical"`
	Observations string `json:"observations"`
	Active       bool   `json:"active"`
	RawJsonData  string `json:"raw_json_data"`
}
