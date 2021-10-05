package odoo

import (
	"fmt"
)

// GamificationBadgeUserWizard represents gamification.badge.user.wizard model.
type GamificationBadgeUserWizard struct {
	BadgeId     *Many2One `xmlrpc:"badge_id,omptempty"`
	Comment     *String   `xmlrpc:"comment,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// GamificationBadgeUserWizards represents array of gamification.badge.user.wizard model.
type GamificationBadgeUserWizards []GamificationBadgeUserWizard

// GamificationBadgeUserWizardModel is the odoo model name.
const GamificationBadgeUserWizardModel = "gamification.badge.user.wizard"

// Many2One convert GamificationBadgeUserWizard to *Many2One.
func (gbuw *GamificationBadgeUserWizard) Many2One() *Many2One {
	return NewMany2One(gbuw.Id.Get(), "")
}

// CreateGamificationBadgeUserWizard creates a new gamification.badge.user.wizard model and returns its id.
func (c *Client) CreateGamificationBadgeUserWizard(gbuw *GamificationBadgeUserWizard) (int64, error) {
	return c.Create(GamificationBadgeUserWizardModel, gbuw)
}

// UpdateGamificationBadgeUserWizard updates an existing gamification.badge.user.wizard record.
func (c *Client) UpdateGamificationBadgeUserWizard(gbuw *GamificationBadgeUserWizard) error {
	return c.UpdateGamificationBadgeUserWizards([]int64{gbuw.Id.Get()}, gbuw)
}

// UpdateGamificationBadgeUserWizards updates existing gamification.badge.user.wizard records.
// All records (represented by ids) will be updated by gbuw values.
func (c *Client) UpdateGamificationBadgeUserWizards(ids []int64, gbuw *GamificationBadgeUserWizard) error {
	return c.Update(GamificationBadgeUserWizardModel, ids, gbuw)
}

// DeleteGamificationBadgeUserWizard deletes an existing gamification.badge.user.wizard record.
func (c *Client) DeleteGamificationBadgeUserWizard(id int64) error {
	return c.DeleteGamificationBadgeUserWizards([]int64{id})
}

// DeleteGamificationBadgeUserWizards deletes existing gamification.badge.user.wizard records.
func (c *Client) DeleteGamificationBadgeUserWizards(ids []int64) error {
	return c.Delete(GamificationBadgeUserWizardModel, ids)
}

// GetGamificationBadgeUserWizard gets gamification.badge.user.wizard existing record.
func (c *Client) GetGamificationBadgeUserWizard(id int64) (*GamificationBadgeUserWizard, error) {
	gbuws, err := c.GetGamificationBadgeUserWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if gbuws != nil && len(*gbuws) > 0 {
		return &((*gbuws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.badge.user.wizard not found", id)
}

// GetGamificationBadgeUserWizards gets gamification.badge.user.wizard existing records.
func (c *Client) GetGamificationBadgeUserWizards(ids []int64) (*GamificationBadgeUserWizards, error) {
	gbuws := &GamificationBadgeUserWizards{}
	if err := c.Read(GamificationBadgeUserWizardModel, ids, nil, gbuws); err != nil {
		return nil, err
	}
	return gbuws, nil
}

// FindGamificationBadgeUserWizard finds gamification.badge.user.wizard record by querying it with criteria.
func (c *Client) FindGamificationBadgeUserWizard(criteria *Criteria) (*GamificationBadgeUserWizard, error) {
	gbuws := &GamificationBadgeUserWizards{}
	if err := c.SearchRead(GamificationBadgeUserWizardModel, criteria, NewOptions().Limit(1), gbuws); err != nil {
		return nil, err
	}
	if gbuws != nil && len(*gbuws) > 0 {
		return &((*gbuws)[0]), nil
	}
	return nil, fmt.Errorf("gamification.badge.user.wizard was not found")
}

// FindGamificationBadgeUserWizards finds gamification.badge.user.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeUserWizards(criteria *Criteria, options *Options) (*GamificationBadgeUserWizards, error) {
	gbuws := &GamificationBadgeUserWizards{}
	if err := c.SearchRead(GamificationBadgeUserWizardModel, criteria, options, gbuws); err != nil {
		return nil, err
	}
	return gbuws, nil
}

// FindGamificationBadgeUserWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeUserWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationBadgeUserWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationBadgeUserWizardId finds record id by querying it with criteria.
func (c *Client) FindGamificationBadgeUserWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationBadgeUserWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.badge.user.wizard was not found")
}
