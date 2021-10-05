package odoo

import (
	"fmt"
)

// SaleOrderTemplateOption represents sale.order.template.option model.
type SaleOrderTemplateOption struct {
	CompanyId            *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	Discount             *Float    `xmlrpc:"discount,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	Name                 *String   `xmlrpc:"name,omptempty"`
	PriceUnit            *Float    `xmlrpc:"price_unit,omptempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omptempty"`
	Quantity             *Float    `xmlrpc:"quantity,omptempty"`
	SaleOrderTemplateId  *Many2One `xmlrpc:"sale_order_template_id,omptempty"`
	UomId                *Many2One `xmlrpc:"uom_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderTemplateOptions represents array of sale.order.template.option model.
type SaleOrderTemplateOptions []SaleOrderTemplateOption

// SaleOrderTemplateOptionModel is the odoo model name.
const SaleOrderTemplateOptionModel = "sale.order.template.option"

// Many2One convert SaleOrderTemplateOption to *Many2One.
func (soto *SaleOrderTemplateOption) Many2One() *Many2One {
	return NewMany2One(soto.Id.Get(), "")
}

// CreateSaleOrderTemplateOption creates a new sale.order.template.option model and returns its id.
func (c *Client) CreateSaleOrderTemplateOption(soto *SaleOrderTemplateOption) (int64, error) {
	return c.Create(SaleOrderTemplateOptionModel, soto)
}

// UpdateSaleOrderTemplateOption updates an existing sale.order.template.option record.
func (c *Client) UpdateSaleOrderTemplateOption(soto *SaleOrderTemplateOption) error {
	return c.UpdateSaleOrderTemplateOptions([]int64{soto.Id.Get()}, soto)
}

// UpdateSaleOrderTemplateOptions updates existing sale.order.template.option records.
// All records (represented by ids) will be updated by soto values.
func (c *Client) UpdateSaleOrderTemplateOptions(ids []int64, soto *SaleOrderTemplateOption) error {
	return c.Update(SaleOrderTemplateOptionModel, ids, soto)
}

// DeleteSaleOrderTemplateOption deletes an existing sale.order.template.option record.
func (c *Client) DeleteSaleOrderTemplateOption(id int64) error {
	return c.DeleteSaleOrderTemplateOptions([]int64{id})
}

// DeleteSaleOrderTemplateOptions deletes existing sale.order.template.option records.
func (c *Client) DeleteSaleOrderTemplateOptions(ids []int64) error {
	return c.Delete(SaleOrderTemplateOptionModel, ids)
}

// GetSaleOrderTemplateOption gets sale.order.template.option existing record.
func (c *Client) GetSaleOrderTemplateOption(id int64) (*SaleOrderTemplateOption, error) {
	sotos, err := c.GetSaleOrderTemplateOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	if sotos != nil && len(*sotos) > 0 {
		return &((*sotos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.template.option not found", id)
}

// GetSaleOrderTemplateOptions gets sale.order.template.option existing records.
func (c *Client) GetSaleOrderTemplateOptions(ids []int64) (*SaleOrderTemplateOptions, error) {
	sotos := &SaleOrderTemplateOptions{}
	if err := c.Read(SaleOrderTemplateOptionModel, ids, nil, sotos); err != nil {
		return nil, err
	}
	return sotos, nil
}

// FindSaleOrderTemplateOption finds sale.order.template.option record by querying it with criteria.
func (c *Client) FindSaleOrderTemplateOption(criteria *Criteria) (*SaleOrderTemplateOption, error) {
	sotos := &SaleOrderTemplateOptions{}
	if err := c.SearchRead(SaleOrderTemplateOptionModel, criteria, NewOptions().Limit(1), sotos); err != nil {
		return nil, err
	}
	if sotos != nil && len(*sotos) > 0 {
		return &((*sotos)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.template.option was not found")
}

// FindSaleOrderTemplateOptions finds sale.order.template.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateOptions(criteria *Criteria, options *Options) (*SaleOrderTemplateOptions, error) {
	sotos := &SaleOrderTemplateOptions{}
	if err := c.SearchRead(SaleOrderTemplateOptionModel, criteria, options, sotos); err != nil {
		return nil, err
	}
	return sotos, nil
}

// FindSaleOrderTemplateOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderTemplateOptionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderTemplateOptionId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderTemplateOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderTemplateOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.template.option was not found")
}
