package odoo

import (
	"fmt"
)

// GamificationChallengeLine represents gamification.challenge.line model.
type GamificationChallengeLine struct {
	ChallengeId          *Many2One  `xmlrpc:"challenge_id,omptempty"`
	Condition            *Selection `xmlrpc:"condition,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefinitionFullSuffix *String    `xmlrpc:"definition_full_suffix,omptempty"`
	DefinitionId         *Many2One  `xmlrpc:"definition_id,omptempty"`
	DefinitionMonetary   *Bool      `xmlrpc:"definition_monetary,omptempty"`
	DefinitionSuffix     *String    `xmlrpc:"definition_suffix,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	Sequence             *Int       `xmlrpc:"sequence,omptempty"`
	TargetGoal           *Float     `xmlrpc:"target_goal,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GamificationChallengeLines represents array of gamification.challenge.line model.
type GamificationChallengeLines []GamificationChallengeLine

// GamificationChallengeLineModel is the odoo model name.
const GamificationChallengeLineModel = "gamification.challenge.line"

// Many2One convert GamificationChallengeLine to *Many2One.
func (gcl *GamificationChallengeLine) Many2One() *Many2One {
	return NewMany2One(gcl.Id.Get(), "")
}

// CreateGamificationChallengeLine creates a new gamification.challenge.line model and returns its id.
func (c *Client) CreateGamificationChallengeLine(gcl *GamificationChallengeLine) (int64, error) {
	return c.Create(GamificationChallengeLineModel, gcl)
}

// UpdateGamificationChallengeLine updates an existing gamification.challenge.line record.
func (c *Client) UpdateGamificationChallengeLine(gcl *GamificationChallengeLine) error {
	return c.UpdateGamificationChallengeLines([]int64{gcl.Id.Get()}, gcl)
}

// UpdateGamificationChallengeLines updates existing gamification.challenge.line records.
// All records (represented by ids) will be updated by gcl values.
func (c *Client) UpdateGamificationChallengeLines(ids []int64, gcl *GamificationChallengeLine) error {
	return c.Update(GamificationChallengeLineModel, ids, gcl)
}

// DeleteGamificationChallengeLine deletes an existing gamification.challenge.line record.
func (c *Client) DeleteGamificationChallengeLine(id int64) error {
	return c.DeleteGamificationChallengeLines([]int64{id})
}

// DeleteGamificationChallengeLines deletes existing gamification.challenge.line records.
func (c *Client) DeleteGamificationChallengeLines(ids []int64) error {
	return c.Delete(GamificationChallengeLineModel, ids)
}

// GetGamificationChallengeLine gets gamification.challenge.line existing record.
func (c *Client) GetGamificationChallengeLine(id int64) (*GamificationChallengeLine, error) {
	gcls, err := c.GetGamificationChallengeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if gcls != nil && len(*gcls) > 0 {
		return &((*gcls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.challenge.line not found", id)
}

// GetGamificationChallengeLines gets gamification.challenge.line existing records.
func (c *Client) GetGamificationChallengeLines(ids []int64) (*GamificationChallengeLines, error) {
	gcls := &GamificationChallengeLines{}
	if err := c.Read(GamificationChallengeLineModel, ids, nil, gcls); err != nil {
		return nil, err
	}
	return gcls, nil
}

// FindGamificationChallengeLine finds gamification.challenge.line record by querying it with criteria.
func (c *Client) FindGamificationChallengeLine(criteria *Criteria) (*GamificationChallengeLine, error) {
	gcls := &GamificationChallengeLines{}
	if err := c.SearchRead(GamificationChallengeLineModel, criteria, NewOptions().Limit(1), gcls); err != nil {
		return nil, err
	}
	if gcls != nil && len(*gcls) > 0 {
		return &((*gcls)[0]), nil
	}
	return nil, fmt.Errorf("gamification.challenge.line was not found")
}

// FindGamificationChallengeLines finds gamification.challenge.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallengeLines(criteria *Criteria, options *Options) (*GamificationChallengeLines, error) {
	gcls := &GamificationChallengeLines{}
	if err := c.SearchRead(GamificationChallengeLineModel, criteria, options, gcls); err != nil {
		return nil, err
	}
	return gcls, nil
}

// FindGamificationChallengeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallengeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationChallengeLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationChallengeLineId finds record id by querying it with criteria.
func (c *Client) FindGamificationChallengeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationChallengeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.challenge.line was not found")
}
