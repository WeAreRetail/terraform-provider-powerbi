package models

// AssignWorkspaceRequest is a struct containing the necessary parameters to assign a workspace to a pipeline.
type AssignWorkspaceRequest struct {
	WorkspaceId string `json:"workspaceId"`
}
