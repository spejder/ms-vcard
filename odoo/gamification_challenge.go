package odoo

import (
	"fmt"
)

// GamificationChallenge represents gamification.challenge model.
type GamificationChallenge struct {
	Category                 *Selection `xmlrpc:"category,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	EndDate                  *Time      `xmlrpc:"end_date,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	InvitedUserIds           *Relation  `xmlrpc:"invited_user_ids,omptempty"`
	LastReportDate           *Time      `xmlrpc:"last_report_date,omptempty"`
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	LineIds                  *Relation  `xmlrpc:"line_ids,omptempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omptempty"`
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
	NextReportDate           *Time      `xmlrpc:"next_report_date,omptempty"`
	Period                   *Selection `xmlrpc:"period,omptempty"`
	RemindUpdateDelay        *Int       `xmlrpc:"remind_update_delay,omptempty"`
	ReportMessageFrequency   *Selection `xmlrpc:"report_message_frequency,omptempty"`
	ReportMessageGroupId     *Many2One  `xmlrpc:"report_message_group_id,omptempty"`
	ReportTemplateId         *Many2One  `xmlrpc:"report_template_id,omptempty"`
	RewardFailure            *Bool      `xmlrpc:"reward_failure,omptempty"`
	RewardFirstId            *Many2One  `xmlrpc:"reward_first_id,omptempty"`
	RewardId                 *Many2One  `xmlrpc:"reward_id,omptempty"`
	RewardRealtime           *Bool      `xmlrpc:"reward_realtime,omptempty"`
	RewardSecondId           *Many2One  `xmlrpc:"reward_second_id,omptempty"`
	RewardThirdId            *Many2One  `xmlrpc:"reward_third_id,omptempty"`
	StartDate                *Time      `xmlrpc:"start_date,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	UserDomain               *String    `xmlrpc:"user_domain,omptempty"`
	UserIds                  *Relation  `xmlrpc:"user_ids,omptempty"`
	VisibilityMode           *Selection `xmlrpc:"visibility_mode,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GamificationChallenges represents array of gamification.challenge model.
type GamificationChallenges []GamificationChallenge

// GamificationChallengeModel is the odoo model name.
const GamificationChallengeModel = "gamification.challenge"

// Many2One convert GamificationChallenge to *Many2One.
func (gc *GamificationChallenge) Many2One() *Many2One {
	return NewMany2One(gc.Id.Get(), "")
}

// CreateGamificationChallenge creates a new gamification.challenge model and returns its id.
func (c *Client) CreateGamificationChallenge(gc *GamificationChallenge) (int64, error) {
	return c.Create(GamificationChallengeModel, gc)
}

// UpdateGamificationChallenge updates an existing gamification.challenge record.
func (c *Client) UpdateGamificationChallenge(gc *GamificationChallenge) error {
	return c.UpdateGamificationChallenges([]int64{gc.Id.Get()}, gc)
}

// UpdateGamificationChallenges updates existing gamification.challenge records.
// All records (represented by ids) will be updated by gc values.
func (c *Client) UpdateGamificationChallenges(ids []int64, gc *GamificationChallenge) error {
	return c.Update(GamificationChallengeModel, ids, gc)
}

// DeleteGamificationChallenge deletes an existing gamification.challenge record.
func (c *Client) DeleteGamificationChallenge(id int64) error {
	return c.DeleteGamificationChallenges([]int64{id})
}

// DeleteGamificationChallenges deletes existing gamification.challenge records.
func (c *Client) DeleteGamificationChallenges(ids []int64) error {
	return c.Delete(GamificationChallengeModel, ids)
}

// GetGamificationChallenge gets gamification.challenge existing record.
func (c *Client) GetGamificationChallenge(id int64) (*GamificationChallenge, error) {
	gcs, err := c.GetGamificationChallenges([]int64{id})
	if err != nil {
		return nil, err
	}
	if gcs != nil && len(*gcs) > 0 {
		return &((*gcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.challenge not found", id)
}

// GetGamificationChallenges gets gamification.challenge existing records.
func (c *Client) GetGamificationChallenges(ids []int64) (*GamificationChallenges, error) {
	gcs := &GamificationChallenges{}
	if err := c.Read(GamificationChallengeModel, ids, nil, gcs); err != nil {
		return nil, err
	}
	return gcs, nil
}

// FindGamificationChallenge finds gamification.challenge record by querying it with criteria.
func (c *Client) FindGamificationChallenge(criteria *Criteria) (*GamificationChallenge, error) {
	gcs := &GamificationChallenges{}
	if err := c.SearchRead(GamificationChallengeModel, criteria, NewOptions().Limit(1), gcs); err != nil {
		return nil, err
	}
	if gcs != nil && len(*gcs) > 0 {
		return &((*gcs)[0]), nil
	}
	return nil, fmt.Errorf("gamification.challenge was not found")
}

// FindGamificationChallenges finds gamification.challenge records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallenges(criteria *Criteria, options *Options) (*GamificationChallenges, error) {
	gcs := &GamificationChallenges{}
	if err := c.SearchRead(GamificationChallengeModel, criteria, options, gcs); err != nil {
		return nil, err
	}
	return gcs, nil
}

// FindGamificationChallengeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallengeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationChallengeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationChallengeId finds record id by querying it with criteria.
func (c *Client) FindGamificationChallengeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationChallengeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.challenge was not found")
}
