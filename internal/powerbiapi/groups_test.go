package powerbiapi

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"terraform-provider-powerbi/internal/powerbiapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAddGroupUser_User is a unit test function that tests the AddGroupUser method of the Client struct.
// It verifies that the correct request URL and method are used, and that the function returns no error.
func TestAddGroupUser_User(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/65d6aaca-2275-4e70-bb4f-91dde4dc6c99/users", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPost, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Prepare Request
	groupUserAccess := &models.GroupUser{
		EmailAddress:         "john.doe@example.com",
		GroupUserAccessRight: models.GroupUserAccessRightAdmin,
		PrincipalType:        models.PrincipalTypeUser,
	}

	// Call the GetGroup function
	err = client.AddGroupUser("65d6aaca-2275-4e70-bb4f-91dde4dc6c99", groupUserAccess)

	// Check the result
	assert.NoError(t, err)
}

// TestAddGroupUser_Group is a unit test function that tests the AddGroupUser method of the Client struct.
// It verifies that the correct request URL and method are used, and that the function returns no error.
func TestAddGroupUser_Group(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/ac653691-1af8-4be1-8468-9d73cdcc1250/users", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPost, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Prepare Request
	groupUserAccess := &models.GroupUser{
		Identifier:           "796131c3-8d85-44e1-bdfc-88ad8ba46520",
		GroupUserAccessRight: models.GroupUserAccessRightAdmin,
		PrincipalType:        models.PrincipalTypeGroup,
	}

	// Call the GetGroup function
	err = client.AddGroupUser("ac653691-1af8-4be1-8468-9d73cdcc1250", groupUserAccess)

	// Check the result
	assert.NoError(t, err)
}

// TestCreateGroup is a unit test function that tests the CreateGroup function of the client.
// It creates a test server, sends a mock request to the server, and checks the response.
func TestCreateGroup(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPost, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"isReadOnly": false,
			"isOnDedicatedCapacity": true,
			"capacityId": "FEFC7A26-4758-41CD-9069-5E73B7E9DB0E",
			"defaultDatasetStorageFormat": "Small",
			"type": "Workspace",
			"id": "abd062a2-a915-4e2d-921a-db799a069f62",
			"name": "UNIT_TEST"
		  }
		  `)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the GetGroup function
	group, err := client.CreateGroup("UNIT_TEST")

	// Check the result
	assert.NoError(t, err)
	assert.Equal(t, "UNIT_TEST", group.Name)
}

// TestDeleteGroup is a unit test function that tests the DeleteGroup function.
// It creates a test server, sends a DELETE request to the server, and checks the response.
func TestDeleteGroup(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/878026dd-3e07-402e-a38f-9a2a0356d83f", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodDelete, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the DeleteGroup function
	err = client.DeleteGroup("878026dd-3e07-402e-a38f-9a2a0356d83f")

	// Check the result
	assert.NoError(t, err)
}

// TestDeleteUserGroup is a unit test function that tests the DeleteUserGroup method of the Client struct.
// It creates a test server, sends a DELETE request to the server, and checks the response.
// The function verifies that the request URL and method are correct, and that the response status code is HTTP 200 OK.
// Finally, it checks if the DeleteUserGroup method returns an error or not.
func TestDeleteUserGroup(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/ac653691-1af8-4be1-8468-9d73cdcc1250/users/796131c3-8d85-44e1-bdfc-88ad8ba46520", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodDelete, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the DeleteGroup function
	err = client.DeleteUserGroup("ac653691-1af8-4be1-8468-9d73cdcc1250", "796131c3-8d85-44e1-bdfc-88ad8ba46520")

	// Check the result
	assert.NoError(t, err)
}

// TestGetGroup is a unit test function that tests the GetGroup function.
// It creates a test server, sends a mock request to the server, and checks the response.
// The function verifies that the request URL, method, and response are correct.
// It also checks if the returned group object has the expected ID and name.
func TestGetGroup(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/465d5aaa-c6a7-4add-a618-dc76d27a00ca", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodGet, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"isReadOnly": false,
			"isOnDedicatedCapacity": true,
			"capacityId": "FEFC7A26-4758-41CD-9069-5E73B7E9DB0E",
			"defaultDatasetStorageFormat": "Small",
			"type": "Workspace",
			"id": "465d5aaa-c6a7-4add-a618-dc76d27a00ca",
			"name": "INFRA_TEST"
		  }
		  `)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the GetGroup function
	group, err := client.GetGroup("465d5aaa-c6a7-4add-a618-dc76d27a00ca")

	log.Printf("group: %v", group)

	// Check the result
	assert.NoError(t, err)
	assert.Equal(t, "465d5aaa-c6a7-4add-a618-dc76d27a00ca", group.Id)
	assert.Equal(t, "INFRA_TEST", group.Name)
}

// TestGetGroupUsers tests the GetGroupUsers function.
func TestGetGroupUsers(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/ac653691-1af8-4be1-8468-9d73cdcc1250/users", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodGet, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"@odata.context": "http://wabi-france-central-a-primary-redirect.analysis.windows.net/v1.0/myorg/groups/ac653691-1af8-4be1-8468-9d73cdcc1250/$metadata#users",
			"value": [
			  {
				"emailAddress": "john.doe@example.com",
				"groupUserAccessRight": "Admin",
				"displayName": "John Doe",
				"identifier": "john.doe@example.com",
				"principalType": "User"
			  }
			]
		  }`)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the GetGroup function
	_, err = client.GetGroupUsers("ac653691-1af8-4be1-8468-9d73cdcc1250")

	// Check the result
	assert.NoError(t, err)
}

// TestGetGroups is a unit test function that tests the GetGroups function of the client.
// It creates a test server, sends a mock response, and checks the result.
func TestGetGroups(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodGet, r.Method)

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"@odata.context": "http://wabi-france-central-a-primary-redirect.analysis.windows.net/v1.0/myorg/$metadata#groups",
			"@odata.count": 6,
			"value": [
			  {
				"id": "86c546e0-31dc-4c3f-92c7-cf7d4aa04757",
				"isReadOnly": false,
				"isOnDedicatedCapacity": false,
				"type": "Workspace",
				"name": "FINOPS - QA"
			  },
			  {
				"id": "8b1b14e7-64dd-4d2f-9383-fbcce9910009",
				"isReadOnly": false,
				"isOnDedicatedCapacity": true,
				"capacityId": "FEFC7A26-4758-41CD-9069-5E73B7E9DB0E",
				"defaultDatasetStorageFormat": "Small",
				"type": "Workspace",
				"name": "DATA_GATEWAY"
			  },
			  {
				"id": "4d8b460c-3b58-4eaf-b1e2-489c1c7b81f7",
				"isReadOnly": false,
				"isOnDedicatedCapacity": false,
				"type": "AdminInsights",
				"name": "Admin monitoring"
			  },
			  {
				"id": "465d5aaa-c6a7-4add-a618-dc76d27a00ca",
				"isReadOnly": false,
				"isOnDedicatedCapacity": true,
				"capacityId": "FEFC7A26-4758-41CD-9069-5E73B7E9DB0E",
				"defaultDatasetStorageFormat": "Small",
				"type": "Workspace",
				"name": "INFRA_TEST"
			  },
			  {
				"id": "70a4c5b6-da4c-44fd-8d69-f11d00ae9148",
				"isReadOnly": false,
				"isOnDedicatedCapacity": false,
				"type": "Workspace",
				"name": "API_WORKSPACE"
			  },
			  {
				"id": "deef788f-0798-4b6c-8b52-e8e2ed39de93",
				"isReadOnly": false,
				"isOnDedicatedCapacity": true,
				"capacityId": "FEFC7A26-4758-41CD-9069-5E73B7E9DB0E",
				"defaultDatasetStorageFormat": "Small",
				"type": "Workspace",
				"name": "INFRA_PRD"
			  }
			]
		  }
		  `)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the GetGroup function
	group, err := client.GetGroups("", 0, 0)

	t.Logf("group: %v", group)

	// Check the result
	assert.NoError(t, err)
}

// TestUpdateGroups is a unit test function that tests the UpdateGroup method of the client.
// It creates a test server, sends a PATCH request to update a group, and checks the response.
// The function verifies that the request URL, method, and response status code are correct.
// It also checks if the client successfully updates the group and returns no error.
func TestUpdateGroups(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/370e64cb-da5a-40df-a85e-4499f074b0cf", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPatch, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Call the GetGroup function
	err = client.UpdateGroup("370e64cb-da5a-40df-a85e-4499f074b0cf", &models.UpdateGroupRequest{Name: "TF_WORKSPACE_POSTMAN"})

	// Check the result
	assert.NoError(t, err)
}

// TestUpdateGroupUser_User is a unit test function that tests the UpdateGroupUser method of the Client struct.
// It verifies that the correct request is made to update a user in a group and checks the response status code.
// The test server is created to mock the API endpoint, and the client is created with the test server URL.
// The group user access details are prepared, and the UpdateGroupUser method is called.
// Finally, it asserts that there is no error returned from the method.
func TestUpdateGroupUser_User(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/ac653691-1af8-4be1-8468-9d73cdcc1250/users", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPut, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Prepare Request
	groupUserAccess := &models.GroupUser{
		EmailAddress:         "cedric.ghiot@weareretail.ai", //"john.doe@example.com",
		GroupUserAccessRight: models.GroupUserAccessRightAdmin,
		PrincipalType:        models.PrincipalTypeUser,
	}

	// Call the GetGroup function
	err = client.UpdateGroupUser("ac653691-1af8-4be1-8468-9d73cdcc1250", groupUserAccess)

	// Check the result
	assert.NoError(t, err)
}

// TestUpdateGroupUser_Group is a unit test function that tests the UpdateGroupUser method of the Client struct.
// It verifies that the correct HTTP request is made to update a group user's access rights in the Power BI API.
// The function creates a test server, sets up the necessary mock response, and checks the request URL and method.
// It then creates a client with the test server URL, prepares the request payload, and calls the UpdateGroupUser method.
// Finally, it asserts that no error occurred during the API call.
func TestUpdateGroupUser_Group(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/v1.0/myorg/groups/ac653691-1af8-4be1-8468-9d73cdcc1250/users", r.URL.Path)

		// Check the request method
		assert.Equal(t, http.MethodPut, r.Method)

		// Send a mock response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a client with the test server URL
	host := server.URL
	//host = "https://api.powerbi.com"
	client, err := NewClient(host)
	assert.NoError(t, err)

	// Prepare Request
	groupUserAccess := &models.GroupUser{
		Identifier:           "796131c3-8d85-44e1-bdfc-88ad8ba46520",
		GroupUserAccessRight: models.GroupUserAccessRightAdmin,
		PrincipalType:        models.PrincipalTypeGroup,
	}

	// Call the GetGroup function
	err = client.UpdateGroupUser("ac653691-1af8-4be1-8468-9d73cdcc1250", groupUserAccess)

	// Check the result
	assert.NoError(t, err)
}
