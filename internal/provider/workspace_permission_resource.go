package provider

import (
	"context"
	"terraform-provider-powerbi/internal/powerbiapi"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = &WorkspacePermissionResource{}                   // Ensure that WorkspacePermissionResource implements the Resource interface.
var _ resource.ResourceWithImportState = &WorkspacePermissionResource{}    // Ensure that WorkspacePermissionResource implements the ResourceWithImportState interface.
var _ resource.ResourceWithConfigure = &WorkspacePermissionResource{}      // Ensure that WorkspacePermissionResource implements the ResourceWithConfigure interface.
var _ resource.ResourceWithValidateConfig = &WorkspacePermissionResource{} // Ensure that WorkspacePermissionResource implements the ResourceWithValidateConfig interface.

// NewWorkspacePermissionResource is a function that creates a new instance of the WorkspacePermissionResource.
func NewWorkspacePermissionResource() resource.Resource {
	return &WorkspacePermissionResource{}
}

// WorkspacePermissionResource is a struct that represents the Power BI workspace resource.
type WorkspacePermissionResource struct {
	client *powerbiapi.Client
}

// ValidateConfig validates the configuration for the WorkspacePermissionResource.
func (*WorkspacePermissionResource) ValidateConfig(context.Context, resource.ValidateConfigRequest, *resource.ValidateConfigResponse) {
	panic("unimplemented")
}

// Configure configures the WorkspacePermissionResource.
func (r *WorkspacePermissionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

// Create creates a new Power BI workspace permission.
func (r *WorkspacePermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	panic("unimplemented")
}

// Delete deletes the Power BI workspace permission.
func (r *WorkspacePermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	panic("unimplemented")
}

// ImportState implements resource.ResourceWithImportState.
func (r *WorkspacePermissionResource) ImportState(_ context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	panic("unimplemented")
}

// Metadata sets the metadata for the WorkspacePermissionResource.
func (r *WorkspacePermissionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace_permission"
}

// Read updates the state with the data from the Power BI service.
func (r *WorkspacePermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	panic("unimplemented")
}

// Schema sets the schema for the WorkspacePermissionResource.
func (r *WorkspacePermissionResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Power BI workspace permission resource",
		Attributes: map[string]schema.Attribute{
			"workspace_id": schema.StringAttribute{
				MarkdownDescription: "The name of the workspace",
				Required:            true,
			},
		},
	}
}

// Update implements resource.ResourceWithConfigure.
func (r *WorkspacePermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	panic("unimplemented")
}
