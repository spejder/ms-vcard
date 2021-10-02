package odoo

import (
	"fmt"
)

// SlideQuestion represents slide.question model.
type SlideQuestion struct {
	AnswerIds            *Relation `xmlrpc:"answer_ids,omptempty"`
	AssessmentCriteriaId *Int      `xmlrpc:"assessment_criteria_id,omptempty"`
	AttemptsAvg          *Float    `xmlrpc:"attempts_avg,omptempty"`
	AttemptsCount        *Int      `xmlrpc:"attempts_count,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt            *Time     `xmlrpc:"deleted_at,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	DoneCount            *Int      `xmlrpc:"done_count,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	Question             *String   `xmlrpc:"question,omptempty"`
	Sequence             *Int      `xmlrpc:"sequence,omptempty"`
	SlideId              *Many2One `xmlrpc:"slide_id,omptempty"`
	Types                *String   `xmlrpc:"types,omptempty"`
	MultipleChoices      *String   `xmlrpc:"multiple_choices,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideQuestions represents array of slide.question model.
type SlideQuestions []SlideQuestion

// SlideQuestionModel is the odoo model name.
const SlideQuestionModel = "slide.question"

// Many2One convert SlideQuestion to *Many2One.
func (sq *SlideQuestion) Many2One() *Many2One {
	return NewMany2One(sq.Id.Get(), "")
}

// CreateSlideQuestion creates a new slide.question model and returns its id.
func (c *Client) CreateSlideQuestion(sq *SlideQuestion) (int64, error) {
	return c.Create(SlideQuestionModel, sq)
}

// UpdateSlideQuestion updates an existing slide.question record.
func (c *Client) UpdateSlideQuestion(sq *SlideQuestion) error {
	return c.UpdateSlideQuestions([]int64{sq.Id.Get()}, sq)
}

// UpdateSlideQuestions updates existing slide.question records.
// All records (represented by ids) will be updated by sq values.
func (c *Client) UpdateSlideQuestions(ids []int64, sq *SlideQuestion) error {
	return c.Update(SlideQuestionModel, ids, sq)
}

// DeleteSlideQuestion deletes an existing slide.question record.
func (c *Client) DeleteSlideQuestion(id int64) error {
	return c.DeleteSlideQuestions([]int64{id})
}

// DeleteSlideQuestions deletes existing slide.question records.
func (c *Client) DeleteSlideQuestions(ids []int64) error {
	return c.Delete(SlideQuestionModel, ids)
}

// GetSlideQuestion gets slide.question existing record.
func (c *Client) GetSlideQuestion(id int64) (*SlideQuestion, error) {
	sqs, err := c.GetSlideQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.question not found", id)
}

// GetSlideQuestions gets slide.question existing records.
func (c *Client) GetSlideQuestions(ids []int64) (*SlideQuestions, error) {
	sqs := &SlideQuestions{}
	if err := c.Read(SlideQuestionModel, ids, nil, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindSlideQuestion finds slide.question record by querying it with criteria.
func (c *Client) FindSlideQuestion(criteria *Criteria) (*SlideQuestion, error) {
	sqs := &SlideQuestions{}
	if err := c.SearchRead(SlideQuestionModel, criteria, NewOptions().Limit(1), sqs); err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("slide.question was not found")
}

// FindSlideQuestions finds slide.question records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideQuestions(criteria *Criteria, options *Options) (*SlideQuestions, error) {
	sqs := &SlideQuestions{}
	if err := c.SearchRead(SlideQuestionModel, criteria, options, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindSlideQuestionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideQuestionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideQuestionId finds record id by querying it with criteria.
func (c *Client) FindSlideQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.question was not found")
}
