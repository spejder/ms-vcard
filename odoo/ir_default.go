package odoo

import (
	"fmt"
)

// IrDefault represents ir.default model.
type IrDefault struct {
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	Condition   *String   `xmlrpc:"condition,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FieldId     *Many2One `xmlrpc:"field_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JsonValue   *String   `xmlrpc:"json_value,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// IrDefaults represents array of ir.default model.
type IrDefaults []IrDefault

// IrDefaultModel is the odoo model name.
const IrDefaultModel = "ir.default"

// Many2One convert IrDefault to *Many2One.
func (ID *IrDefault) Many2One() *Many2One {
	return NewMany2One(ID.Id.Get(), "")
}

// CreateIrDefault creates a new ir.default model and returns its id.
func (c *Client) CreateIrDefault(ID *IrDefault) (int64, error) {
	return c.Create(IrDefaultModel, ID)
}

// UpdateIrDefault updates an existing ir.default record.
func (c *Client) UpdateIrDefault(ID *IrDefault) error {
	return c.UpdateIrDefaults([]int64{ID.Id.Get()}, ID)
}

// UpdateIrDefaults updates existing ir.default records.
// All records (represented by ids) will be updated by ID values.
func (c *Client) UpdateIrDefaults(ids []int64, ID *IrDefault) error {
	return c.Update(IrDefaultModel, ids, ID)
}

// DeleteIrDefault deletes an existing ir.default record.
func (c *Client) DeleteIrDefault(id int64) error {
	return c.DeleteIrDefaults([]int64{id})
}

// DeleteIrDefaults deletes existing ir.default records.
func (c *Client) DeleteIrDefaults(ids []int64) error {
	return c.Delete(IrDefaultModel, ids)
}

// GetIrDefault gets ir.default existing record.
func (c *Client) GetIrDefault(id int64) (*IrDefault, error) {
	IDs, err := c.GetIrDefaults([]int64{id})
	if err != nil {
		return nil, err
	}
	if IDs != nil && len(*IDs) > 0 {
		return &((*IDs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.default not found", id)
}

// GetIrDefaults gets ir.default existing records.
func (c *Client) GetIrDefaults(ids []int64) (*IrDefaults, error) {
	IDs := &IrDefaults{}
	if err := c.Read(IrDefaultModel, ids, nil, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDefault finds ir.default record by querying it with criteria.
func (c *Client) FindIrDefault(criteria *Criteria) (*IrDefault, error) {
	IDs := &IrDefaults{}
	if err := c.SearchRead(IrDefaultModel, criteria, NewOptions().Limit(1), IDs); err != nil {
		return nil, err
	}
	if IDs != nil && len(*IDs) > 0 {
		return &((*IDs)[0]), nil
	}
	return nil, fmt.Errorf("ir.default was not found")
}

// FindIrDefaults finds ir.default records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDefaults(criteria *Criteria, options *Options) (*IrDefaults, error) {
	IDs := &IrDefaults{}
	if err := c.SearchRead(IrDefaultModel, criteria, options, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDefaultIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDefaultIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrDefaultModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrDefaultId finds record id by querying it with criteria.
func (c *Client) FindIrDefaultId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrDefaultModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.default was not found")
}
