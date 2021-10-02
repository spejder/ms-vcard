package odoo

import (
	"fmt"
)

// SurveySurvey represents survey.survey model.
type SurveySurvey struct {
	AccessMode                  *Selection `xmlrpc:"access_mode,omptempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AnswerCount                 *Int       `xmlrpc:"answer_count,omptempty"`
	AnswerDoneCount             *Int       `xmlrpc:"answer_done_count,omptempty"`
	AnswerScoreAvg              *Float     `xmlrpc:"answer_score_avg,omptempty"`
	AttemptsLimit               *Int       `xmlrpc:"attempts_limit,omptempty"`
	Category                    *Selection `xmlrpc:"category,omptempty"`
	Certificate                 *Bool      `xmlrpc:"certificate,omptempty"`
	CertificationBadgeId        *Many2One  `xmlrpc:"certification_badge_id,omptempty"`
	CertificationBadgeIdDummy   *Many2One  `xmlrpc:"certification_badge_id_dummy,omptempty"`
	CertificationGiveBadge      *Bool      `xmlrpc:"certification_give_badge,omptempty"`
	CertificationMailTemplateId *Many2One  `xmlrpc:"certification_mail_template_id,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IsAttemptsLimited           *Bool      `xmlrpc:"is_attempts_limited,omptempty"`
	IsTimeLimited               *Bool      `xmlrpc:"is_time_limited,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	PageIds                     *Relation  `xmlrpc:"page_ids,omptempty"`
	PassingScore                *Float     `xmlrpc:"passing_score,omptempty"`
	PublicUrl                   *String    `xmlrpc:"public_url,omptempty"`
	QuestionAndPageIds          *Relation  `xmlrpc:"question_and_page_ids,omptempty"`
	QuestionIds                 *Relation  `xmlrpc:"question_ids,omptempty"`
	QuestionsLayout             *Selection `xmlrpc:"questions_layout,omptempty"`
	QuestionsSelection          *Selection `xmlrpc:"questions_selection,omptempty"`
	ScoringType                 *Selection `xmlrpc:"scoring_type,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	SuccessCount                *Int       `xmlrpc:"success_count,omptempty"`
	SuccessRatio                *Int       `xmlrpc:"success_ratio,omptempty"`
	ThankYouMessage             *String    `xmlrpc:"thank_you_message,omptempty"`
	TimeLimit                   *Float     `xmlrpc:"time_limit,omptempty"`
	Title                       *String    `xmlrpc:"title,omptempty"`
	UserInputIds                *Relation  `xmlrpc:"user_input_ids,omptempty"`
	UsersCanGoBack              *Bool      `xmlrpc:"users_can_go_back,omptempty"`
	UsersCanSignup              *Bool      `xmlrpc:"users_can_signup,omptempty"`
	UsersLoginRequired          *Bool      `xmlrpc:"users_login_required,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveySurveys represents array of survey.survey model.
type SurveySurveys []SurveySurvey

// SurveySurveyModel is the odoo model name.
const SurveySurveyModel = "survey.survey"

// Many2One convert SurveySurvey to *Many2One.
func (ss *SurveySurvey) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSurveySurvey creates a new survey.survey model and returns its id.
func (c *Client) CreateSurveySurvey(ss *SurveySurvey) (int64, error) {
	return c.Create(SurveySurveyModel, ss)
}

// UpdateSurveySurvey updates an existing survey.survey record.
func (c *Client) UpdateSurveySurvey(ss *SurveySurvey) error {
	return c.UpdateSurveySurveys([]int64{ss.Id.Get()}, ss)
}

// UpdateSurveySurveys updates existing survey.survey records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSurveySurveys(ids []int64, ss *SurveySurvey) error {
	return c.Update(SurveySurveyModel, ids, ss)
}

// DeleteSurveySurvey deletes an existing survey.survey record.
func (c *Client) DeleteSurveySurvey(id int64) error {
	return c.DeleteSurveySurveys([]int64{id})
}

// DeleteSurveySurveys deletes existing survey.survey records.
func (c *Client) DeleteSurveySurveys(ids []int64) error {
	return c.Delete(SurveySurveyModel, ids)
}

// GetSurveySurvey gets survey.survey existing record.
func (c *Client) GetSurveySurvey(id int64) (*SurveySurvey, error) {
	sss, err := c.GetSurveySurveys([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.survey not found", id)
}

// GetSurveySurveys gets survey.survey existing records.
func (c *Client) GetSurveySurveys(ids []int64) (*SurveySurveys, error) {
	sss := &SurveySurveys{}
	if err := c.Read(SurveySurveyModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSurveySurvey finds survey.survey record by querying it with criteria.
func (c *Client) FindSurveySurvey(criteria *Criteria) (*SurveySurvey, error) {
	sss := &SurveySurveys{}
	if err := c.SearchRead(SurveySurveyModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("survey.survey was not found")
}

// FindSurveySurveys finds survey.survey records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveySurveys(criteria *Criteria, options *Options) (*SurveySurveys, error) {
	sss := &SurveySurveys{}
	if err := c.SearchRead(SurveySurveyModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSurveySurveyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveySurveyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveySurveyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveySurveyId finds record id by querying it with criteria.
func (c *Client) FindSurveySurveyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveySurveyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.survey was not found")
}
