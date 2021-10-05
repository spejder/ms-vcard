package odoo

import (
	"fmt"
)

// HrPlanWizard represents hr.plan.wizard model.
type HrPlanWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	PlanId      *Many2One `xmlrpc:"plan_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPlanWizards represents array of hr.plan.wizard model.
type HrPlanWizards []HrPlanWizard

// HrPlanWizardModel is the odoo model name.
const HrPlanWizardModel = "hr.plan.wizard"

// Many2One convert HrPlanWizard to *Many2One.
func (hpw *HrPlanWizard) Many2One() *Many2One {
	return NewMany2One(hpw.Id.Get(), "")
}

// CreateHrPlanWizard creates a new hr.plan.wizard model and returns its id.
func (c *Client) CreateHrPlanWizard(hpw *HrPlanWizard) (int64, error) {
	return c.Create(HrPlanWizardModel, hpw)
}

// UpdateHrPlanWizard updates an existing hr.plan.wizard record.
func (c *Client) UpdateHrPlanWizard(hpw *HrPlanWizard) error {
	return c.UpdateHrPlanWizards([]int64{hpw.Id.Get()}, hpw)
}

// UpdateHrPlanWizards updates existing hr.plan.wizard records.
// All records (represented by ids) will be updated by hpw values.
func (c *Client) UpdateHrPlanWizards(ids []int64, hpw *HrPlanWizard) error {
	return c.Update(HrPlanWizardModel, ids, hpw)
}

// DeleteHrPlanWizard deletes an existing hr.plan.wizard record.
func (c *Client) DeleteHrPlanWizard(id int64) error {
	return c.DeleteHrPlanWizards([]int64{id})
}

// DeleteHrPlanWizards deletes existing hr.plan.wizard records.
func (c *Client) DeleteHrPlanWizards(ids []int64) error {
	return c.Delete(HrPlanWizardModel, ids)
}

// GetHrPlanWizard gets hr.plan.wizard existing record.
func (c *Client) GetHrPlanWizard(id int64) (*HrPlanWizard, error) {
	hpws, err := c.GetHrPlanWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpws != nil && len(*hpws) > 0 {
		return &((*hpws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.plan.wizard not found", id)
}

// GetHrPlanWizards gets hr.plan.wizard existing records.
func (c *Client) GetHrPlanWizards(ids []int64) (*HrPlanWizards, error) {
	hpws := &HrPlanWizards{}
	if err := c.Read(HrPlanWizardModel, ids, nil, hpws); err != nil {
		return nil, err
	}
	return hpws, nil
}

// FindHrPlanWizard finds hr.plan.wizard record by querying it with criteria.
func (c *Client) FindHrPlanWizard(criteria *Criteria) (*HrPlanWizard, error) {
	hpws := &HrPlanWizards{}
	if err := c.SearchRead(HrPlanWizardModel, criteria, NewOptions().Limit(1), hpws); err != nil {
		return nil, err
	}
	if hpws != nil && len(*hpws) > 0 {
		return &((*hpws)[0]), nil
	}
	return nil, fmt.Errorf("hr.plan.wizard was not found")
}

// FindHrPlanWizards finds hr.plan.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanWizards(criteria *Criteria, options *Options) (*HrPlanWizards, error) {
	hpws := &HrPlanWizards{}
	if err := c.SearchRead(HrPlanWizardModel, criteria, options, hpws); err != nil {
		return nil, err
	}
	return hpws, nil
}

// FindHrPlanWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPlanWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPlanWizardId finds record id by querying it with criteria.
func (c *Client) FindHrPlanWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPlanWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.plan.wizard was not found")
}
