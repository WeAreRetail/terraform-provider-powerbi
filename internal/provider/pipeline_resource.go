package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"terraform-provider-powerbi/internal/powerbiapi"
	pbiModels "terraform-provider-powerbi/internal/powerbiapi/models"
	"terraform-provider-powerbi/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &PipelineResource{}                // Ensure that PipelineResource implements the Resource interface.
var _ resource.ResourceWithImportState = &PipelineResource{} // Ensure that PipelineResource implements the ResourceWithImportState interface.
var _ resource.ResourceWithConfigure = &PipelineResource{}   // Ensure that PipelineResource implements the ResourceWithConfigure interface.

// NewPipelineResource is a function that creates a new instance of the PipelineResource.
func NewPipelineResource() resource.Resource {
	return &PipelineResource{}
}

// PipelineResource is a struct that represents the Power BI pipeline resource.
type PipelineResource struct {
	client *powerbiapi.Client
}

// Configure configures the PipelineResource.
func (r *PipelineResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create creates a new Power BI pipeline.
func (r *PipelineResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config models.Pipeline
	var state models.Pipeline
	var pipeline *pbiModels.Pipeline
	var err error

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating pipeline with name: %s", config.DisplayName.ValueString()))

	pipeline, err = r.client.CreatePipeline(config.DisplayName.ValueString(), config.Description.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot create pipeline with name %s", config.DisplayName.ValueString()), err.Error())
		return
	}
	tflog.Debug(ctx, "Pipeline created successfully")

	// As the creation of the pipeline doesn't yield a full json object describing the pipeline
	// We must get the pipeline full json object by using the GetPipeline method.
	pipeline, err = r.client.GetPipeline(pipeline.Id)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve pipeline with Id %s", state.Id.ValueString()), err.Error())
		return
	}
	tflog.Debug(ctx, "Populate the response with the pipeline data")
	state.Id = types.StringValue(pipeline.Id)
	state.DisplayName = types.StringValue(pipeline.DisplayName)
	state.Description = types.StringValue(pipeline.Description)
	var stages []models.PipelineStage
	for _, stage := range pipeline.Stages {
		var pipelineStage models.PipelineStage

		if stage.WorkspaceName != "" {
			pipelineStage.WorkspaceName = types.StringValue(stage.WorkspaceName)
		} else {
			pipelineStage.WorkspaceName = types.StringNull()
		}

		if stage.WorkspaceId != "" {
			pipelineStage.WorkspaceId = types.StringValue(stage.WorkspaceId)
		} else {
			pipelineStage.WorkspaceId = types.StringNull()
		}

		pipelineStage.Order = types.Int64Value(int64(stage.Order))
		stages = append(stages, pipelineStage)
	}

	// Set stages on the state model
	resp.Diagnostics.Append(tfsdk.ValueFrom(ctx, stages, types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"order":          types.Int64Type,
				"workspace_id":   types.StringType,
				"workspace_name": types.StringType,
			},
		},
	}, &state.Stages)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the Power BI pipeline.
func (r *PipelineResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.Pipeline
	var err error

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting pipeline with name: %s", state.DisplayName.ValueString()))
	err = r.client.DeletePipeline(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot delete pipeline with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	tflog.Debug(ctx, "Pipeline deleted successfully")
}

// ImportState implements resource.ResourceWithImportState.
func (r *PipelineResource) ImportState(_ context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	panic("unimplemented")
}

// Metadata sets the metadata for the PipelineResource.
func (r *PipelineResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pipeline"
}

// Read updates the state with the data from the Power BI service.
func (r *PipelineResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.Pipeline
	var pipeline *pbiModels.Pipeline
	var err error

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading pipeline with name: %s", state.DisplayName.ValueString()))
	pipeline, err = r.client.GetPipeline(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve pipeline with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	state.Id = types.StringValue(pipeline.Id)
	state.DisplayName = types.StringValue(pipeline.DisplayName)
	state.Description = types.StringValue(pipeline.Description)
	var stages []models.PipelineStage
	for _, stage := range pipeline.Stages {
		var pipelineStage models.PipelineStage

		if stage.WorkspaceName != "" {
			pipelineStage.WorkspaceName = types.StringValue(stage.WorkspaceName)
		} else {
			pipelineStage.WorkspaceName = types.StringNull()
		}

		if stage.WorkspaceId != "" {
			pipelineStage.WorkspaceId = types.StringValue(stage.WorkspaceId)
		} else {
			pipelineStage.WorkspaceId = types.StringNull()
		}

		pipelineStage.Order = types.Int64Value(int64(stage.Order))
		stages = append(stages, pipelineStage)
	}

	// Set stages on the state model
	resp.Diagnostics.Append(tfsdk.ValueFrom(ctx, stages, types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"order":          types.Int64Type,
				"workspace_id":   types.StringType,
				"workspace_name": types.StringType,
			},
		},
	}, &state.Stages)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Schema sets the schema for the PipelineResource.
func (r *PipelineResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Power BI pipeline resource",
		Attributes: map[string]schema.Attribute{
			"display_name": schema.StringAttribute{
				MarkdownDescription: "The name of the pipeline",
				Required:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the pipeline",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the pipeline",
				Computed:            false,
				Required:            false,
				Optional:            true,
			},
			"stages": schema.ListAttribute{
				MarkdownDescription: "The stages of the pipeline",
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"order":          types.Int64Type,
						"workspace_id":   types.StringType,
						"workspace_name": types.StringType,
					},
				},
			},
		},
	}
}

// Update updates the Power BI pipeline.
func (r *PipelineResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan models.Pipeline
	var state models.Pipeline
	var pipeline *pbiModels.Pipeline
	var updateRequest *pbiModels.UpdatePipelineRequest
	var err error

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	updateRequest = &pbiModels.UpdatePipelineRequest{
		DisplayName: plan.DisplayName.ValueString(),
		Description: plan.Description.ValueString(),
	}
	tflog.Info(ctx, fmt.Sprintf("Update pipeline Request: %s, %s", updateRequest.DisplayName, updateRequest.Description))

	tflog.Debug(ctx, fmt.Sprintf("Updating pipeline with name: %s", state.DisplayName.ValueString()))
	_, err = r.client.UpdatePipeline(state.Id.ValueString(), *updateRequest)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot update pipeline with Id %s %s", state.Id.ValueString(), err), err.Error())
		return
	}

	tflog.Debug(ctx, "Pipeline updated successfully")

	tflog.Debug(ctx, "Populate the response with the pipeline data")
	tflog.Debug(ctx, fmt.Sprintf("Reading pipeline with name: %s", plan.DisplayName.ValueString()))
	pipeline, err = r.client.GetPipeline(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Cannot retrieve pipeline with Id %s", state.Id.ValueString()), err.Error())
		return
	}

	state.Id = types.StringValue(pipeline.Id)
	state.DisplayName = types.StringValue(pipeline.DisplayName)
	state.Description = types.StringValue(pipeline.Description)

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
