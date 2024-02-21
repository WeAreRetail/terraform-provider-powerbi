package models

type GroupUsers struct {
	ODataContext string      `json:"@odata.context"`
	ODataCount   int         `json:"@odata.count"`
	Value        []GroupUser `json:"value"`
}
