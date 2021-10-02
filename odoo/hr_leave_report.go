package odoo

import (
	"fmt"
)

// HrLeaveReport represents hr.leave.report model.
type HrLeaveReport struct {
	CategoryId      *Many2One  `xmlrpc:"category_id,omptempty"`
	DateFrom        *Time      `xmlrpc:"date_from,omptempty"`
	DateTo          *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId    *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId      *Many2One  `xmlrpc:"employee_id,omptempty"`
	HolidayStatusId *Many2One  `xmlrpc:"holiday_status_id,omptempty"`
	HolidayType     *Selection `xmlrpc:"holiday_type,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	LeaveType       *Selection `xmlrpc:"leave_type,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	NumberOfDays    *Float     `xmlrpc:"number_of_days,omptempty"`
	PayslipStatus   *Bool      `xmlrpc:"payslip_status,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
}

// HrLeaveReports represents array of hr.leave.report model.
type HrLeaveReports []HrLeaveReport

// HrLeaveReportModel is the odoo model name.
const HrLeaveReportModel = "hr.leave.report"

// Many2One convert HrLeaveReport to *Many2One.
func (hlr *HrLeaveReport) Many2One() *Many2One {
	return NewMany2One(hlr.Id.Get(), "")
}

// CreateHrLeaveReport creates a new hr.leave.report model and returns its id.
func (c *Client) CreateHrLeaveReport(hlr *HrLeaveReport) (int64, error) {
	return c.Create(HrLeaveReportModel, hlr)
}

// UpdateHrLeaveReport updates an existing hr.leave.report record.
func (c *Client) UpdateHrLeaveReport(hlr *HrLeaveReport) error {
	return c.UpdateHrLeaveReports([]int64{hlr.Id.Get()}, hlr)
}

// UpdateHrLeaveReports updates existing hr.leave.report records.
// All records (represented by ids) will be updated by hlr values.
func (c *Client) UpdateHrLeaveReports(ids []int64, hlr *HrLeaveReport) error {
	return c.Update(HrLeaveReportModel, ids, hlr)
}

// DeleteHrLeaveReport deletes an existing hr.leave.report record.
func (c *Client) DeleteHrLeaveReport(id int64) error {
	return c.DeleteHrLeaveReports([]int64{id})
}

// DeleteHrLeaveReports deletes existing hr.leave.report records.
func (c *Client) DeleteHrLeaveReports(ids []int64) error {
	return c.Delete(HrLeaveReportModel, ids)
}

// GetHrLeaveReport gets hr.leave.report existing record.
func (c *Client) GetHrLeaveReport(id int64) (*HrLeaveReport, error) {
	hlrs, err := c.GetHrLeaveReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlrs != nil && len(*hlrs) > 0 {
		return &((*hlrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.report not found", id)
}

// GetHrLeaveReports gets hr.leave.report existing records.
func (c *Client) GetHrLeaveReports(ids []int64) (*HrLeaveReports, error) {
	hlrs := &HrLeaveReports{}
	if err := c.Read(HrLeaveReportModel, ids, nil, hlrs); err != nil {
		return nil, err
	}
	return hlrs, nil
}

// FindHrLeaveReport finds hr.leave.report record by querying it with criteria.
func (c *Client) FindHrLeaveReport(criteria *Criteria) (*HrLeaveReport, error) {
	hlrs := &HrLeaveReports{}
	if err := c.SearchRead(HrLeaveReportModel, criteria, NewOptions().Limit(1), hlrs); err != nil {
		return nil, err
	}
	if hlrs != nil && len(*hlrs) > 0 {
		return &((*hlrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.report was not found")
}

// FindHrLeaveReports finds hr.leave.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReports(criteria *Criteria, options *Options) (*HrLeaveReports, error) {
	hlrs := &HrLeaveReports{}
	if err := c.SearchRead(HrLeaveReportModel, criteria, options, hlrs); err != nil {
		return nil, err
	}
	return hlrs, nil
}

// FindHrLeaveReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveReportId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.report was not found")
}
