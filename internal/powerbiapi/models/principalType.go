package models

type PrincipalType string

const (
	PrincipalTypeApp   PrincipalType = "App"
	PrincipalTypeGroup PrincipalType = "Group"
	PrincipalTypeNone  PrincipalType = "None"
	PrincipalTypeUser  PrincipalType = "User"
)
