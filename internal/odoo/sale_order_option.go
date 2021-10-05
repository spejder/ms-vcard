package odoo

import (
	"fmt"
)

// SaleOrderOption represents sale.order.option model.
type SaleOrderOption struct {
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	Discount             *Float    `xmlrpc:"discount,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	IsPresent            *Bool     `xmlrpc:"is_present,omptempty"`
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	LineId               *Many2One `xmlrpc:"line_id,omptempty"`
	Name                 *String   `xmlrpc:"name,omptempty"`
	OrderId              *Many2One `xmlrpc:"order_id,omptempty"`
	PriceUnit            *Float    `xmlrpc:"price_unit,omptempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omptempty"`
	Quantity             *Float    `xmlrpc:"quantity,omptempty"`
	Sequence             *Int      `xmlrpc:"sequence,omptempty"`
	UomId                *Many2One `xmlrpc:"uom_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderOptions represents array of sale.order.option model.
type SaleOrderOptions []SaleOrderOption

// SaleOrderOptionModel is the odoo model name.
const SaleOrderOptionModel = "sale.order.option"

// Many2One convert SaleOrderOption to *Many2One.
func (soo *SaleOrderOption) Many2One() *Many2One {
	return NewMany2One(soo.Id.Get(), "")
}

// CreateSaleOrderOption creates a new sale.order.option model and returns its id.
func (c *Client) CreateSaleOrderOption(soo *SaleOrderOption) (int64, error) {
	return c.Create(SaleOrderOptionModel, soo)
}

// UpdateSaleOrderOption updates an existing sale.order.option record.
func (c *Client) UpdateSaleOrderOption(soo *SaleOrderOption) error {
	return c.UpdateSaleOrderOptions([]int64{soo.Id.Get()}, soo)
}

// UpdateSaleOrderOptions updates existing sale.order.option records.
// All records (represented by ids) will be updated by soo values.
func (c *Client) UpdateSaleOrderOptions(ids []int64, soo *SaleOrderOption) error {
	return c.Update(SaleOrderOptionModel, ids, soo)
}

// DeleteSaleOrderOption deletes an existing sale.order.option record.
func (c *Client) DeleteSaleOrderOption(id int64) error {
	return c.DeleteSaleOrderOptions([]int64{id})
}

// DeleteSaleOrderOptions deletes existing sale.order.option records.
func (c *Client) DeleteSaleOrderOptions(ids []int64) error {
	return c.Delete(SaleOrderOptionModel, ids)
}

// GetSaleOrderOption gets sale.order.option existing record.
func (c *Client) GetSaleOrderOption(id int64) (*SaleOrderOption, error) {
	soos, err := c.GetSaleOrderOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	if soos != nil && len(*soos) > 0 {
		return &((*soos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.option not found", id)
}

// GetSaleOrderOptions gets sale.order.option existing records.
func (c *Client) GetSaleOrderOptions(ids []int64) (*SaleOrderOptions, error) {
	soos := &SaleOrderOptions{}
	if err := c.Read(SaleOrderOptionModel, ids, nil, soos); err != nil {
		return nil, err
	}
	return soos, nil
}

// FindSaleOrderOption finds sale.order.option record by querying it with criteria.
func (c *Client) FindSaleOrderOption(criteria *Criteria) (*SaleOrderOption, error) {
	soos := &SaleOrderOptions{}
	if err := c.SearchRead(SaleOrderOptionModel, criteria, NewOptions().Limit(1), soos); err != nil {
		return nil, err
	}
	if soos != nil && len(*soos) > 0 {
		return &((*soos)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.option was not found")
}

// FindSaleOrderOptions finds sale.order.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderOptions(criteria *Criteria, options *Options) (*SaleOrderOptions, error) {
	soos := &SaleOrderOptions{}
	if err := c.SearchRead(SaleOrderOptionModel, criteria, options, soos); err != nil {
		return nil, err
	}
	return soos, nil
}

// FindSaleOrderOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderOptionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderOptionId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.option was not found")
}
