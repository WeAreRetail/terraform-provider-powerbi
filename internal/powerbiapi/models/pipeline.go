package models

// Pipeline is a Power BI pipeline.
type Pipeline struct {
	Description string          `json:"description"` // The deployment pipeline description
	DisplayName string          `json:"displayName"` // The deployment pipeline display name
	Id          string          `json:"id"`          // The deployment pipeline ID
	Stages      []PipelineStage `json:"stages"`      // The collection of deployment pipeline stages. Only returned when $expand is set to stages in the request.
}
