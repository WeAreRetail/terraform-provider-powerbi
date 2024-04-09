package powerbiapi

import (
	"fmt"
	"terraform-provider-powerbi/internal/powerbiapi/models"
)

// GetPipeline returns the specified deployment pipeline.
// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/get-pipeline
func (c *Client) GetPipeline(pipelineId string) (*models.Pipeline, error) {
	// GET https://api.powerbi.com/v1.0/myorg/pipelines/{pipelineId}

	var err error
	pipeline := &models.Pipeline{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for GetPipelines: %v", err)
	}

	// The "expand" query parameter is used to include the stages in the response.
	resp, err := client.SetResult(pipeline).SetQueryParam("$expand", "stages").
		Get(fmt.Sprintf("/v1.0/myorg/pipelines/%s", pipelineId))
	if err != nil {
		return nil, fmt.Errorf("failed to get pipelines: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get pipelines: %v", resp.Error())
	}

	return pipeline, nil
}

// CreatePipeline returns the specified deployment pipeline.
// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/create-pipeline
func (c *Client) CreatePipeline(displayName string, description string) (*models.Pipeline, error) {
	// POST https://api.powerbi.com/v1.0/myorg/pipelines

	var err error
	pipeline := &models.Pipeline{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for CreatePipeline: %v", err)
	}

	resp, err := client.SetResult(pipeline).
		SetBody(&models.PipelineCreationRequest{DisplayName: displayName, Description: description}).
		Post("/v1.0/myorg/pipelines")
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to create pipeline: %v", resp.Error())
	}

	return pipeline, nil
}

// DeletePipeline returns the specified deployment pipeline.
// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/create-pipeline
func (c *Client) DeletePipeline(pipelineId string) error {
	// POST https://api.powerbi.com/v1.0/myorg/pipelines

	var err error

	client, err := c.prepRequest()
	if err != nil {
		return fmt.Errorf("failed to prepare the request for DeletePipeline: %v", err)
	}

	resp, err := client.
		Delete(fmt.Sprintf("/v1.0/myorg/pipelines/%s", pipelineId))
	if err != nil {
		return fmt.Errorf("failed to delete pipeline: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to delete pipeline: %v", resp.Error())
	}

	return nil
}

// UpdatePipeline returns the specified deployment pipeline.
// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/create-pipeline
func (c *Client) UpdatePipeline(pipelineId string, request models.UpdatePipelineRequest) (*models.Pipeline, error) {
	// POST https://api.powerbi.com/v1.0/myorg/pipelines

	var err error
	pipeline := &models.Pipeline{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for DeletePipeline: %v", err)
	}

	resp, err := client.
		SetResult(pipeline).
		SetBody(&models.UpdatePipelineRequest{DisplayName: request.DisplayName, Description: request.Description}).
		Patch(fmt.Sprintf("/v1.0/myorg/pipelines/%s", pipelineId))
	if err != nil {
		return nil, fmt.Errorf("failed to update pipeline: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to update pipeline: %v %v", resp.RawResponse, err)
	}

	return pipeline, nil
}

// This is commented out, to be moved in its own resource to respect the terraform provider philosophy
//// AssignWorkspace assigns a workspace to a PowerBi Pipeline.
//// https://learn.microsoft.com/en-us/rest/api/power-bi/pipelines/assign-workspace
//func (c *Client) AssignWorkspace(pipelineId string, stageOrder int, workspaceId string) error {
//	// POST https://api.powerbi.com/v1.0/myorg/pipelines/{pipelineId}/stages/{stageOrder}/assignWorkspace
//
//	var err error
//	client, err := c.prepRequest()
//	if err != nil {
//		return fmt.Errorf("failed to prepare the request for AssignWorkspace: %v", err)
//	}
//
//	resp, err := client.SetBody(&models.AssignWorkspaceRequest{WorkspaceId: workspaceId}).
//		Post(fmt.Sprintf("/v1.0/myorg/pipelines/%s/stages/%d/assignWorkspace", pipelineId, stageOrder))
//	if err != nil {
//		return fmt.Errorf("failed to assign workspace to pipeline: %v", err)
//	}
//	if resp.IsError() {
//		return fmt.Errorf("failed to assign workspace to pipeline: %v", resp.Error())
//	}
//
//	return nil
//}
