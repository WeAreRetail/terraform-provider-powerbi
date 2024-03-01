package models

import "fmt"

// GroupUser contains all the available properties for a group user.
// But, when assigning a user to a group, the API only accepts an object with the required properties.
// Adding the other properties, even null or empty, to the object will result in an error.
// So, we create separate structs for the different types of group users.

type GroupUser struct {
	DisplayName          string                  `json:"displayName"`          // Display name of the principal.
	EmailAddress         string                  `json:"emailAddress"`         // Email address of the user.
	GraphId              string                  `json:"graphId"`              // Identifier of the principal in Microsoft Graph. Only available for admin APIs.
	GroupUserAccessRight GroupUserAccessRight    `json:"groupUserAccessRight"` // The access right (permission level) that a user has on the workspace.
	Identifier           string                  `json:"identifier"`           // Identifier of the principal.
	PrincipalType        PrincipalType           `json:"principalType"`        // The principal type.
	Profile              ServicePrincipalProfile `json:"profile"`              // A Power BI service principal profile. Only relevant for Power BI Embedded multi-tenancy solution.
	UserType             string                  `json:"userType"`             // Type of the user.
}

// GroupUserEmail is a structure with the required properties to assign a user to a workspace (group).
type GroupUserEmail struct {
	EmailAddress         string               `json:"emailAddress"`         // Email address of the user.
	GroupUserAccessRight GroupUserAccessRight `json:"groupUserAccessRight"` // The access right (permission level) that a user has on the workspace.
	PrincipalType        PrincipalType        `json:"principalType"`        // The principal type.
}

// GroupUserGroup is a structure with the required properties to assign a group or app to a workspace (group).
type GroupUserGroup struct {
	Identifier           string               `json:"identifier"`           // Identifier of the principal.
	GroupUserAccessRight GroupUserAccessRight `json:"groupUserAccessRight"` // The access right (permission level) that a user has on the workspace.
	PrincipalType        PrincipalType        `json:"principalType"`        // The principal type.
}

// Validate checks the properties of the GroupUser and returns the correct struct for the API call.
func (g *GroupUser) Validate() (interface{}, error) {

	if g.EmailAddress != "" {
		if g.PrincipalType != PrincipalTypeUser {
			return nil, fmt.Errorf("email address is only valid for users")
		}
		return GroupUserEmail{
			EmailAddress:         g.EmailAddress,
			GroupUserAccessRight: g.GroupUserAccessRight,
			PrincipalType:        g.PrincipalType,
		}, nil
	}

	if g.Identifier != "" {
		if g.PrincipalType != PrincipalTypeGroup && g.PrincipalType != PrincipalTypeApp {
			return nil, fmt.Errorf("identifier is only valid for groups or apps")
		}
		return GroupUserGroup{
			GroupUserAccessRight: g.GroupUserAccessRight,
			Identifier:           g.Identifier,
			PrincipalType:        g.PrincipalType,
		}, nil

	}
	return nil, nil
}
