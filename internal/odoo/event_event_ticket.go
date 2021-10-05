package odoo

import (
	"fmt"
)

// EventEventTicket represents event.event.ticket model.
type EventEventTicket struct {
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	Deadline          *Time      `xmlrpc:"deadline,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	EventId           *Many2One  `xmlrpc:"event_id,omptempty"`
	EventTypeId       *Many2One  `xmlrpc:"event_type_id,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	IsExpired         *Bool      `xmlrpc:"is_expired,omptempty"`
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	Price             *Float     `xmlrpc:"price,omptempty"`
	PriceReduce       *Float     `xmlrpc:"price_reduce,omptempty"`
	PriceReduceTaxinc *Float     `xmlrpc:"price_reduce_taxinc,omptempty"`
	ProductId         *Many2One  `xmlrpc:"product_id,omptempty"`
	RegistrationIds   *Relation  `xmlrpc:"registration_ids,omptempty"`
	SeatsAvailability *Selection `xmlrpc:"seats_availability,omptempty"`
	SeatsAvailable    *Int       `xmlrpc:"seats_available,omptempty"`
	SeatsMax          *Int       `xmlrpc:"seats_max,omptempty"`
	SeatsReserved     *Int       `xmlrpc:"seats_reserved,omptempty"`
	SeatsUnconfirmed  *Int       `xmlrpc:"seats_unconfirmed,omptempty"`
	SeatsUsed         *Int       `xmlrpc:"seats_used,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventEventTickets represents array of event.event.ticket model.
type EventEventTickets []EventEventTicket

// EventEventTicketModel is the odoo model name.
const EventEventTicketModel = "event.event.ticket"

// Many2One convert EventEventTicket to *Many2One.
func (eet *EventEventTicket) Many2One() *Many2One {
	return NewMany2One(eet.Id.Get(), "")
}

// CreateEventEventTicket creates a new event.event.ticket model and returns its id.
func (c *Client) CreateEventEventTicket(eet *EventEventTicket) (int64, error) {
	return c.Create(EventEventTicketModel, eet)
}

// UpdateEventEventTicket updates an existing event.event.ticket record.
func (c *Client) UpdateEventEventTicket(eet *EventEventTicket) error {
	return c.UpdateEventEventTickets([]int64{eet.Id.Get()}, eet)
}

// UpdateEventEventTickets updates existing event.event.ticket records.
// All records (represented by ids) will be updated by eet values.
func (c *Client) UpdateEventEventTickets(ids []int64, eet *EventEventTicket) error {
	return c.Update(EventEventTicketModel, ids, eet)
}

// DeleteEventEventTicket deletes an existing event.event.ticket record.
func (c *Client) DeleteEventEventTicket(id int64) error {
	return c.DeleteEventEventTickets([]int64{id})
}

// DeleteEventEventTickets deletes existing event.event.ticket records.
func (c *Client) DeleteEventEventTickets(ids []int64) error {
	return c.Delete(EventEventTicketModel, ids)
}

// GetEventEventTicket gets event.event.ticket existing record.
func (c *Client) GetEventEventTicket(id int64) (*EventEventTicket, error) {
	eets, err := c.GetEventEventTickets([]int64{id})
	if err != nil {
		return nil, err
	}
	if eets != nil && len(*eets) > 0 {
		return &((*eets)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.event.ticket not found", id)
}

// GetEventEventTickets gets event.event.ticket existing records.
func (c *Client) GetEventEventTickets(ids []int64) (*EventEventTickets, error) {
	eets := &EventEventTickets{}
	if err := c.Read(EventEventTicketModel, ids, nil, eets); err != nil {
		return nil, err
	}
	return eets, nil
}

// FindEventEventTicket finds event.event.ticket record by querying it with criteria.
func (c *Client) FindEventEventTicket(criteria *Criteria) (*EventEventTicket, error) {
	eets := &EventEventTickets{}
	if err := c.SearchRead(EventEventTicketModel, criteria, NewOptions().Limit(1), eets); err != nil {
		return nil, err
	}
	if eets != nil && len(*eets) > 0 {
		return &((*eets)[0]), nil
	}
	return nil, fmt.Errorf("event.event.ticket was not found")
}

// FindEventEventTickets finds event.event.ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventTickets(criteria *Criteria, options *Options) (*EventEventTickets, error) {
	eets := &EventEventTickets{}
	if err := c.SearchRead(EventEventTicketModel, criteria, options, eets); err != nil {
		return nil, err
	}
	return eets, nil
}

// FindEventEventTicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventTicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventEventTicketModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventEventTicketId finds record id by querying it with criteria.
func (c *Client) FindEventEventTicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventEventTicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.event.ticket was not found")
}
