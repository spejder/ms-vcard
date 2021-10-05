package odoo

import (
	"fmt"
)

// HrApplicantCategory represents hr.applicant.category model.
type HrApplicantCategory struct {
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrApplicantCategorys represents array of hr.applicant.category model.
type HrApplicantCategorys []HrApplicantCategory

// HrApplicantCategoryModel is the odoo model name.
const HrApplicantCategoryModel = "hr.applicant.category"

// Many2One convert HrApplicantCategory to *Many2One.
func (hac *HrApplicantCategory) Many2One() *Many2One {
	return NewMany2One(hac.Id.Get(), "")
}

// CreateHrApplicantCategory creates a new hr.applicant.category model and returns its id.
func (c *Client) CreateHrApplicantCategory(hac *HrApplicantCategory) (int64, error) {
	return c.Create(HrApplicantCategoryModel, hac)
}

// UpdateHrApplicantCategory updates an existing hr.applicant.category record.
func (c *Client) UpdateHrApplicantCategory(hac *HrApplicantCategory) error {
	return c.UpdateHrApplicantCategorys([]int64{hac.Id.Get()}, hac)
}

// UpdateHrApplicantCategorys updates existing hr.applicant.category records.
// All records (represented by ids) will be updated by hac values.
func (c *Client) UpdateHrApplicantCategorys(ids []int64, hac *HrApplicantCategory) error {
	return c.Update(HrApplicantCategoryModel, ids, hac)
}

// DeleteHrApplicantCategory deletes an existing hr.applicant.category record.
func (c *Client) DeleteHrApplicantCategory(id int64) error {
	return c.DeleteHrApplicantCategorys([]int64{id})
}

// DeleteHrApplicantCategorys deletes existing hr.applicant.category records.
func (c *Client) DeleteHrApplicantCategorys(ids []int64) error {
	return c.Delete(HrApplicantCategoryModel, ids)
}

// GetHrApplicantCategory gets hr.applicant.category existing record.
func (c *Client) GetHrApplicantCategory(id int64) (*HrApplicantCategory, error) {
	hacs, err := c.GetHrApplicantCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if hacs != nil && len(*hacs) > 0 {
		return &((*hacs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.applicant.category not found", id)
}

// GetHrApplicantCategorys gets hr.applicant.category existing records.
func (c *Client) GetHrApplicantCategorys(ids []int64) (*HrApplicantCategorys, error) {
	hacs := &HrApplicantCategorys{}
	if err := c.Read(HrApplicantCategoryModel, ids, nil, hacs); err != nil {
		return nil, err
	}
	return hacs, nil
}

// FindHrApplicantCategory finds hr.applicant.category record by querying it with criteria.
func (c *Client) FindHrApplicantCategory(criteria *Criteria) (*HrApplicantCategory, error) {
	hacs := &HrApplicantCategorys{}
	if err := c.SearchRead(HrApplicantCategoryModel, criteria, NewOptions().Limit(1), hacs); err != nil {
		return nil, err
	}
	if hacs != nil && len(*hacs) > 0 {
		return &((*hacs)[0]), nil
	}
	return nil, fmt.Errorf("hr.applicant.category was not found")
}

// FindHrApplicantCategorys finds hr.applicant.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantCategorys(criteria *Criteria, options *Options) (*HrApplicantCategorys, error) {
	hacs := &HrApplicantCategorys{}
	if err := c.SearchRead(HrApplicantCategoryModel, criteria, options, hacs); err != nil {
		return nil, err
	}
	return hacs, nil
}

// FindHrApplicantCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrApplicantCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrApplicantCategoryId finds record id by querying it with criteria.
func (c *Client) FindHrApplicantCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrApplicantCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.applicant.category was not found")
}
