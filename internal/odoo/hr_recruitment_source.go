package odoo

import (
	"fmt"
)

// HrRecruitmentSource represents hr.recruitment.source model.
type HrRecruitmentSource struct {
	AliasId     *Many2One `xmlrpc:"alias_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Email       *String   `xmlrpc:"email,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JobId       *Many2One `xmlrpc:"job_id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	SourceId    *Many2One `xmlrpc:"source_id,omptempty"`
	Url         *String   `xmlrpc:"url,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrRecruitmentSources represents array of hr.recruitment.source model.
type HrRecruitmentSources []HrRecruitmentSource

// HrRecruitmentSourceModel is the odoo model name.
const HrRecruitmentSourceModel = "hr.recruitment.source"

// Many2One convert HrRecruitmentSource to *Many2One.
func (hrs *HrRecruitmentSource) Many2One() *Many2One {
	return NewMany2One(hrs.Id.Get(), "")
}

// CreateHrRecruitmentSource creates a new hr.recruitment.source model and returns its id.
func (c *Client) CreateHrRecruitmentSource(hrs *HrRecruitmentSource) (int64, error) {
	return c.Create(HrRecruitmentSourceModel, hrs)
}

// UpdateHrRecruitmentSource updates an existing hr.recruitment.source record.
func (c *Client) UpdateHrRecruitmentSource(hrs *HrRecruitmentSource) error {
	return c.UpdateHrRecruitmentSources([]int64{hrs.Id.Get()}, hrs)
}

// UpdateHrRecruitmentSources updates existing hr.recruitment.source records.
// All records (represented by ids) will be updated by hrs values.
func (c *Client) UpdateHrRecruitmentSources(ids []int64, hrs *HrRecruitmentSource) error {
	return c.Update(HrRecruitmentSourceModel, ids, hrs)
}

// DeleteHrRecruitmentSource deletes an existing hr.recruitment.source record.
func (c *Client) DeleteHrRecruitmentSource(id int64) error {
	return c.DeleteHrRecruitmentSources([]int64{id})
}

// DeleteHrRecruitmentSources deletes existing hr.recruitment.source records.
func (c *Client) DeleteHrRecruitmentSources(ids []int64) error {
	return c.Delete(HrRecruitmentSourceModel, ids)
}

// GetHrRecruitmentSource gets hr.recruitment.source existing record.
func (c *Client) GetHrRecruitmentSource(id int64) (*HrRecruitmentSource, error) {
	hrss, err := c.GetHrRecruitmentSources([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrss != nil && len(*hrss) > 0 {
		return &((*hrss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.recruitment.source not found", id)
}

// GetHrRecruitmentSources gets hr.recruitment.source existing records.
func (c *Client) GetHrRecruitmentSources(ids []int64) (*HrRecruitmentSources, error) {
	hrss := &HrRecruitmentSources{}
	if err := c.Read(HrRecruitmentSourceModel, ids, nil, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentSource finds hr.recruitment.source record by querying it with criteria.
func (c *Client) FindHrRecruitmentSource(criteria *Criteria) (*HrRecruitmentSource, error) {
	hrss := &HrRecruitmentSources{}
	if err := c.SearchRead(HrRecruitmentSourceModel, criteria, NewOptions().Limit(1), hrss); err != nil {
		return nil, err
	}
	if hrss != nil && len(*hrss) > 0 {
		return &((*hrss)[0]), nil
	}
	return nil, fmt.Errorf("hr.recruitment.source was not found")
}

// FindHrRecruitmentSources finds hr.recruitment.source records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentSources(criteria *Criteria, options *Options) (*HrRecruitmentSources, error) {
	hrss := &HrRecruitmentSources{}
	if err := c.SearchRead(HrRecruitmentSourceModel, criteria, options, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentSourceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentSourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRecruitmentSourceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRecruitmentSourceId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentSourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentSourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.recruitment.source was not found")
}
