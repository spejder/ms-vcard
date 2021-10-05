package odoo

import (
	"fmt"
)

// HrEmployee represents hr.employee model.
type HrEmployee struct {
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalNote              *String    `xmlrpc:"additional_note,omptempty"`
	AddressHomeId               *Many2One  `xmlrpc:"address_home_id,omptempty"`
	AddressId                   *Many2One  `xmlrpc:"address_id,omptempty"`
	AllocationCount             *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay           *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationUsedCount         *Float     `xmlrpc:"allocation_used_count,omptempty"`
	AllocationUsedDisplay       *String    `xmlrpc:"allocation_used_display,omptempty"`
	AttendanceIds               *Relation  `xmlrpc:"attendance_ids,omptempty"`
	AttendanceState             *Selection `xmlrpc:"attendance_state,omptempty"`
	BadgeIds                    *Relation  `xmlrpc:"badge_ids,omptempty"`
	BankAccountId               *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	Barcode                     *String    `xmlrpc:"barcode,omptempty"`
	Birthday                    *Time      `xmlrpc:"birthday,omptempty"`
	CalendarMismatch            *Bool      `xmlrpc:"calendar_mismatch,omptempty"`
	CategoryIds                 *Relation  `xmlrpc:"category_ids,omptempty"`
	Certificate                 *Selection `xmlrpc:"certificate,omptempty"`
	ChildAllCount               *Int       `xmlrpc:"child_all_count,omptempty"`
	ChildIds                    *Relation  `xmlrpc:"child_ids,omptempty"`
	Children                    *Int       `xmlrpc:"children,omptempty"`
	CoachId                     *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	ContractId                  *Many2One  `xmlrpc:"contract_id,omptempty"`
	ContractIds                 *Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractsCount              *Int       `xmlrpc:"contracts_count,omptempty"`
	ContractWarning             *Bool      `xmlrpc:"contract_warning,omptempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CountryOfBirth              *Many2One  `xmlrpc:"country_of_birth,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrentLeaveId              *Many2One  `xmlrpc:"current_leave_id,omptempty"`
	CurrentLeaveState           *Selection `xmlrpc:"current_leave_state,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DepartureDescription        *String    `xmlrpc:"departure_description,omptempty"`
	DepartureReason             *Selection `xmlrpc:"departure_reason,omptempty"`
	DirectBadgeIds              *Relation  `xmlrpc:"direct_badge_ids,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmergencyContact            *String    `xmlrpc:"emergency_contact,omptempty"`
	EmergencyPhone              *String    `xmlrpc:"emergency_phone,omptempty"`
	ExpenseManagerId            *Many2One  `xmlrpc:"expense_manager_id,omptempty"`
	Gender                      *Selection `xmlrpc:"gender,omptempty"`
	GoalIds                     *Relation  `xmlrpc:"goal_ids,omptempty"`
	HasBadges                   *Bool      `xmlrpc:"has_badges,omptempty"`
	HoursLastMonth              *Float     `xmlrpc:"hours_last_month,omptempty"`
	HoursLastMonthDisplay       *String    `xmlrpc:"hours_last_month_display,omptempty"`
	HoursToday                  *Float     `xmlrpc:"hours_today,omptempty"`
	HrPresenceState             *Selection `xmlrpc:"hr_presence_state,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IdentificationId            *String    `xmlrpc:"identification_id,omptempty"`
	Image1024                   *String    `xmlrpc:"image_1024,omptempty"`
	Image128                    *String    `xmlrpc:"image_128,omptempty"`
	Image1920                   *String    `xmlrpc:"image_1920,omptempty"`
	Image256                    *String    `xmlrpc:"image_256,omptempty"`
	Image512                    *String    `xmlrpc:"image_512,omptempty"`
	IsAbsent                    *Bool      `xmlrpc:"is_absent,omptempty"`
	IsAddressHomeACompany       *Bool      `xmlrpc:"is_address_home_a_company,omptempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omptempty"`
	JobTitle                    *String    `xmlrpc:"job_title,omptempty"`
	KmHomeWork                  *Int       `xmlrpc:"km_home_work,omptempty"`
	LastActivity                *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime            *String    `xmlrpc:"last_activity_time,omptempty"`
	LastAttendanceId            *Many2One  `xmlrpc:"last_attendance_id,omptempty"`
	LastCheckIn                 *Time      `xmlrpc:"last_check_in,omptempty"`
	LastCheckOut                *Time      `xmlrpc:"last_check_out,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	LeaveDateFrom               *Time      `xmlrpc:"leave_date_from,omptempty"`
	LeaveDateTo                 *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeaveManagerId              *Many2One  `xmlrpc:"leave_manager_id,omptempty"`
	LeavesCount                 *Float     `xmlrpc:"leaves_count,omptempty"`
	Marital                     *Selection `xmlrpc:"marital,omptempty"`
	MedicExam                   *Time      `xmlrpc:"medic_exam,omptempty"`
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
	MobilePhone                 *String    `xmlrpc:"mobile_phone,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	NewlyHiredEmployee          *Bool      `xmlrpc:"newly_hired_employee,omptempty"`
	Notes                       *String    `xmlrpc:"notes,omptempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	PassportId                  *String    `xmlrpc:"passport_id,omptempty"`
	PermitNo                    *String    `xmlrpc:"permit_no,omptempty"`
	Phone                       *String    `xmlrpc:"phone,omptempty"`
	Pin                         *String    `xmlrpc:"pin,omptempty"`
	PlaceOfBirth                *String    `xmlrpc:"place_of_birth,omptempty"`
	PrivateEmail                *String    `xmlrpc:"private_email,omptempty"`
	RemainingLeaves             *Float     `xmlrpc:"remaining_leaves,omptempty"`
	ResourceCalendarId          *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceId                  *Many2One  `xmlrpc:"resource_id,omptempty"`
	ShowLeaves                  *Bool      `xmlrpc:"show_leaves,omptempty"`
	Sinid                       *String    `xmlrpc:"sinid,omptempty"`
	SpouseBirthdate             *Time      `xmlrpc:"spouse_birthdate,omptempty"`
	SpouseCompleteName          *String    `xmlrpc:"spouse_complete_name,omptempty"`
	Ssnid                       *String    `xmlrpc:"ssnid,omptempty"`
	StudyField                  *String    `xmlrpc:"study_field,omptempty"`
	StudySchool                 *String    `xmlrpc:"study_school,omptempty"`
	SubordinateIds              *Relation  `xmlrpc:"subordinate_ids,omptempty"`
	Tz                          *Selection `xmlrpc:"tz,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	UserPartnerId               *Many2One  `xmlrpc:"user_partner_id,omptempty"`
	Vehicle                     *String    `xmlrpc:"vehicle,omptempty"`
	VisaExpire                  *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                      *String    `xmlrpc:"visa_no,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WorkEmail                   *String    `xmlrpc:"work_email,omptempty"`
	WorkLocation                *String    `xmlrpc:"work_location,omptempty"`
	WorkPhone                   *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrEmployees represents array of hr.employee model.
type HrEmployees []HrEmployee

// HrEmployeeModel is the odoo model name.
const HrEmployeeModel = "hr.employee"

// Many2One convert HrEmployee to *Many2One.
func (he *HrEmployee) Many2One() *Many2One {
	return NewMany2One(he.Id.Get(), "")
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployee(he *HrEmployee) (int64, error) {
	return c.Create(HrEmployeeModel, he)
}

// UpdateHrEmployee updates an existing hr.employee record.
func (c *Client) UpdateHrEmployee(he *HrEmployee) error {
	return c.UpdateHrEmployees([]int64{he.Id.Get()}, he)
}

// UpdateHrEmployees updates existing hr.employee records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrEmployees(ids []int64, he *HrEmployee) error {
	return c.Update(HrEmployeeModel, ids, he)
}

// DeleteHrEmployee deletes an existing hr.employee record.
func (c *Client) DeleteHrEmployee(id int64) error {
	return c.DeleteHrEmployees([]int64{id})
}

// DeleteHrEmployees deletes existing hr.employee records.
func (c *Client) DeleteHrEmployees(ids []int64) error {
	return c.Delete(HrEmployeeModel, ids)
}

// GetHrEmployee gets hr.employee existing record.
func (c *Client) GetHrEmployee(id int64) (*HrEmployee, error) {
	hes, err := c.GetHrEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee not found", id)
}

// GetHrEmployees gets hr.employee existing records.
func (c *Client) GetHrEmployees(ids []int64) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.Read(HrEmployeeModel, ids, nil, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployee finds hr.employee record by querying it with criteria.
func (c *Client) FindHrEmployee(criteria *Criteria) (*HrEmployee, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, NewOptions().Limit(1), hes); err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee was not found")
}

// FindHrEmployees finds hr.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployees(criteria *Criteria, options *Options) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, options, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee was not found")
}
