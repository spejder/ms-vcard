package odoo

import (
	"fmt"
)

// MailResendMessage represents mail.resend.message model.
type MailResendMessage struct {
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	HasCancel       *Bool     `xmlrpc:"has_cancel,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	MailMessageId   *Many2One `xmlrpc:"mail_message_id,omptempty"`
	NotificationIds *Relation `xmlrpc:"notification_ids,omptempty"`
	PartnerIds      *Relation `xmlrpc:"partner_ids,omptempty"`
	PartnerReadonly *Bool     `xmlrpc:"partner_readonly,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailResendMessages represents array of mail.resend.message model.
type MailResendMessages []MailResendMessage

// MailResendMessageModel is the odoo model name.
const MailResendMessageModel = "mail.resend.message"

// Many2One convert MailResendMessage to *Many2One.
func (mrm *MailResendMessage) Many2One() *Many2One {
	return NewMany2One(mrm.Id.Get(), "")
}

// CreateMailResendMessage creates a new mail.resend.message model and returns its id.
func (c *Client) CreateMailResendMessage(mrm *MailResendMessage) (int64, error) {
	return c.Create(MailResendMessageModel, mrm)
}

// UpdateMailResendMessage updates an existing mail.resend.message record.
func (c *Client) UpdateMailResendMessage(mrm *MailResendMessage) error {
	return c.UpdateMailResendMessages([]int64{mrm.Id.Get()}, mrm)
}

// UpdateMailResendMessages updates existing mail.resend.message records.
// All records (represented by ids) will be updated by mrm values.
func (c *Client) UpdateMailResendMessages(ids []int64, mrm *MailResendMessage) error {
	return c.Update(MailResendMessageModel, ids, mrm)
}

// DeleteMailResendMessage deletes an existing mail.resend.message record.
func (c *Client) DeleteMailResendMessage(id int64) error {
	return c.DeleteMailResendMessages([]int64{id})
}

// DeleteMailResendMessages deletes existing mail.resend.message records.
func (c *Client) DeleteMailResendMessages(ids []int64) error {
	return c.Delete(MailResendMessageModel, ids)
}

// GetMailResendMessage gets mail.resend.message existing record.
func (c *Client) GetMailResendMessage(id int64) (*MailResendMessage, error) {
	mrms, err := c.GetMailResendMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mrms != nil && len(*mrms) > 0 {
		return &((*mrms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.resend.message not found", id)
}

// GetMailResendMessages gets mail.resend.message existing records.
func (c *Client) GetMailResendMessages(ids []int64) (*MailResendMessages, error) {
	mrms := &MailResendMessages{}
	if err := c.Read(MailResendMessageModel, ids, nil, mrms); err != nil {
		return nil, err
	}
	return mrms, nil
}

// FindMailResendMessage finds mail.resend.message record by querying it with criteria.
func (c *Client) FindMailResendMessage(criteria *Criteria) (*MailResendMessage, error) {
	mrms := &MailResendMessages{}
	if err := c.SearchRead(MailResendMessageModel, criteria, NewOptions().Limit(1), mrms); err != nil {
		return nil, err
	}
	if mrms != nil && len(*mrms) > 0 {
		return &((*mrms)[0]), nil
	}
	return nil, fmt.Errorf("mail.resend.message was not found")
}

// FindMailResendMessages finds mail.resend.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendMessages(criteria *Criteria, options *Options) (*MailResendMessages, error) {
	mrms := &MailResendMessages{}
	if err := c.SearchRead(MailResendMessageModel, criteria, options, mrms); err != nil {
		return nil, err
	}
	return mrms, nil
}

// FindMailResendMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailResendMessageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailResendMessageId finds record id by querying it with criteria.
func (c *Client) FindMailResendMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailResendMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.resend.message was not found")
}
