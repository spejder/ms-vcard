package odoo

import (
	"fmt"
)

// SlideAnswerUsers represents slide.answer_users model.
type SlideAnswerUsers struct {
	Attachment     *String   `xmlrpc:"attachment,omptempty"`
	AttachmentName *String   `xmlrpc:"attachment_name,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	QuestionId     *Many2One `xmlrpc:"question_id,omptempty"`
	ScheduleId     *Int      `xmlrpc:"schedule_id,omptempty"`
	SlideId        *Many2One `xmlrpc:"slide_id,omptempty"`
	Competent      *Bool     `xmlrpc:"competent,omptempty"`
	Type           *String   `xmlrpc:"type,omptempty"`
	Url            *String   `xmlrpc:"url,omptempty"`
	UserId         *Many2One `xmlrpc:"user_id,omptempty"`
	Value          *String   `xmlrpc:"value,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}



// SlideAnswerUserss represents array of slide.answer_users model.
type SlideAnswerUserss []SlideAnswerUsers

// SlideAnswerUsersModel is the odoo model name.
const SlideAnswerUsersModel = "slide.answer_users"

// Many2One convert SlideAnswerUsers to *Many2One.
func (sa *SlideAnswerUsers) Many2One() *Many2One {
	return NewMany2One(sa.Id.Get(), "")
}

// CreateSlideAnswerUsers creates a new slide.answer_users model and returns its id.
func (c *Client) CreateSlideAnswerUsers(sa *SlideAnswerUsers) (int64, error) {
	return c.Create(SlideAnswerUsersModel, sa)
}

// UpdateSlideAnswerUsers updates an existing slide.answer_users record.
func (c *Client) UpdateSlideAnswerUsers(sa *SlideAnswerUsers) error {
	return c.UpdateSlideAnswerUserss([]int64{sa.Id.Get()}, sa)
}

// UpdateSlideAnswerUserss updates existing slide.answer_users records.
// All records (represented by ids) will be updated by sa values.
func (c *Client) UpdateSlideAnswerUserss(ids []int64, sa *SlideAnswerUsers) error {
	return c.Update(SlideAnswerUsersModel, ids, sa)
}

// DeleteSlideAnswerUsers deletes an existing slide.answer_users record.
func (c *Client) DeleteSlideAnswerUsers(id int64) error {
	return c.DeleteSlideAnswerUserss([]int64{id})
}

// DeleteSlideAnswerUserss deletes existing slide.answer_users records.
func (c *Client) DeleteSlideAnswerUserss(ids []int64) error {
	return c.Delete(SlideAnswerUsersModel, ids)
}

// GetSlideAnswerUsers gets slide.answer_users existing record.
func (c *Client) GetSlideAnswerUsers(id int64) (*SlideAnswerUsers, error) {
	sas, err := c.GetSlideAnswerUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.answer_users not found", id)
}

// GetSlideAnswerUserss gets slide.answer_users existing records.
func (c *Client) GetSlideAnswerUserss(ids []int64) (*SlideAnswerUserss, error) {
	sas := &SlideAnswerUserss{}
	if err := c.Read(SlideAnswerUsersModel, ids, nil, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSlideAnswerUsers finds slide.answer_users record by querying it with criteria.
func (c *Client) FindSlideAnswerUsers(criteria *Criteria) (*SlideAnswerUsers, error) {
	sas := &SlideAnswerUserss{}
	if err := c.SearchRead(SlideAnswerUsersModel, criteria, NewOptions().Limit(1), sas); err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("slide.answer_users was not found")
}

// FindSlideAnswerUserss finds slide.answer_users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideAnswerUserss(criteria *Criteria, options *Options) (*SlideAnswerUserss, error) {
	sas := &SlideAnswerUserss{}
	if err := c.SearchRead(SlideAnswerUsersModel, criteria, options, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSlideAnswerUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideAnswerUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideAnswerUsersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideAnswerUsersId finds record id by querying it with criteria.
func (c *Client) FindSlideAnswerUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideAnswerUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.answer_users was not found")
}
