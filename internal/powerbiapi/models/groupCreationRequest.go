package models

// GroupCreationRequest represents a request to create a Power BI group (workspace).
type GroupCreationRequest struct {
	Name string `json:"name"` // The group (workspace) name
}
