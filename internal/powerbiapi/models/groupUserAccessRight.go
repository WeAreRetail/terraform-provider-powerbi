package models

type GroupUserAccessRight string

const (
	GroupUserAccessRightAdmin       GroupUserAccessRight = "Admin"
	GroupUserAccessRightContributor GroupUserAccessRight = "Contributor"
	GroupUserAccessRightMember      GroupUserAccessRight = "Member"
	GroupUserAccessRightNone        GroupUserAccessRight = "None"
	GroupUserAccessRightViewer      GroupUserAccessRight = "Viewer"
)
