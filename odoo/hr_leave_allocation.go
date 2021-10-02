package odoo

import (
	"fmt"
)

// HrLeaveAllocation represents hr.leave.allocation model.
type HrLeaveAllocation struct {
	AccrualLimit                *Int       `xmlrpc:"accrual_limit,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AllocationType              *Selection `xmlrpc:"allocation_type,omptempty"`
	CanApprove                  *Bool      `xmlrpc:"can_approve,omptempty"`
	CanReset                    *Bool      `xmlrpc:"can_reset,omptempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	DurationDisplay             *String    `xmlrpc:"duration_display,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	FirstApproverId             *Many2One  `xmlrpc:"first_approver_id,omptempty"`
	HolidayStatusId             *Many2One  `xmlrpc:"holiday_status_id,omptempty"`
	HolidayType                 *Selection `xmlrpc:"holiday_type,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IntervalNumber              *Int       `xmlrpc:"interval_number,omptempty"`
	IntervalUnit                *Selection `xmlrpc:"interval_unit,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	LeavesTaken                 *Float     `xmlrpc:"leaves_taken,omptempty"`
	LinkedRequestIds            *Relation  `xmlrpc:"linked_request_ids,omptempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omptempty"`
	MaxLeaves                   *Float     `xmlrpc:"max_leaves,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	ModeCompanyId               *Many2One  `xmlrpc:"mode_company_id,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	Nextcall                    *Time      `xmlrpc:"nextcall,omptempty"`
	Notes                       *String    `xmlrpc:"notes,omptempty"`
	NumberOfDays                *Float     `xmlrpc:"number_of_days,omptempty"`
	NumberOfDaysDisplay         *Float     `xmlrpc:"number_of_days_display,omptempty"`
	NumberOfHoursDisplay        *Float     `xmlrpc:"number_of_hours_display,omptempty"`
	NumberPerInterval           *Float     `xmlrpc:"number_per_interval,omptempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	SecondApproverId            *Many2One  `xmlrpc:"second_approver_id,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	TypeRequestUnit             *Selection `xmlrpc:"type_request_unit,omptempty"`
	UnitPerInterval             *Selection `xmlrpc:"unit_per_interval,omptempty"`
	ValidationType              *Selection `xmlrpc:"validation_type,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrLeaveAllocations represents array of hr.leave.allocation model.
type HrLeaveAllocations []HrLeaveAllocation

// HrLeaveAllocationModel is the odoo model name.
const HrLeaveAllocationModel = "hr.leave.allocation"

// Many2One convert HrLeaveAllocation to *Many2One.
func (hla *HrLeaveAllocation) Many2One() *Many2One {
	return NewMany2One(hla.Id.Get(), "")
}

// CreateHrLeaveAllocation creates a new hr.leave.allocation model and returns its id.
func (c *Client) CreateHrLeaveAllocation(hla *HrLeaveAllocation) (int64, error) {
	return c.Create(HrLeaveAllocationModel, hla)
}

// UpdateHrLeaveAllocation updates an existing hr.leave.allocation record.
func (c *Client) UpdateHrLeaveAllocation(hla *HrLeaveAllocation) error {
	return c.UpdateHrLeaveAllocations([]int64{hla.Id.Get()}, hla)
}

// UpdateHrLeaveAllocations updates existing hr.leave.allocation records.
// All records (represented by ids) will be updated by hla values.
func (c *Client) UpdateHrLeaveAllocations(ids []int64, hla *HrLeaveAllocation) error {
	return c.Update(HrLeaveAllocationModel, ids, hla)
}

// DeleteHrLeaveAllocation deletes an existing hr.leave.allocation record.
func (c *Client) DeleteHrLeaveAllocation(id int64) error {
	return c.DeleteHrLeaveAllocations([]int64{id})
}

// DeleteHrLeaveAllocations deletes existing hr.leave.allocation records.
func (c *Client) DeleteHrLeaveAllocations(ids []int64) error {
	return c.Delete(HrLeaveAllocationModel, ids)
}

// GetHrLeaveAllocation gets hr.leave.allocation existing record.
func (c *Client) GetHrLeaveAllocation(id int64) (*HrLeaveAllocation, error) {
	hlas, err := c.GetHrLeaveAllocations([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlas != nil && len(*hlas) > 0 {
		return &((*hlas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.allocation not found", id)
}

// GetHrLeaveAllocations gets hr.leave.allocation existing records.
func (c *Client) GetHrLeaveAllocations(ids []int64) (*HrLeaveAllocations, error) {
	hlas := &HrLeaveAllocations{}
	if err := c.Read(HrLeaveAllocationModel, ids, nil, hlas); err != nil {
		return nil, err
	}
	return hlas, nil
}

// FindHrLeaveAllocation finds hr.leave.allocation record by querying it with criteria.
func (c *Client) FindHrLeaveAllocation(criteria *Criteria) (*HrLeaveAllocation, error) {
	hlas := &HrLeaveAllocations{}
	if err := c.SearchRead(HrLeaveAllocationModel, criteria, NewOptions().Limit(1), hlas); err != nil {
		return nil, err
	}
	if hlas != nil && len(*hlas) > 0 {
		return &((*hlas)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.allocation was not found")
}

// FindHrLeaveAllocations finds hr.leave.allocation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAllocations(criteria *Criteria, options *Options) (*HrLeaveAllocations, error) {
	hlas := &HrLeaveAllocations{}
	if err := c.SearchRead(HrLeaveAllocationModel, criteria, options, hlas); err != nil {
		return nil, err
	}
	return hlas, nil
}

// FindHrLeaveAllocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAllocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveAllocationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveAllocationId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAllocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAllocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.allocation was not found")
}
