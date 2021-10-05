package odoo

import (
	"fmt"
)

// SlideSlidePartner represents slide.slide.partner model.
type SlideSlidePartner struct {
	ChannelId         *Many2One `xmlrpc:"channel_id,omptempty"`
	Completed         *Bool     `xmlrpc:"completed,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omptempty"`
	QuizAttemptsCount *Int      `xmlrpc:"quiz_attempts_count,omptempty"`
	SlideId           *Many2One `xmlrpc:"slide_id,omptempty"`
	Vote              *Int      `xmlrpc:"vote,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideSlidePartners represents array of slide.slide.partner model.
type SlideSlidePartners []SlideSlidePartner

// SlideSlidePartnerModel is the odoo model name.
const SlideSlidePartnerModel = "slide.slide.partner"

// Many2One convert SlideSlidePartner to *Many2One.
func (ssp *SlideSlidePartner) Many2One() *Many2One {
	return NewMany2One(ssp.Id.Get(), "")
}

// CreateSlideSlidePartner creates a new slide.slide.partner model and returns its id.
func (c *Client) CreateSlideSlidePartner(ssp *SlideSlidePartner) (int64, error) {
	return c.Create(SlideSlidePartnerModel, ssp)
}

// UpdateSlideSlidePartner updates an existing slide.slide.partner record.
func (c *Client) UpdateSlideSlidePartner(ssp *SlideSlidePartner) error {
	return c.UpdateSlideSlidePartners([]int64{ssp.Id.Get()}, ssp)
}

// UpdateSlideSlidePartners updates existing slide.slide.partner records.
// All records (represented by ids) will be updated by ssp values.
func (c *Client) UpdateSlideSlidePartners(ids []int64, ssp *SlideSlidePartner) error {
	return c.Update(SlideSlidePartnerModel, ids, ssp)
}

// DeleteSlideSlidePartner deletes an existing slide.slide.partner record.
func (c *Client) DeleteSlideSlidePartner(id int64) error {
	return c.DeleteSlideSlidePartners([]int64{id})
}

// DeleteSlideSlidePartners deletes existing slide.slide.partner records.
func (c *Client) DeleteSlideSlidePartners(ids []int64) error {
	return c.Delete(SlideSlidePartnerModel, ids)
}

// GetSlideSlidePartner gets slide.slide.partner existing record.
func (c *Client) GetSlideSlidePartner(id int64) (*SlideSlidePartner, error) {
	ssps, err := c.GetSlideSlidePartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssps != nil && len(*ssps) > 0 {
		return &((*ssps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide.partner not found", id)
}

// GetSlideSlidePartners gets slide.slide.partner existing records.
func (c *Client) GetSlideSlidePartners(ids []int64) (*SlideSlidePartners, error) {
	ssps := &SlideSlidePartners{}
	if err := c.Read(SlideSlidePartnerModel, ids, nil, ssps); err != nil {
		return nil, err
	}
	return ssps, nil
}

// FindSlideSlidePartner finds slide.slide.partner record by querying it with criteria.
func (c *Client) FindSlideSlidePartner(criteria *Criteria) (*SlideSlidePartner, error) {
	ssps := &SlideSlidePartners{}
	if err := c.SearchRead(SlideSlidePartnerModel, criteria, NewOptions().Limit(1), ssps); err != nil {
		return nil, err
	}
	if ssps != nil && len(*ssps) > 0 {
		return &((*ssps)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide.partner was not found")
}

// FindSlideSlidePartners finds slide.slide.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlidePartners(criteria *Criteria, options *Options) (*SlideSlidePartners, error) {
	ssps := &SlideSlidePartners{}
	if err := c.SearchRead(SlideSlidePartnerModel, criteria, options, ssps); err != nil {
		return nil, err
	}
	return ssps, nil
}

// FindSlideSlidePartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlidePartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlidePartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlidePartnerId finds record id by querying it with criteria.
func (c *Client) FindSlideSlidePartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlidePartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide.partner was not found")
}
