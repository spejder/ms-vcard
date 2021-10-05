package odoo

import (
	"fmt"
)

// MailingMailing represents mailing.mailing model.
type MailingMailing struct {
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omptempty"`
	BodyArch                    *String    `xmlrpc:"body_arch,omptempty"`
	BodyHtml                    *String    `xmlrpc:"body_html,omptempty"`
	BodyPlaintext               *String    `xmlrpc:"body_plaintext,omptempty"`
	Bounced                     *Int       `xmlrpc:"bounced,omptempty"`
	BouncedRatio                *Int       `xmlrpc:"bounced_ratio,omptempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omptempty"`
	Clicked                     *Int       `xmlrpc:"clicked,omptempty"`
	ClicksRatio                 *Int       `xmlrpc:"clicks_ratio,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	ContactAbPc                 *Int       `xmlrpc:"contact_ab_pc,omptempty"`
	ContactListIds              *Relation  `xmlrpc:"contact_list_ids,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmLeadActivated            *Bool      `xmlrpc:"crm_lead_activated,omptempty"`
	CrmLeadCount                *Int       `xmlrpc:"crm_lead_count,omptempty"`
	CrmOpportunitiesCount       *Int       `xmlrpc:"crm_opportunities_count,omptempty"`
	Delivered                   *Int       `xmlrpc:"delivered,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom                   *String    `xmlrpc:"email_from,omptempty"`
	Expected                    *Int       `xmlrpc:"expected,omptempty"`
	Failed                      *Int       `xmlrpc:"failed,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	Ignored                     *Int       `xmlrpc:"ignored,omptempty"`
	KeepArchives                *Bool      `xmlrpc:"keep_archives,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	MailingDomain               *String    `xmlrpc:"mailing_domain,omptempty"`
	MailingModelId              *Many2One  `xmlrpc:"mailing_model_id,omptempty"`
	MailingModelName            *String    `xmlrpc:"mailing_model_name,omptempty"`
	MailingModelReal            *String    `xmlrpc:"mailing_model_real,omptempty"`
	MailingTraceIds             *Relation  `xmlrpc:"mailing_trace_ids,omptempty"`
	MailingType                 *Selection `xmlrpc:"mailing_type,omptempty"`
	MailServerId                *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	NextDeparture               *Time      `xmlrpc:"next_departure,omptempty"`
	Opened                      *Int       `xmlrpc:"opened,omptempty"`
	OpenedRatio                 *Int       `xmlrpc:"opened_ratio,omptempty"`
	ReceivedRatio               *Int       `xmlrpc:"received_ratio,omptempty"`
	Replied                     *Int       `xmlrpc:"replied,omptempty"`
	RepliedRatio                *Int       `xmlrpc:"replied_ratio,omptempty"`
	ReplyTo                     *String    `xmlrpc:"reply_to,omptempty"`
	ReplyToMode                 *Selection `xmlrpc:"reply_to_mode,omptempty"`
	SaleInvoicedAmount          *Int       `xmlrpc:"sale_invoiced_amount,omptempty"`
	SaleQuotationCount          *Int       `xmlrpc:"sale_quotation_count,omptempty"`
	Scheduled                   *Int       `xmlrpc:"scheduled,omptempty"`
	ScheduleDate                *Time      `xmlrpc:"schedule_date,omptempty"`
	Sent                        *Int       `xmlrpc:"sent,omptempty"`
	SentDate                    *Time      `xmlrpc:"sent_date,omptempty"`
	SmsAllowUnsubscribe         *Bool      `xmlrpc:"sms_allow_unsubscribe,omptempty"`
	SmsForceSend                *Bool      `xmlrpc:"sms_force_send,omptempty"`
	SmsHasInsufficientCredit    *Bool      `xmlrpc:"sms_has_insufficient_credit,omptempty"`
	SmsTemplateId               *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	Subject                     *String    `xmlrpc:"subject,omptempty"`
	Total                       *Int       `xmlrpc:"total,omptempty"`
	UniqueAbTesting             *Bool      `xmlrpc:"unique_ab_testing,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailingMailings represents array of mailing.mailing model.
type MailingMailings []MailingMailing

// MailingMailingModel is the odoo model name.
const MailingMailingModel = "mailing.mailing"

// Many2One convert MailingMailing to *Many2One.
func (mm *MailingMailing) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailingMailing creates a new mailing.mailing model and returns its id.
func (c *Client) CreateMailingMailing(mm *MailingMailing) (int64, error) {
	return c.Create(MailingMailingModel, mm)
}

// UpdateMailingMailing updates an existing mailing.mailing record.
func (c *Client) UpdateMailingMailing(mm *MailingMailing) error {
	return c.UpdateMailingMailings([]int64{mm.Id.Get()}, mm)
}

// UpdateMailingMailings updates existing mailing.mailing records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailingMailings(ids []int64, mm *MailingMailing) error {
	return c.Update(MailingMailingModel, ids, mm)
}

// DeleteMailingMailing deletes an existing mailing.mailing record.
func (c *Client) DeleteMailingMailing(id int64) error {
	return c.DeleteMailingMailings([]int64{id})
}

// DeleteMailingMailings deletes existing mailing.mailing records.
func (c *Client) DeleteMailingMailings(ids []int64) error {
	return c.Delete(MailingMailingModel, ids)
}

// GetMailingMailing gets mailing.mailing existing record.
func (c *Client) GetMailingMailing(id int64) (*MailingMailing, error) {
	mms, err := c.GetMailingMailings([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.mailing not found", id)
}

// GetMailingMailings gets mailing.mailing existing records.
func (c *Client) GetMailingMailings(ids []int64) (*MailingMailings, error) {
	mms := &MailingMailings{}
	if err := c.Read(MailingMailingModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailingMailing finds mailing.mailing record by querying it with criteria.
func (c *Client) FindMailingMailing(criteria *Criteria) (*MailingMailing, error) {
	mms := &MailingMailings{}
	if err := c.SearchRead(MailingMailingModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("mailing.mailing was not found")
}

// FindMailingMailings finds mailing.mailing records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailings(criteria *Criteria, options *Options) (*MailingMailings, error) {
	mms := &MailingMailings{}
	if err := c.SearchRead(MailingMailingModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailingMailingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingMailingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingMailingId finds record id by querying it with criteria.
func (c *Client) FindMailingMailingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingMailingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.mailing was not found")
}
