package odoo

import (
	"fmt"
)

// MailingTrace represents mailing.trace model.
type MailingTrace struct {
	Bounced       *Time      `xmlrpc:"bounced,omptempty"`
	CampaignId    *Many2One  `xmlrpc:"campaign_id,omptempty"`
	Clicked       *Time      `xmlrpc:"clicked,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Email         *String    `xmlrpc:"email,omptempty"`
	Exception     *Time      `xmlrpc:"exception,omptempty"`
	FailureType   *Selection `xmlrpc:"failure_type,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	Ignored       *Time      `xmlrpc:"ignored,omptempty"`
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	LinksClickIds *Relation  `xmlrpc:"links_click_ids,omptempty"`
	MailMailId    *Many2One  `xmlrpc:"mail_mail_id,omptempty"`
	MailMailIdInt *Int       `xmlrpc:"mail_mail_id_int,omptempty"`
	MassMailingId *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	MessageId     *String    `xmlrpc:"message_id,omptempty"`
	Model         *String    `xmlrpc:"model,omptempty"`
	Opened        *Time      `xmlrpc:"opened,omptempty"`
	Replied       *Time      `xmlrpc:"replied,omptempty"`
	ResId         *Int       `xmlrpc:"res_id,omptempty"`
	Scheduled     *Time      `xmlrpc:"scheduled,omptempty"`
	Sent          *Time      `xmlrpc:"sent,omptempty"`
	SmsCode       *String    `xmlrpc:"sms_code,omptempty"`
	SmsNumber     *String    `xmlrpc:"sms_number,omptempty"`
	SmsSmsId      *Many2One  `xmlrpc:"sms_sms_id,omptempty"`
	SmsSmsIdInt   *Int       `xmlrpc:"sms_sms_id_int,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
	StateUpdate   *Time      `xmlrpc:"state_update,omptempty"`
	TraceType     *Selection `xmlrpc:"trace_type,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailingTraces represents array of mailing.trace model.
type MailingTraces []MailingTrace

// MailingTraceModel is the odoo model name.
const MailingTraceModel = "mailing.trace"

// Many2One convert MailingTrace to *Many2One.
func (mt *MailingTrace) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMailingTrace creates a new mailing.trace model and returns its id.
func (c *Client) CreateMailingTrace(mt *MailingTrace) (int64, error) {
	return c.Create(MailingTraceModel, mt)
}

// UpdateMailingTrace updates an existing mailing.trace record.
func (c *Client) UpdateMailingTrace(mt *MailingTrace) error {
	return c.UpdateMailingTraces([]int64{mt.Id.Get()}, mt)
}

// UpdateMailingTraces updates existing mailing.trace records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMailingTraces(ids []int64, mt *MailingTrace) error {
	return c.Update(MailingTraceModel, ids, mt)
}

// DeleteMailingTrace deletes an existing mailing.trace record.
func (c *Client) DeleteMailingTrace(id int64) error {
	return c.DeleteMailingTraces([]int64{id})
}

// DeleteMailingTraces deletes existing mailing.trace records.
func (c *Client) DeleteMailingTraces(ids []int64) error {
	return c.Delete(MailingTraceModel, ids)
}

// GetMailingTrace gets mailing.trace existing record.
func (c *Client) GetMailingTrace(id int64) (*MailingTrace, error) {
	mts, err := c.GetMailingTraces([]int64{id})
	if err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.trace not found", id)
}

// GetMailingTraces gets mailing.trace existing records.
func (c *Client) GetMailingTraces(ids []int64) (*MailingTraces, error) {
	mts := &MailingTraces{}
	if err := c.Read(MailingTraceModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailingTrace finds mailing.trace record by querying it with criteria.
func (c *Client) FindMailingTrace(criteria *Criteria) (*MailingTrace, error) {
	mts := &MailingTraces{}
	if err := c.SearchRead(MailingTraceModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("mailing.trace was not found")
}

// FindMailingTraces finds mailing.trace records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraces(criteria *Criteria, options *Options) (*MailingTraces, error) {
	mts := &MailingTraces{}
	if err := c.SearchRead(MailingTraceModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMailingTraceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingTraceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingTraceId finds record id by querying it with criteria.
func (c *Client) FindMailingTraceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingTraceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.trace was not found")
}
