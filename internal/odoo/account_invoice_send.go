package odoo

import (
	"fmt"
)

// AccountInvoiceSend represents account.invoice.send model.
type AccountInvoiceSend struct {
	ActiveDomain        *String    `xmlrpc:"active_domain,omptempty"`
	AddSign             *Bool      `xmlrpc:"add_sign,omptempty"`
	AttachmentIds       *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorId            *Many2One  `xmlrpc:"author_id,omptempty"`
	AutoDelete          *Bool      `xmlrpc:"auto_delete,omptempty"`
	AutoDeleteMessage   *Bool      `xmlrpc:"auto_delete_message,omptempty"`
	Body                *String    `xmlrpc:"body,omptempty"`
	CampaignId          *Many2One  `xmlrpc:"campaign_id,omptempty"`
	ComposerId          *Many2One  `xmlrpc:"composer_id,omptempty"`
	CompositionMode     *Selection `xmlrpc:"composition_mode,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom           *String    `xmlrpc:"email_from,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InvalidAddresses    *Int       `xmlrpc:"invalid_addresses,omptempty"`
	InvalidInvoiceIds   *Relation  `xmlrpc:"invalid_invoice_ids,omptempty"`
	InvoiceIds          *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceWithoutEmail *String    `xmlrpc:"invoice_without_email,omptempty"`
	IsEmail             *Bool      `xmlrpc:"is_email,omptempty"`
	IsLog               *Bool      `xmlrpc:"is_log,omptempty"`
	IsPrint             *Bool      `xmlrpc:"is_print,omptempty"`
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Layout              *String    `xmlrpc:"layout,omptempty"`
	MailActivityTypeId  *Many2One  `xmlrpc:"mail_activity_type_id,omptempty"`
	MailingListIds      *Relation  `xmlrpc:"mailing_list_ids,omptempty"`
	MailServerId        *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	MassMailingId       *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	MassMailingName     *String    `xmlrpc:"mass_mailing_name,omptempty"`
	MessageType         *Selection `xmlrpc:"message_type,omptempty"`
	Model               *String    `xmlrpc:"model,omptempty"`
	NoAutoThread        *Bool      `xmlrpc:"no_auto_thread,omptempty"`
	Notify              *Bool      `xmlrpc:"notify,omptempty"`
	ParentId            *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerIds          *Relation  `xmlrpc:"partner_ids,omptempty"`
	Printed             *Bool      `xmlrpc:"printed,omptempty"`
	RecordName          *String    `xmlrpc:"record_name,omptempty"`
	ReplyTo             *String    `xmlrpc:"reply_to,omptempty"`
	ResId               *Int       `xmlrpc:"res_id,omptempty"`
	SnailmailCost       *Float     `xmlrpc:"snailmail_cost,omptempty"`
	SnailmailIsLetter   *Bool      `xmlrpc:"snailmail_is_letter,omptempty"`
	Subject             *String    `xmlrpc:"subject,omptempty"`
	SubtypeId           *Many2One  `xmlrpc:"subtype_id,omptempty"`
	TemplateId          *Many2One  `xmlrpc:"template_id,omptempty"`
	UseActiveDomain     *Bool      `xmlrpc:"use_active_domain,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountInvoiceSends represents array of account.invoice.send model.
type AccountInvoiceSends []AccountInvoiceSend

// AccountInvoiceSendModel is the odoo model name.
const AccountInvoiceSendModel = "account.invoice.send"

// Many2One convert AccountInvoiceSend to *Many2One.
func (ais *AccountInvoiceSend) Many2One() *Many2One {
	return NewMany2One(ais.Id.Get(), "")
}

// CreateAccountInvoiceSend creates a new account.invoice.send model and returns its id.
func (c *Client) CreateAccountInvoiceSend(ais *AccountInvoiceSend) (int64, error) {
	return c.Create(AccountInvoiceSendModel, ais)
}

// UpdateAccountInvoiceSend updates an existing account.invoice.send record.
func (c *Client) UpdateAccountInvoiceSend(ais *AccountInvoiceSend) error {
	return c.UpdateAccountInvoiceSends([]int64{ais.Id.Get()}, ais)
}

// UpdateAccountInvoiceSends updates existing account.invoice.send records.
// All records (represented by ids) will be updated by ais values.
func (c *Client) UpdateAccountInvoiceSends(ids []int64, ais *AccountInvoiceSend) error {
	return c.Update(AccountInvoiceSendModel, ids, ais)
}

// DeleteAccountInvoiceSend deletes an existing account.invoice.send record.
func (c *Client) DeleteAccountInvoiceSend(id int64) error {
	return c.DeleteAccountInvoiceSends([]int64{id})
}

// DeleteAccountInvoiceSends deletes existing account.invoice.send records.
func (c *Client) DeleteAccountInvoiceSends(ids []int64) error {
	return c.Delete(AccountInvoiceSendModel, ids)
}

// GetAccountInvoiceSend gets account.invoice.send existing record.
func (c *Client) GetAccountInvoiceSend(id int64) (*AccountInvoiceSend, error) {
	aiss, err := c.GetAccountInvoiceSends([]int64{id})
	if err != nil {
		return nil, err
	}
	if aiss != nil && len(*aiss) > 0 {
		return &((*aiss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice.send not found", id)
}

// GetAccountInvoiceSends gets account.invoice.send existing records.
func (c *Client) GetAccountInvoiceSends(ids []int64) (*AccountInvoiceSends, error) {
	aiss := &AccountInvoiceSends{}
	if err := c.Read(AccountInvoiceSendModel, ids, nil, aiss); err != nil {
		return nil, err
	}
	return aiss, nil
}

// FindAccountInvoiceSend finds account.invoice.send record by querying it with criteria.
func (c *Client) FindAccountInvoiceSend(criteria *Criteria) (*AccountInvoiceSend, error) {
	aiss := &AccountInvoiceSends{}
	if err := c.SearchRead(AccountInvoiceSendModel, criteria, NewOptions().Limit(1), aiss); err != nil {
		return nil, err
	}
	if aiss != nil && len(*aiss) > 0 {
		return &((*aiss)[0]), nil
	}
	return nil, fmt.Errorf("account.invoice.send was not found")
}

// FindAccountInvoiceSends finds account.invoice.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceSends(criteria *Criteria, options *Options) (*AccountInvoiceSends, error) {
	aiss := &AccountInvoiceSends{}
	if err := c.SearchRead(AccountInvoiceSendModel, criteria, options, aiss); err != nil {
		return nil, err
	}
	return aiss, nil
}

// FindAccountInvoiceSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceSendModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceSendId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.invoice.send was not found")
}
