package models

// Groups represents a response detailing a list of Power BI groups (workspaces).
type Groups struct {
	ODataContext string  `json:"@odata.context"` // The OData context
	ODataCount   int     `json:"@odata.count"`   // The OData count
	Value        []Group `json:"value"`          // The list of groups
}
