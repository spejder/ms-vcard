package odoo

import (
	"fmt"
)

// SlideChannelPrices represents slide.channel_prices model.
type SlideChannelPrices struct {
	ChannelId   *Many2One `xmlrpc:"channel_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt   *Time     `xmlrpc:"deleted_at,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Price       *Float    `xmlrpc:"price,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	Title       *String   `xmlrpc:"title,omptempty"`
	Description *String   `xmlrpc:"description,omptempty"`
	CourseCode  *String   `xmlrpc:"course_code,omptempty"`
	Type        *String   `xmlrpc:"type,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelPricess represents array of slide.channel_prices model.
type SlideChannelPricess []SlideChannelPrices

// SlideChannelPricesModel is the odoo model name.
const SlideChannelPricesModel = "slide.channel_prices"

// Many2One convert SlideChannelPrices to *Many2One.
func (sc *SlideChannelPrices) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSlideChannelPrices creates a new slide.channel_prices model and returns its id.
func (c *Client) CreateSlideChannelPrices(sc *SlideChannelPrices) (int64, error) {
	return c.Create(SlideChannelPricesModel, sc)
}

// UpdateSlideChannelPrices updates an existing slide.channel_prices record.
func (c *Client) UpdateSlideChannelPrices(sc *SlideChannelPrices) error {
	return c.UpdateSlideChannelPricess([]int64{sc.Id.Get()}, sc)
}

// UpdateSlideChannelPricess updates existing slide.channel_prices records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSlideChannelPricess(ids []int64, sc *SlideChannelPrices) error {
	return c.Update(SlideChannelPricesModel, ids, sc)
}

// DeleteSlideChannelPrices deletes an existing slide.channel_prices record.
func (c *Client) DeleteSlideChannelPrices(id int64) error {
	return c.DeleteSlideChannelPricess([]int64{id})
}

// DeleteSlideChannelPricess deletes existing slide.channel_prices records.
func (c *Client) DeleteSlideChannelPricess(ids []int64) error {
	return c.Delete(SlideChannelPricesModel, ids)
}

// GetSlideChannelPrices gets slide.channel_prices existing record.
func (c *Client) GetSlideChannelPrices(id int64) (*SlideChannelPrices, error) {
	scs, err := c.GetSlideChannelPricess([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel_prices not found", id)
}

// GetSlideChannelPricess gets slide.channel_prices existing records.
func (c *Client) GetSlideChannelPricess(ids []int64) (*SlideChannelPricess, error) {
	scs := &SlideChannelPricess{}
	if err := c.Read(SlideChannelPricesModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelPrices finds slide.channel_prices record by querying it with criteria.
func (c *Client) FindSlideChannelPrices(criteria *Criteria) (*SlideChannelPrices, error) {
	scs := &SlideChannelPricess{}
	if err := c.SearchRead(SlideChannelPricesModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel_prices was not found")
}

// FindSlideChannelPricess finds slide.channel_prices records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelPricess(criteria *Criteria, options *Options) (*SlideChannelPricess, error) {
	scs := &SlideChannelPricess{}
	if err := c.SearchRead(SlideChannelPricesModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelPricesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelPricesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelPricesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelPricesId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelPricesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelPricesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel_prices was not found")
}
