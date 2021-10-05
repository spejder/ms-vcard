package odoo

import (
	"fmt"
)

// HrRecruitmentStage represents hr.recruitment.stage model.
type HrRecruitmentStage struct {
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Fold          *Bool     `xmlrpc:"fold,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	JobIds        *Relation `xmlrpc:"job_ids,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	LegendBlocked *String   `xmlrpc:"legend_blocked,omptempty"`
	LegendDone    *String   `xmlrpc:"legend_done,omptempty"`
	LegendNormal  *String   `xmlrpc:"legend_normal,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	Requirements  *String   `xmlrpc:"requirements,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	TemplateId    *Many2One `xmlrpc:"template_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrRecruitmentStages represents array of hr.recruitment.stage model.
type HrRecruitmentStages []HrRecruitmentStage

// HrRecruitmentStageModel is the odoo model name.
const HrRecruitmentStageModel = "hr.recruitment.stage"

// Many2One convert HrRecruitmentStage to *Many2One.
func (hrs *HrRecruitmentStage) Many2One() *Many2One {
	return NewMany2One(hrs.Id.Get(), "")
}

// CreateHrRecruitmentStage creates a new hr.recruitment.stage model and returns its id.
func (c *Client) CreateHrRecruitmentStage(hrs *HrRecruitmentStage) (int64, error) {
	return c.Create(HrRecruitmentStageModel, hrs)
}

// UpdateHrRecruitmentStage updates an existing hr.recruitment.stage record.
func (c *Client) UpdateHrRecruitmentStage(hrs *HrRecruitmentStage) error {
	return c.UpdateHrRecruitmentStages([]int64{hrs.Id.Get()}, hrs)
}

// UpdateHrRecruitmentStages updates existing hr.recruitment.stage records.
// All records (represented by ids) will be updated by hrs values.
func (c *Client) UpdateHrRecruitmentStages(ids []int64, hrs *HrRecruitmentStage) error {
	return c.Update(HrRecruitmentStageModel, ids, hrs)
}

// DeleteHrRecruitmentStage deletes an existing hr.recruitment.stage record.
func (c *Client) DeleteHrRecruitmentStage(id int64) error {
	return c.DeleteHrRecruitmentStages([]int64{id})
}

// DeleteHrRecruitmentStages deletes existing hr.recruitment.stage records.
func (c *Client) DeleteHrRecruitmentStages(ids []int64) error {
	return c.Delete(HrRecruitmentStageModel, ids)
}

// GetHrRecruitmentStage gets hr.recruitment.stage existing record.
func (c *Client) GetHrRecruitmentStage(id int64) (*HrRecruitmentStage, error) {
	hrss, err := c.GetHrRecruitmentStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrss != nil && len(*hrss) > 0 {
		return &((*hrss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.recruitment.stage not found", id)
}

// GetHrRecruitmentStages gets hr.recruitment.stage existing records.
func (c *Client) GetHrRecruitmentStages(ids []int64) (*HrRecruitmentStages, error) {
	hrss := &HrRecruitmentStages{}
	if err := c.Read(HrRecruitmentStageModel, ids, nil, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentStage finds hr.recruitment.stage record by querying it with criteria.
func (c *Client) FindHrRecruitmentStage(criteria *Criteria) (*HrRecruitmentStage, error) {
	hrss := &HrRecruitmentStages{}
	if err := c.SearchRead(HrRecruitmentStageModel, criteria, NewOptions().Limit(1), hrss); err != nil {
		return nil, err
	}
	if hrss != nil && len(*hrss) > 0 {
		return &((*hrss)[0]), nil
	}
	return nil, fmt.Errorf("hr.recruitment.stage was not found")
}

// FindHrRecruitmentStages finds hr.recruitment.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentStages(criteria *Criteria, options *Options) (*HrRecruitmentStages, error) {
	hrss := &HrRecruitmentStages{}
	if err := c.SearchRead(HrRecruitmentStageModel, criteria, options, hrss); err != nil {
		return nil, err
	}
	return hrss, nil
}

// FindHrRecruitmentStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRecruitmentStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRecruitmentStageId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.recruitment.stage was not found")
}
