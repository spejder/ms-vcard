package odoo

import (
	"fmt"
)

// GamificationKarmaRank represents gamification.karma.rank model.
type GamificationKarmaRank struct {
	CreateDate              *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omptempty"`
	Description             *String   `xmlrpc:"description,omptempty"`
	DescriptionMotivational *String   `xmlrpc:"description_motivational,omptempty"`
	DisplayName             *String   `xmlrpc:"display_name,omptempty"`
	Id                      *Int      `xmlrpc:"id,omptempty"`
	Image1024               *String   `xmlrpc:"image_1024,omptempty"`
	Image128                *String   `xmlrpc:"image_128,omptempty"`
	Image1920               *String   `xmlrpc:"image_1920,omptempty"`
	Image256                *String   `xmlrpc:"image_256,omptempty"`
	Image512                *String   `xmlrpc:"image_512,omptempty"`
	KarmaMin                *Int      `xmlrpc:"karma_min,omptempty"`
	LastUpdate              *Time     `xmlrpc:"__last_update,omptempty"`
	Name                    *String   `xmlrpc:"name,omptempty"`
	UserIds                 *Relation `xmlrpc:"user_ids,omptempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omptempty"`
}

// GamificationKarmaRanks represents array of gamification.karma.rank model.
type GamificationKarmaRanks []GamificationKarmaRank

// GamificationKarmaRankModel is the odoo model name.
const GamificationKarmaRankModel = "gamification.karma.rank"

// Many2One convert GamificationKarmaRank to *Many2One.
func (gkr *GamificationKarmaRank) Many2One() *Many2One {
	return NewMany2One(gkr.Id.Get(), "")
}

// CreateGamificationKarmaRank creates a new gamification.karma.rank model and returns its id.
func (c *Client) CreateGamificationKarmaRank(gkr *GamificationKarmaRank) (int64, error) {
	return c.Create(GamificationKarmaRankModel, gkr)
}

// UpdateGamificationKarmaRank updates an existing gamification.karma.rank record.
func (c *Client) UpdateGamificationKarmaRank(gkr *GamificationKarmaRank) error {
	return c.UpdateGamificationKarmaRanks([]int64{gkr.Id.Get()}, gkr)
}

// UpdateGamificationKarmaRanks updates existing gamification.karma.rank records.
// All records (represented by ids) will be updated by gkr values.
func (c *Client) UpdateGamificationKarmaRanks(ids []int64, gkr *GamificationKarmaRank) error {
	return c.Update(GamificationKarmaRankModel, ids, gkr)
}

// DeleteGamificationKarmaRank deletes an existing gamification.karma.rank record.
func (c *Client) DeleteGamificationKarmaRank(id int64) error {
	return c.DeleteGamificationKarmaRanks([]int64{id})
}

// DeleteGamificationKarmaRanks deletes existing gamification.karma.rank records.
func (c *Client) DeleteGamificationKarmaRanks(ids []int64) error {
	return c.Delete(GamificationKarmaRankModel, ids)
}

// GetGamificationKarmaRank gets gamification.karma.rank existing record.
func (c *Client) GetGamificationKarmaRank(id int64) (*GamificationKarmaRank, error) {
	gkrs, err := c.GetGamificationKarmaRanks([]int64{id})
	if err != nil {
		return nil, err
	}
	if gkrs != nil && len(*gkrs) > 0 {
		return &((*gkrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.karma.rank not found", id)
}

// GetGamificationKarmaRanks gets gamification.karma.rank existing records.
func (c *Client) GetGamificationKarmaRanks(ids []int64) (*GamificationKarmaRanks, error) {
	gkrs := &GamificationKarmaRanks{}
	if err := c.Read(GamificationKarmaRankModel, ids, nil, gkrs); err != nil {
		return nil, err
	}
	return gkrs, nil
}

// FindGamificationKarmaRank finds gamification.karma.rank record by querying it with criteria.
func (c *Client) FindGamificationKarmaRank(criteria *Criteria) (*GamificationKarmaRank, error) {
	gkrs := &GamificationKarmaRanks{}
	if err := c.SearchRead(GamificationKarmaRankModel, criteria, NewOptions().Limit(1), gkrs); err != nil {
		return nil, err
	}
	if gkrs != nil && len(*gkrs) > 0 {
		return &((*gkrs)[0]), nil
	}
	return nil, fmt.Errorf("gamification.karma.rank was not found")
}

// FindGamificationKarmaRanks finds gamification.karma.rank records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationKarmaRanks(criteria *Criteria, options *Options) (*GamificationKarmaRanks, error) {
	gkrs := &GamificationKarmaRanks{}
	if err := c.SearchRead(GamificationKarmaRankModel, criteria, options, gkrs); err != nil {
		return nil, err
	}
	return gkrs, nil
}

// FindGamificationKarmaRankIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationKarmaRankIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationKarmaRankModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationKarmaRankId finds record id by querying it with criteria.
func (c *Client) FindGamificationKarmaRankId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationKarmaRankModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.karma.rank was not found")
}
