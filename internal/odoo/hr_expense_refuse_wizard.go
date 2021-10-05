package odoo

import (
	"fmt"
)

// HrExpenseRefuseWizard represents hr.expense.refuse.wizard model.
type HrExpenseRefuseWizard struct {
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	HrExpenseIds     *Relation `xmlrpc:"hr_expense_ids,omptempty"`
	HrExpenseSheetId *Many2One `xmlrpc:"hr_expense_sheet_id,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	Reason           *String   `xmlrpc:"reason,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrExpenseRefuseWizards represents array of hr.expense.refuse.wizard model.
type HrExpenseRefuseWizards []HrExpenseRefuseWizard

// HrExpenseRefuseWizardModel is the odoo model name.
const HrExpenseRefuseWizardModel = "hr.expense.refuse.wizard"

// Many2One convert HrExpenseRefuseWizard to *Many2One.
func (herw *HrExpenseRefuseWizard) Many2One() *Many2One {
	return NewMany2One(herw.Id.Get(), "")
}

// CreateHrExpenseRefuseWizard creates a new hr.expense.refuse.wizard model and returns its id.
func (c *Client) CreateHrExpenseRefuseWizard(herw *HrExpenseRefuseWizard) (int64, error) {
	return c.Create(HrExpenseRefuseWizardModel, herw)
}

// UpdateHrExpenseRefuseWizard updates an existing hr.expense.refuse.wizard record.
func (c *Client) UpdateHrExpenseRefuseWizard(herw *HrExpenseRefuseWizard) error {
	return c.UpdateHrExpenseRefuseWizards([]int64{herw.Id.Get()}, herw)
}

// UpdateHrExpenseRefuseWizards updates existing hr.expense.refuse.wizard records.
// All records (represented by ids) will be updated by herw values.
func (c *Client) UpdateHrExpenseRefuseWizards(ids []int64, herw *HrExpenseRefuseWizard) error {
	return c.Update(HrExpenseRefuseWizardModel, ids, herw)
}

// DeleteHrExpenseRefuseWizard deletes an existing hr.expense.refuse.wizard record.
func (c *Client) DeleteHrExpenseRefuseWizard(id int64) error {
	return c.DeleteHrExpenseRefuseWizards([]int64{id})
}

// DeleteHrExpenseRefuseWizards deletes existing hr.expense.refuse.wizard records.
func (c *Client) DeleteHrExpenseRefuseWizards(ids []int64) error {
	return c.Delete(HrExpenseRefuseWizardModel, ids)
}

// GetHrExpenseRefuseWizard gets hr.expense.refuse.wizard existing record.
func (c *Client) GetHrExpenseRefuseWizard(id int64) (*HrExpenseRefuseWizard, error) {
	herws, err := c.GetHrExpenseRefuseWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if herws != nil && len(*herws) > 0 {
		return &((*herws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense.refuse.wizard not found", id)
}

// GetHrExpenseRefuseWizards gets hr.expense.refuse.wizard existing records.
func (c *Client) GetHrExpenseRefuseWizards(ids []int64) (*HrExpenseRefuseWizards, error) {
	herws := &HrExpenseRefuseWizards{}
	if err := c.Read(HrExpenseRefuseWizardModel, ids, nil, herws); err != nil {
		return nil, err
	}
	return herws, nil
}

// FindHrExpenseRefuseWizard finds hr.expense.refuse.wizard record by querying it with criteria.
func (c *Client) FindHrExpenseRefuseWizard(criteria *Criteria) (*HrExpenseRefuseWizard, error) {
	herws := &HrExpenseRefuseWizards{}
	if err := c.SearchRead(HrExpenseRefuseWizardModel, criteria, NewOptions().Limit(1), herws); err != nil {
		return nil, err
	}
	if herws != nil && len(*herws) > 0 {
		return &((*herws)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense.refuse.wizard was not found")
}

// FindHrExpenseRefuseWizards finds hr.expense.refuse.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseRefuseWizards(criteria *Criteria, options *Options) (*HrExpenseRefuseWizards, error) {
	herws := &HrExpenseRefuseWizards{}
	if err := c.SearchRead(HrExpenseRefuseWizardModel, criteria, options, herws); err != nil {
		return nil, err
	}
	return herws, nil
}

// FindHrExpenseRefuseWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseRefuseWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseRefuseWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseRefuseWizardId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseRefuseWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseRefuseWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense.refuse.wizard was not found")
}
