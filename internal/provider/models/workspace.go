package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Workspace is a struct that represents the workspace data model.
type Workspace struct {
	IsReadOnly            types.Bool   `tfsdk:"is_read_only"`
	IsOnDedicatedCapacity types.Bool   `tfsdk:"is_on_dedicated_capacity"`
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
}
