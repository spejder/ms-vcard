package odoo

import (
	"fmt"
)

// SlideChannelPartner represents slide.channel.partner model.
type SlideChannelPartner struct {
	ChannelId    *Many2One `xmlrpc:"channel_id,omptempty"`
	Completed    *Bool     `xmlrpc:"completed,omptempty"`
	Completion   *Int      `xmlrpc:"completion,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	PartnerEmail *String   `xmlrpc:"partner_email,omptempty"`
	PartnerId    *Many2One `xmlrpc:"partner_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelPartners represents array of slide.channel.partner model.
type SlideChannelPartners []SlideChannelPartner

// SlideChannelPartnerModel is the odoo model name.
const SlideChannelPartnerModel = "slide.channel.partner"

// Many2One convert SlideChannelPartner to *Many2One.
func (scp *SlideChannelPartner) Many2One() *Many2One {
	return NewMany2One(scp.Id.Get(), "")
}

// CreateSlideChannelPartner creates a new slide.channel.partner model and returns its id.
func (c *Client) CreateSlideChannelPartner(scp *SlideChannelPartner) (int64, error) {
	return c.Create(SlideChannelPartnerModel, scp)
}

// UpdateSlideChannelPartner updates an existing slide.channel.partner record.
func (c *Client) UpdateSlideChannelPartner(scp *SlideChannelPartner) error {
	return c.UpdateSlideChannelPartners([]int64{scp.Id.Get()}, scp)
}

// UpdateSlideChannelPartners updates existing slide.channel.partner records.
// All records (represented by ids) will be updated by scp values.
func (c *Client) UpdateSlideChannelPartners(ids []int64, scp *SlideChannelPartner) error {
	return c.Update(SlideChannelPartnerModel, ids, scp)
}

// DeleteSlideChannelPartner deletes an existing slide.channel.partner record.
func (c *Client) DeleteSlideChannelPartner(id int64) error {
	return c.DeleteSlideChannelPartners([]int64{id})
}

// DeleteSlideChannelPartners deletes existing slide.channel.partner records.
func (c *Client) DeleteSlideChannelPartners(ids []int64) error {
	return c.Delete(SlideChannelPartnerModel, ids)
}

// GetSlideChannelPartner gets slide.channel.partner existing record.
func (c *Client) GetSlideChannelPartner(id int64) (*SlideChannelPartner, error) {
	scps, err := c.GetSlideChannelPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if scps != nil && len(*scps) > 0 {
		return &((*scps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.partner not found", id)
}

// GetSlideChannelPartners gets slide.channel.partner existing records.
func (c *Client) GetSlideChannelPartners(ids []int64) (*SlideChannelPartners, error) {
	scps := &SlideChannelPartners{}
	if err := c.Read(SlideChannelPartnerModel, ids, nil, scps); err != nil {
		return nil, err
	}
	return scps, nil
}

// FindSlideChannelPartner finds slide.channel.partner record by querying it with criteria.
func (c *Client) FindSlideChannelPartner(criteria *Criteria) (*SlideChannelPartner, error) {
	scps := &SlideChannelPartners{}
	if err := c.SearchRead(SlideChannelPartnerModel, criteria, NewOptions().Limit(1), scps); err != nil {
		return nil, err
	}
	if scps != nil && len(*scps) > 0 {
		return &((*scps)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.partner was not found")
}

// FindSlideChannelPartners finds slide.channel.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelPartners(criteria *Criteria, options *Options) (*SlideChannelPartners, error) {
	scps := &SlideChannelPartners{}
	if err := c.SearchRead(SlideChannelPartnerModel, criteria, options, scps); err != nil {
		return nil, err
	}
	return scps, nil
}

// FindSlideChannelPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelPartnerId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.partner was not found")
}
