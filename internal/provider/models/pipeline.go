package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Pipeline is a struct that represents the pipeline data model.
// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/create-pipeline#pipeline
type Pipeline struct {
	Description types.String `tfsdk:"description"`
	DisplayName types.String `tfsdk:"display_name"`
	Id          types.String `tfsdk:"id"`
	Stages      types.List   `tfsdk:"stages"`
}

// PipelineStage is a struct that represents the stage data model
// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/create-pipeline#pipelinestage
type PipelineStage struct {
	Order         types.Int64  `tfsdk:"order"`
	WorkspaceId   types.String `tfsdk:"workspace_id"`
	WorkspaceName types.String `tfsdk:"workspace_name"`
}
