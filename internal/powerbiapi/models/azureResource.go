package models

// AzureResource.
type AzureResource struct {
	Id             string `json:"id"`
	ResourceGroup  string `json:"resourceGroup"`
	ResourceName   string `json:"resourceName"`
	SubscriptionId string `json:"subscriptionId"`
}
