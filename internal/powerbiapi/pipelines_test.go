package powerbiapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetGroup is a unit test function that tests the GetGroup function.
// It creates a test server, sends a mock request to the server, and checks the response.
// The function verifies that the request URL, method, and response are correct.
// It also checks if the returned group object has the expected ID and name.
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
