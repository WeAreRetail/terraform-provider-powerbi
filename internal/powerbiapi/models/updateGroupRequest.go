package models

type UpdateGroupRequest struct {
	Name                        string                      `json:"name"`
	DefaultDatasetStorageFormat DefaultDatasetStorageFormat `json:"defaultDatasetStorageFormat"`
}

type UpdateGroupRequestName struct {
	Name string `json:"name"`
}

func (g *UpdateGroupRequest) Validate() interface{} {
	if g.DefaultDatasetStorageFormat != "" {
		return g.DefaultDatasetStorageFormat
	} else {
		return UpdateGroupRequestName{Name: g.Name}
	}
}
