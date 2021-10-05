package odoo

import (
	"fmt"
)

// SlideTag represents slide.tag model.
type SlideTag struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideTags represents array of slide.tag model.
type SlideTags []SlideTag

// SlideTagModel is the odoo model name.
const SlideTagModel = "slide.tag"

// Many2One convert SlideTag to *Many2One.
func (st *SlideTag) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSlideTag creates a new slide.tag model and returns its id.
func (c *Client) CreateSlideTag(st *SlideTag) (int64, error) {
	return c.Create(SlideTagModel, st)
}

// UpdateSlideTag updates an existing slide.tag record.
func (c *Client) UpdateSlideTag(st *SlideTag) error {
	return c.UpdateSlideTags([]int64{st.Id.Get()}, st)
}

// UpdateSlideTags updates existing slide.tag records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSlideTags(ids []int64, st *SlideTag) error {
	return c.Update(SlideTagModel, ids, st)
}

// DeleteSlideTag deletes an existing slide.tag record.
func (c *Client) DeleteSlideTag(id int64) error {
	return c.DeleteSlideTags([]int64{id})
}

// DeleteSlideTags deletes existing slide.tag records.
func (c *Client) DeleteSlideTags(ids []int64) error {
	return c.Delete(SlideTagModel, ids)
}

// GetSlideTag gets slide.tag existing record.
func (c *Client) GetSlideTag(id int64) (*SlideTag, error) {
	sts, err := c.GetSlideTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.tag not found", id)
}

// GetSlideTags gets slide.tag existing records.
func (c *Client) GetSlideTags(ids []int64) (*SlideTags, error) {
	sts := &SlideTags{}
	if err := c.Read(SlideTagModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSlideTag finds slide.tag record by querying it with criteria.
func (c *Client) FindSlideTag(criteria *Criteria) (*SlideTag, error) {
	sts := &SlideTags{}
	if err := c.SearchRead(SlideTagModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("slide.tag was not found")
}

// FindSlideTags finds slide.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideTags(criteria *Criteria, options *Options) (*SlideTags, error) {
	sts := &SlideTags{}
	if err := c.SearchRead(SlideTagModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSlideTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideTagId finds record id by querying it with criteria.
func (c *Client) FindSlideTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.tag was not found")
}
