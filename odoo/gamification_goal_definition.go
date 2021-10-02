package odoo

import (
	"fmt"
)

// GamificationGoalDefinition represents gamification.goal.definition model.
type GamificationGoalDefinition struct {
	ActionId              *Many2One  `xmlrpc:"action_id,omptempty"`
	BatchDistinctiveField *Many2One  `xmlrpc:"batch_distinctive_field,omptempty"`
	BatchMode             *Bool      `xmlrpc:"batch_mode,omptempty"`
	BatchUserExpression   *String    `xmlrpc:"batch_user_expression,omptempty"`
	ComputationMode       *Selection `xmlrpc:"computation_mode,omptempty"`
	ComputeCode           *String    `xmlrpc:"compute_code,omptempty"`
	Condition             *Selection `xmlrpc:"condition,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description           *String    `xmlrpc:"description,omptempty"`
	DisplayMode           *Selection `xmlrpc:"display_mode,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Domain                *String    `xmlrpc:"domain,omptempty"`
	FieldDateId           *Many2One  `xmlrpc:"field_date_id,omptempty"`
	FieldId               *Many2One  `xmlrpc:"field_id,omptempty"`
	FullSuffix            *String    `xmlrpc:"full_suffix,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	ModelId               *Many2One  `xmlrpc:"model_id,omptempty"`
	Monetary              *Bool      `xmlrpc:"monetary,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	ResIdField            *String    `xmlrpc:"res_id_field,omptempty"`
	Suffix                *String    `xmlrpc:"suffix,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GamificationGoalDefinitions represents array of gamification.goal.definition model.
type GamificationGoalDefinitions []GamificationGoalDefinition

// GamificationGoalDefinitionModel is the odoo model name.
const GamificationGoalDefinitionModel = "gamification.goal.definition"

// Many2One convert GamificationGoalDefinition to *Many2One.
func (ggd *GamificationGoalDefinition) Many2One() *Many2One {
	return NewMany2One(ggd.Id.Get(), "")
}

// CreateGamificationGoalDefinition creates a new gamification.goal.definition model and returns its id.
func (c *Client) CreateGamificationGoalDefinition(ggd *GamificationGoalDefinition) (int64, error) {
	return c.Create(GamificationGoalDefinitionModel, ggd)
}

// UpdateGamificationGoalDefinition updates an existing gamification.goal.definition record.
func (c *Client) UpdateGamificationGoalDefinition(ggd *GamificationGoalDefinition) error {
	return c.UpdateGamificationGoalDefinitions([]int64{ggd.Id.Get()}, ggd)
}

// UpdateGamificationGoalDefinitions updates existing gamification.goal.definition records.
// All records (represented by ids) will be updated by ggd values.
func (c *Client) UpdateGamificationGoalDefinitions(ids []int64, ggd *GamificationGoalDefinition) error {
	return c.Update(GamificationGoalDefinitionModel, ids, ggd)
}

// DeleteGamificationGoalDefinition deletes an existing gamification.goal.definition record.
func (c *Client) DeleteGamificationGoalDefinition(id int64) error {
	return c.DeleteGamificationGoalDefinitions([]int64{id})
}

// DeleteGamificationGoalDefinitions deletes existing gamification.goal.definition records.
func (c *Client) DeleteGamificationGoalDefinitions(ids []int64) error {
	return c.Delete(GamificationGoalDefinitionModel, ids)
}

// GetGamificationGoalDefinition gets gamification.goal.definition existing record.
func (c *Client) GetGamificationGoalDefinition(id int64) (*GamificationGoalDefinition, error) {
	ggds, err := c.GetGamificationGoalDefinitions([]int64{id})
	if err != nil {
		return nil, err
	}
	if ggds != nil && len(*ggds) > 0 {
		return &((*ggds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.goal.definition not found", id)
}

// GetGamificationGoalDefinitions gets gamification.goal.definition existing records.
func (c *Client) GetGamificationGoalDefinitions(ids []int64) (*GamificationGoalDefinitions, error) {
	ggds := &GamificationGoalDefinitions{}
	if err := c.Read(GamificationGoalDefinitionModel, ids, nil, ggds); err != nil {
		return nil, err
	}
	return ggds, nil
}

// FindGamificationGoalDefinition finds gamification.goal.definition record by querying it with criteria.
func (c *Client) FindGamificationGoalDefinition(criteria *Criteria) (*GamificationGoalDefinition, error) {
	ggds := &GamificationGoalDefinitions{}
	if err := c.SearchRead(GamificationGoalDefinitionModel, criteria, NewOptions().Limit(1), ggds); err != nil {
		return nil, err
	}
	if ggds != nil && len(*ggds) > 0 {
		return &((*ggds)[0]), nil
	}
	return nil, fmt.Errorf("gamification.goal.definition was not found")
}

// FindGamificationGoalDefinitions finds gamification.goal.definition records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalDefinitions(criteria *Criteria, options *Options) (*GamificationGoalDefinitions, error) {
	ggds := &GamificationGoalDefinitions{}
	if err := c.SearchRead(GamificationGoalDefinitionModel, criteria, options, ggds); err != nil {
		return nil, err
	}
	return ggds, nil
}

// FindGamificationGoalDefinitionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalDefinitionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationGoalDefinitionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationGoalDefinitionId finds record id by querying it with criteria.
func (c *Client) FindGamificationGoalDefinitionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationGoalDefinitionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.goal.definition was not found")
}
