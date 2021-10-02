package odoo

import (
	"fmt"
)

// CrmLeadScoringFrequency represents crm.lead.scoring.frequency model.
type CrmLeadScoringFrequency struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	LostCount   *Float    `xmlrpc:"lost_count,omptempty"`
	TeamId      *Many2One `xmlrpc:"team_id,omptempty"`
	Value       *String   `xmlrpc:"value,omptempty"`
	Variable    *String   `xmlrpc:"variable,omptempty"`
	WonCount    *Float    `xmlrpc:"won_count,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CrmLeadScoringFrequencys represents array of crm.lead.scoring.frequency model.
type CrmLeadScoringFrequencys []CrmLeadScoringFrequency

// CrmLeadScoringFrequencyModel is the odoo model name.
const CrmLeadScoringFrequencyModel = "crm.lead.scoring.frequency"

// Many2One convert CrmLeadScoringFrequency to *Many2One.
func (clsf *CrmLeadScoringFrequency) Many2One() *Many2One {
	return NewMany2One(clsf.Id.Get(), "")
}

// CreateCrmLeadScoringFrequency creates a new crm.lead.scoring.frequency model and returns its id.
func (c *Client) CreateCrmLeadScoringFrequency(clsf *CrmLeadScoringFrequency) (int64, error) {
	return c.Create(CrmLeadScoringFrequencyModel, clsf)
}

// UpdateCrmLeadScoringFrequency updates an existing crm.lead.scoring.frequency record.
func (c *Client) UpdateCrmLeadScoringFrequency(clsf *CrmLeadScoringFrequency) error {
	return c.UpdateCrmLeadScoringFrequencys([]int64{clsf.Id.Get()}, clsf)
}

// UpdateCrmLeadScoringFrequencys updates existing crm.lead.scoring.frequency records.
// All records (represented by ids) will be updated by clsf values.
func (c *Client) UpdateCrmLeadScoringFrequencys(ids []int64, clsf *CrmLeadScoringFrequency) error {
	return c.Update(CrmLeadScoringFrequencyModel, ids, clsf)
}

// DeleteCrmLeadScoringFrequency deletes an existing crm.lead.scoring.frequency record.
func (c *Client) DeleteCrmLeadScoringFrequency(id int64) error {
	return c.DeleteCrmLeadScoringFrequencys([]int64{id})
}

// DeleteCrmLeadScoringFrequencys deletes existing crm.lead.scoring.frequency records.
func (c *Client) DeleteCrmLeadScoringFrequencys(ids []int64) error {
	return c.Delete(CrmLeadScoringFrequencyModel, ids)
}

// GetCrmLeadScoringFrequency gets crm.lead.scoring.frequency existing record.
func (c *Client) GetCrmLeadScoringFrequency(id int64) (*CrmLeadScoringFrequency, error) {
	clsfs, err := c.GetCrmLeadScoringFrequencys([]int64{id})
	if err != nil {
		return nil, err
	}
	if clsfs != nil && len(*clsfs) > 0 {
		return &((*clsfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lead.scoring.frequency not found", id)
}

// GetCrmLeadScoringFrequencys gets crm.lead.scoring.frequency existing records.
func (c *Client) GetCrmLeadScoringFrequencys(ids []int64) (*CrmLeadScoringFrequencys, error) {
	clsfs := &CrmLeadScoringFrequencys{}
	if err := c.Read(CrmLeadScoringFrequencyModel, ids, nil, clsfs); err != nil {
		return nil, err
	}
	return clsfs, nil
}

// FindCrmLeadScoringFrequency finds crm.lead.scoring.frequency record by querying it with criteria.
func (c *Client) FindCrmLeadScoringFrequency(criteria *Criteria) (*CrmLeadScoringFrequency, error) {
	clsfs := &CrmLeadScoringFrequencys{}
	if err := c.SearchRead(CrmLeadScoringFrequencyModel, criteria, NewOptions().Limit(1), clsfs); err != nil {
		return nil, err
	}
	if clsfs != nil && len(*clsfs) > 0 {
		return &((*clsfs)[0]), nil
	}
	return nil, fmt.Errorf("crm.lead.scoring.frequency was not found")
}

// FindCrmLeadScoringFrequencys finds crm.lead.scoring.frequency records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadScoringFrequencys(criteria *Criteria, options *Options) (*CrmLeadScoringFrequencys, error) {
	clsfs := &CrmLeadScoringFrequencys{}
	if err := c.SearchRead(CrmLeadScoringFrequencyModel, criteria, options, clsfs); err != nil {
		return nil, err
	}
	return clsfs, nil
}

// FindCrmLeadScoringFrequencyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadScoringFrequencyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLeadScoringFrequencyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLeadScoringFrequencyId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadScoringFrequencyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadScoringFrequencyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.lead.scoring.frequency was not found")
}
