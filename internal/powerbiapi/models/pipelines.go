package models

// Pipeline is a Power BI pipeline.
type Pipelines struct {
	ODataContext string `json:"@odata.context"` // The OData context
	Id           string `json:"id"`
	DisplayName  string `json:"displayName"`
	Description  string `json:"description"`
}
