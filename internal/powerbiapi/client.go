package powerbiapi

import (
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/go-resty/resty/v2"
)

// BaseURL - Default Power BI URL
const BaseURL string = "https://api.powerbi.com"

// Client - Power BI API client
type Client struct {
	BaseURL     string
	RestyClient *resty.Client
	Credentials *azidentity.DefaultAzureCredential
}

// NewClient creates a new instance of the Client struct.
// It takes a host URL as a parameter and returns a pointer to the Client and an error.
// If the host URL is provided, it will be used as the BaseURL for the Client.
// If the host URL is not provided, the default BaseURL will be used.
// It also retrieves a token using the GetToken method and assigns it to the Client's Token field.
// If an error occurs while retrieving the token, an error is returned.
func NewClient(host string) (*Client, error) {

	var err error

	c := Client{
		BaseURL:     BaseURL,
		RestyClient: resty.New(),
	}

	if host != "" {
		c.BaseURL = host
	}

	c.RestyClient.SetBaseURL(c.BaseURL).AddRetryCondition(
		func(r *resty.Response, err error) bool {
			// Including "err != nil" emulates the default retry behavior for errors encountered during the request.
			return err != nil || r.StatusCode() == http.StatusTooManyRequests
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the client: %v", err)
	}

	return &c, nil
}

// prepRequest - Prepares a request for the Power BI API
// It sets the global request parameters and returns a pointer to a resty.Request.
// It returns a pointer to a resty.Request and an error.
func (c *Client) prepRequest() (*resty.Request, error) {
	token, err := c.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token while preparing the request: %v", err)
	}
	return c.RestyClient.R().SetAuthToken(token), nil
}
