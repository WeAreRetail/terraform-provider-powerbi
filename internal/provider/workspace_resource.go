package provider

import (
	"context"
	"fmt"
	"terraform-provider-powerbi/internal/powerbiapi"
	pbiModels "terraform-provider-powerbi/internal/powerbiapi/models"
	"terraform-provider-powerbi/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &PowerBIWorkspaceResource{}                // Ensure that PowerBIWorkspaceResource implements the Resource interface.
var _ resource.ResourceWithImportState = &PowerBIWorkspaceResource{} // Ensure that PowerBIWorkspaceResource implements the ResourceWithImportState interface.
var _ resource.ResourceWithConfigure = &PowerBIWorkspaceResource{}   // Ensure that PowerBIWorkspaceResource implements the ResourceWithConfigure interface.

// NewPowerBIWorkspaceResource is a function that creates a new instance of the PowerBIWorkspaceResource.
func NewPowerBIWorkspaceResource() resource.Resource {
	return &PowerBIWorkspaceResource{}
}

// PowerBIWorkspaceResource is a struct that represents the Power BI workspace resource.
type PowerBIWorkspaceResource struct {
	client *powerbiapi.Client
}

// Configure configures the PowerBIWorkspaceResource.
func (r *PowerBIWorkspaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create creates a new Power BI workspace.
func (r *PowerBIWorkspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config models.PowerBIWorkspaceModel
	var state models.PowerBIWorkspaceModel
	var workspace *pbiModels.Group
	var err error

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating workspace with name: %s", config.Name.ValueString()))

	workspace, err = r.client.CreateGroup(config.Name.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot create workspace with name %s", config.Name.ValueString()), err.Error())
		return
	}
	tflog.Debug(ctx, "Workspace created successfully")

	tflog.Debug(ctx, "Populate the response with the workspace data")
	state.Id = types.StringValue(workspace.Id)
	state.Name = types.StringValue(workspace.Name)
	state.IsReadOnly = types.BoolValue(workspace.IsReadOnly)
	state.IsOnDedicatedCapacity = types.BoolValue(workspace.IsOnDedicatedCapacity)

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the Power BI workspace.
func (r *PowerBIWorkspaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.PowerBIWorkspaceModel
	var err error

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting workspace with name: %s", state.Name.ValueString()))
	err = r.client.DeleteGroup(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot delete group with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	tflog.Debug(ctx, "Workspace deleted successfully")
}

// ImportState implements resource.ResourceWithImportState.
func (r *PowerBIWorkspaceResource) ImportState(_ context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	panic("unimplemented")
}

// Metadata sets the metadata for the PowerBIWorkspaceResource.
func (r *PowerBIWorkspaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace"
}

// Read updates the state with the data from the Power BI service.
func (r *PowerBIWorkspaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.PowerBIWorkspaceModel
	var workspace *pbiModels.Group
	var err error

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading workspace with name: %s", state.Name.ValueString()))
	workspace, err = r.client.GetGroup(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve group with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	state.Id = types.StringValue(workspace.Id)
	state.Name = types.StringValue(workspace.Name)
	state.IsReadOnly = types.BoolValue(workspace.IsReadOnly)
	state.IsOnDedicatedCapacity = types.BoolValue(workspace.IsOnDedicatedCapacity)

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Schema sets the schema for the PowerBIWorkspaceResource.
func (r *PowerBIWorkspaceResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Power BI workspace resource",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the workspace",
				Required:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the workspace",
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

// Update implements resource.ResourceWithConfigure.
func (r *PowerBIWorkspaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan models.PowerBIWorkspaceModel
	var state models.PowerBIWorkspaceModel
	var workspace *pbiModels.Group
	var updateRequest *pbiModels.UpdateGroupRequest
	var err error

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	updateRequest = &pbiModels.UpdateGroupRequest{
		Name: plan.Name.ValueString(),
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating workspace with name: %s", state.Name.ValueString()))
	err = r.client.UpdateGroup(state.Id.ValueString(), updateRequest)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot update group with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	tflog.Debug(ctx, "Workspace updated successfully")

	tflog.Debug(ctx, "Populate the response with the workspace data")
	tflog.Debug(ctx, fmt.Sprintf("Reading workspace with name: %s", plan.Name.ValueString()))
	workspace, err = r.client.GetGroup(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve group with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	state.Id = types.StringValue(workspace.Id)
	state.Name = types.StringValue(workspace.Name)
	state.IsReadOnly = types.BoolValue(workspace.IsReadOnly)
	state.IsOnDedicatedCapacity = types.BoolValue(workspace.IsOnDedicatedCapacity)

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
