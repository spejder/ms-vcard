package odoo

import (
	"fmt"
)

// HrExpenseSheetRegisterPaymentWizard represents hr.expense.sheet.register.payment.wizard model.
type HrExpenseSheetRegisterPaymentWizard struct {
	Amount                    *Float    `xmlrpc:"amount,omptempty"`
	Communication             *String   `xmlrpc:"communication,omptempty"`
	CompanyId                 *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId                *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	ExpenseSheetId            *Many2One `xmlrpc:"expense_sheet_id,omptempty"`
	HidePaymentMethod         *Bool     `xmlrpc:"hide_payment_method,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	JournalId                 *Many2One `xmlrpc:"journal_id,omptempty"`
	LastUpdate                *Time     `xmlrpc:"__last_update,omptempty"`
	PartnerBankAccountId      *Many2One `xmlrpc:"partner_bank_account_id,omptempty"`
	PartnerId                 *Many2One `xmlrpc:"partner_id,omptempty"`
	PaymentDate               *Time     `xmlrpc:"payment_date,omptempty"`
	PaymentMethodId           *Many2One `xmlrpc:"payment_method_id,omptempty"`
	RequirePartnerBankAccount *Bool     `xmlrpc:"require_partner_bank_account,omptempty"`
	ShowPartnerBankAccount    *Bool     `xmlrpc:"show_partner_bank_account,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrExpenseSheetRegisterPaymentWizards represents array of hr.expense.sheet.register.payment.wizard model.
type HrExpenseSheetRegisterPaymentWizards []HrExpenseSheetRegisterPaymentWizard

// HrExpenseSheetRegisterPaymentWizardModel is the odoo model name.
const HrExpenseSheetRegisterPaymentWizardModel = "hr.expense.sheet.register.payment.wizard"

// Many2One convert HrExpenseSheetRegisterPaymentWizard to *Many2One.
func (hesrpw *HrExpenseSheetRegisterPaymentWizard) Many2One() *Many2One {
	return NewMany2One(hesrpw.Id.Get(), "")
}

// CreateHrExpenseSheetRegisterPaymentWizard creates a new hr.expense.sheet.register.payment.wizard model and returns its id.
func (c *Client) CreateHrExpenseSheetRegisterPaymentWizard(hesrpw *HrExpenseSheetRegisterPaymentWizard) (int64, error) {
	return c.Create(HrExpenseSheetRegisterPaymentWizardModel, hesrpw)
}

// UpdateHrExpenseSheetRegisterPaymentWizard updates an existing hr.expense.sheet.register.payment.wizard record.
func (c *Client) UpdateHrExpenseSheetRegisterPaymentWizard(hesrpw *HrExpenseSheetRegisterPaymentWizard) error {
	return c.UpdateHrExpenseSheetRegisterPaymentWizards([]int64{hesrpw.Id.Get()}, hesrpw)
}

// UpdateHrExpenseSheetRegisterPaymentWizards updates existing hr.expense.sheet.register.payment.wizard records.
// All records (represented by ids) will be updated by hesrpw values.
func (c *Client) UpdateHrExpenseSheetRegisterPaymentWizards(ids []int64, hesrpw *HrExpenseSheetRegisterPaymentWizard) error {
	return c.Update(HrExpenseSheetRegisterPaymentWizardModel, ids, hesrpw)
}

// DeleteHrExpenseSheetRegisterPaymentWizard deletes an existing hr.expense.sheet.register.payment.wizard record.
func (c *Client) DeleteHrExpenseSheetRegisterPaymentWizard(id int64) error {
	return c.DeleteHrExpenseSheetRegisterPaymentWizards([]int64{id})
}

// DeleteHrExpenseSheetRegisterPaymentWizards deletes existing hr.expense.sheet.register.payment.wizard records.
func (c *Client) DeleteHrExpenseSheetRegisterPaymentWizards(ids []int64) error {
	return c.Delete(HrExpenseSheetRegisterPaymentWizardModel, ids)
}

// GetHrExpenseSheetRegisterPaymentWizard gets hr.expense.sheet.register.payment.wizard existing record.
func (c *Client) GetHrExpenseSheetRegisterPaymentWizard(id int64) (*HrExpenseSheetRegisterPaymentWizard, error) {
	hesrpws, err := c.GetHrExpenseSheetRegisterPaymentWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hesrpws != nil && len(*hesrpws) > 0 {
		return &((*hesrpws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense.sheet.register.payment.wizard not found", id)
}

// GetHrExpenseSheetRegisterPaymentWizards gets hr.expense.sheet.register.payment.wizard existing records.
func (c *Client) GetHrExpenseSheetRegisterPaymentWizards(ids []int64) (*HrExpenseSheetRegisterPaymentWizards, error) {
	hesrpws := &HrExpenseSheetRegisterPaymentWizards{}
	if err := c.Read(HrExpenseSheetRegisterPaymentWizardModel, ids, nil, hesrpws); err != nil {
		return nil, err
	}
	return hesrpws, nil
}

// FindHrExpenseSheetRegisterPaymentWizard finds hr.expense.sheet.register.payment.wizard record by querying it with criteria.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizard(criteria *Criteria) (*HrExpenseSheetRegisterPaymentWizard, error) {
	hesrpws := &HrExpenseSheetRegisterPaymentWizards{}
	if err := c.SearchRead(HrExpenseSheetRegisterPaymentWizardModel, criteria, NewOptions().Limit(1), hesrpws); err != nil {
		return nil, err
	}
	if hesrpws != nil && len(*hesrpws) > 0 {
		return &((*hesrpws)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense.sheet.register.payment.wizard was not found")
}

// FindHrExpenseSheetRegisterPaymentWizards finds hr.expense.sheet.register.payment.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizards(criteria *Criteria, options *Options) (*HrExpenseSheetRegisterPaymentWizards, error) {
	hesrpws := &HrExpenseSheetRegisterPaymentWizards{}
	if err := c.SearchRead(HrExpenseSheetRegisterPaymentWizardModel, criteria, options, hesrpws); err != nil {
		return nil, err
	}
	return hesrpws, nil
}

// FindHrExpenseSheetRegisterPaymentWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseSheetRegisterPaymentWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseSheetRegisterPaymentWizardId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSheetRegisterPaymentWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSheetRegisterPaymentWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense.sheet.register.payment.wizard was not found")
}
