package odoo

import (
	"fmt"
)

// MailThreadPhone represents mail.thread.phone model.
type MailThreadPhone struct {
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	PhoneBlacklisted         *Bool     `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneSanitized           *String   `xmlrpc:"phone_sanitized,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
}

// MailThreadPhones represents array of mail.thread.phone model.
type MailThreadPhones []MailThreadPhone

// MailThreadPhoneModel is the odoo model name.
const MailThreadPhoneModel = "mail.thread.phone"

// Many2One convert MailThreadPhone to *Many2One.
func (mtp *MailThreadPhone) Many2One() *Many2One {
	return NewMany2One(mtp.Id.Get(), "")
}

// CreateMailThreadPhone creates a new mail.thread.phone model and returns its id.
func (c *Client) CreateMailThreadPhone(mtp *MailThreadPhone) (int64, error) {
	return c.Create(MailThreadPhoneModel, mtp)
}

// UpdateMailThreadPhone updates an existing mail.thread.phone record.
func (c *Client) UpdateMailThreadPhone(mtp *MailThreadPhone) error {
	return c.UpdateMailThreadPhones([]int64{mtp.Id.Get()}, mtp)
}

// UpdateMailThreadPhones updates existing mail.thread.phone records.
// All records (represented by ids) will be updated by mtp values.
func (c *Client) UpdateMailThreadPhones(ids []int64, mtp *MailThreadPhone) error {
	return c.Update(MailThreadPhoneModel, ids, mtp)
}

// DeleteMailThreadPhone deletes an existing mail.thread.phone record.
func (c *Client) DeleteMailThreadPhone(id int64) error {
	return c.DeleteMailThreadPhones([]int64{id})
}

// DeleteMailThreadPhones deletes existing mail.thread.phone records.
func (c *Client) DeleteMailThreadPhones(ids []int64) error {
	return c.Delete(MailThreadPhoneModel, ids)
}

// GetMailThreadPhone gets mail.thread.phone existing record.
func (c *Client) GetMailThreadPhone(id int64) (*MailThreadPhone, error) {
	mtps, err := c.GetMailThreadPhones([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtps != nil && len(*mtps) > 0 {
		return &((*mtps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.thread.phone not found", id)
}

// GetMailThreadPhones gets mail.thread.phone existing records.
func (c *Client) GetMailThreadPhones(ids []int64) (*MailThreadPhones, error) {
	mtps := &MailThreadPhones{}
	if err := c.Read(MailThreadPhoneModel, ids, nil, mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// FindMailThreadPhone finds mail.thread.phone record by querying it with criteria.
func (c *Client) FindMailThreadPhone(criteria *Criteria) (*MailThreadPhone, error) {
	mtps := &MailThreadPhones{}
	if err := c.SearchRead(MailThreadPhoneModel, criteria, NewOptions().Limit(1), mtps); err != nil {
		return nil, err
	}
	if mtps != nil && len(*mtps) > 0 {
		return &((*mtps)[0]), nil
	}
	return nil, fmt.Errorf("mail.thread.phone was not found")
}

// FindMailThreadPhones finds mail.thread.phone records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadPhones(criteria *Criteria, options *Options) (*MailThreadPhones, error) {
	mtps := &MailThreadPhones{}
	if err := c.SearchRead(MailThreadPhoneModel, criteria, options, mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// FindMailThreadPhoneIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadPhoneIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailThreadPhoneModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailThreadPhoneId finds record id by querying it with criteria.
func (c *Client) FindMailThreadPhoneId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailThreadPhoneModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.thread.phone was not found")
}
