package powerbiapi

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// scopes - Power BI API scopes
var scopes = []string{"https://analysis.windows.net/powerbi/api/.default"}

// Authenticate - Authenticates the client
func (c *Client) Authenticate() error {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("failed to get credentials: %v", err)
	}

	c.Credentials = creds

	return nil
}

// GetToken retrieves an access token for the Power BI API.
// It uses the default Azure credentials to authenticate and obtain the token.
// Returns the access token as a string or an error if the token retrieval fails.
func (c *Client) GetToken() (string, error) {

	var err error

	creds := c.Credentials
	if creds == nil {
		err = c.Authenticate()
		if err != nil {
			return "", fmt.Errorf("failed to authenticate: %v", err)
		}
		creds = c.Credentials
	}

	token, err := creds.GetToken(context.TODO(), policy.TokenRequestOptions{Scopes: scopes})
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	return token.Token, nil
}
