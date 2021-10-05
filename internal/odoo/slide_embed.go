package odoo

import (
	"fmt"
)

// SlideEmbed represents slide.embed model.
type SlideEmbed struct {
	CountViews  *Int      `xmlrpc:"count_views,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	SlideId     *Many2One `xmlrpc:"slide_id,omptempty"`
	Url         *String   `xmlrpc:"url,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideEmbeds represents array of slide.embed model.
type SlideEmbeds []SlideEmbed

// SlideEmbedModel is the odoo model name.
const SlideEmbedModel = "slide.embed"

// Many2One convert SlideEmbed to *Many2One.
func (se *SlideEmbed) Many2One() *Many2One {
	return NewMany2One(se.Id.Get(), "")
}

// CreateSlideEmbed creates a new slide.embed model and returns its id.
func (c *Client) CreateSlideEmbed(se *SlideEmbed) (int64, error) {
	return c.Create(SlideEmbedModel, se)
}

// UpdateSlideEmbed updates an existing slide.embed record.
func (c *Client) UpdateSlideEmbed(se *SlideEmbed) error {
	return c.UpdateSlideEmbeds([]int64{se.Id.Get()}, se)
}

// UpdateSlideEmbeds updates existing slide.embed records.
// All records (represented by ids) will be updated by se values.
func (c *Client) UpdateSlideEmbeds(ids []int64, se *SlideEmbed) error {
	return c.Update(SlideEmbedModel, ids, se)
}

// DeleteSlideEmbed deletes an existing slide.embed record.
func (c *Client) DeleteSlideEmbed(id int64) error {
	return c.DeleteSlideEmbeds([]int64{id})
}

// DeleteSlideEmbeds deletes existing slide.embed records.
func (c *Client) DeleteSlideEmbeds(ids []int64) error {
	return c.Delete(SlideEmbedModel, ids)
}

// GetSlideEmbed gets slide.embed existing record.
func (c *Client) GetSlideEmbed(id int64) (*SlideEmbed, error) {
	ses, err := c.GetSlideEmbeds([]int64{id})
	if err != nil {
		return nil, err
	}
	if ses != nil && len(*ses) > 0 {
		return &((*ses)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.embed not found", id)
}

// GetSlideEmbeds gets slide.embed existing records.
func (c *Client) GetSlideEmbeds(ids []int64) (*SlideEmbeds, error) {
	ses := &SlideEmbeds{}
	if err := c.Read(SlideEmbedModel, ids, nil, ses); err != nil {
		return nil, err
	}
	return ses, nil
}

// FindSlideEmbed finds slide.embed record by querying it with criteria.
func (c *Client) FindSlideEmbed(criteria *Criteria) (*SlideEmbed, error) {
	ses := &SlideEmbeds{}
	if err := c.SearchRead(SlideEmbedModel, criteria, NewOptions().Limit(1), ses); err != nil {
		return nil, err
	}
	if ses != nil && len(*ses) > 0 {
		return &((*ses)[0]), nil
	}
	return nil, fmt.Errorf("slide.embed was not found")
}

// FindSlideEmbeds finds slide.embed records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideEmbeds(criteria *Criteria, options *Options) (*SlideEmbeds, error) {
	ses := &SlideEmbeds{}
	if err := c.SearchRead(SlideEmbedModel, criteria, options, ses); err != nil {
		return nil, err
	}
	return ses, nil
}

// FindSlideEmbedIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideEmbedIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideEmbedModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideEmbedId finds record id by querying it with criteria.
func (c *Client) FindSlideEmbedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideEmbedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.embed was not found")
}
