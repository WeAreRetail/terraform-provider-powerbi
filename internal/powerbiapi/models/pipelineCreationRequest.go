package models

// PipelineCreationRequest is a Power BI pipeline creation object.
type PipelineCreationRequest struct {
	Description string `json:"description"` // The deployment pipeline description
	DisplayName string `json:"displayName"` // The deployment pipeline display name
}
