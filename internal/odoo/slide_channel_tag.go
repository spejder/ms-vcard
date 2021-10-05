package odoo

import (
	"fmt"
)

// SlideChannelTag represents slide.channel.tag model.
type SlideChannelTag struct {
	ChannelIds    *Relation `xmlrpc:"channel_ids,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	GroupId       *Many2One `xmlrpc:"group_id,omptempty"`
	GroupSequence *Int      `xmlrpc:"group_sequence,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelTags represents array of slide.channel.tag model.
type SlideChannelTags []SlideChannelTag

// SlideChannelTagModel is the odoo model name.
const SlideChannelTagModel = "slide.channel.tag"

// Many2One convert SlideChannelTag to *Many2One.
func (sct *SlideChannelTag) Many2One() *Many2One {
	return NewMany2One(sct.Id.Get(), "")
}

// CreateSlideChannelTag creates a new slide.channel.tag model and returns its id.
func (c *Client) CreateSlideChannelTag(sct *SlideChannelTag) (int64, error) {
	return c.Create(SlideChannelTagModel, sct)
}

// UpdateSlideChannelTag updates an existing slide.channel.tag record.
func (c *Client) UpdateSlideChannelTag(sct *SlideChannelTag) error {
	return c.UpdateSlideChannelTags([]int64{sct.Id.Get()}, sct)
}

// UpdateSlideChannelTags updates existing slide.channel.tag records.
// All records (represented by ids) will be updated by sct values.
func (c *Client) UpdateSlideChannelTags(ids []int64, sct *SlideChannelTag) error {
	return c.Update(SlideChannelTagModel, ids, sct)
}

// DeleteSlideChannelTag deletes an existing slide.channel.tag record.
func (c *Client) DeleteSlideChannelTag(id int64) error {
	return c.DeleteSlideChannelTags([]int64{id})
}

// DeleteSlideChannelTags deletes existing slide.channel.tag records.
func (c *Client) DeleteSlideChannelTags(ids []int64) error {
	return c.Delete(SlideChannelTagModel, ids)
}

// GetSlideChannelTag gets slide.channel.tag existing record.
func (c *Client) GetSlideChannelTag(id int64) (*SlideChannelTag, error) {
	scts, err := c.GetSlideChannelTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if scts != nil && len(*scts) > 0 {
		return &((*scts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.tag not found", id)
}

// GetSlideChannelTags gets slide.channel.tag existing records.
func (c *Client) GetSlideChannelTags(ids []int64) (*SlideChannelTags, error) {
	scts := &SlideChannelTags{}
	if err := c.Read(SlideChannelTagModel, ids, nil, scts); err != nil {
		return nil, err
	}
	return scts, nil
}

// FindSlideChannelTag finds slide.channel.tag record by querying it with criteria.
func (c *Client) FindSlideChannelTag(criteria *Criteria) (*SlideChannelTag, error) {
	scts := &SlideChannelTags{}
	if err := c.SearchRead(SlideChannelTagModel, criteria, NewOptions().Limit(1), scts); err != nil {
		return nil, err
	}
	if scts != nil && len(*scts) > 0 {
		return &((*scts)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.tag was not found")
}

// FindSlideChannelTags finds slide.channel.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelTags(criteria *Criteria, options *Options) (*SlideChannelTags, error) {
	scts := &SlideChannelTags{}
	if err := c.SearchRead(SlideChannelTagModel, criteria, options, scts); err != nil {
		return nil, err
	}
	return scts, nil
}

// FindSlideChannelTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelTagId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.tag was not found")
}
