package odoo

import (
	"fmt"
)

// RegistrationEditor represents registration.editor model.
type RegistrationEditor struct {
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	EventRegistrationIds *Relation `xmlrpc:"event_registration_ids,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	SaleOrderId          *Many2One `xmlrpc:"sale_order_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// RegistrationEditors represents array of registration.editor model.
type RegistrationEditors []RegistrationEditor

// RegistrationEditorModel is the odoo model name.
const RegistrationEditorModel = "registration.editor"

// Many2One convert RegistrationEditor to *Many2One.
func (re *RegistrationEditor) Many2One() *Many2One {
	return NewMany2One(re.Id.Get(), "")
}

// CreateRegistrationEditor creates a new registration.editor model and returns its id.
func (c *Client) CreateRegistrationEditor(re *RegistrationEditor) (int64, error) {
	return c.Create(RegistrationEditorModel, re)
}

// UpdateRegistrationEditor updates an existing registration.editor record.
func (c *Client) UpdateRegistrationEditor(re *RegistrationEditor) error {
	return c.UpdateRegistrationEditors([]int64{re.Id.Get()}, re)
}

// UpdateRegistrationEditors updates existing registration.editor records.
// All records (represented by ids) will be updated by re values.
func (c *Client) UpdateRegistrationEditors(ids []int64, re *RegistrationEditor) error {
	return c.Update(RegistrationEditorModel, ids, re)
}

// DeleteRegistrationEditor deletes an existing registration.editor record.
func (c *Client) DeleteRegistrationEditor(id int64) error {
	return c.DeleteRegistrationEditors([]int64{id})
}

// DeleteRegistrationEditors deletes existing registration.editor records.
func (c *Client) DeleteRegistrationEditors(ids []int64) error {
	return c.Delete(RegistrationEditorModel, ids)
}

// GetRegistrationEditor gets registration.editor existing record.
func (c *Client) GetRegistrationEditor(id int64) (*RegistrationEditor, error) {
	res, err := c.GetRegistrationEditors([]int64{id})
	if err != nil {
		return nil, err
	}
	if res != nil && len(*res) > 0 {
		return &((*res)[0]), nil
	}
	return nil, fmt.Errorf("id %v of registration.editor not found", id)
}

// GetRegistrationEditors gets registration.editor existing records.
func (c *Client) GetRegistrationEditors(ids []int64) (*RegistrationEditors, error) {
	res := &RegistrationEditors{}
	if err := c.Read(RegistrationEditorModel, ids, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}

// FindRegistrationEditor finds registration.editor record by querying it with criteria.
func (c *Client) FindRegistrationEditor(criteria *Criteria) (*RegistrationEditor, error) {
	res := &RegistrationEditors{}
	if err := c.SearchRead(RegistrationEditorModel, criteria, NewOptions().Limit(1), res); err != nil {
		return nil, err
	}
	if res != nil && len(*res) > 0 {
		return &((*res)[0]), nil
	}
	return nil, fmt.Errorf("registration.editor was not found")
}

// FindRegistrationEditors finds registration.editor records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditors(criteria *Criteria, options *Options) (*RegistrationEditors, error) {
	res := &RegistrationEditors{}
	if err := c.SearchRead(RegistrationEditorModel, criteria, options, res); err != nil {
		return nil, err
	}
	return res, nil
}

// FindRegistrationEditorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RegistrationEditorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRegistrationEditorId finds record id by querying it with criteria.
func (c *Client) FindRegistrationEditorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RegistrationEditorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("registration.editor was not found")
}
