---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-order_order Data Source - terraform-provider-epilot-order"
subcategory: ""
description: |-
  Order DataSource
---

# epilot-order_order (Data Source)

Order DataSource

## Example Usage

```terraform
data "epilot-order_order" "my_order" {
  hydrate = false
  id      = "123e4567-e89b-12d3-a456-426614174000"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The order id

### Optional

- `hydrate` (Boolean) Hydrates entities in relations when passed true

### Read-Only

- `acl` (Attributes) Access control list (ACL) for an entity. Defines sharing access to external orgs or users. (see [below for nested schema](#nestedatt--acl))
- `additional_addresses` (Attributes List) Addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--additional_addresses))
- `billing_address` (Attributes List) Addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--billing_address))
- `billing_contact` (Attributes) (see [below for nested schema](#nestedatt--billing_contact))
- `created_at` (String)
- `customer` (Attributes) (see [below for nested schema](#nestedatt--customer))
- `delivery_address` (Attributes List) Addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--delivery_address))
- `expires_at` (String)
- `mapped_entities` (Attributes) (see [below for nested schema](#nestedatt--mapped_entities))
- `order_number` (String)
- `org` (String) Organization Id the entity belongs to
- `owners` (Attributes List) (see [below for nested schema](#nestedatt--owners))
- `prices` (Attributes) (see [below for nested schema](#nestedatt--prices))
- `products` (Attributes) (see [below for nested schema](#nestedatt--products))
- `schema` (String)
- `source` (Attributes) (see [below for nested schema](#nestedatt--source))
- `source_type` (String) Default: "manual"
- `status` (String) must be one of ["draft", "quote", "placed", "complete", "cancelled", "open_for_acceptance"]; Default: "draft"
- `tags` (List of String)
- `title` (String)
- `totals_detail` (Attributes) (see [below for nested schema](#nestedatt--totals_detail))
- `updated_at` (String)

<a id="nestedatt--acl"></a>
### Nested Schema for `acl`

Read-Only:

- `additional_properties` (String) Parsed as JSON.
- `delete` (List of String)
- `edit` (List of String)
- `view` (List of String)


<a id="nestedatt--additional_addresses"></a>
### Nested Schema for `additional_addresses`

Read-Only:

- `additional_info` (String)
- `city` (String)
- `country` (String) must be one of ["DE", "AT", "CH"]
- `id` (String)
- `postal_code` (String)
- `street` (String)
- `street_number` (String)
- `tags` (List of String)


<a id="nestedatt--billing_address"></a>
### Nested Schema for `billing_address`

Read-Only:

- `additional_info` (String)
- `city` (String)
- `country` (String) must be one of ["DE", "AT", "CH"]
- `id` (String)
- `postal_code` (String)
- `street` (String)
- `street_number` (String)
- `tags` (List of String)


<a id="nestedatt--billing_contact"></a>
### Nested Schema for `billing_contact`

Read-Only:

- `dollar_relation` (Attributes List) (see [below for nested schema](#nestedatt--billing_contact--dollar_relation))

<a id="nestedatt--billing_contact--dollar_relation"></a>
### Nested Schema for `billing_contact.dollar_relation`

Read-Only:

- `entity_id` (String)
- `tags` (List of String)



<a id="nestedatt--customer"></a>
### Nested Schema for `customer`

Read-Only:

- `dollar_relation` (Attributes List) (see [below for nested schema](#nestedatt--customer--dollar_relation))

<a id="nestedatt--customer--dollar_relation"></a>
### Nested Schema for `customer.dollar_relation`

Read-Only:

- `entity_id` (String)
- `tags` (List of String)



<a id="nestedatt--delivery_address"></a>
### Nested Schema for `delivery_address`

Read-Only:

- `additional_info` (String)
- `city` (String)
- `country` (String) must be one of ["DE", "AT", "CH"]
- `id` (String)
- `postal_code` (String)
- `street` (String)
- `street_number` (String)
- `tags` (List of String)


<a id="nestedatt--mapped_entities"></a>
### Nested Schema for `mapped_entities`

Read-Only:

- `dollar_relation` (Attributes List) (see [below for nested schema](#nestedatt--mapped_entities--dollar_relation))

<a id="nestedatt--mapped_entities--dollar_relation"></a>
### Nested Schema for `mapped_entities.dollar_relation`

Read-Only:

- `entity_id` (String)
- `tags` (List of String)



<a id="nestedatt--owners"></a>
### Nested Schema for `owners`

Read-Only:

- `org_id` (String)
- `user_id` (String)


<a id="nestedatt--prices"></a>
### Nested Schema for `prices`

Read-Only:

- `dollar_relation` (Attributes List) (see [below for nested schema](#nestedatt--prices--dollar_relation))

<a id="nestedatt--prices--dollar_relation"></a>
### Nested Schema for `prices.dollar_relation`

Read-Only:

- `entity_id` (String)
- `tags` (List of String)



<a id="nestedatt--products"></a>
### Nested Schema for `products`

Read-Only:

- `dollar_relation` (Attributes List) (see [below for nested schema](#nestedatt--products--dollar_relation))

<a id="nestedatt--products--dollar_relation"></a>
### Nested Schema for `products.dollar_relation`

Read-Only:

- `entity_id` (String)
- `tags` (List of String)



<a id="nestedatt--source"></a>
### Nested Schema for `source`

Read-Only:

- `href` (String) Default: null
- `title` (String) Default: "manual"


<a id="nestedatt--totals_detail"></a>
### Nested Schema for `totals_detail`

Read-Only:

- `amount_tax` (Number)
- `breakdown` (Attributes) (see [below for nested schema](#nestedatt--totals_detail--breakdown))

<a id="nestedatt--totals_detail--breakdown"></a>
### Nested Schema for `totals_detail.breakdown`

Read-Only:

- `recurrences` (Attributes List) (see [below for nested schema](#nestedatt--totals_detail--breakdown--recurrences))
- `taxes` (Attributes List) (see [below for nested schema](#nestedatt--totals_detail--breakdown--taxes))

<a id="nestedatt--totals_detail--breakdown--recurrences"></a>
### Nested Schema for `totals_detail.breakdown.recurrences`

Read-Only:

- `amount_subtotal` (Number)
- `amount_subtotal_decimal` (String)
- `amount_tax` (Number)
- `amount_total` (Number)
- `amount_total_decimal` (String)
- `type` (String) One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase. must be one of ["one_time", "recurring"]; Default: "one_time"


<a id="nestedatt--totals_detail--breakdown--taxes"></a>
### Nested Schema for `totals_detail.breakdown.taxes`

Read-Only:

- `amount` (Number)
- `tax` (Attributes List) (see [below for nested schema](#nestedatt--totals_detail--breakdown--taxes--tax))

<a id="nestedatt--totals_detail--breakdown--taxes--tax"></a>
### Nested Schema for `totals_detail.breakdown.taxes.tax`

Read-Only:

- `id` (String)
- `rate` (String)
- `type` (String) must be one of ["VAT", "Custom"]

