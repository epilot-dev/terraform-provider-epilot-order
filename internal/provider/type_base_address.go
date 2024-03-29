// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type BaseAddress struct {
	ID             types.String   `tfsdk:"id"`
	Tags           []types.String `tfsdk:"tags"`
	AdditionalInfo types.String   `tfsdk:"additional_info"`
	City           types.String   `tfsdk:"city"`
	Country        types.String   `tfsdk:"country"`
	PostalCode     types.String   `tfsdk:"postal_code"`
	Street         types.String   `tfsdk:"street"`
	StreetNumber   types.String   `tfsdk:"street_number"`
}
