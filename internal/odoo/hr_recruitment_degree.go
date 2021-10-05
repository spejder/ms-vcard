package odoo

import (
	"fmt"
)

// HrRecruitmentDegree represents hr.recruitment.degree model.
type HrRecruitmentDegree struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrRecruitmentDegrees represents array of hr.recruitment.degree model.
type HrRecruitmentDegrees []HrRecruitmentDegree

// HrRecruitmentDegreeModel is the odoo model name.
const HrRecruitmentDegreeModel = "hr.recruitment.degree"

// Many2One convert HrRecruitmentDegree to *Many2One.
func (hrd *HrRecruitmentDegree) Many2One() *Many2One {
	return NewMany2One(hrd.Id.Get(), "")
}

// CreateHrRecruitmentDegree creates a new hr.recruitment.degree model and returns its id.
func (c *Client) CreateHrRecruitmentDegree(hrd *HrRecruitmentDegree) (int64, error) {
	return c.Create(HrRecruitmentDegreeModel, hrd)
}

// UpdateHrRecruitmentDegree updates an existing hr.recruitment.degree record.
func (c *Client) UpdateHrRecruitmentDegree(hrd *HrRecruitmentDegree) error {
	return c.UpdateHrRecruitmentDegrees([]int64{hrd.Id.Get()}, hrd)
}

// UpdateHrRecruitmentDegrees updates existing hr.recruitment.degree records.
// All records (represented by ids) will be updated by hrd values.
func (c *Client) UpdateHrRecruitmentDegrees(ids []int64, hrd *HrRecruitmentDegree) error {
	return c.Update(HrRecruitmentDegreeModel, ids, hrd)
}

// DeleteHrRecruitmentDegree deletes an existing hr.recruitment.degree record.
func (c *Client) DeleteHrRecruitmentDegree(id int64) error {
	return c.DeleteHrRecruitmentDegrees([]int64{id})
}

// DeleteHrRecruitmentDegrees deletes existing hr.recruitment.degree records.
func (c *Client) DeleteHrRecruitmentDegrees(ids []int64) error {
	return c.Delete(HrRecruitmentDegreeModel, ids)
}

// GetHrRecruitmentDegree gets hr.recruitment.degree existing record.
func (c *Client) GetHrRecruitmentDegree(id int64) (*HrRecruitmentDegree, error) {
	hrds, err := c.GetHrRecruitmentDegrees([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrds != nil && len(*hrds) > 0 {
		return &((*hrds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.recruitment.degree not found", id)
}

// GetHrRecruitmentDegrees gets hr.recruitment.degree existing records.
func (c *Client) GetHrRecruitmentDegrees(ids []int64) (*HrRecruitmentDegrees, error) {
	hrds := &HrRecruitmentDegrees{}
	if err := c.Read(HrRecruitmentDegreeModel, ids, nil, hrds); err != nil {
		return nil, err
	}
	return hrds, nil
}

// FindHrRecruitmentDegree finds hr.recruitment.degree record by querying it with criteria.
func (c *Client) FindHrRecruitmentDegree(criteria *Criteria) (*HrRecruitmentDegree, error) {
	hrds := &HrRecruitmentDegrees{}
	if err := c.SearchRead(HrRecruitmentDegreeModel, criteria, NewOptions().Limit(1), hrds); err != nil {
		return nil, err
	}
	if hrds != nil && len(*hrds) > 0 {
		return &((*hrds)[0]), nil
	}
	return nil, fmt.Errorf("hr.recruitment.degree was not found")
}

// FindHrRecruitmentDegrees finds hr.recruitment.degree records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentDegrees(criteria *Criteria, options *Options) (*HrRecruitmentDegrees, error) {
	hrds := &HrRecruitmentDegrees{}
	if err := c.SearchRead(HrRecruitmentDegreeModel, criteria, options, hrds); err != nil {
		return nil, err
	}
	return hrds, nil
}

// FindHrRecruitmentDegreeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentDegreeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRecruitmentDegreeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRecruitmentDegreeId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentDegreeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentDegreeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.recruitment.degree was not found")
}
