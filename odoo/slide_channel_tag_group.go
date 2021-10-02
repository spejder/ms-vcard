package odoo

import (
	"fmt"
)

// SlideChannelTagGroup represents slide.channel.tag.group model.
type SlideChannelTagGroup struct {
	CanPublish       *Bool     `xmlrpc:"can_publish,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	IsPublished      *Bool     `xmlrpc:"is_published,omptempty"`
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	Sequence         *Int      `xmlrpc:"sequence,omptempty"`
	TagIds           *Relation `xmlrpc:"tag_ids,omptempty"`
	WebsitePublished *Bool     `xmlrpc:"website_published,omptempty"`
	WebsiteUrl       *String   `xmlrpc:"website_url,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelTagGroups represents array of slide.channel.tag.group model.
type SlideChannelTagGroups []SlideChannelTagGroup

// SlideChannelTagGroupModel is the odoo model name.
const SlideChannelTagGroupModel = "slide.channel.tag.group"

// Many2One convert SlideChannelTagGroup to *Many2One.
func (sctg *SlideChannelTagGroup) Many2One() *Many2One {
	return NewMany2One(sctg.Id.Get(), "")
}

// CreateSlideChannelTagGroup creates a new slide.channel.tag.group model and returns its id.
func (c *Client) CreateSlideChannelTagGroup(sctg *SlideChannelTagGroup) (int64, error) {
	return c.Create(SlideChannelTagGroupModel, sctg)
}

// UpdateSlideChannelTagGroup updates an existing slide.channel.tag.group record.
func (c *Client) UpdateSlideChannelTagGroup(sctg *SlideChannelTagGroup) error {
	return c.UpdateSlideChannelTagGroups([]int64{sctg.Id.Get()}, sctg)
}

// UpdateSlideChannelTagGroups updates existing slide.channel.tag.group records.
// All records (represented by ids) will be updated by sctg values.
func (c *Client) UpdateSlideChannelTagGroups(ids []int64, sctg *SlideChannelTagGroup) error {
	return c.Update(SlideChannelTagGroupModel, ids, sctg)
}

// DeleteSlideChannelTagGroup deletes an existing slide.channel.tag.group record.
func (c *Client) DeleteSlideChannelTagGroup(id int64) error {
	return c.DeleteSlideChannelTagGroups([]int64{id})
}

// DeleteSlideChannelTagGroups deletes existing slide.channel.tag.group records.
func (c *Client) DeleteSlideChannelTagGroups(ids []int64) error {
	return c.Delete(SlideChannelTagGroupModel, ids)
}

// GetSlideChannelTagGroup gets slide.channel.tag.group existing record.
func (c *Client) GetSlideChannelTagGroup(id int64) (*SlideChannelTagGroup, error) {
	sctgs, err := c.GetSlideChannelTagGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if sctgs != nil && len(*sctgs) > 0 {
		return &((*sctgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.tag.group not found", id)
}

// GetSlideChannelTagGroups gets slide.channel.tag.group existing records.
func (c *Client) GetSlideChannelTagGroups(ids []int64) (*SlideChannelTagGroups, error) {
	sctgs := &SlideChannelTagGroups{}
	if err := c.Read(SlideChannelTagGroupModel, ids, nil, sctgs); err != nil {
		return nil, err
	}
	return sctgs, nil
}

// FindSlideChannelTagGroup finds slide.channel.tag.group record by querying it with criteria.
func (c *Client) FindSlideChannelTagGroup(criteria *Criteria) (*SlideChannelTagGroup, error) {
	sctgs := &SlideChannelTagGroups{}
	if err := c.SearchRead(SlideChannelTagGroupModel, criteria, NewOptions().Limit(1), sctgs); err != nil {
		return nil, err
	}
	if sctgs != nil && len(*sctgs) > 0 {
		return &((*sctgs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.tag.group was not found")
}

// FindSlideChannelTagGroups finds slide.channel.tag.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelTagGroups(criteria *Criteria, options *Options) (*SlideChannelTagGroups, error) {
	sctgs := &SlideChannelTagGroups{}
	if err := c.SearchRead(SlideChannelTagGroupModel, criteria, options, sctgs); err != nil {
		return nil, err
	}
	return sctgs, nil
}

// FindSlideChannelTagGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelTagGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelTagGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelTagGroupId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelTagGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelTagGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.tag.group was not found")
}
