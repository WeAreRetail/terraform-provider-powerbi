package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PowerBIWorkspaceModel is a struct that represents the workspace permissions data model.
type PowerBIWorkspacePermissionsModel struct {
	WorkspaceId   types.String                      `tfsdk:"workspace_id"`
	WorkspaceName types.String                      `tfsdk:"workspace_name"`
	Permissions   []PowerBIWorkspacePermissionModel `tfsdk:"permissions"`
}

type PowerBIWorkspacePermissionModel struct {
	DisplayName   types.String            `tfsdk:"display_name"`
	EmailAddress  types.String            `tfsdk:"email_address"`
	GraphId       types.String            `tfsdk:"graph_id"`
	AccessRight   types.String            `tfsdk:"access_right"`
	Identifier    types.String            `tfsdk:"identifier"`
	PrincipalType types.String            `tfsdk:"principal_type"`
	Profile       ServicePrincipalProfile `tfsdk:"profile"`
	UserType      types.String            `tfsdk:"user_type"`
}

type ServicePrincipalProfile struct {
	DisplayName types.String `tfsdk:"display_name"`
	Id          types.String `tfsdk:"id"`
}
