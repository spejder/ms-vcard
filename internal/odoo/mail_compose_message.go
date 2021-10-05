package odoo

import (
	"fmt"
)

// MailComposeMessage represents mail.compose.message model.
type MailComposeMessage struct {
	ActiveDomain       *String    `xmlrpc:"active_domain,omptempty"`
	AddSign            *Bool      `xmlrpc:"add_sign,omptempty"`
	AttachmentIds      *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorId           *Many2One  `xmlrpc:"author_id,omptempty"`
	AutoDelete         *Bool      `xmlrpc:"auto_delete,omptempty"`
	AutoDeleteMessage  *Bool      `xmlrpc:"auto_delete_message,omptempty"`
	Body               *String    `xmlrpc:"body,omptempty"`
	CampaignId         *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CompositionMode    *Selection `xmlrpc:"composition_mode,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom          *String    `xmlrpc:"email_from,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	IsLog              *Bool      `xmlrpc:"is_log,omptempty"`
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	Layout             *String    `xmlrpc:"layout,omptempty"`
	MailActivityTypeId *Many2One  `xmlrpc:"mail_activity_type_id,omptempty"`
	MailingListIds     *Relation  `xmlrpc:"mailing_list_ids,omptempty"`
	MailServerId       *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	MassMailingId      *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	MassMailingName    *String    `xmlrpc:"mass_mailing_name,omptempty"`
	MessageType        *Selection `xmlrpc:"message_type,omptempty"`
	Model              *String    `xmlrpc:"model,omptempty"`
	NoAutoThread       *Bool      `xmlrpc:"no_auto_thread,omptempty"`
	Notify             *Bool      `xmlrpc:"notify,omptempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerIds         *Relation  `xmlrpc:"partner_ids,omptempty"`
	RecordName         *String    `xmlrpc:"record_name,omptempty"`
	ReplyTo            *String    `xmlrpc:"reply_to,omptempty"`
	ResId              *Int       `xmlrpc:"res_id,omptempty"`
	Subject            *String    `xmlrpc:"subject,omptempty"`
	SubtypeId          *Many2One  `xmlrpc:"subtype_id,omptempty"`
	TemplateId         *Many2One  `xmlrpc:"template_id,omptempty"`
	UseActiveDomain    *Bool      `xmlrpc:"use_active_domain,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailComposeMessages represents array of mail.compose.message model.
type MailComposeMessages []MailComposeMessage

// MailComposeMessageModel is the odoo model name.
const MailComposeMessageModel = "mail.compose.message"

// Many2One convert MailComposeMessage to *Many2One.
func (mcm *MailComposeMessage) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailComposeMessage creates a new mail.compose.message model and returns its id.
func (c *Client) CreateMailComposeMessage(mcm *MailComposeMessage) (int64, error) {
	return c.Create(MailComposeMessageModel, mcm)
}

// UpdateMailComposeMessage updates an existing mail.compose.message record.
func (c *Client) UpdateMailComposeMessage(mcm *MailComposeMessage) error {
	return c.UpdateMailComposeMessages([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailComposeMessages updates existing mail.compose.message records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailComposeMessages(ids []int64, mcm *MailComposeMessage) error {
	return c.Update(MailComposeMessageModel, ids, mcm)
}

// DeleteMailComposeMessage deletes an existing mail.compose.message record.
func (c *Client) DeleteMailComposeMessage(id int64) error {
	return c.DeleteMailComposeMessages([]int64{id})
}

// DeleteMailComposeMessages deletes existing mail.compose.message records.
func (c *Client) DeleteMailComposeMessages(ids []int64) error {
	return c.Delete(MailComposeMessageModel, ids)
}

// GetMailComposeMessage gets mail.compose.message existing record.
func (c *Client) GetMailComposeMessage(id int64) (*MailComposeMessage, error) {
	mcms, err := c.GetMailComposeMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.compose.message not found", id)
}

// GetMailComposeMessages gets mail.compose.message existing records.
func (c *Client) GetMailComposeMessages(ids []int64) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.Read(MailComposeMessageModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessage finds mail.compose.message record by querying it with criteria.
func (c *Client) FindMailComposeMessage(criteria *Criteria) (*MailComposeMessage, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("mail.compose.message was not found")
}

// FindMailComposeMessages finds mail.compose.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessages(criteria *Criteria, options *Options) (*MailComposeMessages, error) {
	mcms := &MailComposeMessages{}
	if err := c.SearchRead(MailComposeMessageModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposeMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposeMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailComposeMessageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailComposeMessageId finds record id by querying it with criteria.
func (c *Client) FindMailComposeMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailComposeMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.compose.message was not found")
}
