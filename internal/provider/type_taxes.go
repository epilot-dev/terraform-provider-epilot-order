// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Taxes struct {
	Amount types.Number `tfsdk:"amount"`
	Tax    []Tax        `tfsdk:"tax"`
}
