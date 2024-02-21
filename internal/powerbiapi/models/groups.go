package models

type Groups struct {
	ODataContext string  `json:"@odata.context"`
	ODataCount   int     `json:"@odata.count"`
	Value        []Group `json:"value"`
}
