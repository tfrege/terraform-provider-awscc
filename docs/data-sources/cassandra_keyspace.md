---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cassandra_keyspace Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Cassandra::Keyspace
---

# awscc_cassandra_keyspace (Data Source)

Data Source schema for AWS::Cassandra::Keyspace



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `keyspace_name` (String) Name for Cassandra keyspace
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
