package odoo

import (
	"fmt"
)

// GamificationBadgeUser represents gamification.badge.user model.
type GamificationBadgeUser struct {
	BadgeId     *Many2One  `xmlrpc:"badge_id,omptempty"`
	BadgeName   *String    `xmlrpc:"badge_name,omptempty"`
	ChallengeId *Many2One  `xmlrpc:"challenge_id,omptempty"`
	Comment     *String    `xmlrpc:"comment,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId  *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Level       *Selection `xmlrpc:"level,omptempty"`
	SenderId    *Many2One  `xmlrpc:"sender_id,omptempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GamificationBadgeUsers represents array of gamification.badge.user model.
type GamificationBadgeUsers []GamificationBadgeUser

// GamificationBadgeUserModel is the odoo model name.
const GamificationBadgeUserModel = "gamification.badge.user"

// Many2One convert GamificationBadgeUser to *Many2One.
func (gbu *GamificationBadgeUser) Many2One() *Many2One {
	return NewMany2One(gbu.Id.Get(), "")
}

// CreateGamificationBadgeUser creates a new gamification.badge.user model and returns its id.
func (c *Client) CreateGamificationBadgeUser(gbu *GamificationBadgeUser) (int64, error) {
	return c.Create(GamificationBadgeUserModel, gbu)
}

// UpdateGamificationBadgeUser updates an existing gamification.badge.user record.
func (c *Client) UpdateGamificationBadgeUser(gbu *GamificationBadgeUser) error {
	return c.UpdateGamificationBadgeUsers([]int64{gbu.Id.Get()}, gbu)
}

// UpdateGamificationBadgeUsers updates existing gamification.badge.user records.
// All records (represented by ids) will be updated by gbu values.
func (c *Client) UpdateGamificationBadgeUsers(ids []int64, gbu *GamificationBadgeUser) error {
	return c.Update(GamificationBadgeUserModel, ids, gbu)
}

// DeleteGamificationBadgeUser deletes an existing gamification.badge.user record.
func (c *Client) DeleteGamificationBadgeUser(id int64) error {
	return c.DeleteGamificationBadgeUsers([]int64{id})
}

// DeleteGamificationBadgeUsers deletes existing gamification.badge.user records.
func (c *Client) DeleteGamificationBadgeUsers(ids []int64) error {
	return c.Delete(GamificationBadgeUserModel, ids)
}

// GetGamificationBadgeUser gets gamification.badge.user existing record.
func (c *Client) GetGamificationBadgeUser(id int64) (*GamificationBadgeUser, error) {
	gbus, err := c.GetGamificationBadgeUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	if gbus != nil && len(*gbus) > 0 {
		return &((*gbus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.badge.user not found", id)
}

// GetGamificationBadgeUsers gets gamification.badge.user existing records.
func (c *Client) GetGamificationBadgeUsers(ids []int64) (*GamificationBadgeUsers, error) {
	gbus := &GamificationBadgeUsers{}
	if err := c.Read(GamificationBadgeUserModel, ids, nil, gbus); err != nil {
		return nil, err
	}
	return gbus, nil
}

// FindGamificationBadgeUser finds gamification.badge.user record by querying it with criteria.
func (c *Client) FindGamificationBadgeUser(criteria *Criteria) (*GamificationBadgeUser, error) {
	gbus := &GamificationBadgeUsers{}
	if err := c.SearchRead(GamificationBadgeUserModel, criteria, NewOptions().Limit(1), gbus); err != nil {
		return nil, err
	}
	if gbus != nil && len(*gbus) > 0 {
		return &((*gbus)[0]), nil
	}
	return nil, fmt.Errorf("gamification.badge.user was not found")
}

// FindGamificationBadgeUsers finds gamification.badge.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeUsers(criteria *Criteria, options *Options) (*GamificationBadgeUsers, error) {
	gbus := &GamificationBadgeUsers{}
	if err := c.SearchRead(GamificationBadgeUserModel, criteria, options, gbus); err != nil {
		return nil, err
	}
	return gbus, nil
}

// FindGamificationBadgeUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationBadgeUserModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationBadgeUserId finds record id by querying it with criteria.
func (c *Client) FindGamificationBadgeUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationBadgeUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.badge.user was not found")
}
