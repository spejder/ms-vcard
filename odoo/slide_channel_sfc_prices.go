package odoo

import (
	"fmt"
)

// SlideChannelSfcPrices represents slide.channel_sfc_prices model.
type SlideChannelSfcPrices struct {
	ChannelId    *Many2One `xmlrpc:"channel_id,omptempty"`
	CourseCode   *String   `xmlrpc:"course_code,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt    *Time     `xmlrpc:"deleted_at,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	Price        *Float    `xmlrpc:"price,omptempty"`
	Sequence     *Int      `xmlrpc:"sequence,omptempty"`
	Title        *String   `xmlrpc:"title,omptempty"`
	Type         *String   `xmlrpc:"type,omptempty"`
	PriceOptions *String   `xmlrpc:"price_options,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelSfcPricess represents array of slide.channel_sfc_prices model.
type SlideChannelSfcPricess []SlideChannelSfcPrices

// SlideChannelSfcPricesModel is the odoo model name.
const SlideChannelSfcPricesModel = "slide.channel_sfc_prices"

// Many2One convert SlideChannelSfcPrices to *Many2One.
func (sc *SlideChannelSfcPrices) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSlideChannelSfcPrices creates a new slide.channel_sfc_prices model and returns its id.
func (c *Client) CreateSlideChannelSfcPrices(sc *SlideChannelSfcPrices) (int64, error) {
	return c.Create(SlideChannelSfcPricesModel, sc)
}

// UpdateSlideChannelSfcPrices updates an existing slide.channel_sfc_prices record.
func (c *Client) UpdateSlideChannelSfcPrices(sc *SlideChannelSfcPrices) error {
	return c.UpdateSlideChannelSfcPricess([]int64{sc.Id.Get()}, sc)
}

// UpdateSlideChannelSfcPricess updates existing slide.channel_sfc_prices records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSlideChannelSfcPricess(ids []int64, sc *SlideChannelSfcPrices) error {
	return c.Update(SlideChannelSfcPricesModel, ids, sc)
}

// DeleteSlideChannelSfcPrices deletes an existing slide.channel_sfc_prices record.
func (c *Client) DeleteSlideChannelSfcPrices(id int64) error {
	return c.DeleteSlideChannelSfcPricess([]int64{id})
}

// DeleteSlideChannelSfcPricess deletes existing slide.channel_sfc_prices records.
func (c *Client) DeleteSlideChannelSfcPricess(ids []int64) error {
	return c.Delete(SlideChannelSfcPricesModel, ids)
}

// GetSlideChannelSfcPrices gets slide.channel_sfc_prices existing record.
func (c *Client) GetSlideChannelSfcPrices(id int64) (*SlideChannelSfcPrices, error) {
	scs, err := c.GetSlideChannelSfcPricess([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel_sfc_prices not found", id)
}

// GetSlideChannelSfcPricess gets slide.channel_sfc_prices existing records.
func (c *Client) GetSlideChannelSfcPricess(ids []int64) (*SlideChannelSfcPricess, error) {
	scs := &SlideChannelSfcPricess{}
	if err := c.Read(SlideChannelSfcPricesModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelSfcPrices finds slide.channel_sfc_prices record by querying it with criteria.
func (c *Client) FindSlideChannelSfcPrices(criteria *Criteria) (*SlideChannelSfcPrices, error) {
	scs := &SlideChannelSfcPricess{}
	if err := c.SearchRead(SlideChannelSfcPricesModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel_sfc_prices was not found")
}

// FindSlideChannelSfcPricess finds slide.channel_sfc_prices records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelSfcPricess(criteria *Criteria, options *Options) (*SlideChannelSfcPricess, error) {
	scs := &SlideChannelSfcPricess{}
	if err := c.SearchRead(SlideChannelSfcPricesModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelSfcPricesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelSfcPricesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelSfcPricesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelSfcPricesId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelSfcPricesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelSfcPricesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel_sfc_prices was not found")
}
