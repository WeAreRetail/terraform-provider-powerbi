package powerbiapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"terraform-provider-powerbi/internal/powerbiapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetPipeline is a unit test function that tests the GetPipeline function.
// It creates a test server, sends a mock request to the server, and checks the response.
// The function verifies that the request URL, method, and response are correct.
// It also checks if the returned pipeline object has the expected ID and name.
func TestGetPipeline(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/pipelines/57eb01e2-2803-4d0d-ae65-8fd112ae5b7c", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodGet, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"@odata.context":"http://wabi-france-central-a-primary-redirect.analysis.windows.net/v1.0/myorg/$metadata#pipelines/$entity",
			"id":"57eb01e2-2803-4d0d-ae65-8fd112ae5b7c",
			"displayName":"wksPipeline","stages":[
			  {
				"order":0,"workspaceId":"6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4","workspaceName":"wks_DEV"
			  },{
				"order":1,"workspaceId":"20eb5f9d-8467-4f19-ba65-1a7bc64da6a9","workspaceName":"wks_PRD"
			  }
			]
		  }
		  `)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host := "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the GetGroup function
	pipeline, err := client.GetPipeline("57eb01e2-2803-4d0d-ae65-8fd112ae5b7c")

	// Check the result
	assert.NoError(t, err)
	assert.Equal(t, "57eb01e2-2803-4d0d-ae65-8fd112ae5b7c", pipeline.Id)
	assert.Equal(t, "wksPipeline", pipeline.DisplayName)
}

// TestCreatePipeline is a unit test function that tests the CreatePipeline function.
// It creates a test server, sends a mock request to the server, and checks the response.
// The function verifies that the request URL, method, and response are correct.
// It also checks if the returned pipeline  object has the expected ID and name.
func TestCreatePipeline(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/pipelines", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPost, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"@odata.context": "http://wabi-france-central-a-primary-redirect.analysis.windows.net/v1.0/myorg/$metadata#pipelines/$entity",
			"id": "70ab2a0e-77ec-43d1-a473-efb6058ba37d",
			"displayName": "test_pipeline",
			"description": "test Pipeline",
			"stages":[
				{
					"order": 0,
					"workspaceId":"6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4",
					"workspaceName":"Development"
				},
				{
					"order": 1,
					"workspaceId":"6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4",
					"workspaceName":"Test"
				},
				{
					"order": 2,
					"workspaceId":"20eb5f9d-8467-4f19-ba65-1a7bc64da6a9",
					"workspaceName":"Production"
				}
			]
		  }
		  `)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host := "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the CreatePipeline function
	pipeline, err := client.CreatePipeline("test_pipeline", "test Pipeline")

	// Check the result
	assert.NoError(t, err)
	assert.Equal(t, "test_pipeline", pipeline.DisplayName)
	assert.Equal(t, "test Pipeline", pipeline.Description)
}

// TestDeletePipeline is a unit test function that tests the DeletePipeline function.
// It creates a test server, sends a mock request to the server, and checks the response.
// The function verifies that the request URL, method, and response are correct.
func TestDeletePipeline(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/pipelines/70ab2a0e-77ec-43d1-a473-efb6058ba37d", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodDelete, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host := "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the CreatePipeline function
	err = client.DeletePipeline("70ab2a0e-77ec-43d1-a473-efb6058ba37d")

	// Check the result
	assert.NoError(t, err)
}

// TestUpdatePipeline is a unit test function that tests the UpdatePipeline function.
// It creates a test server, sends a mock request to the server, and checks the response.
// The function verifies that the request URL, method, and response are correct.
func TestUpdatePipeline(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/pipelines/70ab2a0e-77ec-43d1-a473-efb6058ba37d", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPatch, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"@odata.context": "http://wabi-france-central-a-primary-redirect.analysis.windows.net/v1.0/myorg/$metadata#pipelines/$entity",
			"id": "70ab2a0e-77ec-43d1-a473-efb6058ba37d",
			"displayName": "test_pipeline_rename",
			"description": "description_rename",
			"stages":[
				{
					"order": 0,
					"workspaceId":"6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4",
					"workspaceName":"Development"
				},
				{
					"order": 1,
					"workspaceId":"6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4",
					"workspaceName":"Test"
				},
				{
					"order": 2,
					"workspaceId":"20eb5f9d-8467-4f19-ba65-1a7bc64da6a9",
					"workspaceName":"Production"
				}
			]
		  }
		  `)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host := "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	updatePipelineRequest := &models.UpdatePipelineRequest{DisplayName: "test_pipeline_rename", Description: "description_rename"}
	// Call the CreatePipeline function
	pipeline, err := client.UpdatePipeline("70ab2a0e-77ec-43d1-a473-efb6058ba37d", *updatePipelineRequest)

	// Check the result
	assert.NoError(t, err)
	assert.Equal(t, "test_pipeline_rename", pipeline.DisplayName)
	assert.Equal(t, "description_rename", pipeline.Description)
}

//// TestAssignWorkspace is a unit test function that tests the AssignToWorkspace function.
//// It creates a test server, sends a mock request to the server, and checks the response.
//// The function verifies that the request URL, method, and response are correct.
//// It checks if the pipeline is correctly assigned to the workspace.
//func TestAssignWorkspace(t *testing.T) {
//	// Create a test server
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Check the request URL
//		assert.Equal(t, "/v1.0/myorg/pipelines/70ab2a0e-77ec-43d1-a473-efb6058ba37d/stages/0/assignWorkspace", r.URL.Path)
//
//		// Check the request method
//		assert.Equal(t, http.MethodPost, r.Method)
//
//		// Send a mock response
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//	}))
//	defer server.Close()
//
//	// Create a client with the test server URL
//	host := server.URL
//	//host := "https://api.powerbi.com"
//	client, err := NewClient(host)
//	assert.NoError(t, err)
//
//	// Call the CreatePipeline function
//	err = client.AssignWorkspace("70ab2a0e-77ec-43d1-a473-efb6058ba37d", 0, "6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4")
//
//	// Check the result
//	assert.NoError(t, err)
//}
