package models

// AzureResource represents a response detailing a user-owned Azure resource such as a Log Analytics workspace.
type AzureResource struct {
	Id             string `json:"id"`             // An identifier for the resource within Power BI.
	ResourceGroup  string `json:"resourceGroup"`  // The resource group within the subscription where the resource resides.
	ResourceName   string `json:"resourceName"`   // The name of the resource.
	SubscriptionId string `json:"subscriptionId"` // The Azure subscription where the resource resides.
}
