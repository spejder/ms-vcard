package odoo

import (
	"fmt"
)

// ResourceCalendarAttendance represents resource.calendar.attendance model.
type ResourceCalendarAttendance struct {
	CalendarId       *Many2One  `xmlrpc:"calendar_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom         *Time      `xmlrpc:"date_from,omptempty"`
	DateTo           *Time      `xmlrpc:"date_to,omptempty"`
	Dayofweek        *Selection `xmlrpc:"dayofweek,omptempty"`
	DayPeriod        *Selection `xmlrpc:"day_period,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	DisplayType      *Selection `xmlrpc:"display_type,omptempty"`
	HourFrom         *Float     `xmlrpc:"hour_from,omptempty"`
	HourTo           *Float     `xmlrpc:"hour_to,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	ResourceId       *Many2One  `xmlrpc:"resource_id,omptempty"`
	Sequence         *Int       `xmlrpc:"sequence,omptempty"`
	TwoWeeksCalendar *Bool      `xmlrpc:"two_weeks_calendar,omptempty"`
	WeekType         *Selection `xmlrpc:"week_type,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResourceCalendarAttendances represents array of resource.calendar.attendance model.
type ResourceCalendarAttendances []ResourceCalendarAttendance

// ResourceCalendarAttendanceModel is the odoo model name.
const ResourceCalendarAttendanceModel = "resource.calendar.attendance"

// Many2One convert ResourceCalendarAttendance to *Many2One.
func (rca *ResourceCalendarAttendance) Many2One() *Many2One {
	return NewMany2One(rca.Id.Get(), "")
}

// CreateResourceCalendarAttendance creates a new resource.calendar.attendance model and returns its id.
func (c *Client) CreateResourceCalendarAttendance(rca *ResourceCalendarAttendance) (int64, error) {
	return c.Create(ResourceCalendarAttendanceModel, rca)
}

// UpdateResourceCalendarAttendance updates an existing resource.calendar.attendance record.
func (c *Client) UpdateResourceCalendarAttendance(rca *ResourceCalendarAttendance) error {
	return c.UpdateResourceCalendarAttendances([]int64{rca.Id.Get()}, rca)
}

// UpdateResourceCalendarAttendances updates existing resource.calendar.attendance records.
// All records (represented by ids) will be updated by rca values.
func (c *Client) UpdateResourceCalendarAttendances(ids []int64, rca *ResourceCalendarAttendance) error {
	return c.Update(ResourceCalendarAttendanceModel, ids, rca)
}

// DeleteResourceCalendarAttendance deletes an existing resource.calendar.attendance record.
func (c *Client) DeleteResourceCalendarAttendance(id int64) error {
	return c.DeleteResourceCalendarAttendances([]int64{id})
}

// DeleteResourceCalendarAttendances deletes existing resource.calendar.attendance records.
func (c *Client) DeleteResourceCalendarAttendances(ids []int64) error {
	return c.Delete(ResourceCalendarAttendanceModel, ids)
}

// GetResourceCalendarAttendance gets resource.calendar.attendance existing record.
func (c *Client) GetResourceCalendarAttendance(id int64) (*ResourceCalendarAttendance, error) {
	rcas, err := c.GetResourceCalendarAttendances([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcas != nil && len(*rcas) > 0 {
		return &((*rcas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of resource.calendar.attendance not found", id)
}

// GetResourceCalendarAttendances gets resource.calendar.attendance existing records.
func (c *Client) GetResourceCalendarAttendances(ids []int64) (*ResourceCalendarAttendances, error) {
	rcas := &ResourceCalendarAttendances{}
	if err := c.Read(ResourceCalendarAttendanceModel, ids, nil, rcas); err != nil {
		return nil, err
	}
	return rcas, nil
}

// FindResourceCalendarAttendance finds resource.calendar.attendance record by querying it with criteria.
func (c *Client) FindResourceCalendarAttendance(criteria *Criteria) (*ResourceCalendarAttendance, error) {
	rcas := &ResourceCalendarAttendances{}
	if err := c.SearchRead(ResourceCalendarAttendanceModel, criteria, NewOptions().Limit(1), rcas); err != nil {
		return nil, err
	}
	if rcas != nil && len(*rcas) > 0 {
		return &((*rcas)[0]), nil
	}
	return nil, fmt.Errorf("resource.calendar.attendance was not found")
}

// FindResourceCalendarAttendances finds resource.calendar.attendance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarAttendances(criteria *Criteria, options *Options) (*ResourceCalendarAttendances, error) {
	rcas := &ResourceCalendarAttendances{}
	if err := c.SearchRead(ResourceCalendarAttendanceModel, criteria, options, rcas); err != nil {
		return nil, err
	}
	return rcas, nil
}

// FindResourceCalendarAttendanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResourceCalendarAttendanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResourceCalendarAttendanceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResourceCalendarAttendanceId finds record id by querying it with criteria.
func (c *Client) FindResourceCalendarAttendanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResourceCalendarAttendanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("resource.calendar.attendance was not found")
}
