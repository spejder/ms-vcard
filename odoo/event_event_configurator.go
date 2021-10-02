package odoo

import (
	"fmt"
)

// EventEventConfigurator represents event.event.configurator model.
type EventEventConfigurator struct {
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	EventId       *Many2One `xmlrpc:"event_id,omptempty"`
	EventTicketId *Many2One `xmlrpc:"event_ticket_id,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	ProductId     *Many2One `xmlrpc:"product_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventEventConfigurators represents array of event.event.configurator model.
type EventEventConfigurators []EventEventConfigurator

// EventEventConfiguratorModel is the odoo model name.
const EventEventConfiguratorModel = "event.event.configurator"

// Many2One convert EventEventConfigurator to *Many2One.
func (eec *EventEventConfigurator) Many2One() *Many2One {
	return NewMany2One(eec.Id.Get(), "")
}

// CreateEventEventConfigurator creates a new event.event.configurator model and returns its id.
func (c *Client) CreateEventEventConfigurator(eec *EventEventConfigurator) (int64, error) {
	return c.Create(EventEventConfiguratorModel, eec)
}

// UpdateEventEventConfigurator updates an existing event.event.configurator record.
func (c *Client) UpdateEventEventConfigurator(eec *EventEventConfigurator) error {
	return c.UpdateEventEventConfigurators([]int64{eec.Id.Get()}, eec)
}

// UpdateEventEventConfigurators updates existing event.event.configurator records.
// All records (represented by ids) will be updated by eec values.
func (c *Client) UpdateEventEventConfigurators(ids []int64, eec *EventEventConfigurator) error {
	return c.Update(EventEventConfiguratorModel, ids, eec)
}

// DeleteEventEventConfigurator deletes an existing event.event.configurator record.
func (c *Client) DeleteEventEventConfigurator(id int64) error {
	return c.DeleteEventEventConfigurators([]int64{id})
}

// DeleteEventEventConfigurators deletes existing event.event.configurator records.
func (c *Client) DeleteEventEventConfigurators(ids []int64) error {
	return c.Delete(EventEventConfiguratorModel, ids)
}

// GetEventEventConfigurator gets event.event.configurator existing record.
func (c *Client) GetEventEventConfigurator(id int64) (*EventEventConfigurator, error) {
	eecs, err := c.GetEventEventConfigurators([]int64{id})
	if err != nil {
		return nil, err
	}
	if eecs != nil && len(*eecs) > 0 {
		return &((*eecs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.event.configurator not found", id)
}

// GetEventEventConfigurators gets event.event.configurator existing records.
func (c *Client) GetEventEventConfigurators(ids []int64) (*EventEventConfigurators, error) {
	eecs := &EventEventConfigurators{}
	if err := c.Read(EventEventConfiguratorModel, ids, nil, eecs); err != nil {
		return nil, err
	}
	return eecs, nil
}

// FindEventEventConfigurator finds event.event.configurator record by querying it with criteria.
func (c *Client) FindEventEventConfigurator(criteria *Criteria) (*EventEventConfigurator, error) {
	eecs := &EventEventConfigurators{}
	if err := c.SearchRead(EventEventConfiguratorModel, criteria, NewOptions().Limit(1), eecs); err != nil {
		return nil, err
	}
	if eecs != nil && len(*eecs) > 0 {
		return &((*eecs)[0]), nil
	}
	return nil, fmt.Errorf("event.event.configurator was not found")
}

// FindEventEventConfigurators finds event.event.configurator records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventConfigurators(criteria *Criteria, options *Options) (*EventEventConfigurators, error) {
	eecs := &EventEventConfigurators{}
	if err := c.SearchRead(EventEventConfiguratorModel, criteria, options, eecs); err != nil {
		return nil, err
	}
	return eecs, nil
}

// FindEventEventConfiguratorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventConfiguratorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventEventConfiguratorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventEventConfiguratorId finds record id by querying it with criteria.
func (c *Client) FindEventEventConfiguratorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventEventConfiguratorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.event.configurator was not found")
}
