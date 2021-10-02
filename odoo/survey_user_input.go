package odoo

import (
	"fmt"
)

// SurveyUserInput represents survey.user_input model.
type SurveyUserInput struct {
	AttemptNumber       *Int       `xmlrpc:"attempt_number,omptempty"`
	AttemptsLimit       *Int       `xmlrpc:"attempts_limit,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	Deadline            *Time      `xmlrpc:"deadline,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Email               *String    `xmlrpc:"email,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InputType           *Selection `xmlrpc:"input_type,omptempty"`
	InviteToken         *String    `xmlrpc:"invite_token,omptempty"`
	IsAttemptsLimited   *Bool      `xmlrpc:"is_attempts_limited,omptempty"`
	IsTimeLimitReached  *Bool      `xmlrpc:"is_time_limit_reached,omptempty"`
	LastDisplayedPageId *Many2One  `xmlrpc:"last_displayed_page_id,omptempty"`
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	QuestionIds         *Relation  `xmlrpc:"question_ids,omptempty"`
	QuizzPassed         *Bool      `xmlrpc:"quizz_passed,omptempty"`
	QuizzScore          *Float     `xmlrpc:"quizz_score,omptempty"`
	ScoringType         *Selection `xmlrpc:"scoring_type,omptempty"`
	StartDatetime       *Time      `xmlrpc:"start_datetime,omptempty"`
	State               *Selection `xmlrpc:"state,omptempty"`
	SurveyId            *Many2One  `xmlrpc:"survey_id,omptempty"`
	TestEntry           *Bool      `xmlrpc:"test_entry,omptempty"`
	Token               *String    `xmlrpc:"token,omptempty"`
	UserInputLineIds    *Relation  `xmlrpc:"user_input_line_ids,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyUserInputs represents array of survey.user_input model.
type SurveyUserInputs []SurveyUserInput

// SurveyUserInputModel is the odoo model name.
const SurveyUserInputModel = "survey.user_input"

// Many2One convert SurveyUserInput to *Many2One.
func (su *SurveyUserInput) Many2One() *Many2One {
	return NewMany2One(su.Id.Get(), "")
}

// CreateSurveyUserInput creates a new survey.user_input model and returns its id.
func (c *Client) CreateSurveyUserInput(su *SurveyUserInput) (int64, error) {
	return c.Create(SurveyUserInputModel, su)
}

// UpdateSurveyUserInput updates an existing survey.user_input record.
func (c *Client) UpdateSurveyUserInput(su *SurveyUserInput) error {
	return c.UpdateSurveyUserInputs([]int64{su.Id.Get()}, su)
}

// UpdateSurveyUserInputs updates existing survey.user_input records.
// All records (represented by ids) will be updated by su values.
func (c *Client) UpdateSurveyUserInputs(ids []int64, su *SurveyUserInput) error {
	return c.Update(SurveyUserInputModel, ids, su)
}

// DeleteSurveyUserInput deletes an existing survey.user_input record.
func (c *Client) DeleteSurveyUserInput(id int64) error {
	return c.DeleteSurveyUserInputs([]int64{id})
}

// DeleteSurveyUserInputs deletes existing survey.user_input records.
func (c *Client) DeleteSurveyUserInputs(ids []int64) error {
	return c.Delete(SurveyUserInputModel, ids)
}

// GetSurveyUserInput gets survey.user_input existing record.
func (c *Client) GetSurveyUserInput(id int64) (*SurveyUserInput, error) {
	sus, err := c.GetSurveyUserInputs([]int64{id})
	if err != nil {
		return nil, err
	}
	if sus != nil && len(*sus) > 0 {
		return &((*sus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.user_input not found", id)
}

// GetSurveyUserInputs gets survey.user_input existing records.
func (c *Client) GetSurveyUserInputs(ids []int64) (*SurveyUserInputs, error) {
	sus := &SurveyUserInputs{}
	if err := c.Read(SurveyUserInputModel, ids, nil, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInput finds survey.user_input record by querying it with criteria.
func (c *Client) FindSurveyUserInput(criteria *Criteria) (*SurveyUserInput, error) {
	sus := &SurveyUserInputs{}
	if err := c.SearchRead(SurveyUserInputModel, criteria, NewOptions().Limit(1), sus); err != nil {
		return nil, err
	}
	if sus != nil && len(*sus) > 0 {
		return &((*sus)[0]), nil
	}
	return nil, fmt.Errorf("survey.user_input was not found")
}

// FindSurveyUserInputs finds survey.user_input records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputs(criteria *Criteria, options *Options) (*SurveyUserInputs, error) {
	sus := &SurveyUserInputs{}
	if err := c.SearchRead(SurveyUserInputModel, criteria, options, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInputIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyUserInputModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyUserInputId finds record id by querying it with criteria.
func (c *Client) FindSurveyUserInputId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyUserInputModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.user_input was not found")
}
