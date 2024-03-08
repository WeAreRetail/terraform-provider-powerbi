package powerbiapi

import (
	"fmt"
	"terraform-provider-powerbi/internal/powerbiapi/models"
)

// GetPipelines returns the specified deployment pipeline.
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

func (c *Client) CreatePipeline(displayName string, description string) (*models.Pipelines, error) {
	// POST https://api.powerbi.com/v1.0/myorg/groups

	var err error
	pipeline := &models.Pipelines{}

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
