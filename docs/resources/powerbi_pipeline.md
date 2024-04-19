---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "powerbi_pipeline Resource - terraform-provider-pbi"
subcategory: ""
description: |-
  Power BI pipeline resource
---

# powerbi_pipeline (Resource)

Power BI pipeline resource

## Example Usage

```terraform
resource "powerbi_pipeline" "example" {
  display_name = "TF_PIPELINE_TEST"
  description  = "TEST PIPELINE VIA TERRAFORM PROVIDER"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The name of the pipeline

### Optional

- `description` (String) The description of the pipeline

### Read-Only

- `id` (String) The id of the pipeline
- `stages` (List of Object) The stages of the pipeline (see [below for nested schema](#nestedatt--stages))

<a id="nestedatt--stages"></a>
### Nested Schema for `stages`

Read-Only:

- `order` (Number)
- `workspace_id` (String)
- `workspace_name` (String)