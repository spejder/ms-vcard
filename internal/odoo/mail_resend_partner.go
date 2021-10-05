package odoo

import (
	"fmt"
)

// MailResendPartner represents mail.resend.partner model.
type MailResendPartner struct {
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Email          *String   `xmlrpc:"email,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	Message        *String   `xmlrpc:"message,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	PartnerId      *Many2One `xmlrpc:"partner_id,omptempty"`
	Resend         *Bool     `xmlrpc:"resend,omptempty"`
	ResendWizardId *Many2One `xmlrpc:"resend_wizard_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailResendPartners represents array of mail.resend.partner model.
type MailResendPartners []MailResendPartner

// MailResendPartnerModel is the odoo model name.
const MailResendPartnerModel = "mail.resend.partner"

// Many2One convert MailResendPartner to *Many2One.
func (mrp *MailResendPartner) Many2One() *Many2One {
	return NewMany2One(mrp.Id.Get(), "")
}

// CreateMailResendPartner creates a new mail.resend.partner model and returns its id.
func (c *Client) CreateMailResendPartner(mrp *MailResendPartner) (int64, error) {
	return c.Create(MailResendPartnerModel, mrp)
}

// UpdateMailResendPartner updates an existing mail.resend.partner record.
func (c *Client) UpdateMailResendPartner(mrp *MailResendPartner) error {
	return c.UpdateMailResendPartners([]int64{mrp.Id.Get()}, mrp)
}

// UpdateMailResendPartners updates existing mail.resend.partner records.
// All records (represented by ids) will be updated by mrp values.
func (c *Client) UpdateMailResendPartners(ids []int64, mrp *MailResendPartner) error {
	return c.Update(MailResendPartnerModel, ids, mrp)
}

// DeleteMailResendPartner deletes an existing mail.resend.partner record.
func (c *Client) DeleteMailResendPartner(id int64) error {
	return c.DeleteMailResendPartners([]int64{id})
}

// DeleteMailResendPartners deletes existing mail.resend.partner records.
func (c *Client) DeleteMailResendPartners(ids []int64) error {
	return c.Delete(MailResendPartnerModel, ids)
}

// GetMailResendPartner gets mail.resend.partner existing record.
func (c *Client) GetMailResendPartner(id int64) (*MailResendPartner, error) {
	mrps, err := c.GetMailResendPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if mrps != nil && len(*mrps) > 0 {
		return &((*mrps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.resend.partner not found", id)
}

// GetMailResendPartners gets mail.resend.partner existing records.
func (c *Client) GetMailResendPartners(ids []int64) (*MailResendPartners, error) {
	mrps := &MailResendPartners{}
	if err := c.Read(MailResendPartnerModel, ids, nil, mrps); err != nil {
		return nil, err
	}
	return mrps, nil
}

// FindMailResendPartner finds mail.resend.partner record by querying it with criteria.
func (c *Client) FindMailResendPartner(criteria *Criteria) (*MailResendPartner, error) {
	mrps := &MailResendPartners{}
	if err := c.SearchRead(MailResendPartnerModel, criteria, NewOptions().Limit(1), mrps); err != nil {
		return nil, err
	}
	if mrps != nil && len(*mrps) > 0 {
		return &((*mrps)[0]), nil
	}
	return nil, fmt.Errorf("mail.resend.partner was not found")
}

// FindMailResendPartners finds mail.resend.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendPartners(criteria *Criteria, options *Options) (*MailResendPartners, error) {
	mrps := &MailResendPartners{}
	if err := c.SearchRead(MailResendPartnerModel, criteria, options, mrps); err != nil {
		return nil, err
	}
	return mrps, nil
}

// FindMailResendPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailResendPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailResendPartnerId finds record id by querying it with criteria.
func (c *Client) FindMailResendPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailResendPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.resend.partner was not found")
}
