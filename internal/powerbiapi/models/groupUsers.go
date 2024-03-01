package models

// The OData response wrapper for a list of Power BI users with access to a workspace.
type GroupUsers struct {
	ODataContext string      `json:"@odata.context"` //The OData context
	ODataCount   int         `json:"@odata.count"`   //The OData count
	Value        []GroupUser `json:"value"`          //The list of users with access to a workspace
}
