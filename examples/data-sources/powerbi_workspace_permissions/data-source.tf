data "powerbi_workspace_permissions" "example_id" {
  workspace_id = "6ac9aad1-88c9-47d1-baa2-6c4d469fe7d4"
}

output "permissions" {
  value = data.powerbi_workspace_permissions.example_id.permissions
}
