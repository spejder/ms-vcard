package odoo

import (
	"fmt"
)

// HrAttendance represents hr.attendance model.
type HrAttendance struct {
	CheckIn      *Time     `xmlrpc:"check_in,omptempty"`
	CheckOut     *Time     `xmlrpc:"check_out,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DepartmentId *Many2One `xmlrpc:"department_id,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId   *Many2One `xmlrpc:"employee_id,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	WorkedHours  *Float    `xmlrpc:"worked_hours,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrAttendances represents array of hr.attendance model.
type HrAttendances []HrAttendance

// HrAttendanceModel is the odoo model name.
const HrAttendanceModel = "hr.attendance"

// Many2One convert HrAttendance to *Many2One.
func (ha *HrAttendance) Many2One() *Many2One {
	return NewMany2One(ha.Id.Get(), "")
}

// CreateHrAttendance creates a new hr.attendance model and returns its id.
func (c *Client) CreateHrAttendance(ha *HrAttendance) (int64, error) {
	return c.Create(HrAttendanceModel, ha)
}

// UpdateHrAttendance updates an existing hr.attendance record.
func (c *Client) UpdateHrAttendance(ha *HrAttendance) error {
	return c.UpdateHrAttendances([]int64{ha.Id.Get()}, ha)
}

// UpdateHrAttendances updates existing hr.attendance records.
// All records (represented by ids) will be updated by ha values.
func (c *Client) UpdateHrAttendances(ids []int64, ha *HrAttendance) error {
	return c.Update(HrAttendanceModel, ids, ha)
}

// DeleteHrAttendance deletes an existing hr.attendance record.
func (c *Client) DeleteHrAttendance(id int64) error {
	return c.DeleteHrAttendances([]int64{id})
}

// DeleteHrAttendances deletes existing hr.attendance records.
func (c *Client) DeleteHrAttendances(ids []int64) error {
	return c.Delete(HrAttendanceModel, ids)
}

// GetHrAttendance gets hr.attendance existing record.
func (c *Client) GetHrAttendance(id int64) (*HrAttendance, error) {
	has, err := c.GetHrAttendances([]int64{id})
	if err != nil {
		return nil, err
	}
	if has != nil && len(*has) > 0 {
		return &((*has)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.attendance not found", id)
}

// GetHrAttendances gets hr.attendance existing records.
func (c *Client) GetHrAttendances(ids []int64) (*HrAttendances, error) {
	has := &HrAttendances{}
	if err := c.Read(HrAttendanceModel, ids, nil, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrAttendance finds hr.attendance record by querying it with criteria.
func (c *Client) FindHrAttendance(criteria *Criteria) (*HrAttendance, error) {
	has := &HrAttendances{}
	if err := c.SearchRead(HrAttendanceModel, criteria, NewOptions().Limit(1), has); err != nil {
		return nil, err
	}
	if has != nil && len(*has) > 0 {
		return &((*has)[0]), nil
	}
	return nil, fmt.Errorf("hr.attendance was not found")
}

// FindHrAttendances finds hr.attendance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendances(criteria *Criteria, options *Options) (*HrAttendances, error) {
	has := &HrAttendances{}
	if err := c.SearchRead(HrAttendanceModel, criteria, options, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrAttendanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAttendanceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAttendanceId finds record id by querying it with criteria.
func (c *Client) FindHrAttendanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAttendanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.attendance was not found")
}
