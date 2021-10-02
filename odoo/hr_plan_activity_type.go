package odoo

import (
	"fmt"
)

// HrPlanActivityType represents hr.plan.activity.type model.
type HrPlanActivityType struct {
	ActivityTypeId *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Note           *String    `xmlrpc:"note,omptempty"`
	Responsible    *Selection `xmlrpc:"responsible,omptempty"`
	ResponsibleId  *Many2One  `xmlrpc:"responsible_id,omptempty"`
	Summary        *String    `xmlrpc:"summary,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrPlanActivityTypes represents array of hr.plan.activity.type model.
type HrPlanActivityTypes []HrPlanActivityType

// HrPlanActivityTypeModel is the odoo model name.
const HrPlanActivityTypeModel = "hr.plan.activity.type"

// Many2One convert HrPlanActivityType to *Many2One.
func (hpat *HrPlanActivityType) Many2One() *Many2One {
	return NewMany2One(hpat.Id.Get(), "")
}

// CreateHrPlanActivityType creates a new hr.plan.activity.type model and returns its id.
func (c *Client) CreateHrPlanActivityType(hpat *HrPlanActivityType) (int64, error) {
	return c.Create(HrPlanActivityTypeModel, hpat)
}

// UpdateHrPlanActivityType updates an existing hr.plan.activity.type record.
func (c *Client) UpdateHrPlanActivityType(hpat *HrPlanActivityType) error {
	return c.UpdateHrPlanActivityTypes([]int64{hpat.Id.Get()}, hpat)
}

// UpdateHrPlanActivityTypes updates existing hr.plan.activity.type records.
// All records (represented by ids) will be updated by hpat values.
func (c *Client) UpdateHrPlanActivityTypes(ids []int64, hpat *HrPlanActivityType) error {
	return c.Update(HrPlanActivityTypeModel, ids, hpat)
}

// DeleteHrPlanActivityType deletes an existing hr.plan.activity.type record.
func (c *Client) DeleteHrPlanActivityType(id int64) error {
	return c.DeleteHrPlanActivityTypes([]int64{id})
}

// DeleteHrPlanActivityTypes deletes existing hr.plan.activity.type records.
func (c *Client) DeleteHrPlanActivityTypes(ids []int64) error {
	return c.Delete(HrPlanActivityTypeModel, ids)
}

// GetHrPlanActivityType gets hr.plan.activity.type existing record.
func (c *Client) GetHrPlanActivityType(id int64) (*HrPlanActivityType, error) {
	hpats, err := c.GetHrPlanActivityTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpats != nil && len(*hpats) > 0 {
		return &((*hpats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.plan.activity.type not found", id)
}

// GetHrPlanActivityTypes gets hr.plan.activity.type existing records.
func (c *Client) GetHrPlanActivityTypes(ids []int64) (*HrPlanActivityTypes, error) {
	hpats := &HrPlanActivityTypes{}
	if err := c.Read(HrPlanActivityTypeModel, ids, nil, hpats); err != nil {
		return nil, err
	}
	return hpats, nil
}

// FindHrPlanActivityType finds hr.plan.activity.type record by querying it with criteria.
func (c *Client) FindHrPlanActivityType(criteria *Criteria) (*HrPlanActivityType, error) {
	hpats := &HrPlanActivityTypes{}
	if err := c.SearchRead(HrPlanActivityTypeModel, criteria, NewOptions().Limit(1), hpats); err != nil {
		return nil, err
	}
	if hpats != nil && len(*hpats) > 0 {
		return &((*hpats)[0]), nil
	}
	return nil, fmt.Errorf("hr.plan.activity.type was not found")
}

// FindHrPlanActivityTypes finds hr.plan.activity.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanActivityTypes(criteria *Criteria, options *Options) (*HrPlanActivityTypes, error) {
	hpats := &HrPlanActivityTypes{}
	if err := c.SearchRead(HrPlanActivityTypeModel, criteria, options, hpats); err != nil {
		return nil, err
	}
	return hpats, nil
}

// FindHrPlanActivityTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPlanActivityTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPlanActivityTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPlanActivityTypeId finds record id by querying it with criteria.
func (c *Client) FindHrPlanActivityTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPlanActivityTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.plan.activity.type was not found")
}
