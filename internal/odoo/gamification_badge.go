package odoo

import (
	"fmt"
)

// GamificationBadge represents gamification.badge model.
type GamificationBadge struct {
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	CanPublish               *Bool      `xmlrpc:"can_publish,omptempty"`
	ChallengeIds             *Relation  `xmlrpc:"challenge_ids,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	GoalDefinitionIds        *Relation  `xmlrpc:"goal_definition_ids,omptempty"`
	GrantedEmployeesCount    *Int       `xmlrpc:"granted_employees_count,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Image1024                *String    `xmlrpc:"image_1024,omptempty"`
	Image128                 *String    `xmlrpc:"image_128,omptempty"`
	Image1920                *String    `xmlrpc:"image_1920,omptempty"`
	Image256                 *String    `xmlrpc:"image_256,omptempty"`
	Image512                 *String    `xmlrpc:"image_512,omptempty"`
	IsPublished              *Bool      `xmlrpc:"is_published,omptempty"`
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Level                    *Selection `xmlrpc:"level,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	OwnerIds                 *Relation  `xmlrpc:"owner_ids,omptempty"`
	RemainingSending         *Int       `xmlrpc:"remaining_sending,omptempty"`
	RuleAuth                 *Selection `xmlrpc:"rule_auth,omptempty"`
	RuleAuthBadgeIds         *Relation  `xmlrpc:"rule_auth_badge_ids,omptempty"`
	RuleAuthUserIds          *Relation  `xmlrpc:"rule_auth_user_ids,omptempty"`
	RuleMax                  *Bool      `xmlrpc:"rule_max,omptempty"`
	RuleMaxNumber            *Int       `xmlrpc:"rule_max_number,omptempty"`
	StatCount                *Int       `xmlrpc:"stat_count,omptempty"`
	StatCountDistinct        *Int       `xmlrpc:"stat_count_distinct,omptempty"`
	StatMy                   *Int       `xmlrpc:"stat_my,omptempty"`
	StatMyMonthlySending     *Int       `xmlrpc:"stat_my_monthly_sending,omptempty"`
	StatMyThisMonth          *Int       `xmlrpc:"stat_my_this_month,omptempty"`
	StatThisMonth            *Int       `xmlrpc:"stat_this_month,omptempty"`
	SurveyId                 *Many2One  `xmlrpc:"survey_id,omptempty"`
	SurveyIds                *Relation  `xmlrpc:"survey_ids,omptempty"`
	UniqueOwnerIds           *Relation  `xmlrpc:"unique_owner_ids,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsitePublished         *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl               *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GamificationBadges represents array of gamification.badge model.
type GamificationBadges []GamificationBadge

// GamificationBadgeModel is the odoo model name.
const GamificationBadgeModel = "gamification.badge"

// Many2One convert GamificationBadge to *Many2One.
func (gb *GamificationBadge) Many2One() *Many2One {
	return NewMany2One(gb.Id.Get(), "")
}

// CreateGamificationBadge creates a new gamification.badge model and returns its id.
func (c *Client) CreateGamificationBadge(gb *GamificationBadge) (int64, error) {
	return c.Create(GamificationBadgeModel, gb)
}

// UpdateGamificationBadge updates an existing gamification.badge record.
func (c *Client) UpdateGamificationBadge(gb *GamificationBadge) error {
	return c.UpdateGamificationBadges([]int64{gb.Id.Get()}, gb)
}

// UpdateGamificationBadges updates existing gamification.badge records.
// All records (represented by ids) will be updated by gb values.
func (c *Client) UpdateGamificationBadges(ids []int64, gb *GamificationBadge) error {
	return c.Update(GamificationBadgeModel, ids, gb)
}

// DeleteGamificationBadge deletes an existing gamification.badge record.
func (c *Client) DeleteGamificationBadge(id int64) error {
	return c.DeleteGamificationBadges([]int64{id})
}

// DeleteGamificationBadges deletes existing gamification.badge records.
func (c *Client) DeleteGamificationBadges(ids []int64) error {
	return c.Delete(GamificationBadgeModel, ids)
}

// GetGamificationBadge gets gamification.badge existing record.
func (c *Client) GetGamificationBadge(id int64) (*GamificationBadge, error) {
	gbs, err := c.GetGamificationBadges([]int64{id})
	if err != nil {
		return nil, err
	}
	if gbs != nil && len(*gbs) > 0 {
		return &((*gbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.badge not found", id)
}

// GetGamificationBadges gets gamification.badge existing records.
func (c *Client) GetGamificationBadges(ids []int64) (*GamificationBadges, error) {
	gbs := &GamificationBadges{}
	if err := c.Read(GamificationBadgeModel, ids, nil, gbs); err != nil {
		return nil, err
	}
	return gbs, nil
}

// FindGamificationBadge finds gamification.badge record by querying it with criteria.
func (c *Client) FindGamificationBadge(criteria *Criteria) (*GamificationBadge, error) {
	gbs := &GamificationBadges{}
	if err := c.SearchRead(GamificationBadgeModel, criteria, NewOptions().Limit(1), gbs); err != nil {
		return nil, err
	}
	if gbs != nil && len(*gbs) > 0 {
		return &((*gbs)[0]), nil
	}
	return nil, fmt.Errorf("gamification.badge was not found")
}

// FindGamificationBadges finds gamification.badge records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadges(criteria *Criteria, options *Options) (*GamificationBadges, error) {
	gbs := &GamificationBadges{}
	if err := c.SearchRead(GamificationBadgeModel, criteria, options, gbs); err != nil {
		return nil, err
	}
	return gbs, nil
}

// FindGamificationBadgeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationBadgeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationBadgeId finds record id by querying it with criteria.
func (c *Client) FindGamificationBadgeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationBadgeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.badge was not found")
}
