package models

import "fmt"

// GroupUser contains all the available properties for a group user.
// But, when assigning a user to a group, the API only accepts an object with the required properties.
// Adding the other properties, even null or empty, to the object will result in an error.
// So, we create separate structs for the different types of group users.

type GroupUser struct {
	DisplayName          string                  `json:"displayName"`
	EmailAddress         string                  `json:"emailAddress"`
	GraphId              string                  `json:"graphId"`
	GroupUserAccessRight GroupUserAccessRight    `json:"groupUserAccessRight"`
	Identifier           string                  `json:"identifier"`
	PrincipalType        PrincipalType           `json:"principalType"`
	Profile              ServicePrincipalProfile `json:"profile"`
	UserType             string                  `json:"userType"`
}

type GroupUserEmail struct {
	EmailAddress         string               `json:"emailAddress"`
	GroupUserAccessRight GroupUserAccessRight `json:"groupUserAccessRight"`
	PrincipalType        PrincipalType        `json:"principalType"`
}

type GroupUserGroup struct {
	GroupUserAccessRight GroupUserAccessRight `json:"groupUserAccessRight"`
	Identifier           string               `json:"identifier"`
	PrincipalType        PrincipalType        `json:"principalType"`
}

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
