package odoo

import (
	"fmt"
)

// HrExpenseSheet represents hr.expense.sheet model.
type HrExpenseSheet struct {
	AccountingDate              *Time      `xmlrpc:"accounting_date,omptempty"`
	AccountMoveId               *Many2One  `xmlrpc:"account_move_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AddressId                   *Many2One  `xmlrpc:"address_id,omptempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omptempty"`
	BankJournalId               *Many2One  `xmlrpc:"bank_journal_id,omptempty"`
	CanReset                    *Bool      `xmlrpc:"can_reset,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	ExpenseLineIds              *Relation  `xmlrpc:"expense_line_ids,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IsMultipleCurrency          *Bool      `xmlrpc:"is_multiple_currency,omptempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omptempty"`
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
	PaymentMode                 *Selection `xmlrpc:"payment_mode,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	TotalAmount                 *Float     `xmlrpc:"total_amount,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrExpenseSheets represents array of hr.expense.sheet model.
type HrExpenseSheets []HrExpenseSheet

// HrExpenseSheetModel is the odoo model name.
const HrExpenseSheetModel = "hr.expense.sheet"

// Many2One convert HrExpenseSheet to *Many2One.
func (hes *HrExpenseSheet) Many2One() *Many2One {
	return NewMany2One(hes.Id.Get(), "")
}

// CreateHrExpenseSheet creates a new hr.expense.sheet model and returns its id.
func (c *Client) CreateHrExpenseSheet(hes *HrExpenseSheet) (int64, error) {
	return c.Create(HrExpenseSheetModel, hes)
}

// UpdateHrExpenseSheet updates an existing hr.expense.sheet record.
func (c *Client) UpdateHrExpenseSheet(hes *HrExpenseSheet) error {
	return c.UpdateHrExpenseSheets([]int64{hes.Id.Get()}, hes)
}

// UpdateHrExpenseSheets updates existing hr.expense.sheet records.
// All records (represented by ids) will be updated by hes values.
func (c *Client) UpdateHrExpenseSheets(ids []int64, hes *HrExpenseSheet) error {
	return c.Update(HrExpenseSheetModel, ids, hes)
}

// DeleteHrExpenseSheet deletes an existing hr.expense.sheet record.
func (c *Client) DeleteHrExpenseSheet(id int64) error {
	return c.DeleteHrExpenseSheets([]int64{id})
}

// DeleteHrExpenseSheets deletes existing hr.expense.sheet records.
func (c *Client) DeleteHrExpenseSheets(ids []int64) error {
	return c.Delete(HrExpenseSheetModel, ids)
}

// GetHrExpenseSheet gets hr.expense.sheet existing record.
func (c *Client) GetHrExpenseSheet(id int64) (*HrExpenseSheet, error) {
	hess, err := c.GetHrExpenseSheets([]int64{id})
	if err != nil {
		return nil, err
	}
	if hess != nil && len(*hess) > 0 {
		return &((*hess)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense.sheet not found", id)
}

// GetHrExpenseSheets gets hr.expense.sheet existing records.
func (c *Client) GetHrExpenseSheets(ids []int64) (*HrExpenseSheets, error) {
	hess := &HrExpenseSheets{}
	if err := c.Read(HrExpenseSheetModel, ids, nil, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSheet finds hr.expense.sheet record by querying it with criteria.
func (c *Client) FindHrExpenseSheet(criteria *Criteria) (*HrExpenseSheet, error) {
	hess := &HrExpenseSheets{}
	if err := c.SearchRead(HrExpenseSheetModel, criteria, NewOptions().Limit(1), hess); err != nil {
		return nil, err
	}
	if hess != nil && len(*hess) > 0 {
		return &((*hess)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense.sheet was not found")
}

// FindHrExpenseSheets finds hr.expense.sheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheets(criteria *Criteria, options *Options) (*HrExpenseSheets, error) {
	hess := &HrExpenseSheets{}
	if err := c.SearchRead(HrExpenseSheetModel, criteria, options, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseSheetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseSheetId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense.sheet was not found")
}
