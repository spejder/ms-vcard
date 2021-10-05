package odoo

import (
	"fmt"
)

// OpenacademySession represents openacademy.session model.
type OpenacademySession struct {
	Active         *Bool     `xmlrpc:"active,omptempty"`
	AttendeeIds    *Relation `xmlrpc:"attendee_ids,omptempty"`
	AttendeesCount *Int      `xmlrpc:"attendees_count,omptempty"`
	BundleName     *Many2One `xmlrpc:"bundle_name,omptempty"`
	Color          *Int      `xmlrpc:"color,omptempty"`
	CourseId       *Many2One `xmlrpc:"course_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Duration       *Float    `xmlrpc:"duration,omptempty"`
	EndDate        *Time     `xmlrpc:"end_date,omptempty"`
	Hours          *Float    `xmlrpc:"hours,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	InstructorId   *Many2One `xmlrpc:"instructor_id,omptempty"`
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	Seats          *Int      `xmlrpc:"seats,omptempty"`
	StartDate      *Time     `xmlrpc:"start_date,omptempty"`
	TakenSeats     *Float    `xmlrpc:"taken_seats,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OpenacademySessions represents array of openacademy.session model.
type OpenacademySessions []OpenacademySession

// OpenacademySessionModel is the odoo model name.
const OpenacademySessionModel = "openacademy.session"

// Many2One convert OpenacademySession to *Many2One.
func (os *OpenacademySession) Many2One() *Many2One {
	return NewMany2One(os.Id.Get(), "")
}

// CreateOpenacademySession creates a new openacademy.session model and returns its id.
func (c *Client) CreateOpenacademySession(os *OpenacademySession) (int64, error) {
	return c.Create(OpenacademySessionModel, os)
}

// UpdateOpenacademySession updates an existing openacademy.session record.
func (c *Client) UpdateOpenacademySession(os *OpenacademySession) error {
	return c.UpdateOpenacademySessions([]int64{os.Id.Get()}, os)
}

// UpdateOpenacademySessions updates existing openacademy.session records.
// All records (represented by ids) will be updated by os values.
func (c *Client) UpdateOpenacademySessions(ids []int64, os *OpenacademySession) error {
	return c.Update(OpenacademySessionModel, ids, os)
}

// DeleteOpenacademySession deletes an existing openacademy.session record.
func (c *Client) DeleteOpenacademySession(id int64) error {
	return c.DeleteOpenacademySessions([]int64{id})
}

// DeleteOpenacademySessions deletes existing openacademy.session records.
func (c *Client) DeleteOpenacademySessions(ids []int64) error {
	return c.Delete(OpenacademySessionModel, ids)
}

// GetOpenacademySession gets openacademy.session existing record.
func (c *Client) GetOpenacademySession(id int64) (*OpenacademySession, error) {
	oss, err := c.GetOpenacademySessions([]int64{id})
	if err != nil {
		return nil, err
	}
	if oss != nil && len(*oss) > 0 {
		return &((*oss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of openacademy.session not found", id)
}

// GetOpenacademySessions gets openacademy.session existing records.
func (c *Client) GetOpenacademySessions(ids []int64) (*OpenacademySessions, error) {
	oss := &OpenacademySessions{}
	if err := c.Read(OpenacademySessionModel, ids, nil, oss); err != nil {
		return nil, err
	}
	return oss, nil
}

// FindOpenacademySession finds openacademy.session record by querying it with criteria.
func (c *Client) FindOpenacademySession(criteria *Criteria) (*OpenacademySession, error) {
	oss := &OpenacademySessions{}
	if err := c.SearchRead(OpenacademySessionModel, criteria, NewOptions().Limit(1), oss); err != nil {
		return nil, err
	}
	if oss != nil && len(*oss) > 0 {
		return &((*oss)[0]), nil
	}
	return nil, fmt.Errorf("openacademy.session was not found")
}

// FindOpenacademySessions finds openacademy.session records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademySessions(criteria *Criteria, options *Options) (*OpenacademySessions, error) {
	oss := &OpenacademySessions{}
	if err := c.SearchRead(OpenacademySessionModel, criteria, options, oss); err != nil {
		return nil, err
	}
	return oss, nil
}

// FindOpenacademySessionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademySessionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OpenacademySessionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOpenacademySessionId finds record id by querying it with criteria.
func (c *Client) FindOpenacademySessionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OpenacademySessionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("openacademy.session was not found")
}
