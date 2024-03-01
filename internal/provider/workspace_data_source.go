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
	_ datasource.DataSource                   = &WorkspaceDataSource{} // Ensure that WorkspaceDataSource implements the DataSource interface.
	_ datasource.DataSourceWithValidateConfig = &WorkspaceDataSource{} // Ensure that WorkspaceDataSource implements the DataSourceWithValidateConfig interface.
	_ datasource.DataSourceWithConfigure      = &WorkspaceDataSource{} // Ensure that WorkspaceDataSource implements the DataSourceWithConfigure interface.
)

// NewWorkspaceDataSource is a function that creates a new instance of the WorkspaceDataSource.
func NewWorkspaceDataSource() datasource.DataSource {
	return &WorkspaceDataSource{}
}

// WorkspaceDataSource is a struct that represents the Power BI workspace data source.
type WorkspaceDataSource struct {
	client *powerbiapi.Client
}

// Metadata is a method that sets the metadata for the WorkspaceDataSource.
func (d *WorkspaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

// Schema is a method that sets the schema for the WorkspaceDataSource.
func (d *WorkspaceDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Power BI workspace data source",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the workspace",
				Optional:            true,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the workspace",
				Optional:            true,
				Computed:            true,
			},
			"is_read_only": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether the workspace is read-only",
				Computed:            true,
			},
			"is_on_dedicated_capacity": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether the workspace is on dedicated capacity",
				Computed:            true,
			},
		},
	}
}

// ValidateConfig validates the configuration for the WorkspaceDataSource.
// It checks if either the 'name' or 'id' attribute is set, and adds an error to the response diagnostics if both are empty or both are non-empty.
// If there are any errors in the response diagnostics, the function returns without further processing.
// Parameters:
//   - ctx: The context.Context object for the request.
//   - req: The ValidateConfigRequest object containing the configuration to validate.
//   - resp: The ValidateConfigResponse object to store the validation results.
//
// Returns: None.
func (d *WorkspaceDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var data models.Workspace

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bothNull := data.Name.IsNull() && data.Id.IsNull()
	bothSet := !data.Name.IsNull() && !data.Id.IsNull()

	if bothNull || bothSet {
		resp.Diagnostics.AddAttributeError(
			path.Root("name"),
			"Invalid attribute configuration",
			"one of 'name' or 'id' must be set",
		)
	}
}

// Configure is a method that configures the WorkspaceDataSource.
// It creates a new instance of the powerbiapi.Client with the specified base URL.
// If an error occurs while creating the client, an error is added to the response diagnostics.
// Parameters:
//   - ctx: The context.Context object for the request.
//   - req: The ConfigureRequest object containing the configuration to configure.
//   - resp: The ConfigureResponse object to store the configuration results.
//
// Returns: None.
func (d *WorkspaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *WorkspaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.Workspace
	var workspace *pbiModels.Group
	var err error

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Id.IsNull() {
		workspace, err = d.client.GetGroup(data.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with Id %s", data.Id.ValueString()), err.Error())
			return
		}
	}

	if !data.Name.IsNull() {
		workspaces, err := d.client.GetGroups(fmt.Sprintf("name eq '%s'", data.Name.ValueString()), 0, 0)
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with name %s", data.Name.ValueString()), err.Error())
			return
		}

		if len(workspaces.Value) == 0 {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with name %s", data.Name.ValueString()), "No groups found")
			return
		}

		if len(workspaces.Value) > 1 {
			resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve workspace with name %s", data.Name.ValueString()), "Multiple groups found")
			return
		}

		workspace = &workspaces.Value[0]
	}

	if workspace == nil {
		resp.Diagnostics.AddError("No workspace found", "No workspace found with the specified name or id.")
		return
	}

	data.Id = types.StringValue(workspace.Id)
	data.Name = types.StringValue(workspace.Name)
	data.IsReadOnly = types.BoolValue(workspace.IsReadOnly)
	data.IsOnDedicatedCapacity = types.BoolValue(workspace.IsOnDedicatedCapacity)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
