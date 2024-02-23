---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "powerbi_workspace Data Source - terraform-provider-pbi"
subcategory: ""
description: |-
  Power BI workspace data source
---

# powerbi_workspace (Data Source)

Power BI workspace data source

## Example Usage

```terraform
data "powerbi_workspace" "example_id" {
  id = "ac653691-1af8-4be1-8468-9d73cdcc1250"
}

data "powerbi_workspace" "example_name" {
  name = "UNIT_TEST"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the workspace
- `name` (String) The name of the workspace

### Read-Only

- `is_on_dedicated_capacity` (Boolean) Indicates whether the workspace is on dedicated capacity
- `is_read_only` (Boolean) Indicates whether the workspace is read-only