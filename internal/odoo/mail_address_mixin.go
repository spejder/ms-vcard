package odoo

import (
	"fmt"
)

// MailAddressMixin represents mail.address.mixin model.
type MailAddressMixin struct {
	DisplayName     *String `xmlrpc:"display_name,omptempty"`
	EmailNormalized *String `xmlrpc:"email_normalized,omptempty"`
	Id              *Int    `xmlrpc:"id,omptempty"`
	LastUpdate      *Time   `xmlrpc:"__last_update,omptempty"`
}

// MailAddressMixins represents array of mail.address.mixin model.
type MailAddressMixins []MailAddressMixin

// MailAddressMixinModel is the odoo model name.
const MailAddressMixinModel = "mail.address.mixin"

// Many2One convert MailAddressMixin to *Many2One.
func (mam *MailAddressMixin) Many2One() *Many2One {
	return NewMany2One(mam.Id.Get(), "")
}

// CreateMailAddressMixin creates a new mail.address.mixin model and returns its id.
func (c *Client) CreateMailAddressMixin(mam *MailAddressMixin) (int64, error) {
	return c.Create(MailAddressMixinModel, mam)
}

// UpdateMailAddressMixin updates an existing mail.address.mixin record.
func (c *Client) UpdateMailAddressMixin(mam *MailAddressMixin) error {
	return c.UpdateMailAddressMixins([]int64{mam.Id.Get()}, mam)
}

// UpdateMailAddressMixins updates existing mail.address.mixin records.
// All records (represented by ids) will be updated by mam values.
func (c *Client) UpdateMailAddressMixins(ids []int64, mam *MailAddressMixin) error {
	return c.Update(MailAddressMixinModel, ids, mam)
}

// DeleteMailAddressMixin deletes an existing mail.address.mixin record.
func (c *Client) DeleteMailAddressMixin(id int64) error {
	return c.DeleteMailAddressMixins([]int64{id})
}

// DeleteMailAddressMixins deletes existing mail.address.mixin records.
func (c *Client) DeleteMailAddressMixins(ids []int64) error {
	return c.Delete(MailAddressMixinModel, ids)
}

// GetMailAddressMixin gets mail.address.mixin existing record.
func (c *Client) GetMailAddressMixin(id int64) (*MailAddressMixin, error) {
	mams, err := c.GetMailAddressMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if mams != nil && len(*mams) > 0 {
		return &((*mams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.address.mixin not found", id)
}

// GetMailAddressMixins gets mail.address.mixin existing records.
func (c *Client) GetMailAddressMixins(ids []int64) (*MailAddressMixins, error) {
	mams := &MailAddressMixins{}
	if err := c.Read(MailAddressMixinModel, ids, nil, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailAddressMixin finds mail.address.mixin record by querying it with criteria.
func (c *Client) FindMailAddressMixin(criteria *Criteria) (*MailAddressMixin, error) {
	mams := &MailAddressMixins{}
	if err := c.SearchRead(MailAddressMixinModel, criteria, NewOptions().Limit(1), mams); err != nil {
		return nil, err
	}
	if mams != nil && len(*mams) > 0 {
		return &((*mams)[0]), nil
	}
	return nil, fmt.Errorf("mail.address.mixin was not found")
}

// FindMailAddressMixins finds mail.address.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAddressMixins(criteria *Criteria, options *Options) (*MailAddressMixins, error) {
	mams := &MailAddressMixins{}
	if err := c.SearchRead(MailAddressMixinModel, criteria, options, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailAddressMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAddressMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailAddressMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailAddressMixinId finds record id by querying it with criteria.
func (c *Client) FindMailAddressMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailAddressMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.address.mixin was not found")
}
