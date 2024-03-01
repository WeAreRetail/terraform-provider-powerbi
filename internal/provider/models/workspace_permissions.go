package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// WorkspacePermissionsData is a struct that represents the workspace permissions data model.
type WorkspacePermissionsData struct {
	WorkspaceId   types.String          `tfsdk:"workspace_id"`   // The workspace id.
	WorkspaceName types.String          `tfsdk:"workspace_name"` // The workspace name.
	Permissions   []WorkspacePermission `tfsdk:"permissions"`    // A list of permissions.
}

// WorkspacePermission is a struct that represents a single workspace permission data model.
type WorkspacePermission struct {
	DisplayName   types.String            `tfsdk:"display_name"`   // The display name of the user, group, or service principal.
	EmailAddress  types.String            `tfsdk:"email_address"`  // The email address of the user.
	GraphId       types.String            `tfsdk:"graph_id"`       // Identifier of the principal in Microsoft Graph. Only available for admin APIs.
	AccessRight   types.String            `tfsdk:"access_right"`   // The type of access right. Possible values include: "None", "Viewer", "Member", "Contributor", "Admin"
	Identifier    types.String            `tfsdk:"identifier"`     // Identifier of the principal.
	PrincipalType types.String            `tfsdk:"principal_type"` // The principal type. Possible values include: "User", "Group", "App"
	Profile       ServicePrincipalProfile `tfsdk:"profile"`        // A Power BI service principal profile. Only relevant for Power BI Embedded multi-tenancy solution.
	UserType      types.String            `tfsdk:"user_type"`      // Type of the user.
}

// ServicePrincipalProfile is a struct that represents a service principal profile data model.
type ServicePrincipalProfile struct {
	DisplayName types.String `tfsdk:"display_name"` // The display name of the service principal.
	Id          types.String `tfsdk:"id"`           // The identifier of the service principal.
}
