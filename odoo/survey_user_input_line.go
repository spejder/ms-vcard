package odoo

import (
	"fmt"
)

// SurveyUserInputLine represents survey.user_input_line model.
type SurveyUserInputLine struct {
	AnswerIsCorrect   *Bool      `xmlrpc:"answer_is_correct,omptempty"`
	AnswerScore       *Float     `xmlrpc:"answer_score,omptempty"`
	AnswerType        *Selection `xmlrpc:"answer_type,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	PageId            *Many2One  `xmlrpc:"page_id,omptempty"`
	QuestionId        *Many2One  `xmlrpc:"question_id,omptempty"`
	QuestionSequence  *Int       `xmlrpc:"question_sequence,omptempty"`
	Skipped           *Bool      `xmlrpc:"skipped,omptempty"`
	SurveyId          *Many2One  `xmlrpc:"survey_id,omptempty"`
	UserInputId       *Many2One  `xmlrpc:"user_input_id,omptempty"`
	ValueDate         *Time      `xmlrpc:"value_date,omptempty"`
	ValueDatetime     *Time      `xmlrpc:"value_datetime,omptempty"`
	ValueFreeText     *String    `xmlrpc:"value_free_text,omptempty"`
	ValueNumber       *Float     `xmlrpc:"value_number,omptempty"`
	ValueSuggested    *Many2One  `xmlrpc:"value_suggested,omptempty"`
	ValueSuggestedRow *Many2One  `xmlrpc:"value_suggested_row,omptempty"`
	ValueText         *String    `xmlrpc:"value_text,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyUserInputLines represents array of survey.user_input_line model.
type SurveyUserInputLines []SurveyUserInputLine

// SurveyUserInputLineModel is the odoo model name.
const SurveyUserInputLineModel = "survey.user_input_line"

// Many2One convert SurveyUserInputLine to *Many2One.
func (su *SurveyUserInputLine) Many2One() *Many2One {
	return NewMany2One(su.Id.Get(), "")
}

// CreateSurveyUserInputLine creates a new survey.user_input_line model and returns its id.
func (c *Client) CreateSurveyUserInputLine(su *SurveyUserInputLine) (int64, error) {
	return c.Create(SurveyUserInputLineModel, su)
}

// UpdateSurveyUserInputLine updates an existing survey.user_input_line record.
func (c *Client) UpdateSurveyUserInputLine(su *SurveyUserInputLine) error {
	return c.UpdateSurveyUserInputLines([]int64{su.Id.Get()}, su)
}

// UpdateSurveyUserInputLines updates existing survey.user_input_line records.
// All records (represented by ids) will be updated by su values.
func (c *Client) UpdateSurveyUserInputLines(ids []int64, su *SurveyUserInputLine) error {
	return c.Update(SurveyUserInputLineModel, ids, su)
}

// DeleteSurveyUserInputLine deletes an existing survey.user_input_line record.
func (c *Client) DeleteSurveyUserInputLine(id int64) error {
	return c.DeleteSurveyUserInputLines([]int64{id})
}

// DeleteSurveyUserInputLines deletes existing survey.user_input_line records.
func (c *Client) DeleteSurveyUserInputLines(ids []int64) error {
	return c.Delete(SurveyUserInputLineModel, ids)
}

// GetSurveyUserInputLine gets survey.user_input_line existing record.
func (c *Client) GetSurveyUserInputLine(id int64) (*SurveyUserInputLine, error) {
	sus, err := c.GetSurveyUserInputLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if sus != nil && len(*sus) > 0 {
		return &((*sus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.user_input_line not found", id)
}

// GetSurveyUserInputLines gets survey.user_input_line existing records.
func (c *Client) GetSurveyUserInputLines(ids []int64) (*SurveyUserInputLines, error) {
	sus := &SurveyUserInputLines{}
	if err := c.Read(SurveyUserInputLineModel, ids, nil, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInputLine finds survey.user_input_line record by querying it with criteria.
func (c *Client) FindSurveyUserInputLine(criteria *Criteria) (*SurveyUserInputLine, error) {
	sus := &SurveyUserInputLines{}
	if err := c.SearchRead(SurveyUserInputLineModel, criteria, NewOptions().Limit(1), sus); err != nil {
		return nil, err
	}
	if sus != nil && len(*sus) > 0 {
		return &((*sus)[0]), nil
	}
	return nil, fmt.Errorf("survey.user_input_line was not found")
}

// FindSurveyUserInputLines finds survey.user_input_line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputLines(criteria *Criteria, options *Options) (*SurveyUserInputLines, error) {
	sus := &SurveyUserInputLines{}
	if err := c.SearchRead(SurveyUserInputLineModel, criteria, options, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInputLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyUserInputLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyUserInputLineId finds record id by querying it with criteria.
func (c *Client) FindSurveyUserInputLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyUserInputLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.user_input_line was not found")
}
