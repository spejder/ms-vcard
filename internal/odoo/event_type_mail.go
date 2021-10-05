package odoo

import (
	"fmt"
)

// EventTypeMail represents event.type.mail model.
type EventTypeMail struct {
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	EventTypeId      *Many2One  `xmlrpc:"event_type_id,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	IntervalNbr      *Int       `xmlrpc:"interval_nbr,omptempty"`
	IntervalType     *Selection `xmlrpc:"interval_type,omptempty"`
	IntervalUnit     *Selection `xmlrpc:"interval_unit,omptempty"`
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	NotificationType *Selection `xmlrpc:"notification_type,omptempty"`
	SmsTemplateId    *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	TemplateId       *Many2One  `xmlrpc:"template_id,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventTypeMails represents array of event.type.mail model.
type EventTypeMails []EventTypeMail

// EventTypeMailModel is the odoo model name.
const EventTypeMailModel = "event.type.mail"

// Many2One convert EventTypeMail to *Many2One.
func (etm *EventTypeMail) Many2One() *Many2One {
	return NewMany2One(etm.Id.Get(), "")
}

// CreateEventTypeMail creates a new event.type.mail model and returns its id.
func (c *Client) CreateEventTypeMail(etm *EventTypeMail) (int64, error) {
	return c.Create(EventTypeMailModel, etm)
}

// UpdateEventTypeMail updates an existing event.type.mail record.
func (c *Client) UpdateEventTypeMail(etm *EventTypeMail) error {
	return c.UpdateEventTypeMails([]int64{etm.Id.Get()}, etm)
}

// UpdateEventTypeMails updates existing event.type.mail records.
// All records (represented by ids) will be updated by etm values.
func (c *Client) UpdateEventTypeMails(ids []int64, etm *EventTypeMail) error {
	return c.Update(EventTypeMailModel, ids, etm)
}

// DeleteEventTypeMail deletes an existing event.type.mail record.
func (c *Client) DeleteEventTypeMail(id int64) error {
	return c.DeleteEventTypeMails([]int64{id})
}

// DeleteEventTypeMails deletes existing event.type.mail records.
func (c *Client) DeleteEventTypeMails(ids []int64) error {
	return c.Delete(EventTypeMailModel, ids)
}

// GetEventTypeMail gets event.type.mail existing record.
func (c *Client) GetEventTypeMail(id int64) (*EventTypeMail, error) {
	etms, err := c.GetEventTypeMails([]int64{id})
	if err != nil {
		return nil, err
	}
	if etms != nil && len(*etms) > 0 {
		return &((*etms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.type.mail not found", id)
}

// GetEventTypeMails gets event.type.mail existing records.
func (c *Client) GetEventTypeMails(ids []int64) (*EventTypeMails, error) {
	etms := &EventTypeMails{}
	if err := c.Read(EventTypeMailModel, ids, nil, etms); err != nil {
		return nil, err
	}
	return etms, nil
}

// FindEventTypeMail finds event.type.mail record by querying it with criteria.
func (c *Client) FindEventTypeMail(criteria *Criteria) (*EventTypeMail, error) {
	etms := &EventTypeMails{}
	if err := c.SearchRead(EventTypeMailModel, criteria, NewOptions().Limit(1), etms); err != nil {
		return nil, err
	}
	if etms != nil && len(*etms) > 0 {
		return &((*etms)[0]), nil
	}
	return nil, fmt.Errorf("event.type.mail was not found")
}

// FindEventTypeMails finds event.type.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeMails(criteria *Criteria, options *Options) (*EventTypeMails, error) {
	etms := &EventTypeMails{}
	if err := c.SearchRead(EventTypeMailModel, criteria, options, etms); err != nil {
		return nil, err
	}
	return etms, nil
}

// FindEventTypeMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventTypeMailModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventTypeMailId finds record id by querying it with criteria.
func (c *Client) FindEventTypeMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTypeMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.type.mail was not found")
}
