package odoo

import (
	"fmt"
)

// HrLeaveReportCalendar represents hr.leave.report.calendar model.
type HrLeaveReportCalendar struct {
	CompanyId     *Many2One  `xmlrpc:"company_id,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Duration      *Float     `xmlrpc:"duration,omptempty"`
	EmployeeId    *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	StartDatetime *Time      `xmlrpc:"start_datetime,omptempty"`
	StopDatetime  *Time      `xmlrpc:"stop_datetime,omptempty"`
	Tz            *Selection `xmlrpc:"tz,omptempty"`
}

// HrLeaveReportCalendars represents array of hr.leave.report.calendar model.
type HrLeaveReportCalendars []HrLeaveReportCalendar

// HrLeaveReportCalendarModel is the odoo model name.
const HrLeaveReportCalendarModel = "hr.leave.report.calendar"

// Many2One convert HrLeaveReportCalendar to *Many2One.
func (hlrc *HrLeaveReportCalendar) Many2One() *Many2One {
	return NewMany2One(hlrc.Id.Get(), "")
}

// CreateHrLeaveReportCalendar creates a new hr.leave.report.calendar model and returns its id.
func (c *Client) CreateHrLeaveReportCalendar(hlrc *HrLeaveReportCalendar) (int64, error) {
	return c.Create(HrLeaveReportCalendarModel, hlrc)
}

// UpdateHrLeaveReportCalendar updates an existing hr.leave.report.calendar record.
func (c *Client) UpdateHrLeaveReportCalendar(hlrc *HrLeaveReportCalendar) error {
	return c.UpdateHrLeaveReportCalendars([]int64{hlrc.Id.Get()}, hlrc)
}

// UpdateHrLeaveReportCalendars updates existing hr.leave.report.calendar records.
// All records (represented by ids) will be updated by hlrc values.
func (c *Client) UpdateHrLeaveReportCalendars(ids []int64, hlrc *HrLeaveReportCalendar) error {
	return c.Update(HrLeaveReportCalendarModel, ids, hlrc)
}

// DeleteHrLeaveReportCalendar deletes an existing hr.leave.report.calendar record.
func (c *Client) DeleteHrLeaveReportCalendar(id int64) error {
	return c.DeleteHrLeaveReportCalendars([]int64{id})
}

// DeleteHrLeaveReportCalendars deletes existing hr.leave.report.calendar records.
func (c *Client) DeleteHrLeaveReportCalendars(ids []int64) error {
	return c.Delete(HrLeaveReportCalendarModel, ids)
}

// GetHrLeaveReportCalendar gets hr.leave.report.calendar existing record.
func (c *Client) GetHrLeaveReportCalendar(id int64) (*HrLeaveReportCalendar, error) {
	hlrcs, err := c.GetHrLeaveReportCalendars([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlrcs != nil && len(*hlrcs) > 0 {
		return &((*hlrcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.report.calendar not found", id)
}

// GetHrLeaveReportCalendars gets hr.leave.report.calendar existing records.
func (c *Client) GetHrLeaveReportCalendars(ids []int64) (*HrLeaveReportCalendars, error) {
	hlrcs := &HrLeaveReportCalendars{}
	if err := c.Read(HrLeaveReportCalendarModel, ids, nil, hlrcs); err != nil {
		return nil, err
	}
	return hlrcs, nil
}

// FindHrLeaveReportCalendar finds hr.leave.report.calendar record by querying it with criteria.
func (c *Client) FindHrLeaveReportCalendar(criteria *Criteria) (*HrLeaveReportCalendar, error) {
	hlrcs := &HrLeaveReportCalendars{}
	if err := c.SearchRead(HrLeaveReportCalendarModel, criteria, NewOptions().Limit(1), hlrcs); err != nil {
		return nil, err
	}
	if hlrcs != nil && len(*hlrcs) > 0 {
		return &((*hlrcs)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.report.calendar was not found")
}

// FindHrLeaveReportCalendars finds hr.leave.report.calendar records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReportCalendars(criteria *Criteria, options *Options) (*HrLeaveReportCalendars, error) {
	hlrcs := &HrLeaveReportCalendars{}
	if err := c.SearchRead(HrLeaveReportCalendarModel, criteria, options, hlrcs); err != nil {
		return nil, err
	}
	return hlrcs, nil
}

// FindHrLeaveReportCalendarIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReportCalendarIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveReportCalendarModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveReportCalendarId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveReportCalendarId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveReportCalendarModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.report.calendar was not found")
}
