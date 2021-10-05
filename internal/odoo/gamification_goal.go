package odoo

import (
	"fmt"
)

// GamificationGoal represents gamification.goal model.
type GamificationGoal struct {
	ChallengeId           *Many2One  `xmlrpc:"challenge_id,omptempty"`
	Closed                *Bool      `xmlrpc:"closed,omptempty"`
	Completeness          *Float     `xmlrpc:"completeness,omptempty"`
	ComputationMode       *Selection `xmlrpc:"computation_mode,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	Current               *Float     `xmlrpc:"current,omptempty"`
	DefinitionCondition   *Selection `xmlrpc:"definition_condition,omptempty"`
	DefinitionDescription *String    `xmlrpc:"definition_description,omptempty"`
	DefinitionDisplay     *Selection `xmlrpc:"definition_display,omptempty"`
	DefinitionId          *Many2One  `xmlrpc:"definition_id,omptempty"`
	DefinitionSuffix      *String    `xmlrpc:"definition_suffix,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	EndDate               *Time      `xmlrpc:"end_date,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	LastUpdate            *Time      `xmlrpc:"last_update,omptempty"`
	LineId                *Many2One  `xmlrpc:"line_id,omptempty"`
	RemindUpdateDelay     *Int       `xmlrpc:"remind_update_delay,omptempty"`
	StartDate             *Time      `xmlrpc:"start_date,omptempty"`
	State                 *Selection `xmlrpc:"state,omptempty"`
	TargetGoal            *Float     `xmlrpc:"target_goal,omptempty"`
	ToUpdate              *Bool      `xmlrpc:"to_update,omptempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GamificationGoals represents array of gamification.goal model.
type GamificationGoals []GamificationGoal

// GamificationGoalModel is the odoo model name.
const GamificationGoalModel = "gamification.goal"

// Many2One convert GamificationGoal to *Many2One.
func (gg *GamificationGoal) Many2One() *Many2One {
	return NewMany2One(gg.Id.Get(), "")
}

// CreateGamificationGoal creates a new gamification.goal model and returns its id.
func (c *Client) CreateGamificationGoal(gg *GamificationGoal) (int64, error) {
	return c.Create(GamificationGoalModel, gg)
}

// UpdateGamificationGoal updates an existing gamification.goal record.
func (c *Client) UpdateGamificationGoal(gg *GamificationGoal) error {
	return c.UpdateGamificationGoals([]int64{gg.Id.Get()}, gg)
}

// UpdateGamificationGoals updates existing gamification.goal records.
// All records (represented by ids) will be updated by gg values.
func (c *Client) UpdateGamificationGoals(ids []int64, gg *GamificationGoal) error {
	return c.Update(GamificationGoalModel, ids, gg)
}

// DeleteGamificationGoal deletes an existing gamification.goal record.
func (c *Client) DeleteGamificationGoal(id int64) error {
	return c.DeleteGamificationGoals([]int64{id})
}

// DeleteGamificationGoals deletes existing gamification.goal records.
func (c *Client) DeleteGamificationGoals(ids []int64) error {
	return c.Delete(GamificationGoalModel, ids)
}

// GetGamificationGoal gets gamification.goal existing record.
func (c *Client) GetGamificationGoal(id int64) (*GamificationGoal, error) {
	ggs, err := c.GetGamificationGoals([]int64{id})
	if err != nil {
		return nil, err
	}
	if ggs != nil && len(*ggs) > 0 {
		return &((*ggs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.goal not found", id)
}

// GetGamificationGoals gets gamification.goal existing records.
func (c *Client) GetGamificationGoals(ids []int64) (*GamificationGoals, error) {
	ggs := &GamificationGoals{}
	if err := c.Read(GamificationGoalModel, ids, nil, ggs); err != nil {
		return nil, err
	}
	return ggs, nil
}

// FindGamificationGoal finds gamification.goal record by querying it with criteria.
func (c *Client) FindGamificationGoal(criteria *Criteria) (*GamificationGoal, error) {
	ggs := &GamificationGoals{}
	if err := c.SearchRead(GamificationGoalModel, criteria, NewOptions().Limit(1), ggs); err != nil {
		return nil, err
	}
	if ggs != nil && len(*ggs) > 0 {
		return &((*ggs)[0]), nil
	}
	return nil, fmt.Errorf("gamification.goal was not found")
}

// FindGamificationGoals finds gamification.goal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoals(criteria *Criteria, options *Options) (*GamificationGoals, error) {
	ggs := &GamificationGoals{}
	if err := c.SearchRead(GamificationGoalModel, criteria, options, ggs); err != nil {
		return nil, err
	}
	return ggs, nil
}

// FindGamificationGoalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationGoalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationGoalId finds record id by querying it with criteria.
func (c *Client) FindGamificationGoalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationGoalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.goal was not found")
}
