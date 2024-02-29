package models

// GroupUserAccessRight is a type that represents the access rights of a user in a Power BI workspace (group).
type GroupUserAccessRight string

const (
	GroupUserAccessRightAdmin       GroupUserAccessRight = "Admin"
	GroupUserAccessRightContributor GroupUserAccessRight = "Contributor"
	GroupUserAccessRightMember      GroupUserAccessRight = "Member"
	GroupUserAccessRightNone        GroupUserAccessRight = "None"
	GroupUserAccessRightViewer      GroupUserAccessRight = "Viewer"
)
