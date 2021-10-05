package odoo

import (
	"fmt"
)

// SurveyLabel represents survey.label model.
type SurveyLabel struct {
	AnswerScore *Float    `xmlrpc:"answer_score,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	IsCorrect   *Bool     `xmlrpc:"is_correct,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	QuestionId  *Many2One `xmlrpc:"question_id,omptempty"`
	QuestionId2 *Many2One `xmlrpc:"question_id_2,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	Value       *String   `xmlrpc:"value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SurveyLabels represents array of survey.label model.
type SurveyLabels []SurveyLabel

// SurveyLabelModel is the odoo model name.
const SurveyLabelModel = "survey.label"

// Many2One convert SurveyLabel to *Many2One.
func (sl *SurveyLabel) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateSurveyLabel creates a new survey.label model and returns its id.
func (c *Client) CreateSurveyLabel(sl *SurveyLabel) (int64, error) {
	return c.Create(SurveyLabelModel, sl)
}

// UpdateSurveyLabel updates an existing survey.label record.
func (c *Client) UpdateSurveyLabel(sl *SurveyLabel) error {
	return c.UpdateSurveyLabels([]int64{sl.Id.Get()}, sl)
}

// UpdateSurveyLabels updates existing survey.label records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateSurveyLabels(ids []int64, sl *SurveyLabel) error {
	return c.Update(SurveyLabelModel, ids, sl)
}

// DeleteSurveyLabel deletes an existing survey.label record.
func (c *Client) DeleteSurveyLabel(id int64) error {
	return c.DeleteSurveyLabels([]int64{id})
}

// DeleteSurveyLabels deletes existing survey.label records.
func (c *Client) DeleteSurveyLabels(ids []int64) error {
	return c.Delete(SurveyLabelModel, ids)
}

// GetSurveyLabel gets survey.label existing record.
func (c *Client) GetSurveyLabel(id int64) (*SurveyLabel, error) {
	sls, err := c.GetSurveyLabels([]int64{id})
	if err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.label not found", id)
}

// GetSurveyLabels gets survey.label existing records.
func (c *Client) GetSurveyLabels(ids []int64) (*SurveyLabels, error) {
	sls := &SurveyLabels{}
	if err := c.Read(SurveyLabelModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSurveyLabel finds survey.label record by querying it with criteria.
func (c *Client) FindSurveyLabel(criteria *Criteria) (*SurveyLabel, error) {
	sls := &SurveyLabels{}
	if err := c.SearchRead(SurveyLabelModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("survey.label was not found")
}

// FindSurveyLabels finds survey.label records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyLabels(criteria *Criteria, options *Options) (*SurveyLabels, error) {
	sls := &SurveyLabels{}
	if err := c.SearchRead(SurveyLabelModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSurveyLabelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyLabelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyLabelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyLabelId finds record id by querying it with criteria.
func (c *Client) FindSurveyLabelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyLabelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.label was not found")
}
