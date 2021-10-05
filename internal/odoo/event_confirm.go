package odoo

import (
	"fmt"
)

// EventConfirm represents event.confirm model.
type EventConfirm struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventConfirms represents array of event.confirm model.
type EventConfirms []EventConfirm

// EventConfirmModel is the odoo model name.
const EventConfirmModel = "event.confirm"

// Many2One convert EventConfirm to *Many2One.
func (ec *EventConfirm) Many2One() *Many2One {
	return NewMany2One(ec.Id.Get(), "")
}

// CreateEventConfirm creates a new event.confirm model and returns its id.
func (c *Client) CreateEventConfirm(ec *EventConfirm) (int64, error) {
	return c.Create(EventConfirmModel, ec)
}

// UpdateEventConfirm updates an existing event.confirm record.
func (c *Client) UpdateEventConfirm(ec *EventConfirm) error {
	return c.UpdateEventConfirms([]int64{ec.Id.Get()}, ec)
}

// UpdateEventConfirms updates existing event.confirm records.
// All records (represented by ids) will be updated by ec values.
func (c *Client) UpdateEventConfirms(ids []int64, ec *EventConfirm) error {
	return c.Update(EventConfirmModel, ids, ec)
}

// DeleteEventConfirm deletes an existing event.confirm record.
func (c *Client) DeleteEventConfirm(id int64) error {
	return c.DeleteEventConfirms([]int64{id})
}

// DeleteEventConfirms deletes existing event.confirm records.
func (c *Client) DeleteEventConfirms(ids []int64) error {
	return c.Delete(EventConfirmModel, ids)
}

// GetEventConfirm gets event.confirm existing record.
func (c *Client) GetEventConfirm(id int64) (*EventConfirm, error) {
	ecs, err := c.GetEventConfirms([]int64{id})
	if err != nil {
		return nil, err
	}
	if ecs != nil && len(*ecs) > 0 {
		return &((*ecs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.confirm not found", id)
}

// GetEventConfirms gets event.confirm existing records.
func (c *Client) GetEventConfirms(ids []int64) (*EventConfirms, error) {
	ecs := &EventConfirms{}
	if err := c.Read(EventConfirmModel, ids, nil, ecs); err != nil {
		return nil, err
	}
	return ecs, nil
}

// FindEventConfirm finds event.confirm record by querying it with criteria.
func (c *Client) FindEventConfirm(criteria *Criteria) (*EventConfirm, error) {
	ecs := &EventConfirms{}
	if err := c.SearchRead(EventConfirmModel, criteria, NewOptions().Limit(1), ecs); err != nil {
		return nil, err
	}
	if ecs != nil && len(*ecs) > 0 {
		return &((*ecs)[0]), nil
	}
	return nil, fmt.Errorf("event.confirm was not found")
}

// FindEventConfirms finds event.confirm records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventConfirms(criteria *Criteria, options *Options) (*EventConfirms, error) {
	ecs := &EventConfirms{}
	if err := c.SearchRead(EventConfirmModel, criteria, options, ecs); err != nil {
		return nil, err
	}
	return ecs, nil
}

// FindEventConfirmIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventConfirmIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventConfirmModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventConfirmId finds record id by querying it with criteria.
func (c *Client) FindEventConfirmId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventConfirmModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.confirm was not found")
}
