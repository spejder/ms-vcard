package odoo

import (
	"fmt"
)

// MailBlacklist represents mail.blacklist model.
type MailBlacklist struct {
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Email                    *String   `xmlrpc:"email,omptempty"`
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
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailBlacklists represents array of mail.blacklist model.
type MailBlacklists []MailBlacklist

// MailBlacklistModel is the odoo model name.
const MailBlacklistModel = "mail.blacklist"

// Many2One convert MailBlacklist to *Many2One.
func (mb *MailBlacklist) Many2One() *Many2One {
	return NewMany2One(mb.Id.Get(), "")
}

// CreateMailBlacklist creates a new mail.blacklist model and returns its id.
func (c *Client) CreateMailBlacklist(mb *MailBlacklist) (int64, error) {
	return c.Create(MailBlacklistModel, mb)
}

// UpdateMailBlacklist updates an existing mail.blacklist record.
func (c *Client) UpdateMailBlacklist(mb *MailBlacklist) error {
	return c.UpdateMailBlacklists([]int64{mb.Id.Get()}, mb)
}

// UpdateMailBlacklists updates existing mail.blacklist records.
// All records (represented by ids) will be updated by mb values.
func (c *Client) UpdateMailBlacklists(ids []int64, mb *MailBlacklist) error {
	return c.Update(MailBlacklistModel, ids, mb)
}

// DeleteMailBlacklist deletes an existing mail.blacklist record.
func (c *Client) DeleteMailBlacklist(id int64) error {
	return c.DeleteMailBlacklists([]int64{id})
}

// DeleteMailBlacklists deletes existing mail.blacklist records.
func (c *Client) DeleteMailBlacklists(ids []int64) error {
	return c.Delete(MailBlacklistModel, ids)
}

// GetMailBlacklist gets mail.blacklist existing record.
func (c *Client) GetMailBlacklist(id int64) (*MailBlacklist, error) {
	mbs, err := c.GetMailBlacklists([]int64{id})
	if err != nil {
		return nil, err
	}
	if mbs != nil && len(*mbs) > 0 {
		return &((*mbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.blacklist not found", id)
}

// GetMailBlacklists gets mail.blacklist existing records.
func (c *Client) GetMailBlacklists(ids []int64) (*MailBlacklists, error) {
	mbs := &MailBlacklists{}
	if err := c.Read(MailBlacklistModel, ids, nil, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMailBlacklist finds mail.blacklist record by querying it with criteria.
func (c *Client) FindMailBlacklist(criteria *Criteria) (*MailBlacklist, error) {
	mbs := &MailBlacklists{}
	if err := c.SearchRead(MailBlacklistModel, criteria, NewOptions().Limit(1), mbs); err != nil {
		return nil, err
	}
	if mbs != nil && len(*mbs) > 0 {
		return &((*mbs)[0]), nil
	}
	return nil, fmt.Errorf("mail.blacklist was not found")
}

// FindMailBlacklists finds mail.blacklist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBlacklists(criteria *Criteria, options *Options) (*MailBlacklists, error) {
	mbs := &MailBlacklists{}
	if err := c.SearchRead(MailBlacklistModel, criteria, options, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMailBlacklistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBlacklistIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailBlacklistModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailBlacklistId finds record id by querying it with criteria.
func (c *Client) FindMailBlacklistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailBlacklistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.blacklist was not found")
}
