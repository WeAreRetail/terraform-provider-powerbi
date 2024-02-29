package models

// Pipeline is a Power BI pipeline.
type Pipeline struct {
	Description string          `json:"description"`
	DisplayName string          `json:"displayName"`
	Id          string          `json:"id"`
	Stages      []PipelineStage `json:"stages"`
}
