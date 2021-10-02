package odoo

import (
	"fmt"
)

// SlideSlideLink represents slide.slide.link model.
type SlideSlideLink struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Link        *String   `xmlrpc:"link,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	SlideId     *Many2One `xmlrpc:"slide_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideSlideLinks represents array of slide.slide.link model.
type SlideSlideLinks []SlideSlideLink

// SlideSlideLinkModel is the odoo model name.
const SlideSlideLinkModel = "slide.slide.link"

// Many2One convert SlideSlideLink to *Many2One.
func (ssl *SlideSlideLink) Many2One() *Many2One {
	return NewMany2One(ssl.Id.Get(), "")
}

// CreateSlideSlideLink creates a new slide.slide.link model and returns its id.
func (c *Client) CreateSlideSlideLink(ssl *SlideSlideLink) (int64, error) {
	return c.Create(SlideSlideLinkModel, ssl)
}

// UpdateSlideSlideLink updates an existing slide.slide.link record.
func (c *Client) UpdateSlideSlideLink(ssl *SlideSlideLink) error {
	return c.UpdateSlideSlideLinks([]int64{ssl.Id.Get()}, ssl)
}

// UpdateSlideSlideLinks updates existing slide.slide.link records.
// All records (represented by ids) will be updated by ssl values.
func (c *Client) UpdateSlideSlideLinks(ids []int64, ssl *SlideSlideLink) error {
	return c.Update(SlideSlideLinkModel, ids, ssl)
}

// DeleteSlideSlideLink deletes an existing slide.slide.link record.
func (c *Client) DeleteSlideSlideLink(id int64) error {
	return c.DeleteSlideSlideLinks([]int64{id})
}

// DeleteSlideSlideLinks deletes existing slide.slide.link records.
func (c *Client) DeleteSlideSlideLinks(ids []int64) error {
	return c.Delete(SlideSlideLinkModel, ids)
}

// GetSlideSlideLink gets slide.slide.link existing record.
func (c *Client) GetSlideSlideLink(id int64) (*SlideSlideLink, error) {
	ssls, err := c.GetSlideSlideLinks([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssls != nil && len(*ssls) > 0 {
		return &((*ssls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide.link not found", id)
}

// GetSlideSlideLinks gets slide.slide.link existing records.
func (c *Client) GetSlideSlideLinks(ids []int64) (*SlideSlideLinks, error) {
	ssls := &SlideSlideLinks{}
	if err := c.Read(SlideSlideLinkModel, ids, nil, ssls); err != nil {
		return nil, err
	}
	return ssls, nil
}

// FindSlideSlideLink finds slide.slide.link record by querying it with criteria.
func (c *Client) FindSlideSlideLink(criteria *Criteria) (*SlideSlideLink, error) {
	ssls := &SlideSlideLinks{}
	if err := c.SearchRead(SlideSlideLinkModel, criteria, NewOptions().Limit(1), ssls); err != nil {
		return nil, err
	}
	if ssls != nil && len(*ssls) > 0 {
		return &((*ssls)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide.link was not found")
}

// FindSlideSlideLinks finds slide.slide.link records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideLinks(criteria *Criteria, options *Options) (*SlideSlideLinks, error) {
	ssls := &SlideSlideLinks{}
	if err := c.SearchRead(SlideSlideLinkModel, criteria, options, ssls); err != nil {
		return nil, err
	}
	return ssls, nil
}

// FindSlideSlideLinkIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideLinkIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideLinkModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlideLinkId finds record id by querying it with criteria.
func (c *Client) FindSlideSlideLinkId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideLinkModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide.link was not found")
}
