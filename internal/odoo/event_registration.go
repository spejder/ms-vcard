package odoo

import (
	"fmt"
)

// EventRegistration represents event.registration model.
type EventRegistration struct {
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateClosed                  *Time      `xmlrpc:"date_closed,omptempty"`
	DateOpen                    *Time      `xmlrpc:"date_open,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Email                       *String    `xmlrpc:"email,omptempty"`
	EventBeginDate              *Time      `xmlrpc:"event_begin_date,omptempty"`
	EventEndDate                *Time      `xmlrpc:"event_end_date,omptempty"`
	EventId                     *Many2One  `xmlrpc:"event_id,omptempty"`
	EventTicketId               *Many2One  `xmlrpc:"event_ticket_id,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
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
	Mobile                      *String    `xmlrpc:"mobile,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	Origin                      *String    `xmlrpc:"origin,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	Phone                       *String    `xmlrpc:"phone,omptempty"`
	SaleOrderId                 *Many2One  `xmlrpc:"sale_order_id,omptempty"`
	SaleOrderLineId             *Many2One  `xmlrpc:"sale_order_line_id,omptempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventRegistrations represents array of event.registration model.
type EventRegistrations []EventRegistration

// EventRegistrationModel is the odoo model name.
const EventRegistrationModel = "event.registration"

// Many2One convert EventRegistration to *Many2One.
func (er *EventRegistration) Many2One() *Many2One {
	return NewMany2One(er.Id.Get(), "")
}

// CreateEventRegistration creates a new event.registration model and returns its id.
func (c *Client) CreateEventRegistration(er *EventRegistration) (int64, error) {
	return c.Create(EventRegistrationModel, er)
}

// UpdateEventRegistration updates an existing event.registration record.
func (c *Client) UpdateEventRegistration(er *EventRegistration) error {
	return c.UpdateEventRegistrations([]int64{er.Id.Get()}, er)
}

// UpdateEventRegistrations updates existing event.registration records.
// All records (represented by ids) will be updated by er values.
func (c *Client) UpdateEventRegistrations(ids []int64, er *EventRegistration) error {
	return c.Update(EventRegistrationModel, ids, er)
}

// DeleteEventRegistration deletes an existing event.registration record.
func (c *Client) DeleteEventRegistration(id int64) error {
	return c.DeleteEventRegistrations([]int64{id})
}

// DeleteEventRegistrations deletes existing event.registration records.
func (c *Client) DeleteEventRegistrations(ids []int64) error {
	return c.Delete(EventRegistrationModel, ids)
}

// GetEventRegistration gets event.registration existing record.
func (c *Client) GetEventRegistration(id int64) (*EventRegistration, error) {
	ers, err := c.GetEventRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	if ers != nil && len(*ers) > 0 {
		return &((*ers)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.registration not found", id)
}

// GetEventRegistrations gets event.registration existing records.
func (c *Client) GetEventRegistrations(ids []int64) (*EventRegistrations, error) {
	ers := &EventRegistrations{}
	if err := c.Read(EventRegistrationModel, ids, nil, ers); err != nil {
		return nil, err
	}
	return ers, nil
}

// FindEventRegistration finds event.registration record by querying it with criteria.
func (c *Client) FindEventRegistration(criteria *Criteria) (*EventRegistration, error) {
	ers := &EventRegistrations{}
	if err := c.SearchRead(EventRegistrationModel, criteria, NewOptions().Limit(1), ers); err != nil {
		return nil, err
	}
	if ers != nil && len(*ers) > 0 {
		return &((*ers)[0]), nil
	}
	return nil, fmt.Errorf("event.registration was not found")
}

// FindEventRegistrations finds event.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrations(criteria *Criteria, options *Options) (*EventRegistrations, error) {
	ers := &EventRegistrations{}
	if err := c.SearchRead(EventRegistrationModel, criteria, options, ers); err != nil {
		return nil, err
	}
	return ers, nil
}

// FindEventRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventRegistrationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventRegistrationId finds record id by querying it with criteria.
func (c *Client) FindEventRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.registration was not found")
}
