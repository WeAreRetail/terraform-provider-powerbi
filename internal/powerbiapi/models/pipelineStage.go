package models

// PipelineStage is a Power BI deployment pipeline stage.
type PipelineStage struct {
	Order         int    `json:"order"`
	WorkspaceId   string `json:"workspaceId"`
	WorkspaceName string `json:"workspaceName"`
}
