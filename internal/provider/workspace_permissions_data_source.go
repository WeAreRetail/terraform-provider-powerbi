package provider

import (
	"context"
	"fmt"
	"terraform-provider-powerbi/internal/powerbiapi"
	pbiModels "terraform-provider-powerbi/internal/powerbiapi/models"
	"terraform-provider-powerbi/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource                   = &PowerBIWorkspacePermissionsDataSource{} // Ensure that PowerBIWorkspacePermissionsDataSource implements the DataSource interface.
	_ datasource.DataSourceWithValidateConfig = &PowerBIWorkspacePermissionsDataSource{} // Ensure that PowerBIWorkspacePermissionsDataSource implements the DataSourceWithValidateConfig interface.
	_ datasource.DataSourceWithConfigure      = &PowerBIWorkspacePermissionsDataSource{} // Ensure that PowerBIWorkspacePermissionsDataSource implements the DataSourceWithConfigure interface.
)

// NewPowerBIWorkspacePermissionsDataSource is a function that creates a new instance of the PowerBIWorkspacePermissionsDataSource.
func NewPowerBIWorkspacePermissionsDataSource() datasource.DataSource {
	return &PowerBIWorkspacePermissionsDataSource{}
}

// PowerBIWorkspacePermissionsDataSource is a struct that represents the Power BI Workspace Permissions data source.
type PowerBIWorkspacePermissionsDataSource struct {
	client *powerbiapi.Client
}

// Metadata is a method that sets the metadata for the PowerBIWorkspacePermissionsDataSource.
func (d *PowerBIWorkspacePermissionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace_permissions"
}

// Schema is a method that sets the schema for the PowerBIWorkspacePermissionsDataSource.
func (d *PowerBIWorkspacePermissionsDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Power BI workspace permissions data source",

		Attributes: map[string]schema.Attribute{
			"workspace_id": schema.StringAttribute{
				MarkdownDescription: "The name of the workspace",
				Optional:            true,
				Computed:            true,
			},
			"workspace_name": schema.StringAttribute{
				MarkdownDescription: "The id of the workspace",
				Optional:            true,
				Computed:            true,
			},
			"permissions": schema.ListNestedAttribute{
				MarkdownDescription: "The permissions of the workspace",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_name": schema.StringAttribute{
							MarkdownDescription: "The display name of the principal",
							Computed:            true,
						},
						"email_address": schema.StringAttribute{
							MarkdownDescription: "The email address of the user",
							Computed:            true,
						},
						"graph_id": schema.StringAttribute{
							MarkdownDescription: "Identifier of the principal in Microsoft Graph",
							Computed:            true,
						},
						"access_right": schema.StringAttribute{
							MarkdownDescription: "The access right (permission level) that a user has on the workspace",
							Computed:            true,
						},
						"identifier": schema.StringAttribute{
							MarkdownDescription: "Identifier of the principal",
							Computed:            true,
						},
						"principal_type": schema.StringAttribute{
							MarkdownDescription: "The principal type",
							Computed:            true,
							CustomType:          types.StringType,
						},
						"profile": schema.SingleNestedAttribute{
							MarkdownDescription: "A Power BI service principal profile. Only relevant for Power BI Embedded multi-tenancy solution.",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"display_name": schema.StringAttribute{
									MarkdownDescription: "The service principal profile name",
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: "The service principal profile ID",
									Computed:            true,
								},
							},
						},
						"user_type": schema.StringAttribute{
							MarkdownDescription: "Type of the user",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

// ValidateConfig validates the configuration for the PowerBIWorkspacePermissionsDataSource.
// It checks if either the 'name' or 'id' attribute is set, and adds an error to the response diagnostics if both are empty or both are non-empty.
// If there are any errors in the response diagnostics, the function returns without further processing.
// Parameters:
//   - ctx: The context.Context object for the request.
//   - req: The ValidateConfigRequest object containing the configuration to validate.
//   - resp: The ValidateConfigResponse object to store the validation results.
//
// Returns: None.
func (d *PowerBIWorkspacePermissionsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data models.PowerBIWorkspacePermissionsModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bothNull := data.WorkspaceName.IsNull() && data.WorkspaceId.IsNull()
	bothSet := !data.WorkspaceName.IsNull() && !data.WorkspaceId.IsNull()

	if bothNull || bothSet {
		resp.Diagnostics.AddAttributeError(
			path.Root("name"),
			"Invalid attribute configuration",
			"one of 'name' or 'id' must be set",
		)
	}
}

// Configure is a method that configures the PowerBIWorkspacePermissionsDataSource.
// It creates a new instance of the powerbiapi.Client with the specified base URL.
// If an error occurs while creating the client, an error is added to the response diagnostics.
// Parameters:
//   - ctx: The context.Context object for the request.
//   - req: The ConfigureRequest object containing the configuration to configure.
//   - resp: The ConfigureResponse object to store the configuration results.
//
// Returns: None.
func (d *PowerBIWorkspacePermissionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*powerbiapi.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *powerbiapi.Client, got: %T. Please report this issue to the provider developers.",
		)

		return
	}

	d.client = client
}

// Read is a method that reads the data from the Power BI service and returns the result.
// It takes a ReadRequest and a ReadResponse as input and output parameters, respectively.
// It reads the data from the Power BI service and returns the result in the response.
// Parameters:
//   - ctx: The context.Context object for the request.
//   - req: The ReadRequest object containing the configuration to read.
//   - resp: The ReadResponse object to store the read results.
//
// Returns: None.
func (d *PowerBIWorkspacePermissionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.PowerBIWorkspacePermissionsModel
	var workspace *pbiModels.Group
	var workspaceUsers *pbiModels.GroupUsers
	var err error

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.WorkspaceId.IsNull() {

		workspace, err = d.client.GetGroup(data.WorkspaceId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with Id %s", data.WorkspaceId.ValueString()), err.Error())
			return
		}

		workspaceUsers, err = d.client.GetGroupUsers(data.WorkspaceId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve permissions for workspace with Id %s", data.WorkspaceId.ValueString()), err.Error())
			return
		}
	}

	if !data.WorkspaceName.IsNull() {
		workspaces, err := d.client.GetGroups(fmt.Sprintf("name eq '%s'", data.WorkspaceName.ValueString()), 0, 0)

		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with name %s", data.WorkspaceName.ValueString()), err.Error())
			return
		}

		if len(workspaces.Value) == 0 {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with name %s", data.WorkspaceName.ValueString()), "No groups found")
			return
		}

		if len(workspaces.Value) > 1 {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with name %s", data.WorkspaceName.ValueString()), "Multiple groups found")
			return
		}

		workspace, err = d.client.GetGroup(workspaces.Value[0].Id)
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with Id %s", data.WorkspaceId.ValueString()), err.Error())
			return
		}

		workspaceUsers, err = d.client.GetGroupUsers(workspaces.Value[0].Id)

		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve permissions for workspace %s", data.WorkspaceName.ValueString()), err.Error())
			return
		}
	}

	if workspaceUsers == nil {
		resp.Diagnostics.AddError("No workspace users found", "No workspace users found with the specified name or id.")
		return
	}

	// Loop on the workspace users and populate the data
	var permissions []models.PowerBIWorkspacePermissionModel

	for _, user := range workspaceUsers.Value {
		var permission models.PowerBIWorkspacePermissionModel

		permission.DisplayName = types.StringValue(user.DisplayName)

		if user.EmailAddress != "" {
			permission.EmailAddress = types.StringValue(user.EmailAddress)
		} else {
			permission.EmailAddress = types.StringNull()
		}

		if user.GraphId != "" {
			permission.GraphId = types.StringValue(user.GraphId)
		} else {
			permission.GraphId = types.StringNull()
		}

		if user.GroupUserAccessRight != "" {
			permission.AccessRight = types.StringValue(string(user.GroupUserAccessRight))
		} else {
			permission.AccessRight = types.StringNull()
		}

		//TODO: manage nil
		permission.Identifier = types.StringValue(user.Identifier)

		permission.PrincipalType = types.StringValue(string(user.PrincipalType))

		permission.Profile = models.ServicePrincipalProfile{
			DisplayName: types.StringValue(user.Profile.DisplayName),
			Id:          types.StringValue(user.Profile.Id),
		}

		permission.UserType = types.StringValue(user.UserType)

		permissions = append(permissions, permission)
	}

	data.WorkspaceId = types.StringValue(workspace.Id)
	data.WorkspaceName = types.StringValue(workspace.Name)
	data.Permissions = permissions

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
