package models

// ServicePrincipalProfile is a struct that represents a Power BI service principal profile. Only relevant for Power BI Embedded multi-tenancy solution.
type ServicePrincipalProfile struct {
	DisplayName string `json:"displayName"` // The service principal profile name
	Id          string `json:"id"`          // The service principal profile ID
}
