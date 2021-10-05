package odoo

import (
	"fmt"
)

// EventMailRegistration represents event.mail.registration model.
type EventMailRegistration struct {
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	MailSent       *Bool     `xmlrpc:"mail_sent,omptempty"`
	RegistrationId *Many2One `xmlrpc:"registration_id,omptempty"`
	ScheduledDate  *Time     `xmlrpc:"scheduled_date,omptempty"`
	SchedulerId    *Many2One `xmlrpc:"scheduler_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventMailRegistrations represents array of event.mail.registration model.
type EventMailRegistrations []EventMailRegistration

// EventMailRegistrationModel is the odoo model name.
const EventMailRegistrationModel = "event.mail.registration"

// Many2One convert EventMailRegistration to *Many2One.
func (emr *EventMailRegistration) Many2One() *Many2One {
	return NewMany2One(emr.Id.Get(), "")
}

// CreateEventMailRegistration creates a new event.mail.registration model and returns its id.
func (c *Client) CreateEventMailRegistration(emr *EventMailRegistration) (int64, error) {
	return c.Create(EventMailRegistrationModel, emr)
}

// UpdateEventMailRegistration updates an existing event.mail.registration record.
func (c *Client) UpdateEventMailRegistration(emr *EventMailRegistration) error {
	return c.UpdateEventMailRegistrations([]int64{emr.Id.Get()}, emr)
}

// UpdateEventMailRegistrations updates existing event.mail.registration records.
// All records (represented by ids) will be updated by emr values.
func (c *Client) UpdateEventMailRegistrations(ids []int64, emr *EventMailRegistration) error {
	return c.Update(EventMailRegistrationModel, ids, emr)
}

// DeleteEventMailRegistration deletes an existing event.mail.registration record.
func (c *Client) DeleteEventMailRegistration(id int64) error {
	return c.DeleteEventMailRegistrations([]int64{id})
}

// DeleteEventMailRegistrations deletes existing event.mail.registration records.
func (c *Client) DeleteEventMailRegistrations(ids []int64) error {
	return c.Delete(EventMailRegistrationModel, ids)
}

// GetEventMailRegistration gets event.mail.registration existing record.
func (c *Client) GetEventMailRegistration(id int64) (*EventMailRegistration, error) {
	emrs, err := c.GetEventMailRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	if emrs != nil && len(*emrs) > 0 {
		return &((*emrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.mail.registration not found", id)
}

// GetEventMailRegistrations gets event.mail.registration existing records.
func (c *Client) GetEventMailRegistrations(ids []int64) (*EventMailRegistrations, error) {
	emrs := &EventMailRegistrations{}
	if err := c.Read(EventMailRegistrationModel, ids, nil, emrs); err != nil {
		return nil, err
	}
	return emrs, nil
}

// FindEventMailRegistration finds event.mail.registration record by querying it with criteria.
func (c *Client) FindEventMailRegistration(criteria *Criteria) (*EventMailRegistration, error) {
	emrs := &EventMailRegistrations{}
	if err := c.SearchRead(EventMailRegistrationModel, criteria, NewOptions().Limit(1), emrs); err != nil {
		return nil, err
	}
	if emrs != nil && len(*emrs) > 0 {
		return &((*emrs)[0]), nil
	}
	return nil, fmt.Errorf("event.mail.registration was not found")
}

// FindEventMailRegistrations finds event.mail.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMailRegistrations(criteria *Criteria, options *Options) (*EventMailRegistrations, error) {
	emrs := &EventMailRegistrations{}
	if err := c.SearchRead(EventMailRegistrationModel, criteria, options, emrs); err != nil {
		return nil, err
	}
	return emrs, nil
}

// FindEventMailRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMailRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventMailRegistrationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventMailRegistrationId finds record id by querying it with criteria.
func (c *Client) FindEventMailRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventMailRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.mail.registration was not found")
}
