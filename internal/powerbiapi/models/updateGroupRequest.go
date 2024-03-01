package models

// UpdateGroupRequest represents a request to update a Power BI group (workspace).
type UpdateGroupRequest struct {
	Name                        string                      `json:"name"`                        // The group (workspace) name
	DefaultDatasetStorageFormat DefaultDatasetStorageFormat `json:"defaultDatasetStorageFormat"` //The default dataset storage format in the group
}

// UpdateGroupRequestName represents a request to update a Power BI group (workspace) name.
// This is a helper struct to allow the UpdateGroupRequest to be validated.
type UpdateGroupRequestName struct {
	Name string `json:"name"`
}

// Validate ensures that the UpdateGroupRequest is valid.
// If the DefaultDatasetStorageFormat is not empty, it is returned.
func (g *UpdateGroupRequest) Validate() interface{} {
	if g.DefaultDatasetStorageFormat != "" {
		return g.DefaultDatasetStorageFormat
	} else {
		return UpdateGroupRequestName{Name: g.Name}
	}
}
