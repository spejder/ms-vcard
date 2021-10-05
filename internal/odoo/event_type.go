package odoo

import (
	"fmt"
)

// EventType represents event.type model.
type EventType struct {
	AutoConfirm            *Bool      `xmlrpc:"auto_confirm,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultHashtag         *String    `xmlrpc:"default_hashtag,omptempty"`
	DefaultRegistrationMax *Int       `xmlrpc:"default_registration_max,omptempty"`
	DefaultRegistrationMin *Int       `xmlrpc:"default_registration_min,omptempty"`
	DefaultTimezone        *Selection `xmlrpc:"default_timezone,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	EventTicketIds         *Relation  `xmlrpc:"event_ticket_ids,omptempty"`
	EventTypeMailIds       *Relation  `xmlrpc:"event_type_mail_ids,omptempty"`
	HasSeatsLimitation     *Bool      `xmlrpc:"has_seats_limitation,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	IsOnline               *Bool      `xmlrpc:"is_online,omptempty"`
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	UseHashtag             *Bool      `xmlrpc:"use_hashtag,omptempty"`
	UseMailSchedule        *Bool      `xmlrpc:"use_mail_schedule,omptempty"`
	UseTicketing           *Bool      `xmlrpc:"use_ticketing,omptempty"`
	UseTimezone            *Bool      `xmlrpc:"use_timezone,omptempty"`
	WebsiteMenu            *Bool      `xmlrpc:"website_menu,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventTypes represents array of event.type model.
type EventTypes []EventType

// EventTypeModel is the odoo model name.
const EventTypeModel = "event.type"

// Many2One convert EventType to *Many2One.
func (et *EventType) Many2One() *Many2One {
	return NewMany2One(et.Id.Get(), "")
}

// CreateEventType creates a new event.type model and returns its id.
func (c *Client) CreateEventType(et *EventType) (int64, error) {
	return c.Create(EventTypeModel, et)
}

// UpdateEventType updates an existing event.type record.
func (c *Client) UpdateEventType(et *EventType) error {
	return c.UpdateEventTypes([]int64{et.Id.Get()}, et)
}

// UpdateEventTypes updates existing event.type records.
// All records (represented by ids) will be updated by et values.
func (c *Client) UpdateEventTypes(ids []int64, et *EventType) error {
	return c.Update(EventTypeModel, ids, et)
}

// DeleteEventType deletes an existing event.type record.
func (c *Client) DeleteEventType(id int64) error {
	return c.DeleteEventTypes([]int64{id})
}

// DeleteEventTypes deletes existing event.type records.
func (c *Client) DeleteEventTypes(ids []int64) error {
	return c.Delete(EventTypeModel, ids)
}

// GetEventType gets event.type existing record.
func (c *Client) GetEventType(id int64) (*EventType, error) {
	ets, err := c.GetEventTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if ets != nil && len(*ets) > 0 {
		return &((*ets)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.type not found", id)
}

// GetEventTypes gets event.type existing records.
func (c *Client) GetEventTypes(ids []int64) (*EventTypes, error) {
	ets := &EventTypes{}
	if err := c.Read(EventTypeModel, ids, nil, ets); err != nil {
		return nil, err
	}
	return ets, nil
}

// FindEventType finds event.type record by querying it with criteria.
func (c *Client) FindEventType(criteria *Criteria) (*EventType, error) {
	ets := &EventTypes{}
	if err := c.SearchRead(EventTypeModel, criteria, NewOptions().Limit(1), ets); err != nil {
		return nil, err
	}
	if ets != nil && len(*ets) > 0 {
		return &((*ets)[0]), nil
	}
	return nil, fmt.Errorf("event.type was not found")
}

// FindEventTypes finds event.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypes(criteria *Criteria, options *Options) (*EventTypes, error) {
	ets := &EventTypes{}
	if err := c.SearchRead(EventTypeModel, criteria, options, ets); err != nil {
		return nil, err
	}
	return ets, nil
}

// FindEventTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventTypeId finds record id by querying it with criteria.
func (c *Client) FindEventTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.type was not found")
}
