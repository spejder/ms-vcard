package odoo

import (
	"fmt"
)

// MailMessage represents mail.message model.
type MailMessage struct {
	AddSign            *Bool      `xmlrpc:"add_sign,omptempty"`
	AttachmentIds      *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorAvatar       *String    `xmlrpc:"author_avatar,omptempty"`
	AuthorId           *Many2One  `xmlrpc:"author_id,omptempty"`
	Body               *String    `xmlrpc:"body,omptempty"`
	CannedResponseIds  *Relation  `xmlrpc:"canned_response_ids,omptempty"`
	ChannelIds         *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds           *Relation  `xmlrpc:"child_ids,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date               *Time      `xmlrpc:"date,omptempty"`
	Description        *String    `xmlrpc:"description,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom          *String    `xmlrpc:"email_from,omptempty"`
	EmailLayoutXmlid   *String    `xmlrpc:"email_layout_xmlid,omptempty"`
	HasError           *Bool      `xmlrpc:"has_error,omptempty"`
	HasSmsError        *Bool      `xmlrpc:"has_sms_error,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	LetterIds          *Relation  `xmlrpc:"letter_ids,omptempty"`
	MailActivityTypeId *Many2One  `xmlrpc:"mail_activity_type_id,omptempty"`
	MailIds            *Relation  `xmlrpc:"mail_ids,omptempty"`
	MailServerId       *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	MessageId          *String    `xmlrpc:"message_id,omptempty"`
	MessageType        *Selection `xmlrpc:"message_type,omptempty"`
	Model              *String    `xmlrpc:"model,omptempty"`
	ModerationStatus   *Selection `xmlrpc:"moderation_status,omptempty"`
	ModeratorId        *Many2One  `xmlrpc:"moderator_id,omptempty"`
	Needaction         *Bool      `xmlrpc:"needaction,omptempty"`
	NeedModeration     *Bool      `xmlrpc:"need_moderation,omptempty"`
	NoAutoThread       *Bool      `xmlrpc:"no_auto_thread,omptempty"`
	NotificationIds    *Relation  `xmlrpc:"notification_ids,omptempty"`
	NotifiedPartnerIds *Relation  `xmlrpc:"notified_partner_ids,omptempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerIds         *Relation  `xmlrpc:"partner_ids,omptempty"`
	RatingIds          *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingValue        *Float     `xmlrpc:"rating_value,omptempty"`
	RecordName         *String    `xmlrpc:"record_name,omptempty"`
	ReplyTo            *String    `xmlrpc:"reply_to,omptempty"`
	ResId              *Many2One  `xmlrpc:"res_id,omptempty"`
	SnailmailError     *Bool      `xmlrpc:"snailmail_error,omptempty"`
	SnailmailStatus    *String    `xmlrpc:"snailmail_status,omptempty"`
	Starred            *Bool      `xmlrpc:"starred,omptempty"`
	StarredPartnerIds  *Relation  `xmlrpc:"starred_partner_ids,omptempty"`
	Subject            *String    `xmlrpc:"subject,omptempty"`
	SubtypeId          *Many2One  `xmlrpc:"subtype_id,omptempty"`
	TrackingValueIds   *Relation  `xmlrpc:"tracking_value_ids,omptempty"`
	WebsitePublished   *Bool      `xmlrpc:"website_published,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailMessages represents array of mail.message model.
type MailMessages []MailMessage

// MailMessageModel is the odoo model name.
const MailMessageModel = "mail.message"

// Many2One convert MailMessage to *Many2One.
func (mm *MailMessage) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMessage creates a new mail.message model and returns its id.
func (c *Client) CreateMailMessage(mm *MailMessage) (int64, error) {
	return c.Create(MailMessageModel, mm)
}

// UpdateMailMessage updates an existing mail.message record.
func (c *Client) UpdateMailMessage(mm *MailMessage) error {
	return c.UpdateMailMessages([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMessages updates existing mail.message records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMessages(ids []int64, mm *MailMessage) error {
	return c.Update(MailMessageModel, ids, mm)
}

// DeleteMailMessage deletes an existing mail.message record.
func (c *Client) DeleteMailMessage(id int64) error {
	return c.DeleteMailMessages([]int64{id})
}

// DeleteMailMessages deletes existing mail.message records.
func (c *Client) DeleteMailMessages(ids []int64) error {
	return c.Delete(MailMessageModel, ids)
}

// GetMailMessage gets mail.message existing record.
func (c *Client) GetMailMessage(id int64) (*MailMessage, error) {
	mms, err := c.GetMailMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.message not found", id)
}

// GetMailMessages gets mail.message existing records.
func (c *Client) GetMailMessages(ids []int64) (*MailMessages, error) {
	mms := &MailMessages{}
	if err := c.Read(MailMessageModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMessage finds mail.message record by querying it with criteria.
func (c *Client) FindMailMessage(criteria *Criteria) (*MailMessage, error) {
	mms := &MailMessages{}
	if err := c.SearchRead(MailMessageModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("mail.message was not found")
}

// FindMailMessages finds mail.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessages(criteria *Criteria, options *Options) (*MailMessages, error) {
	mms := &MailMessages{}
	if err := c.SearchRead(MailMessageModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMessageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMessageId finds record id by querying it with criteria.
func (c *Client) FindMailMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.message was not found")
}
