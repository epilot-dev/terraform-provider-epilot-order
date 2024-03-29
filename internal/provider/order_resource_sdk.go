// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-order/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"time"
)

func (r *OrderResourceModel) ToSharedOrderCreate() *shared.OrderCreate {
	var additionalAddresses []shared.BaseAddress = nil
	for _, additionalAddressesItem := range r.AdditionalAddresses {
		id := new(string)
		if !additionalAddressesItem.ID.IsUnknown() && !additionalAddressesItem.ID.IsNull() {
			*id = additionalAddressesItem.ID.ValueString()
		} else {
			id = nil
		}
		var tags []string = nil
		for _, tagsItem := range additionalAddressesItem.Tags {
			tags = append(tags, tagsItem.ValueString())
		}
		additionalInfo := new(string)
		if !additionalAddressesItem.AdditionalInfo.IsUnknown() && !additionalAddressesItem.AdditionalInfo.IsNull() {
			*additionalInfo = additionalAddressesItem.AdditionalInfo.ValueString()
		} else {
			additionalInfo = nil
		}
		city := new(string)
		if !additionalAddressesItem.City.IsUnknown() && !additionalAddressesItem.City.IsNull() {
			*city = additionalAddressesItem.City.ValueString()
		} else {
			city = nil
		}
		country := new(shared.Country)
		if !additionalAddressesItem.Country.IsUnknown() && !additionalAddressesItem.Country.IsNull() {
			*country = shared.Country(additionalAddressesItem.Country.ValueString())
		} else {
			country = nil
		}
		postalCode := new(string)
		if !additionalAddressesItem.PostalCode.IsUnknown() && !additionalAddressesItem.PostalCode.IsNull() {
			*postalCode = additionalAddressesItem.PostalCode.ValueString()
		} else {
			postalCode = nil
		}
		street := new(string)
		if !additionalAddressesItem.Street.IsUnknown() && !additionalAddressesItem.Street.IsNull() {
			*street = additionalAddressesItem.Street.ValueString()
		} else {
			street = nil
		}
		streetNumber := new(string)
		if !additionalAddressesItem.StreetNumber.IsUnknown() && !additionalAddressesItem.StreetNumber.IsNull() {
			*streetNumber = additionalAddressesItem.StreetNumber.ValueString()
		} else {
			streetNumber = nil
		}
		additionalAddresses = append(additionalAddresses, shared.BaseAddress{
			ID:             id,
			Tags:           tags,
			AdditionalInfo: additionalInfo,
			City:           city,
			Country:        country,
			PostalCode:     postalCode,
			Street:         street,
			StreetNumber:   streetNumber,
		})
	}
	var billingAddress []shared.BaseAddress = nil
	for _, billingAddressItem := range r.BillingAddress {
		id1 := new(string)
		if !billingAddressItem.ID.IsUnknown() && !billingAddressItem.ID.IsNull() {
			*id1 = billingAddressItem.ID.ValueString()
		} else {
			id1 = nil
		}
		var tags1 []string = nil
		for _, tagsItem1 := range billingAddressItem.Tags {
			tags1 = append(tags1, tagsItem1.ValueString())
		}
		additionalInfo1 := new(string)
		if !billingAddressItem.AdditionalInfo.IsUnknown() && !billingAddressItem.AdditionalInfo.IsNull() {
			*additionalInfo1 = billingAddressItem.AdditionalInfo.ValueString()
		} else {
			additionalInfo1 = nil
		}
		city1 := new(string)
		if !billingAddressItem.City.IsUnknown() && !billingAddressItem.City.IsNull() {
			*city1 = billingAddressItem.City.ValueString()
		} else {
			city1 = nil
		}
		country1 := new(shared.Country)
		if !billingAddressItem.Country.IsUnknown() && !billingAddressItem.Country.IsNull() {
			*country1 = shared.Country(billingAddressItem.Country.ValueString())
		} else {
			country1 = nil
		}
		postalCode1 := new(string)
		if !billingAddressItem.PostalCode.IsUnknown() && !billingAddressItem.PostalCode.IsNull() {
			*postalCode1 = billingAddressItem.PostalCode.ValueString()
		} else {
			postalCode1 = nil
		}
		street1 := new(string)
		if !billingAddressItem.Street.IsUnknown() && !billingAddressItem.Street.IsNull() {
			*street1 = billingAddressItem.Street.ValueString()
		} else {
			street1 = nil
		}
		streetNumber1 := new(string)
		if !billingAddressItem.StreetNumber.IsUnknown() && !billingAddressItem.StreetNumber.IsNull() {
			*streetNumber1 = billingAddressItem.StreetNumber.ValueString()
		} else {
			streetNumber1 = nil
		}
		billingAddress = append(billingAddress, shared.BaseAddress{
			ID:             id1,
			Tags:           tags1,
			AdditionalInfo: additionalInfo1,
			City:           city1,
			Country:        country1,
			PostalCode:     postalCode1,
			Street:         street1,
			StreetNumber:   streetNumber1,
		})
	}
	var billingContact *shared.BaseRelation
	if r.BillingContact != nil {
		var dollarRelation []shared.DollarRelation = nil
		for _, dollarRelationItem := range r.BillingContact.DollarRelation {
			var tags2 []string = nil
			for _, tagsItem2 := range dollarRelationItem.Tags {
				tags2 = append(tags2, tagsItem2.ValueString())
			}
			entityID := new(string)
			if !dollarRelationItem.EntityID.IsUnknown() && !dollarRelationItem.EntityID.IsNull() {
				*entityID = dollarRelationItem.EntityID.ValueString()
			} else {
				entityID = nil
			}
			dollarRelation = append(dollarRelation, shared.DollarRelation{
				Tags:     tags2,
				EntityID: entityID,
			})
		}
		billingContact = &shared.BaseRelation{
			DollarRelation: dollarRelation,
		}
	}
	var customer *shared.BaseRelation
	if r.Customer != nil {
		var dollarRelation1 []shared.DollarRelation = nil
		for _, dollarRelationItem1 := range r.Customer.DollarRelation {
			var tags3 []string = nil
			for _, tagsItem3 := range dollarRelationItem1.Tags {
				tags3 = append(tags3, tagsItem3.ValueString())
			}
			entityId1 := new(string)
			if !dollarRelationItem1.EntityID.IsUnknown() && !dollarRelationItem1.EntityID.IsNull() {
				*entityId1 = dollarRelationItem1.EntityID.ValueString()
			} else {
				entityId1 = nil
			}
			dollarRelation1 = append(dollarRelation1, shared.DollarRelation{
				Tags:     tags3,
				EntityID: entityId1,
			})
		}
		customer = &shared.BaseRelation{
			DollarRelation: dollarRelation1,
		}
	}
	var deliveryAddress []shared.BaseAddress = nil
	for _, deliveryAddressItem := range r.DeliveryAddress {
		id2 := new(string)
		if !deliveryAddressItem.ID.IsUnknown() && !deliveryAddressItem.ID.IsNull() {
			*id2 = deliveryAddressItem.ID.ValueString()
		} else {
			id2 = nil
		}
		var tags4 []string = nil
		for _, tagsItem4 := range deliveryAddressItem.Tags {
			tags4 = append(tags4, tagsItem4.ValueString())
		}
		additionalInfo2 := new(string)
		if !deliveryAddressItem.AdditionalInfo.IsUnknown() && !deliveryAddressItem.AdditionalInfo.IsNull() {
			*additionalInfo2 = deliveryAddressItem.AdditionalInfo.ValueString()
		} else {
			additionalInfo2 = nil
		}
		city2 := new(string)
		if !deliveryAddressItem.City.IsUnknown() && !deliveryAddressItem.City.IsNull() {
			*city2 = deliveryAddressItem.City.ValueString()
		} else {
			city2 = nil
		}
		country2 := new(shared.Country)
		if !deliveryAddressItem.Country.IsUnknown() && !deliveryAddressItem.Country.IsNull() {
			*country2 = shared.Country(deliveryAddressItem.Country.ValueString())
		} else {
			country2 = nil
		}
		postalCode2 := new(string)
		if !deliveryAddressItem.PostalCode.IsUnknown() && !deliveryAddressItem.PostalCode.IsNull() {
			*postalCode2 = deliveryAddressItem.PostalCode.ValueString()
		} else {
			postalCode2 = nil
		}
		street2 := new(string)
		if !deliveryAddressItem.Street.IsUnknown() && !deliveryAddressItem.Street.IsNull() {
			*street2 = deliveryAddressItem.Street.ValueString()
		} else {
			street2 = nil
		}
		streetNumber2 := new(string)
		if !deliveryAddressItem.StreetNumber.IsUnknown() && !deliveryAddressItem.StreetNumber.IsNull() {
			*streetNumber2 = deliveryAddressItem.StreetNumber.ValueString()
		} else {
			streetNumber2 = nil
		}
		deliveryAddress = append(deliveryAddress, shared.BaseAddress{
			ID:             id2,
			Tags:           tags4,
			AdditionalInfo: additionalInfo2,
			City:           city2,
			Country:        country2,
			PostalCode:     postalCode2,
			Street:         street2,
			StreetNumber:   streetNumber2,
		})
	}
	expiresAt := new(time.Time)
	if !r.ExpiresAt.IsUnknown() && !r.ExpiresAt.IsNull() {
		*expiresAt, _ = time.Parse(time.RFC3339Nano, r.ExpiresAt.ValueString())
	} else {
		expiresAt = nil
	}
	var mappedEntities *shared.BaseRelation
	if r.MappedEntities != nil {
		var dollarRelation2 []shared.DollarRelation = nil
		for _, dollarRelationItem2 := range r.MappedEntities.DollarRelation {
			var tags5 []string = nil
			for _, tagsItem5 := range dollarRelationItem2.Tags {
				tags5 = append(tags5, tagsItem5.ValueString())
			}
			entityId2 := new(string)
			if !dollarRelationItem2.EntityID.IsUnknown() && !dollarRelationItem2.EntityID.IsNull() {
				*entityId2 = dollarRelationItem2.EntityID.ValueString()
			} else {
				entityId2 = nil
			}
			dollarRelation2 = append(dollarRelation2, shared.DollarRelation{
				Tags:     tags5,
				EntityID: entityId2,
			})
		}
		mappedEntities = &shared.BaseRelation{
			DollarRelation: dollarRelation2,
		}
	}
	orderNumber := new(string)
	if !r.OrderNumber.IsUnknown() && !r.OrderNumber.IsNull() {
		*orderNumber = r.OrderNumber.ValueString()
	} else {
		orderNumber = nil
	}
	var prices *shared.BaseRelation
	if r.Prices != nil {
		var dollarRelation3 []shared.DollarRelation = nil
		for _, dollarRelationItem3 := range r.Prices.DollarRelation {
			var tags6 []string = nil
			for _, tagsItem6 := range dollarRelationItem3.Tags {
				tags6 = append(tags6, tagsItem6.ValueString())
			}
			entityId3 := new(string)
			if !dollarRelationItem3.EntityID.IsUnknown() && !dollarRelationItem3.EntityID.IsNull() {
				*entityId3 = dollarRelationItem3.EntityID.ValueString()
			} else {
				entityId3 = nil
			}
			dollarRelation3 = append(dollarRelation3, shared.DollarRelation{
				Tags:     tags6,
				EntityID: entityId3,
			})
		}
		prices = &shared.BaseRelation{
			DollarRelation: dollarRelation3,
		}
	}
	var products *shared.BaseRelation
	if r.Products != nil {
		var dollarRelation4 []shared.DollarRelation = nil
		for _, dollarRelationItem4 := range r.Products.DollarRelation {
			var tags7 []string = nil
			for _, tagsItem7 := range dollarRelationItem4.Tags {
				tags7 = append(tags7, tagsItem7.ValueString())
			}
			entityId4 := new(string)
			if !dollarRelationItem4.EntityID.IsUnknown() && !dollarRelationItem4.EntityID.IsNull() {
				*entityId4 = dollarRelationItem4.EntityID.ValueString()
			} else {
				entityId4 = nil
			}
			dollarRelation4 = append(dollarRelation4, shared.DollarRelation{
				Tags:     tags7,
				EntityID: entityId4,
			})
		}
		products = &shared.BaseRelation{
			DollarRelation: dollarRelation4,
		}
	}
	var source *shared.OrderCreateSource
	if r.Source != nil {
		href := new(string)
		if !r.Source.Href.IsUnknown() && !r.Source.Href.IsNull() {
			*href = r.Source.Href.ValueString()
		} else {
			href = nil
		}
		title := new(string)
		if !r.Source.Title.IsUnknown() && !r.Source.Title.IsNull() {
			*title = r.Source.Title.ValueString()
		} else {
			title = nil
		}
		source = &shared.OrderCreateSource{
			Href:  href,
			Title: title,
		}
	}
	sourceType := new(string)
	if !r.SourceType.IsUnknown() && !r.SourceType.IsNull() {
		*sourceType = r.SourceType.ValueString()
	} else {
		sourceType = nil
	}
	status := new(shared.OrderCreateStatus)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.OrderCreateStatus(r.Status.ValueString())
	} else {
		status = nil
	}
	var totalsDetail *shared.OrderCreateTotalsDetail
	if r.TotalsDetail != nil {
		amountTax := new(float64)
		if !r.TotalsDetail.AmountTax.IsUnknown() && !r.TotalsDetail.AmountTax.IsNull() {
			*amountTax, _ = r.TotalsDetail.AmountTax.ValueBigFloat().Float64()
		} else {
			amountTax = nil
		}
		var breakdown *shared.OrderCreateBreakdown
		if r.TotalsDetail.Breakdown != nil {
			var recurrences []shared.OrderCreateRecurrences = nil
			for _, recurrencesItem := range r.TotalsDetail.Breakdown.Recurrences {
				amountSubtotal := new(float64)
				if !recurrencesItem.AmountSubtotal.IsUnknown() && !recurrencesItem.AmountSubtotal.IsNull() {
					*amountSubtotal, _ = recurrencesItem.AmountSubtotal.ValueBigFloat().Float64()
				} else {
					amountSubtotal = nil
				}
				amountSubtotalDecimal := new(string)
				if !recurrencesItem.AmountSubtotalDecimal.IsUnknown() && !recurrencesItem.AmountSubtotalDecimal.IsNull() {
					*amountSubtotalDecimal = recurrencesItem.AmountSubtotalDecimal.ValueString()
				} else {
					amountSubtotalDecimal = nil
				}
				amountTax1 := new(float64)
				if !recurrencesItem.AmountTax.IsUnknown() && !recurrencesItem.AmountTax.IsNull() {
					*amountTax1, _ = recurrencesItem.AmountTax.ValueBigFloat().Float64()
				} else {
					amountTax1 = nil
				}
				amountTotal := new(float64)
				if !recurrencesItem.AmountTotal.IsUnknown() && !recurrencesItem.AmountTotal.IsNull() {
					*amountTotal, _ = recurrencesItem.AmountTotal.ValueBigFloat().Float64()
				} else {
					amountTotal = nil
				}
				amountTotalDecimal := new(string)
				if !recurrencesItem.AmountTotalDecimal.IsUnknown() && !recurrencesItem.AmountTotalDecimal.IsNull() {
					*amountTotalDecimal = recurrencesItem.AmountTotalDecimal.ValueString()
				} else {
					amountTotalDecimal = nil
				}
				typeVar := new(shared.OrderCreateType)
				if !recurrencesItem.Type.IsUnknown() && !recurrencesItem.Type.IsNull() {
					*typeVar = shared.OrderCreateType(recurrencesItem.Type.ValueString())
				} else {
					typeVar = nil
				}
				recurrences = append(recurrences, shared.OrderCreateRecurrences{
					AmountSubtotal:        amountSubtotal,
					AmountSubtotalDecimal: amountSubtotalDecimal,
					AmountTax:             amountTax1,
					AmountTotal:           amountTotal,
					AmountTotalDecimal:    amountTotalDecimal,
					Type:                  typeVar,
				})
			}
			var taxes []shared.OrderCreateTaxes = nil
			for _, taxesItem := range r.TotalsDetail.Breakdown.Taxes {
				amount := new(float64)
				if !taxesItem.Amount.IsUnknown() && !taxesItem.Amount.IsNull() {
					*amount, _ = taxesItem.Amount.ValueBigFloat().Float64()
				} else {
					amount = nil
				}
				var tax []shared.OrderCreateTax = nil
				for _, taxItem := range taxesItem.Tax {
					id3 := new(string)
					if !taxItem.ID.IsUnknown() && !taxItem.ID.IsNull() {
						*id3 = taxItem.ID.ValueString()
					} else {
						id3 = nil
					}
					rate := new(string)
					if !taxItem.Rate.IsUnknown() && !taxItem.Rate.IsNull() {
						*rate = taxItem.Rate.ValueString()
					} else {
						rate = nil
					}
					type1 := new(shared.OrderCreateSchemasType)
					if !taxItem.Type.IsUnknown() && !taxItem.Type.IsNull() {
						*type1 = shared.OrderCreateSchemasType(taxItem.Type.ValueString())
					} else {
						type1 = nil
					}
					tax = append(tax, shared.OrderCreateTax{
						ID:   id3,
						Rate: rate,
						Type: type1,
					})
				}
				taxes = append(taxes, shared.OrderCreateTaxes{
					Amount: amount,
					Tax:    tax,
				})
			}
			breakdown = &shared.OrderCreateBreakdown{
				Recurrences: recurrences,
				Taxes:       taxes,
			}
		}
		totalsDetail = &shared.OrderCreateTotalsDetail{
			AmountTax: amountTax,
			Breakdown: breakdown,
		}
	}
	out := shared.OrderCreate{
		AdditionalAddresses: additionalAddresses,
		BillingAddress:      billingAddress,
		BillingContact:      billingContact,
		Customer:            customer,
		DeliveryAddress:     deliveryAddress,
		ExpiresAt:           expiresAt,
		MappedEntities:      mappedEntities,
		OrderNumber:         orderNumber,
		Prices:              prices,
		Products:            products,
		Source:              source,
		SourceType:          sourceType,
		Status:              status,
		TotalsDetail:        totalsDetail,
	}
	return &out
}

func (r *OrderResourceModel) RefreshFromSharedOrder(resp *shared.Order) {
	if resp.ACL.AdditionalProperties == nil {
		r.ACL.AdditionalProperties = types.StringNull()
	} else {
		additionalPropertiesResult, _ := json.Marshal(resp.ACL.AdditionalProperties)
		r.ACL.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
	}
	r.ACL.Delete = nil
	for _, v := range resp.ACL.Delete {
		r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
	}
	r.ACL.Edit = nil
	for _, v := range resp.ACL.Edit {
		r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
	}
	r.ACL.View = nil
	for _, v := range resp.ACL.View {
		r.ACL.View = append(r.ACL.View, types.StringValue(v))
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	r.ID = types.StringValue(resp.ID)
	r.Org = types.StringValue(resp.Org)
	if len(r.Owners) > len(resp.Owners) {
		r.Owners = r.Owners[:len(resp.Owners)]
	}
	for ownersCount, ownersItem := range resp.Owners {
		var owners1 BaseEntityOwner
		owners1.OrgID = types.StringValue(ownersItem.OrgID)
		owners1.UserID = types.StringPointerValue(ownersItem.UserID)
		if ownersCount+1 > len(r.Owners) {
			r.Owners = append(r.Owners, owners1)
		} else {
			r.Owners[ownersCount].OrgID = owners1.OrgID
			r.Owners[ownersCount].UserID = owners1.UserID
		}
	}
	r.Schema = types.StringValue(resp.Schema)
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
	r.Title = types.StringValue(resp.Title)
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	if len(r.AdditionalAddresses) > len(resp.AdditionalAddresses) {
		r.AdditionalAddresses = r.AdditionalAddresses[:len(resp.AdditionalAddresses)]
	}
	for additionalAddressesCount, additionalAddressesItem := range resp.AdditionalAddresses {
		var additionalAddresses1 BaseAddress
		additionalAddresses1.ID = types.StringPointerValue(additionalAddressesItem.ID)
		additionalAddresses1.Tags = nil
		for _, v := range additionalAddressesItem.Tags {
			additionalAddresses1.Tags = append(additionalAddresses1.Tags, types.StringValue(v))
		}
		additionalAddresses1.AdditionalInfo = types.StringPointerValue(additionalAddressesItem.AdditionalInfo)
		additionalAddresses1.City = types.StringPointerValue(additionalAddressesItem.City)
		if additionalAddressesItem.Country != nil {
			additionalAddresses1.Country = types.StringValue(string(*additionalAddressesItem.Country))
		} else {
			additionalAddresses1.Country = types.StringNull()
		}
		additionalAddresses1.PostalCode = types.StringPointerValue(additionalAddressesItem.PostalCode)
		additionalAddresses1.Street = types.StringPointerValue(additionalAddressesItem.Street)
		additionalAddresses1.StreetNumber = types.StringPointerValue(additionalAddressesItem.StreetNumber)
		if additionalAddressesCount+1 > len(r.AdditionalAddresses) {
			r.AdditionalAddresses = append(r.AdditionalAddresses, additionalAddresses1)
		} else {
			r.AdditionalAddresses[additionalAddressesCount].ID = additionalAddresses1.ID
			r.AdditionalAddresses[additionalAddressesCount].Tags = additionalAddresses1.Tags
			r.AdditionalAddresses[additionalAddressesCount].AdditionalInfo = additionalAddresses1.AdditionalInfo
			r.AdditionalAddresses[additionalAddressesCount].City = additionalAddresses1.City
			r.AdditionalAddresses[additionalAddressesCount].Country = additionalAddresses1.Country
			r.AdditionalAddresses[additionalAddressesCount].PostalCode = additionalAddresses1.PostalCode
			r.AdditionalAddresses[additionalAddressesCount].Street = additionalAddresses1.Street
			r.AdditionalAddresses[additionalAddressesCount].StreetNumber = additionalAddresses1.StreetNumber
		}
	}
	if len(r.BillingAddress) > len(resp.BillingAddress) {
		r.BillingAddress = r.BillingAddress[:len(resp.BillingAddress)]
	}
	for billingAddressCount, billingAddressItem := range resp.BillingAddress {
		var billingAddress1 BaseAddress
		billingAddress1.ID = types.StringPointerValue(billingAddressItem.ID)
		billingAddress1.Tags = nil
		for _, v := range billingAddressItem.Tags {
			billingAddress1.Tags = append(billingAddress1.Tags, types.StringValue(v))
		}
		billingAddress1.AdditionalInfo = types.StringPointerValue(billingAddressItem.AdditionalInfo)
		billingAddress1.City = types.StringPointerValue(billingAddressItem.City)
		if billingAddressItem.Country != nil {
			billingAddress1.Country = types.StringValue(string(*billingAddressItem.Country))
		} else {
			billingAddress1.Country = types.StringNull()
		}
		billingAddress1.PostalCode = types.StringPointerValue(billingAddressItem.PostalCode)
		billingAddress1.Street = types.StringPointerValue(billingAddressItem.Street)
		billingAddress1.StreetNumber = types.StringPointerValue(billingAddressItem.StreetNumber)
		if billingAddressCount+1 > len(r.BillingAddress) {
			r.BillingAddress = append(r.BillingAddress, billingAddress1)
		} else {
			r.BillingAddress[billingAddressCount].ID = billingAddress1.ID
			r.BillingAddress[billingAddressCount].Tags = billingAddress1.Tags
			r.BillingAddress[billingAddressCount].AdditionalInfo = billingAddress1.AdditionalInfo
			r.BillingAddress[billingAddressCount].City = billingAddress1.City
			r.BillingAddress[billingAddressCount].Country = billingAddress1.Country
			r.BillingAddress[billingAddressCount].PostalCode = billingAddress1.PostalCode
			r.BillingAddress[billingAddressCount].Street = billingAddress1.Street
			r.BillingAddress[billingAddressCount].StreetNumber = billingAddress1.StreetNumber
		}
	}
	if resp.BillingContact == nil {
		r.BillingContact = nil
	} else {
		r.BillingContact = &BaseRelation{}
		if len(r.BillingContact.DollarRelation) > len(resp.BillingContact.DollarRelation) {
			r.BillingContact.DollarRelation = r.BillingContact.DollarRelation[:len(resp.BillingContact.DollarRelation)]
		}
		for dollarRelationCount, dollarRelationItem := range resp.BillingContact.DollarRelation {
			var dollarRelation1 DollarRelation
			dollarRelation1.Tags = nil
			for _, v := range dollarRelationItem.Tags {
				dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
			}
			dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
			if dollarRelationCount+1 > len(r.BillingContact.DollarRelation) {
				r.BillingContact.DollarRelation = append(r.BillingContact.DollarRelation, dollarRelation1)
			} else {
				r.BillingContact.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
				r.BillingContact.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
			}
		}
	}
	if resp.Customer == nil {
		r.Customer = nil
	} else {
		r.Customer = &BaseRelation{}
		if len(r.Customer.DollarRelation) > len(resp.Customer.DollarRelation) {
			r.Customer.DollarRelation = r.Customer.DollarRelation[:len(resp.Customer.DollarRelation)]
		}
		for dollarRelationCount1, dollarRelationItem1 := range resp.Customer.DollarRelation {
			var dollarRelation3 DollarRelation
			dollarRelation3.Tags = nil
			for _, v := range dollarRelationItem1.Tags {
				dollarRelation3.Tags = append(dollarRelation3.Tags, types.StringValue(v))
			}
			dollarRelation3.EntityID = types.StringPointerValue(dollarRelationItem1.EntityID)
			if dollarRelationCount1+1 > len(r.Customer.DollarRelation) {
				r.Customer.DollarRelation = append(r.Customer.DollarRelation, dollarRelation3)
			} else {
				r.Customer.DollarRelation[dollarRelationCount1].Tags = dollarRelation3.Tags
				r.Customer.DollarRelation[dollarRelationCount1].EntityID = dollarRelation3.EntityID
			}
		}
	}
	if len(r.DeliveryAddress) > len(resp.DeliveryAddress) {
		r.DeliveryAddress = r.DeliveryAddress[:len(resp.DeliveryAddress)]
	}
	for deliveryAddressCount, deliveryAddressItem := range resp.DeliveryAddress {
		var deliveryAddress1 BaseAddress
		deliveryAddress1.ID = types.StringPointerValue(deliveryAddressItem.ID)
		deliveryAddress1.Tags = nil
		for _, v := range deliveryAddressItem.Tags {
			deliveryAddress1.Tags = append(deliveryAddress1.Tags, types.StringValue(v))
		}
		deliveryAddress1.AdditionalInfo = types.StringPointerValue(deliveryAddressItem.AdditionalInfo)
		deliveryAddress1.City = types.StringPointerValue(deliveryAddressItem.City)
		if deliveryAddressItem.Country != nil {
			deliveryAddress1.Country = types.StringValue(string(*deliveryAddressItem.Country))
		} else {
			deliveryAddress1.Country = types.StringNull()
		}
		deliveryAddress1.PostalCode = types.StringPointerValue(deliveryAddressItem.PostalCode)
		deliveryAddress1.Street = types.StringPointerValue(deliveryAddressItem.Street)
		deliveryAddress1.StreetNumber = types.StringPointerValue(deliveryAddressItem.StreetNumber)
		if deliveryAddressCount+1 > len(r.DeliveryAddress) {
			r.DeliveryAddress = append(r.DeliveryAddress, deliveryAddress1)
		} else {
			r.DeliveryAddress[deliveryAddressCount].ID = deliveryAddress1.ID
			r.DeliveryAddress[deliveryAddressCount].Tags = deliveryAddress1.Tags
			r.DeliveryAddress[deliveryAddressCount].AdditionalInfo = deliveryAddress1.AdditionalInfo
			r.DeliveryAddress[deliveryAddressCount].City = deliveryAddress1.City
			r.DeliveryAddress[deliveryAddressCount].Country = deliveryAddress1.Country
			r.DeliveryAddress[deliveryAddressCount].PostalCode = deliveryAddress1.PostalCode
			r.DeliveryAddress[deliveryAddressCount].Street = deliveryAddress1.Street
			r.DeliveryAddress[deliveryAddressCount].StreetNumber = deliveryAddress1.StreetNumber
		}
	}
	if resp.ExpiresAt != nil {
		r.ExpiresAt = types.StringValue(resp.ExpiresAt.Format(time.RFC3339Nano))
	} else {
		r.ExpiresAt = types.StringNull()
	}
	if resp.MappedEntities == nil {
		r.MappedEntities = nil
	} else {
		r.MappedEntities = &BaseRelation{}
		if len(r.MappedEntities.DollarRelation) > len(resp.MappedEntities.DollarRelation) {
			r.MappedEntities.DollarRelation = r.MappedEntities.DollarRelation[:len(resp.MappedEntities.DollarRelation)]
		}
		for dollarRelationCount2, dollarRelationItem2 := range resp.MappedEntities.DollarRelation {
			var dollarRelation5 DollarRelation
			dollarRelation5.Tags = nil
			for _, v := range dollarRelationItem2.Tags {
				dollarRelation5.Tags = append(dollarRelation5.Tags, types.StringValue(v))
			}
			dollarRelation5.EntityID = types.StringPointerValue(dollarRelationItem2.EntityID)
			if dollarRelationCount2+1 > len(r.MappedEntities.DollarRelation) {
				r.MappedEntities.DollarRelation = append(r.MappedEntities.DollarRelation, dollarRelation5)
			} else {
				r.MappedEntities.DollarRelation[dollarRelationCount2].Tags = dollarRelation5.Tags
				r.MappedEntities.DollarRelation[dollarRelationCount2].EntityID = dollarRelation5.EntityID
			}
		}
	}
	r.OrderNumber = types.StringPointerValue(resp.OrderNumber)
	if resp.Prices == nil {
		r.Prices = nil
	} else {
		r.Prices = &BaseRelation{}
		if len(r.Prices.DollarRelation) > len(resp.Prices.DollarRelation) {
			r.Prices.DollarRelation = r.Prices.DollarRelation[:len(resp.Prices.DollarRelation)]
		}
		for dollarRelationCount3, dollarRelationItem3 := range resp.Prices.DollarRelation {
			var dollarRelation7 DollarRelation
			dollarRelation7.Tags = nil
			for _, v := range dollarRelationItem3.Tags {
				dollarRelation7.Tags = append(dollarRelation7.Tags, types.StringValue(v))
			}
			dollarRelation7.EntityID = types.StringPointerValue(dollarRelationItem3.EntityID)
			if dollarRelationCount3+1 > len(r.Prices.DollarRelation) {
				r.Prices.DollarRelation = append(r.Prices.DollarRelation, dollarRelation7)
			} else {
				r.Prices.DollarRelation[dollarRelationCount3].Tags = dollarRelation7.Tags
				r.Prices.DollarRelation[dollarRelationCount3].EntityID = dollarRelation7.EntityID
			}
		}
	}
	if resp.Products == nil {
		r.Products = nil
	} else {
		r.Products = &BaseRelation{}
		if len(r.Products.DollarRelation) > len(resp.Products.DollarRelation) {
			r.Products.DollarRelation = r.Products.DollarRelation[:len(resp.Products.DollarRelation)]
		}
		for dollarRelationCount4, dollarRelationItem4 := range resp.Products.DollarRelation {
			var dollarRelation9 DollarRelation
			dollarRelation9.Tags = nil
			for _, v := range dollarRelationItem4.Tags {
				dollarRelation9.Tags = append(dollarRelation9.Tags, types.StringValue(v))
			}
			dollarRelation9.EntityID = types.StringPointerValue(dollarRelationItem4.EntityID)
			if dollarRelationCount4+1 > len(r.Products.DollarRelation) {
				r.Products.DollarRelation = append(r.Products.DollarRelation, dollarRelation9)
			} else {
				r.Products.DollarRelation[dollarRelationCount4].Tags = dollarRelation9.Tags
				r.Products.DollarRelation[dollarRelationCount4].EntityID = dollarRelation9.EntityID
			}
		}
	}
	if resp.Source == nil {
		r.Source = nil
	} else {
		r.Source = &OrderCreateSource{}
		r.Source.Href = types.StringPointerValue(resp.Source.Href)
		r.Source.Title = types.StringPointerValue(resp.Source.Title)
	}
	r.SourceType = types.StringPointerValue(resp.SourceType)
	if resp.Status != nil {
		r.Status = types.StringValue(string(*resp.Status))
	} else {
		r.Status = types.StringNull()
	}
	if resp.TotalsDetail == nil {
		r.TotalsDetail = nil
	} else {
		r.TotalsDetail = &OrderCreateTotalsDetail{}
		if resp.TotalsDetail.AmountTax != nil {
			r.TotalsDetail.AmountTax = types.NumberValue(big.NewFloat(float64(*resp.TotalsDetail.AmountTax)))
		} else {
			r.TotalsDetail.AmountTax = types.NumberNull()
		}
		if resp.TotalsDetail.Breakdown == nil {
			r.TotalsDetail.Breakdown = nil
		} else {
			r.TotalsDetail.Breakdown = &OrderCreateBreakdown{}
			if len(r.TotalsDetail.Breakdown.Recurrences) > len(resp.TotalsDetail.Breakdown.Recurrences) {
				r.TotalsDetail.Breakdown.Recurrences = r.TotalsDetail.Breakdown.Recurrences[:len(resp.TotalsDetail.Breakdown.Recurrences)]
			}
			for recurrencesCount, recurrencesItem := range resp.TotalsDetail.Breakdown.Recurrences {
				var recurrences1 Recurrences
				if recurrencesItem.AmountSubtotal != nil {
					recurrences1.AmountSubtotal = types.NumberValue(big.NewFloat(float64(*recurrencesItem.AmountSubtotal)))
				} else {
					recurrences1.AmountSubtotal = types.NumberNull()
				}
				recurrences1.AmountSubtotalDecimal = types.StringPointerValue(recurrencesItem.AmountSubtotalDecimal)
				if recurrencesItem.AmountTax != nil {
					recurrences1.AmountTax = types.NumberValue(big.NewFloat(float64(*recurrencesItem.AmountTax)))
				} else {
					recurrences1.AmountTax = types.NumberNull()
				}
				if recurrencesItem.AmountTotal != nil {
					recurrences1.AmountTotal = types.NumberValue(big.NewFloat(float64(*recurrencesItem.AmountTotal)))
				} else {
					recurrences1.AmountTotal = types.NumberNull()
				}
				recurrences1.AmountTotalDecimal = types.StringPointerValue(recurrencesItem.AmountTotalDecimal)
				if recurrencesItem.Type != nil {
					recurrences1.Type = types.StringValue(string(*recurrencesItem.Type))
				} else {
					recurrences1.Type = types.StringNull()
				}
				if recurrencesCount+1 > len(r.TotalsDetail.Breakdown.Recurrences) {
					r.TotalsDetail.Breakdown.Recurrences = append(r.TotalsDetail.Breakdown.Recurrences, recurrences1)
				} else {
					r.TotalsDetail.Breakdown.Recurrences[recurrencesCount].AmountSubtotal = recurrences1.AmountSubtotal
					r.TotalsDetail.Breakdown.Recurrences[recurrencesCount].AmountSubtotalDecimal = recurrences1.AmountSubtotalDecimal
					r.TotalsDetail.Breakdown.Recurrences[recurrencesCount].AmountTax = recurrences1.AmountTax
					r.TotalsDetail.Breakdown.Recurrences[recurrencesCount].AmountTotal = recurrences1.AmountTotal
					r.TotalsDetail.Breakdown.Recurrences[recurrencesCount].AmountTotalDecimal = recurrences1.AmountTotalDecimal
					r.TotalsDetail.Breakdown.Recurrences[recurrencesCount].Type = recurrences1.Type
				}
			}
			if len(r.TotalsDetail.Breakdown.Taxes) > len(resp.TotalsDetail.Breakdown.Taxes) {
				r.TotalsDetail.Breakdown.Taxes = r.TotalsDetail.Breakdown.Taxes[:len(resp.TotalsDetail.Breakdown.Taxes)]
			}
			for taxesCount, taxesItem := range resp.TotalsDetail.Breakdown.Taxes {
				var taxes1 Taxes
				if taxesItem.Amount != nil {
					taxes1.Amount = types.NumberValue(big.NewFloat(float64(*taxesItem.Amount)))
				} else {
					taxes1.Amount = types.NumberNull()
				}
				for taxCount, taxItem := range taxesItem.Tax {
					var tax1 Tax
					tax1.ID = types.StringPointerValue(taxItem.ID)
					tax1.Rate = types.StringPointerValue(taxItem.Rate)
					if taxItem.Type != nil {
						tax1.Type = types.StringValue(string(*taxItem.Type))
					} else {
						tax1.Type = types.StringNull()
					}
					if taxCount+1 > len(taxes1.Tax) {
						taxes1.Tax = append(taxes1.Tax, tax1)
					} else {
						taxes1.Tax[taxCount].ID = tax1.ID
						taxes1.Tax[taxCount].Rate = tax1.Rate
						taxes1.Tax[taxCount].Type = tax1.Type
					}
				}
				if taxesCount+1 > len(r.TotalsDetail.Breakdown.Taxes) {
					r.TotalsDetail.Breakdown.Taxes = append(r.TotalsDetail.Breakdown.Taxes, taxes1)
				} else {
					r.TotalsDetail.Breakdown.Taxes[taxesCount].Amount = taxes1.Amount
					r.TotalsDetail.Breakdown.Taxes[taxesCount].Tax = taxes1.Tax
				}
			}
		}
	}
}
