package odoo

import (
	"fmt"
)

// HrContract represents hr.contract model.
type HrContract struct {
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	Advantages                  *String    `xmlrpc:"advantages,omptempty"`
	CalendarMismatch            *Bool      `xmlrpc:"calendar_mismatch,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omptempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	HrResponsibleId             *Many2One  `xmlrpc:"hr_responsible_id,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omptempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
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
	Name                        *String    `xmlrpc:"name,omptempty"`
	Notes                       *String    `xmlrpc:"notes,omptempty"`
	PermitNo                    *String    `xmlrpc:"permit_no,omptempty"`
	ResourceCalendarId          *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	TrialDateEnd                *Time      `xmlrpc:"trial_date_end,omptempty"`
	VisaExpire                  *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                      *String    `xmlrpc:"visa_no,omptempty"`
	Wage                        *Float     `xmlrpc:"wage,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrContracts represents array of hr.contract model.
type HrContracts []HrContract

// HrContractModel is the odoo model name.
const HrContractModel = "hr.contract"

// Many2One convert HrContract to *Many2One.
func (hc *HrContract) Many2One() *Many2One {
	return NewMany2One(hc.Id.Get(), "")
}

// CreateHrContract creates a new hr.contract model and returns its id.
func (c *Client) CreateHrContract(hc *HrContract) (int64, error) {
	return c.Create(HrContractModel, hc)
}

// UpdateHrContract updates an existing hr.contract record.
func (c *Client) UpdateHrContract(hc *HrContract) error {
	return c.UpdateHrContracts([]int64{hc.Id.Get()}, hc)
}

// UpdateHrContracts updates existing hr.contract records.
// All records (represented by ids) will be updated by hc values.
func (c *Client) UpdateHrContracts(ids []int64, hc *HrContract) error {
	return c.Update(HrContractModel, ids, hc)
}

// DeleteHrContract deletes an existing hr.contract record.
func (c *Client) DeleteHrContract(id int64) error {
	return c.DeleteHrContracts([]int64{id})
}

// DeleteHrContracts deletes existing hr.contract records.
func (c *Client) DeleteHrContracts(ids []int64) error {
	return c.Delete(HrContractModel, ids)
}

// GetHrContract gets hr.contract existing record.
func (c *Client) GetHrContract(id int64) (*HrContract, error) {
	hcs, err := c.GetHrContracts([]int64{id})
	if err != nil {
		return nil, err
	}
	if hcs != nil && len(*hcs) > 0 {
		return &((*hcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.contract not found", id)
}

// GetHrContracts gets hr.contract existing records.
func (c *Client) GetHrContracts(ids []int64) (*HrContracts, error) {
	hcs := &HrContracts{}
	if err := c.Read(HrContractModel, ids, nil, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrContract finds hr.contract record by querying it with criteria.
func (c *Client) FindHrContract(criteria *Criteria) (*HrContract, error) {
	hcs := &HrContracts{}
	if err := c.SearchRead(HrContractModel, criteria, NewOptions().Limit(1), hcs); err != nil {
		return nil, err
	}
	if hcs != nil && len(*hcs) > 0 {
		return &((*hcs)[0]), nil
	}
	return nil, fmt.Errorf("hr.contract was not found")
}

// FindHrContracts finds hr.contract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContracts(criteria *Criteria, options *Options) (*HrContracts, error) {
	hcs := &HrContracts{}
	if err := c.SearchRead(HrContractModel, criteria, options, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrContractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrContractModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrContractId finds record id by querying it with criteria.
func (c *Client) FindHrContractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.contract was not found")
}
