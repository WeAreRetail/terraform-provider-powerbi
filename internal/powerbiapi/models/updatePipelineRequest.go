package models

// UpdatePipelineRequest represents a request to update a Power BI pipeline (workspace).
type UpdatePipelineRequest struct {
	DisplayName string `json:"displayName"` // The pipeline (workspace) name
	Description string `json:"description"` //The default dataset storage format in the pipeline
}

// UpdatePipelineRequestName represents a request to update a Power BI pipeline (workspace) name.
// This is a helper struct to allow the UpdatePipelineRequest to be validated.
type UpdatePipelineRequestName struct {
	DisplayName string `json:"displayName"`
}

// UpdatePipelineRequestDescription represents a request to update a Power BI pipeline (workspace) name.
// This is a helper struct to allow the UpdatePipelineRequest to be validated.
type UpdatePipelineRequestDescription struct {
	Description string `json:"description"`
}

// Validate ensures that the UpdatePipelineRequest is valid.
// If the DefaultDatasetStorageFormat is not empty, it is returned.
func (p *UpdatePipelineRequest) Validate() interface{} {
	if p.DisplayName != "" && p.Description != "" {
		return p
	} else if p.DisplayName != "" {
		return UpdatePipelineRequestName{DisplayName: p.DisplayName}
	} else if p.Description != "" {
		return UpdatePipelineRequestDescription{p.Description}
	}
	return p
}
