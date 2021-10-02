package odoo

import (
	"fmt"
)

// MailingContact represents mailing.contact model.
type MailingContact struct {
	CompanyName              *String   `xmlrpc:"company_name,omptempty"`
	CountryId                *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Email                    *String   `xmlrpc:"email,omptempty"`
	EmailNormalized          *String   `xmlrpc:"email_normalized,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsBlacklisted            *Bool     `xmlrpc:"is_blacklisted,omptempty"`
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	ListIds                  *Relation `xmlrpc:"list_ids,omptempty"`
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
	Mobile                   *String   `xmlrpc:"mobile,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	OptOut                   *Bool     `xmlrpc:"opt_out,omptempty"`
	PhoneBlacklisted         *Bool     `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneSanitized           *String   `xmlrpc:"phone_sanitized,omptempty"`
	SubscriptionListIds      *Relation `xmlrpc:"subscription_list_ids,omptempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omptempty"`
	TitleId                  *Many2One `xmlrpc:"title_id,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailingContacts represents array of mailing.contact model.
type MailingContacts []MailingContact

// MailingContactModel is the odoo model name.
const MailingContactModel = "mailing.contact"

// Many2One convert MailingContact to *Many2One.
func (mc *MailingContact) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMailingContact creates a new mailing.contact model and returns its id.
func (c *Client) CreateMailingContact(mc *MailingContact) (int64, error) {
	return c.Create(MailingContactModel, mc)
}

// UpdateMailingContact updates an existing mailing.contact record.
func (c *Client) UpdateMailingContact(mc *MailingContact) error {
	return c.UpdateMailingContacts([]int64{mc.Id.Get()}, mc)
}

// UpdateMailingContacts updates existing mailing.contact records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMailingContacts(ids []int64, mc *MailingContact) error {
	return c.Update(MailingContactModel, ids, mc)
}

// DeleteMailingContact deletes an existing mailing.contact record.
func (c *Client) DeleteMailingContact(id int64) error {
	return c.DeleteMailingContacts([]int64{id})
}

// DeleteMailingContacts deletes existing mailing.contact records.
func (c *Client) DeleteMailingContacts(ids []int64) error {
	return c.Delete(MailingContactModel, ids)
}

// GetMailingContact gets mailing.contact existing record.
func (c *Client) GetMailingContact(id int64) (*MailingContact, error) {
	mcs, err := c.GetMailingContacts([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.contact not found", id)
}

// GetMailingContacts gets mailing.contact existing records.
func (c *Client) GetMailingContacts(ids []int64) (*MailingContacts, error) {
	mcs := &MailingContacts{}
	if err := c.Read(MailingContactModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailingContact finds mailing.contact record by querying it with criteria.
func (c *Client) FindMailingContact(criteria *Criteria) (*MailingContact, error) {
	mcs := &MailingContacts{}
	if err := c.SearchRead(MailingContactModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("mailing.contact was not found")
}

// FindMailingContacts finds mailing.contact records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContacts(criteria *Criteria, options *Options) (*MailingContacts, error) {
	mcs := &MailingContacts{}
	if err := c.SearchRead(MailingContactModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailingContactIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingContactModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingContactId finds record id by querying it with criteria.
func (c *Client) FindMailingContactId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingContactModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.contact was not found")
}
