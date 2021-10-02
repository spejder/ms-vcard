package odoo

import (
	"fmt"
)

// SaleOrderTemplateLine represents sale.order.template.line model.
type SaleOrderTemplateLine struct {
	CompanyId            *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	Discount             *Float     `xmlrpc:"discount,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	DisplayType          *Selection `xmlrpc:"display_type,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	PriceUnit            *Float     `xmlrpc:"price_unit,omptempty"`
	ProductId            *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId *Many2One  `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId         *Many2One  `xmlrpc:"product_uom_id,omptempty"`
	ProductUomQty        *Float     `xmlrpc:"product_uom_qty,omptempty"`
	SaleOrderTemplateId  *Many2One  `xmlrpc:"sale_order_template_id,omptempty"`
	Sequence             *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderTemplateLines represents array of sale.order.template.line model.
type SaleOrderTemplateLines []SaleOrderTemplateLine

// SaleOrderTemplateLineModel is the odoo model name.
const SaleOrderTemplateLineModel = "sale.order.template.line"

// Many2One convert SaleOrderTemplateLine to *Many2One.
func (sotl *SaleOrderTemplateLine) Many2One() *Many2One {
	return NewMany2One(sotl.Id.Get(), "")
}

// CreateSaleOrderTemplateLine creates a new sale.order.template.line model and returns its id.
func (c *Client) CreateSaleOrderTemplateLine(sotl *SaleOrderTemplateLine) (int64, error) {
	return c.Create(SaleOrderTemplateLineModel, sotl)
}

// UpdateSaleOrderTemplateLine updates an existing sale.order.template.line record.
func (c *Client) UpdateSaleOrderTemplateLine(sotl *SaleOrderTemplateLine) error {
	return c.UpdateSaleOrderTemplateLines([]int64{sotl.Id.Get()}, sotl)
}

// UpdateSaleOrderTemplateLines updates existing sale.order.template.line records.
// All records (represented by ids) will be updated by sotl values.
func (c *Client) UpdateSaleOrderTemplateLines(ids []int64, sotl *SaleOrderTemplateLine) error {
	return c.Update(SaleOrderTemplateLineModel, ids, sotl)
}

// DeleteSaleOrderTemplateLine deletes an existing sale.order.template.line record.
func (c *Client) DeleteSaleOrderTemplateLine(id int64) error {
	return c.DeleteSaleOrderTemplateLines([]int64{id})
}

// DeleteSaleOrderTemplateLines deletes existing sale.order.template.line records.
func (c *Client) DeleteSaleOrderTemplateLines(ids []int64) error {
	return c.Delete(SaleOrderTemplateLineModel, ids)
}

// GetSaleOrderTemplateLine gets sale.order.template.line existing record.
func (c *Client) GetSaleOrderTemplateLine(id int64) (*SaleOrderTemplateLine, error) {
	sotls, err := c.GetSaleOrderTemplateLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if sotls != nil && len(*sotls) > 0 {
		return &((*sotls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.template.line not found", id)
}

// GetSaleOrderTemplateLines gets sale.order.template.line existing records.
func (c *Client) GetSaleOrderTemplateLines(ids []int64) (*SaleOrderTemplateLines, error) {
	sotls := &SaleOrderTemplateLines{}
	if err := c.Read(SaleOrderTemplateLineModel, ids, nil, sotls); err != nil {
		return nil, err
	}
	return sotls, nil
}

// FindSaleOrderTemplateLine finds sale.order.template.line record by querying it with criteria.
func (c *Client) FindSaleOrderTemplateLine(criteria *Criteria) (*SaleOrderTemplateLine, error) {
	sotls := &SaleOrderTemplateLines{}
	if err := c.SearchRead(SaleOrderTemplateLineModel, criteria, NewOptions().Limit(1), sotls); err != nil {
		return nil, err
	}
	if sotls != nil && len(*sotls) > 0 {
		return &((*sotls)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.template.line was not found")
}

// FindSaleOrderTemplateLines finds sale.order.template.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateLines(criteria *Criteria, options *Options) (*SaleOrderTemplateLines, error) {
	sotls := &SaleOrderTemplateLines{}
	if err := c.SearchRead(SaleOrderTemplateLineModel, criteria, options, sotls); err != nil {
		return nil, err
	}
	return sotls, nil
}

// FindSaleOrderTemplateLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderTemplateLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderTemplateLineId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderTemplateLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderTemplateLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.template.line was not found")
}
