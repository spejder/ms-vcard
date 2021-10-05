package odoo

import (
	"fmt"
)

// CrmLeadScoringFrequencyField represents crm.lead.scoring.frequency.field model.
type CrmLeadScoringFrequencyField struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FieldId     *Many2One `xmlrpc:"field_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CrmLeadScoringFrequencyFields represents array of crm.lead.scoring.frequency.field model.
type CrmLeadScoringFrequencyFields []CrmLeadScoringFrequencyField

// CrmLeadScoringFrequencyFieldModel is the odoo model name.
const CrmLeadScoringFrequencyFieldModel = "crm.lead.scoring.frequency.field"

// Many2One convert CrmLeadScoringFrequencyField to *Many2One.
func (clsff *CrmLeadScoringFrequencyField) Many2One() *Many2One {
	return NewMany2One(clsff.Id.Get(), "")
}

// CreateCrmLeadScoringFrequencyField creates a new crm.lead.scoring.frequency.field model and returns its id.
func (c *Client) CreateCrmLeadScoringFrequencyField(clsff *CrmLeadScoringFrequencyField) (int64, error) {
	return c.Create(CrmLeadScoringFrequencyFieldModel, clsff)
}

// UpdateCrmLeadScoringFrequencyField updates an existing crm.lead.scoring.frequency.field record.
func (c *Client) UpdateCrmLeadScoringFrequencyField(clsff *CrmLeadScoringFrequencyField) error {
	return c.UpdateCrmLeadScoringFrequencyFields([]int64{clsff.Id.Get()}, clsff)
}

// UpdateCrmLeadScoringFrequencyFields updates existing crm.lead.scoring.frequency.field records.
// All records (represented by ids) will be updated by clsff values.
func (c *Client) UpdateCrmLeadScoringFrequencyFields(ids []int64, clsff *CrmLeadScoringFrequencyField) error {
	return c.Update(CrmLeadScoringFrequencyFieldModel, ids, clsff)
}

// DeleteCrmLeadScoringFrequencyField deletes an existing crm.lead.scoring.frequency.field record.
func (c *Client) DeleteCrmLeadScoringFrequencyField(id int64) error {
	return c.DeleteCrmLeadScoringFrequencyFields([]int64{id})
}

// DeleteCrmLeadScoringFrequencyFields deletes existing crm.lead.scoring.frequency.field records.
func (c *Client) DeleteCrmLeadScoringFrequencyFields(ids []int64) error {
	return c.Delete(CrmLeadScoringFrequencyFieldModel, ids)
}

// GetCrmLeadScoringFrequencyField gets crm.lead.scoring.frequency.field existing record.
func (c *Client) GetCrmLeadScoringFrequencyField(id int64) (*CrmLeadScoringFrequencyField, error) {
	clsffs, err := c.GetCrmLeadScoringFrequencyFields([]int64{id})
	if err != nil {
		return nil, err
	}
	if clsffs != nil && len(*clsffs) > 0 {
		return &((*clsffs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lead.scoring.frequency.field not found", id)
}

// GetCrmLeadScoringFrequencyFields gets crm.lead.scoring.frequency.field existing records.
func (c *Client) GetCrmLeadScoringFrequencyFields(ids []int64) (*CrmLeadScoringFrequencyFields, error) {
	clsffs := &CrmLeadScoringFrequencyFields{}
	if err := c.Read(CrmLeadScoringFrequencyFieldModel, ids, nil, clsffs); err != nil {
		return nil, err
	}
	return clsffs, nil
}

// FindCrmLeadScoringFrequencyField finds crm.lead.scoring.frequency.field record by querying it with criteria.
func (c *Client) FindCrmLeadScoringFrequencyField(criteria *Criteria) (*CrmLeadScoringFrequencyField, error) {
	clsffs := &CrmLeadScoringFrequencyFields{}
	if err := c.SearchRead(CrmLeadScoringFrequencyFieldModel, criteria, NewOptions().Limit(1), clsffs); err != nil {
		return nil, err
	}
	if clsffs != nil && len(*clsffs) > 0 {
		return &((*clsffs)[0]), nil
	}
	return nil, fmt.Errorf("crm.lead.scoring.frequency.field was not found")
}

// FindCrmLeadScoringFrequencyFields finds crm.lead.scoring.frequency.field records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadScoringFrequencyFields(criteria *Criteria, options *Options) (*CrmLeadScoringFrequencyFields, error) {
	clsffs := &CrmLeadScoringFrequencyFields{}
	if err := c.SearchRead(CrmLeadScoringFrequencyFieldModel, criteria, options, clsffs); err != nil {
		return nil, err
	}
	return clsffs, nil
}

// FindCrmLeadScoringFrequencyFieldIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadScoringFrequencyFieldIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLeadScoringFrequencyFieldModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLeadScoringFrequencyFieldId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadScoringFrequencyFieldId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadScoringFrequencyFieldModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.lead.scoring.frequency.field was not found")
}
