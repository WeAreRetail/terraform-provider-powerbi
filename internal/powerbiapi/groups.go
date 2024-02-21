package powerbiapi

import (
	"fmt"
	"strconv"
	"terraform-provider-powerbi/internal/powerbiapi/models"
)

// AddGroupUser adds a user/group/app to a group.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/add-group-user
func (c *Client) AddGroupUser(groupId string, groupUserAccessRight *models.GroupUser) error {
	// POST https://api.powerbi.com/v1.0/myorg/groups/{groupId}/users

	client, err := c.prepRequest()
	if err != nil {
		return fmt.Errorf("failed to prepare the request for AddGroupUser: %v", err)
	}

	body, err := groupUserAccessRight.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate group user: %v", err)
	}

	resp, err := client.SetBody(body).Post(fmt.Sprintf("/v1.0/myorg/groups/%s/users", groupId))
	if err != nil {
		return fmt.Errorf("failed to add group user: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to add group user: %v", resp.Error())
	}

	return nil
}

// CreateGroup creates a new group.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/create-group
func (c *Client) CreateGroup(groupName string) (*models.Group, error) {
	// POST https://api.powerbi.com/v1.0/myorg/groups

	var err error
	group := &models.Group{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for CreateGroup: %v", err)
	}

	resp, err := client.SetResult(group).
		SetQueryParam("workspaceV2", "True").
		SetBody(&models.GroupCreationRequest{Name: groupName}).
		Post("/v1.0/myorg/groups")
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to create group: %v", resp.Error())
	}

	return group, nil
}

// DeleteGroup deletes a group by its ID.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/delete-group
func (c *Client) DeleteGroup(groupId string) error {
	// DELETE https://api.powerbi.com/v1.0/myorg/groups/{groupId}
	var err error

	client, err := c.prepRequest()
	if err != nil {
		return fmt.Errorf("failed to prepare the request for DeleteGroup: %v", err)
	}

	resp, err := client.Delete(fmt.Sprintf("/v1.0/myorg/groups/%s", groupId))
	if err != nil {
		return fmt.Errorf("failed to delete group: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to delete group: %v", resp.Error())
	}

	return nil
}

// DeleteUserGroup deletes a user from a group.
// user is the email address of the user or object ID of the service principal to delete
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/delete-user-in-group
func (c *Client) DeleteUserGroup(groupId string, user string) error {

	client, err := c.prepRequest()
	if err != nil {
		return fmt.Errorf("failed to prepare the request for DeleteUserGroup: %v", err)
	}

	resp, err := client.Delete(fmt.Sprintf("/v1.0/myorg/groups/%s/users/%s", groupId, user))
	if err != nil {
		return fmt.Errorf("failed to delete user from group: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to delete user from group: %v", resp.Error())
	}

	return nil
}

// GetGroup retrieves a group by its ID.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/get-group
func (c *Client) GetGroup(groupId string) (*models.Group, error) {
	// GET https://api.powerbi.com/v1.0/myorg/groups/{groupId}

	var err error
	group := &models.Group{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for GetGroups: %v", err)
	}

	resp, err := client.SetResult(group).Get(fmt.Sprintf("/v1.0/myorg/groups/%s", groupId))
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get group: %v", resp.Error())
	}

	return group, nil
}

// GetGroupUsers retrieves a list of users, groups, and service principals in a group.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/get-group-users
func (c *Client) GetGroupUsers(groupId string) (*models.GroupUsers, error) {
	// GET https://api.powerbi.com/v1.0/myorg/groups/{groupId}/users

	var err error
	groupUsers := &models.GroupUsers{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for GetGroupUsers: %v", err)
	}

	resp, err := client.SetResult(groupUsers).Get(fmt.Sprintf("/v1.0/myorg/groups/%s/users", groupId))
	if err != nil {
		return nil, fmt.Errorf("failed to get group users: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get group users: %v", resp.Error())
	}

	return groupUsers, nil
}

// GetGroups retrieves a list of groups.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/get-groups
func (c *Client) GetGroups(filter string, top int, skip int) (*models.Groups, error) {
	// GET https://api.powerbi.com/v1.0/myorg/groups

	var err error
	groups := &models.Groups{}

	client, err := c.prepRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request for GetGroups: %v", err)
	}

	if filter != "" {
		client.SetQueryParam("$filter", filter)
	}
	if top > 0 {
		client.SetQueryParam("$top", strconv.Itoa(top))
	}
	if skip > 0 {
		client.SetQueryParam("$skip", strconv.Itoa(skip))
	}

	resp, err := client.SetResult(&groups).Get("/v1.0/myorg/groups")
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get groups: %v", resp.Error())
	}

	return groups, nil
}

// UpdateGroup updates a specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/update-group
func (c *Client) UpdateGroup(groupId string, updateGroupRequest *models.UpdateGroupRequest) error {
	// PATCH https://api.powerbi.com/v1.0/myorg/groups/{groupId}

	var err error

	client, err := c.prepRequest()
	if err != nil {
		return fmt.Errorf("failed to prepare the request for GetGroups: %v", err)
	}

	body := updateGroupRequest.Validate()

	resp, err := client.
		SetBody(body).
		Patch(fmt.Sprintf("/v1.0/myorg/groups/%s", groupId))
	if err != nil {
		return fmt.Errorf("failed to update groups: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to get groups: [%v] %s", resp.StatusCode(), resp.String())
	}

	return nil
}

// UpdateGroupUser updates the specified user permissions to the specified workspace.
// https://learn.microsoft.com/en-us/rest/api/power-bi/groups/update-group-user
func (c *Client) UpdateGroupUser(groupId string, groupUserAccessRight *models.GroupUser) error {
	// PUT https://api.powerbi.com/v1.0/myorg/groups/{groupId}/users

	var err error

	client, err := c.prepRequest()
	if err != nil {
		return fmt.Errorf("failed to prepare the request for GetGroups: %v", err)
	}

	body, err := groupUserAccessRight.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate group user: %v", err)
	}

	resp, err := client.
		SetBody(body).
		Put(fmt.Sprintf("/v1.0/myorg/groups/%s/users", groupId))
	if err != nil {
		return fmt.Errorf("failed to update groups users: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to get groups users: %v", resp.Error())
	}

	return nil
}
