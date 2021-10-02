package odoo

import (
	"fmt"
)

// SlideAnswer represents slide.answer model.
type SlideAnswer struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	IsCorrect   *Bool     `xmlrpc:"is_correct,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	QuestionId  *Many2One `xmlrpc:"question_id,omptempty"`
	TextValue   *String   `xmlrpc:"text_value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideAnswers represents array of slide.answer model.
type SlideAnswers []SlideAnswer

// SlideAnswerModel is the odoo model name.
const SlideAnswerModel = "slide.answer"

// Many2One convert SlideAnswer to *Many2One.
func (sa *SlideAnswer) Many2One() *Many2One {
	return NewMany2One(sa.Id.Get(), "")
}

// CreateSlideAnswer creates a new slide.answer model and returns its id.
func (c *Client) CreateSlideAnswer(sa *SlideAnswer) (int64, error) {
	return c.Create(SlideAnswerModel, sa)
}

// UpdateSlideAnswer updates an existing slide.answer record.
func (c *Client) UpdateSlideAnswer(sa *SlideAnswer) error {
	return c.UpdateSlideAnswers([]int64{sa.Id.Get()}, sa)
}

// UpdateSlideAnswers updates existing slide.answer records.
// All records (represented by ids) will be updated by sa values.
func (c *Client) UpdateSlideAnswers(ids []int64, sa *SlideAnswer) error {
	return c.Update(SlideAnswerModel, ids, sa)
}

// DeleteSlideAnswer deletes an existing slide.answer record.
func (c *Client) DeleteSlideAnswer(id int64) error {
	return c.DeleteSlideAnswers([]int64{id})
}

// DeleteSlideAnswers deletes existing slide.answer records.
func (c *Client) DeleteSlideAnswers(ids []int64) error {
	return c.Delete(SlideAnswerModel, ids)
}

// GetSlideAnswer gets slide.answer existing record.
func (c *Client) GetSlideAnswer(id int64) (*SlideAnswer, error) {
	sas, err := c.GetSlideAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.answer not found", id)
}

// GetSlideAnswers gets slide.answer existing records.
func (c *Client) GetSlideAnswers(ids []int64) (*SlideAnswers, error) {
	sas := &SlideAnswers{}
	if err := c.Read(SlideAnswerModel, ids, nil, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSlideAnswer finds slide.answer record by querying it with criteria.
func (c *Client) FindSlideAnswer(criteria *Criteria) (*SlideAnswer, error) {
	sas := &SlideAnswers{}
	if err := c.SearchRead(SlideAnswerModel, criteria, NewOptions().Limit(1), sas); err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("slide.answer was not found")
}

// FindSlideAnswers finds slide.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideAnswers(criteria *Criteria, options *Options) (*SlideAnswers, error) {
	sas := &SlideAnswers{}
	if err := c.SearchRead(SlideAnswerModel, criteria, options, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSlideAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideAnswerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideAnswerId finds record id by querying it with criteria.
func (c *Client) FindSlideAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.answer was not found")
}
