// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-order/internal/sdk/pkg/utils"
)

// BaseEntityACL - Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
type BaseEntityACL struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	Delete               []string    `json:"delete,omitempty"`
	Edit                 []string    `json:"edit,omitempty"`
	View                 []string    `json:"view,omitempty"`
}

func (b BaseEntityACL) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BaseEntityACL) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BaseEntityACL) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *BaseEntityACL) GetDelete() []string {
	if o == nil {
		return nil
	}
	return o.Delete
}

func (o *BaseEntityACL) GetEdit() []string {
	if o == nil {
		return nil
	}
	return o.Edit
}

func (o *BaseEntityACL) GetView() []string {
	if o == nil {
		return nil
	}
	return o.View
}
