package odoo

import (
	"fmt"
)

// HrDepartureWizard represents hr.departure.wizard model.
type HrDepartureWizard struct {
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	DepartureDescription *String    `xmlrpc:"departure_description,omptempty"`
	DepartureReason      *Selection `xmlrpc:"departure_reason,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId           *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	PlanId               *Many2One  `xmlrpc:"plan_id,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrDepartureWizards represents array of hr.departure.wizard model.
type HrDepartureWizards []HrDepartureWizard

// HrDepartureWizardModel is the odoo model name.
const HrDepartureWizardModel = "hr.departure.wizard"

// Many2One convert HrDepartureWizard to *Many2One.
func (hdw *HrDepartureWizard) Many2One() *Many2One {
	return NewMany2One(hdw.Id.Get(), "")
}

// CreateHrDepartureWizard creates a new hr.departure.wizard model and returns its id.
func (c *Client) CreateHrDepartureWizard(hdw *HrDepartureWizard) (int64, error) {
	return c.Create(HrDepartureWizardModel, hdw)
}

// UpdateHrDepartureWizard updates an existing hr.departure.wizard record.
func (c *Client) UpdateHrDepartureWizard(hdw *HrDepartureWizard) error {
	return c.UpdateHrDepartureWizards([]int64{hdw.Id.Get()}, hdw)
}

// UpdateHrDepartureWizards updates existing hr.departure.wizard records.
// All records (represented by ids) will be updated by hdw values.
func (c *Client) UpdateHrDepartureWizards(ids []int64, hdw *HrDepartureWizard) error {
	return c.Update(HrDepartureWizardModel, ids, hdw)
}

// DeleteHrDepartureWizard deletes an existing hr.departure.wizard record.
func (c *Client) DeleteHrDepartureWizard(id int64) error {
	return c.DeleteHrDepartureWizards([]int64{id})
}

// DeleteHrDepartureWizards deletes existing hr.departure.wizard records.
func (c *Client) DeleteHrDepartureWizards(ids []int64) error {
	return c.Delete(HrDepartureWizardModel, ids)
}

// GetHrDepartureWizard gets hr.departure.wizard existing record.
func (c *Client) GetHrDepartureWizard(id int64) (*HrDepartureWizard, error) {
	hdws, err := c.GetHrDepartureWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hdws != nil && len(*hdws) > 0 {
		return &((*hdws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.departure.wizard not found", id)
}

// GetHrDepartureWizards gets hr.departure.wizard existing records.
func (c *Client) GetHrDepartureWizards(ids []int64) (*HrDepartureWizards, error) {
	hdws := &HrDepartureWizards{}
	if err := c.Read(HrDepartureWizardModel, ids, nil, hdws); err != nil {
		return nil, err
	}
	return hdws, nil
}

// FindHrDepartureWizard finds hr.departure.wizard record by querying it with criteria.
func (c *Client) FindHrDepartureWizard(criteria *Criteria) (*HrDepartureWizard, error) {
	hdws := &HrDepartureWizards{}
	if err := c.SearchRead(HrDepartureWizardModel, criteria, NewOptions().Limit(1), hdws); err != nil {
		return nil, err
	}
	if hdws != nil && len(*hdws) > 0 {
		return &((*hdws)[0]), nil
	}
	return nil, fmt.Errorf("hr.departure.wizard was not found")
}

// FindHrDepartureWizards finds hr.departure.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartureWizards(criteria *Criteria, options *Options) (*HrDepartureWizards, error) {
	hdws := &HrDepartureWizards{}
	if err := c.SearchRead(HrDepartureWizardModel, criteria, options, hdws); err != nil {
		return nil, err
	}
	return hdws, nil
}

// FindHrDepartureWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartureWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrDepartureWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrDepartureWizardId finds record id by querying it with criteria.
func (c *Client) FindHrDepartureWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrDepartureWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.departure.wizard was not found")
}
