package odoo

import (
	"fmt"
)

// MailThreadBlacklist represents mail.thread.blacklist model.
type MailThreadBlacklist struct {
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	EmailNormalized          *String   `xmlrpc:"email_normalized,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsBlacklisted            *Bool     `xmlrpc:"is_blacklisted,omptempty"`
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce            *Int      `xmlrpc:"message_bounce,omptempty"`
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
}

// MailThreadBlacklists represents array of mail.thread.blacklist model.
type MailThreadBlacklists []MailThreadBlacklist

// MailThreadBlacklistModel is the odoo model name.
const MailThreadBlacklistModel = "mail.thread.blacklist"

// Many2One convert MailThreadBlacklist to *Many2One.
func (mtb *MailThreadBlacklist) Many2One() *Many2One {
	return NewMany2One(mtb.Id.Get(), "")
}

// CreateMailThreadBlacklist creates a new mail.thread.blacklist model and returns its id.
func (c *Client) CreateMailThreadBlacklist(mtb *MailThreadBlacklist) (int64, error) {
	return c.Create(MailThreadBlacklistModel, mtb)
}

// UpdateMailThreadBlacklist updates an existing mail.thread.blacklist record.
func (c *Client) UpdateMailThreadBlacklist(mtb *MailThreadBlacklist) error {
	return c.UpdateMailThreadBlacklists([]int64{mtb.Id.Get()}, mtb)
}

// UpdateMailThreadBlacklists updates existing mail.thread.blacklist records.
// All records (represented by ids) will be updated by mtb values.
func (c *Client) UpdateMailThreadBlacklists(ids []int64, mtb *MailThreadBlacklist) error {
	return c.Update(MailThreadBlacklistModel, ids, mtb)
}

// DeleteMailThreadBlacklist deletes an existing mail.thread.blacklist record.
func (c *Client) DeleteMailThreadBlacklist(id int64) error {
	return c.DeleteMailThreadBlacklists([]int64{id})
}

// DeleteMailThreadBlacklists deletes existing mail.thread.blacklist records.
func (c *Client) DeleteMailThreadBlacklists(ids []int64) error {
	return c.Delete(MailThreadBlacklistModel, ids)
}

// GetMailThreadBlacklist gets mail.thread.blacklist existing record.
func (c *Client) GetMailThreadBlacklist(id int64) (*MailThreadBlacklist, error) {
	mtbs, err := c.GetMailThreadBlacklists([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtbs != nil && len(*mtbs) > 0 {
		return &((*mtbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.thread.blacklist not found", id)
}

// GetMailThreadBlacklists gets mail.thread.blacklist existing records.
func (c *Client) GetMailThreadBlacklists(ids []int64) (*MailThreadBlacklists, error) {
	mtbs := &MailThreadBlacklists{}
	if err := c.Read(MailThreadBlacklistModel, ids, nil, mtbs); err != nil {
		return nil, err
	}
	return mtbs, nil
}

// FindMailThreadBlacklist finds mail.thread.blacklist record by querying it with criteria.
func (c *Client) FindMailThreadBlacklist(criteria *Criteria) (*MailThreadBlacklist, error) {
	mtbs := &MailThreadBlacklists{}
	if err := c.SearchRead(MailThreadBlacklistModel, criteria, NewOptions().Limit(1), mtbs); err != nil {
		return nil, err
	}
	if mtbs != nil && len(*mtbs) > 0 {
		return &((*mtbs)[0]), nil
	}
	return nil, fmt.Errorf("mail.thread.blacklist was not found")
}

// FindMailThreadBlacklists finds mail.thread.blacklist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadBlacklists(criteria *Criteria, options *Options) (*MailThreadBlacklists, error) {
	mtbs := &MailThreadBlacklists{}
	if err := c.SearchRead(MailThreadBlacklistModel, criteria, options, mtbs); err != nil {
		return nil, err
	}
	return mtbs, nil
}

// FindMailThreadBlacklistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadBlacklistIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailThreadBlacklistModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailThreadBlacklistId finds record id by querying it with criteria.
func (c *Client) FindMailThreadBlacklistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailThreadBlacklistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.thread.blacklist was not found")
}
