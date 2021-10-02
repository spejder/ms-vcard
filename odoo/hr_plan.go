package odoo

import (
	"fmt"
)

// HrPlan represents hr.plan model.
type HrPlan struct {
	Active              *Bool     `xmlrpc:"active,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	PlanActivityTypeIds *Relation `xmlrpc:"plan_activity_type_ids,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPlans represents array of hr.plan model.
type HrPlans []HrPlan

// HrPlanModel is the odoo model name.
const HrPlanModel = "hr.plan"

// Many2One convert HrPlan to *Many2One.
func (hp *HrPlan) Many2One() *Many2One {
	return NewMany2One(hp.Id.Get(), "")
}

// CreateHrPlan creates a new hr.plan model and returns its id.
func (c *Client) CreateHrPlan(hp *HrPlan) (int64, error) {
	return c.Create(HrPlanModel, hp)
}

// UpdateHrPlan updates an existing hr.plan record.
func (c *Client) UpdateHrPlan(hp *HrPlan) error {
	return c.UpdateHrPlans([]int64{hp.Id.Get()}, hp)
}

// UpdateHrPlans updates existing hr.plan records.
// All records (represented by ids) will be updated by hp values.
func (c *Client) UpdateHrPlans(ids []int64, hp *HrPlan) error {
	return c.Update(HrPlanModel, ids, hp)
}

// DeleteHrPlan deletes an existing hr.plan record.
func (c *Client) DeleteHrPlan(id int64) error {
	return c.DeleteHrPlans([]int64{id})
}

// DeleteHrPlans deletes existing hr.plan records.
func (c *Client) DeleteHrPlans(ids []int64) error {
	return c.Delete(HrPlanModel, ids)
}

// GetHrPlan gets hr.plan existing record.
func (c *Client) GetHrPlan(id int64) (*HrPlan, error) {
	hps, err := c.GetHrPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	if hps != nil && len(*hps) > 0 {
		return &((*hps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.plan not found", id)
}

// GetHrPlans gets hr.plan existing records.
func (c *Client) GetHrPlans(ids []int64) (*HrPlans, error) {
	hps := &HrPlans{}
	if err := c.Read(HrPlanModel, ids, nil, hps); err != nil {
		return nil, err
	}
	return hps, nil
}

// FindHrPlan finds hr.plan record by querying it with criteria.
func (c *Client) FindHrPlan(criteria *Criteria) (*HrPlan, error) {
	hps := &HrPlans{}
	if err := c.SearchRead(HrPlanModel, criteria, NewOptions().Limit(1), hps); err != nil {
		return nil, err
	}
	if hps != nil && len(*hps) > 0 {
		return &((*hps)[0]), nil
	}
	return nil, fmt.Errorf("hr.plan was not found")
}

// FindHrPlans finds hr.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlans(criteria *Criteria, options *Options) (*HrPlans, error) {
	hps := &HrPlans{}
	if err := c.SearchRead(HrPlanModel, criteria, options, hps); err != nil {
		return nil, err
	}
	return hps, nil
}

// FindHrPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPlanModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPlanId finds record id by querying it with criteria.
func (c *Client) FindHrPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.plan was not found")
}
