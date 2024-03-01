package models

// PipelineStage is a Power BI deployment pipeline stage.
type PipelineStage struct {
	Order         int    `json:"order"`         // The stage order, starting from zero.
	WorkspaceId   string `json:"workspaceId"`   // The assigned workspace ID. Only applicable when there's an assigned workspace.
	WorkspaceName string `json:"workspaceName"` // The assigned workspace name. Only applicable when there's an assigned workspace and the user has access to the workspace.
}
