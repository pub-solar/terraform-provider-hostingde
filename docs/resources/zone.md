---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hostingde_zone Resource - hostingde"
subcategory: ""
description: |-
  
---

# hostingde_zone (Resource)



## Example Usage

```terraform
# Manage example DNS zone.
resource "hostingde_zone" "sample" {
  name = "example.test"
  type = "NATIVE"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Domain name (top-level domain) of the zone.
- `type` (String) The zone type. Valid types are NATIVE, MASTER, and SLAVE. Changing this forces re-creation of the zone.

### Optional

- `email` (String) The hostmaster email address. Only relevant if the type is NATIVE or MASTER. If the field is left empty, the default is hostmaster@name.

### Read-Only

- `id` (String) Numeric identifier of the zone.

## Import

Import is supported using the following syntax:

```shell
# DNS zone can be imported by specifying the zone id.
terraform import hostingde_zone.example 171029aw8802239
```
