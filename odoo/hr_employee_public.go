package odoo

import (
	"fmt"
)

// HrEmployeePublic represents hr.employee.public model.
type HrEmployeePublic struct {
	Active                *Bool      `xmlrpc:"active,omptempty"`
	AddressId             *Many2One  `xmlrpc:"address_id,omptempty"`
	AllocationCount       *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay     *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationUsedCount   *Float     `xmlrpc:"allocation_used_count,omptempty"`
	AllocationUsedDisplay *String    `xmlrpc:"allocation_used_display,omptempty"`
	AttendanceIds         *Relation  `xmlrpc:"attendance_ids,omptempty"`
	AttendanceState       *Selection `xmlrpc:"attendance_state,omptempty"`
	BadgeIds              *Relation  `xmlrpc:"badge_ids,omptempty"`
	ChildAllCount         *Int       `xmlrpc:"child_all_count,omptempty"`
	ChildIds              *Relation  `xmlrpc:"child_ids,omptempty"`
	CoachId               *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                 *Int       `xmlrpc:"color,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrentLeaveId        *Many2One  `xmlrpc:"current_leave_id,omptempty"`
	CurrentLeaveState     *Selection `xmlrpc:"current_leave_state,omptempty"`
	DepartmentId          *Many2One  `xmlrpc:"department_id,omptempty"`
	DirectBadgeIds        *Relation  `xmlrpc:"direct_badge_ids,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	ExpenseManagerId      *Many2One  `xmlrpc:"expense_manager_id,omptempty"`
	GoalIds               *Relation  `xmlrpc:"goal_ids,omptempty"`
	HasBadges             *Bool      `xmlrpc:"has_badges,omptempty"`
	HoursLastMonth        *Float     `xmlrpc:"hours_last_month,omptempty"`
	HoursLastMonthDisplay *String    `xmlrpc:"hours_last_month_display,omptempty"`
	HoursToday            *Float     `xmlrpc:"hours_today,omptempty"`
	HrPresenceState       *Selection `xmlrpc:"hr_presence_state,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	Image1024             *String    `xmlrpc:"image_1024,omptempty"`
	Image128              *String    `xmlrpc:"image_128,omptempty"`
	Image1920             *String    `xmlrpc:"image_1920,omptempty"`
	Image256              *String    `xmlrpc:"image_256,omptempty"`
	Image512              *String    `xmlrpc:"image_512,omptempty"`
	IsAbsent              *Bool      `xmlrpc:"is_absent,omptempty"`
	JobId                 *Many2One  `xmlrpc:"job_id,omptempty"`
	JobTitle              *String    `xmlrpc:"job_title,omptempty"`
	LastActivity          *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime      *String    `xmlrpc:"last_activity_time,omptempty"`
	LastAttendanceId      *Many2One  `xmlrpc:"last_attendance_id,omptempty"`
	LastCheckIn           *Time      `xmlrpc:"last_check_in,omptempty"`
	LastCheckOut          *Time      `xmlrpc:"last_check_out,omptempty"`
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	LeaveDateFrom         *Time      `xmlrpc:"leave_date_from,omptempty"`
	LeaveDateTo           *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeaveManagerId        *Many2One  `xmlrpc:"leave_manager_id,omptempty"`
	LeavesCount           *Float     `xmlrpc:"leaves_count,omptempty"`
	MobilePhone           *String    `xmlrpc:"mobile_phone,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	ParentId              *Many2One  `xmlrpc:"parent_id,omptempty"`
	RemainingLeaves       *Float     `xmlrpc:"remaining_leaves,omptempty"`
	ResourceCalendarId    *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceId            *Many2One  `xmlrpc:"resource_id,omptempty"`
	ShowLeaves            *Bool      `xmlrpc:"show_leaves,omptempty"`
	SubordinateIds        *Relation  `xmlrpc:"subordinate_ids,omptempty"`
	Tz                    *Selection `xmlrpc:"tz,omptempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omptempty"`
	WorkEmail             *String    `xmlrpc:"work_email,omptempty"`
	WorkLocation          *String    `xmlrpc:"work_location,omptempty"`
	WorkPhone             *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrEmployeePublics represents array of hr.employee.public model.
type HrEmployeePublics []HrEmployeePublic

// HrEmployeePublicModel is the odoo model name.
const HrEmployeePublicModel = "hr.employee.public"

// Many2One convert HrEmployeePublic to *Many2One.
func (hep *HrEmployeePublic) Many2One() *Many2One {
	return NewMany2One(hep.Id.Get(), "")
}

// CreateHrEmployeePublic creates a new hr.employee.public model and returns its id.
func (c *Client) CreateHrEmployeePublic(hep *HrEmployeePublic) (int64, error) {
	return c.Create(HrEmployeePublicModel, hep)
}

// UpdateHrEmployeePublic updates an existing hr.employee.public record.
func (c *Client) UpdateHrEmployeePublic(hep *HrEmployeePublic) error {
	return c.UpdateHrEmployeePublics([]int64{hep.Id.Get()}, hep)
}

// UpdateHrEmployeePublics updates existing hr.employee.public records.
// All records (represented by ids) will be updated by hep values.
func (c *Client) UpdateHrEmployeePublics(ids []int64, hep *HrEmployeePublic) error {
	return c.Update(HrEmployeePublicModel, ids, hep)
}

// DeleteHrEmployeePublic deletes an existing hr.employee.public record.
func (c *Client) DeleteHrEmployeePublic(id int64) error {
	return c.DeleteHrEmployeePublics([]int64{id})
}

// DeleteHrEmployeePublics deletes existing hr.employee.public records.
func (c *Client) DeleteHrEmployeePublics(ids []int64) error {
	return c.Delete(HrEmployeePublicModel, ids)
}

// GetHrEmployeePublic gets hr.employee.public existing record.
func (c *Client) GetHrEmployeePublic(id int64) (*HrEmployeePublic, error) {
	heps, err := c.GetHrEmployeePublics([]int64{id})
	if err != nil {
		return nil, err
	}
	if heps != nil && len(*heps) > 0 {
		return &((*heps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee.public not found", id)
}

// GetHrEmployeePublics gets hr.employee.public existing records.
func (c *Client) GetHrEmployeePublics(ids []int64) (*HrEmployeePublics, error) {
	heps := &HrEmployeePublics{}
	if err := c.Read(HrEmployeePublicModel, ids, nil, heps); err != nil {
		return nil, err
	}
	return heps, nil
}

// FindHrEmployeePublic finds hr.employee.public record by querying it with criteria.
func (c *Client) FindHrEmployeePublic(criteria *Criteria) (*HrEmployeePublic, error) {
	heps := &HrEmployeePublics{}
	if err := c.SearchRead(HrEmployeePublicModel, criteria, NewOptions().Limit(1), heps); err != nil {
		return nil, err
	}
	if heps != nil && len(*heps) > 0 {
		return &((*heps)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee.public was not found")
}

// FindHrEmployeePublics finds hr.employee.public records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeePublics(criteria *Criteria, options *Options) (*HrEmployeePublics, error) {
	heps := &HrEmployeePublics{}
	if err := c.SearchRead(HrEmployeePublicModel, criteria, options, heps); err != nil {
		return nil, err
	}
	return heps, nil
}

// FindHrEmployeePublicIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeePublicIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeePublicModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeePublicId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeePublicId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeePublicModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee.public was not found")
}
