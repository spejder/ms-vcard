package odoo

import (
	"fmt"
)

// RegistrationEditorLine represents registration.editor.line model.
type RegistrationEditorLine struct {
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	EditorId        *Many2One `xmlrpc:"editor_id,omptempty"`
	Email           *String   `xmlrpc:"email,omptempty"`
	EventId         *Many2One `xmlrpc:"event_id,omptempty"`
	EventTicketId   *Many2One `xmlrpc:"event_ticket_id,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Mobile          *String   `xmlrpc:"mobile,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	Phone           *String   `xmlrpc:"phone,omptempty"`
	RegistrationId  *Many2One `xmlrpc:"registration_id,omptempty"`
	SaleOrderLineId *Many2One `xmlrpc:"sale_order_line_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// RegistrationEditorLines represents array of registration.editor.line model.
type RegistrationEditorLines []RegistrationEditorLine

// RegistrationEditorLineModel is the odoo model name.
const RegistrationEditorLineModel = "registration.editor.line"

// Many2One convert RegistrationEditorLine to *Many2One.
func (rel *RegistrationEditorLine) Many2One() *Many2One {
	return NewMany2One(rel.Id.Get(), "")
}

// CreateRegistrationEditorLine creates a new registration.editor.line model and returns its id.
func (c *Client) CreateRegistrationEditorLine(rel *RegistrationEditorLine) (int64, error) {
	return c.Create(RegistrationEditorLineModel, rel)
}

// UpdateRegistrationEditorLine updates an existing registration.editor.line record.
func (c *Client) UpdateRegistrationEditorLine(rel *RegistrationEditorLine) error {
	return c.UpdateRegistrationEditorLines([]int64{rel.Id.Get()}, rel)
}

// UpdateRegistrationEditorLines updates existing registration.editor.line records.
// All records (represented by ids) will be updated by rel values.
func (c *Client) UpdateRegistrationEditorLines(ids []int64, rel *RegistrationEditorLine) error {
	return c.Update(RegistrationEditorLineModel, ids, rel)
}

// DeleteRegistrationEditorLine deletes an existing registration.editor.line record.
func (c *Client) DeleteRegistrationEditorLine(id int64) error {
	return c.DeleteRegistrationEditorLines([]int64{id})
}

// DeleteRegistrationEditorLines deletes existing registration.editor.line records.
func (c *Client) DeleteRegistrationEditorLines(ids []int64) error {
	return c.Delete(RegistrationEditorLineModel, ids)
}

// GetRegistrationEditorLine gets registration.editor.line existing record.
func (c *Client) GetRegistrationEditorLine(id int64) (*RegistrationEditorLine, error) {
	rels, err := c.GetRegistrationEditorLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if rels != nil && len(*rels) > 0 {
		return &((*rels)[0]), nil
	}
	return nil, fmt.Errorf("id %v of registration.editor.line not found", id)
}

// GetRegistrationEditorLines gets registration.editor.line existing records.
func (c *Client) GetRegistrationEditorLines(ids []int64) (*RegistrationEditorLines, error) {
	rels := &RegistrationEditorLines{}
	if err := c.Read(RegistrationEditorLineModel, ids, nil, rels); err != nil {
		return nil, err
	}
	return rels, nil
}

// FindRegistrationEditorLine finds registration.editor.line record by querying it with criteria.
func (c *Client) FindRegistrationEditorLine(criteria *Criteria) (*RegistrationEditorLine, error) {
	rels := &RegistrationEditorLines{}
	if err := c.SearchRead(RegistrationEditorLineModel, criteria, NewOptions().Limit(1), rels); err != nil {
		return nil, err
	}
	if rels != nil && len(*rels) > 0 {
		return &((*rels)[0]), nil
	}
	return nil, fmt.Errorf("registration.editor.line was not found")
}

// FindRegistrationEditorLines finds registration.editor.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditorLines(criteria *Criteria, options *Options) (*RegistrationEditorLines, error) {
	rels := &RegistrationEditorLines{}
	if err := c.SearchRead(RegistrationEditorLineModel, criteria, options, rels); err != nil {
		return nil, err
	}
	return rels, nil
}

// FindRegistrationEditorLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRegistrationEditorLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RegistrationEditorLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRegistrationEditorLineId finds record id by querying it with criteria.
func (c *Client) FindRegistrationEditorLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RegistrationEditorLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("registration.editor.line was not found")
}
