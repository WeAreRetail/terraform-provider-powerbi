package provider

import "terraform-provider-powerbi/internal/powerbiapi"

// getClient returns a new instance of the powerbiapi.Client with the specified base URL.
// The base URL is used to establish the connection to the Power BI service.
func getClient(baseUrl string) (*powerbiapi.Client, error) {
	return powerbiapi.NewClient(baseUrl)
}
