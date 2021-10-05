package odoo

import (
	"fmt"
)

// SmsComposer represents sms.composer model.
type SmsComposer struct {
	ActiveDomain            *String    `xmlrpc:"active_domain,omptempty"`
	ActiveDomainCount       *Int       `xmlrpc:"active_domain_count,omptempty"`
	Body                    *String    `xmlrpc:"body,omptempty"`
	CompositionMode         *Selection `xmlrpc:"composition_mode,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	MailingId               *Many2One  `xmlrpc:"mailing_id,omptempty"`
	MassForceSend           *Bool      `xmlrpc:"mass_force_send,omptempty"`
	MassKeepLog             *Bool      `xmlrpc:"mass_keep_log,omptempty"`
	MassSmsAllowUnsubscribe *Bool      `xmlrpc:"mass_sms_allow_unsubscribe,omptempty"`
	MassUseBlacklist        *Bool      `xmlrpc:"mass_use_blacklist,omptempty"`
	NumberFieldName         *String    `xmlrpc:"number_field_name,omptempty"`
	Numbers                 *String    `xmlrpc:"numbers,omptempty"`
	PartnerIds              *Relation  `xmlrpc:"partner_ids,omptempty"`
	RecipientCount          *Int       `xmlrpc:"recipient_count,omptempty"`
	RecipientDescription    *String    `xmlrpc:"recipient_description,omptempty"`
	RecipientInvalidCount   *Int       `xmlrpc:"recipient_invalid_count,omptempty"`
	ResId                   *Int       `xmlrpc:"res_id,omptempty"`
	ResIds                  *String    `xmlrpc:"res_ids,omptempty"`
	ResIdsCount             *Int       `xmlrpc:"res_ids_count,omptempty"`
	ResModel                *String    `xmlrpc:"res_model,omptempty"`
	SanitizedNumbers        *String    `xmlrpc:"sanitized_numbers,omptempty"`
	TemplateId              *Many2One  `xmlrpc:"template_id,omptempty"`
	UseActiveDomain         *Bool      `xmlrpc:"use_active_domain,omptempty"`
	UtmCampaignId           *Many2One  `xmlrpc:"utm_campaign_id,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SmsComposers represents array of sms.composer model.
type SmsComposers []SmsComposer

// SmsComposerModel is the odoo model name.
const SmsComposerModel = "sms.composer"

// Many2One convert SmsComposer to *Many2One.
func (sc *SmsComposer) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSmsComposer creates a new sms.composer model and returns its id.
func (c *Client) CreateSmsComposer(sc *SmsComposer) (int64, error) {
	return c.Create(SmsComposerModel, sc)
}

// UpdateSmsComposer updates an existing sms.composer record.
func (c *Client) UpdateSmsComposer(sc *SmsComposer) error {
	return c.UpdateSmsComposers([]int64{sc.Id.Get()}, sc)
}

// UpdateSmsComposers updates existing sms.composer records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSmsComposers(ids []int64, sc *SmsComposer) error {
	return c.Update(SmsComposerModel, ids, sc)
}

// DeleteSmsComposer deletes an existing sms.composer record.
func (c *Client) DeleteSmsComposer(id int64) error {
	return c.DeleteSmsComposers([]int64{id})
}

// DeleteSmsComposers deletes existing sms.composer records.
func (c *Client) DeleteSmsComposers(ids []int64) error {
	return c.Delete(SmsComposerModel, ids)
}

// GetSmsComposer gets sms.composer existing record.
func (c *Client) GetSmsComposer(id int64) (*SmsComposer, error) {
	scs, err := c.GetSmsComposers([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.composer not found", id)
}

// GetSmsComposers gets sms.composer existing records.
func (c *Client) GetSmsComposers(ids []int64) (*SmsComposers, error) {
	scs := &SmsComposers{}
	if err := c.Read(SmsComposerModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSmsComposer finds sms.composer record by querying it with criteria.
func (c *Client) FindSmsComposer(criteria *Criteria) (*SmsComposer, error) {
	scs := &SmsComposers{}
	if err := c.SearchRead(SmsComposerModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("sms.composer was not found")
}

// FindSmsComposers finds sms.composer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsComposers(criteria *Criteria, options *Options) (*SmsComposers, error) {
	scs := &SmsComposers{}
	if err := c.SearchRead(SmsComposerModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSmsComposerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsComposerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsComposerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsComposerId finds record id by querying it with criteria.
func (c *Client) FindSmsComposerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsComposerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.composer was not found")
}
