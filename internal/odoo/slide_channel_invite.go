package odoo

import (
	"fmt"
)

// SlideChannelInvite represents slide.channel.invite model.
type SlideChannelInvite struct {
	AttachmentIds *Relation `xmlrpc:"attachment_ids,omptempty"`
	AuthorId      *Many2One `xmlrpc:"author_id,omptempty"`
	Body          *String   `xmlrpc:"body,omptempty"`
	ChannelId     *Many2One `xmlrpc:"channel_id,omptempty"`
	ChannelUrl    *String   `xmlrpc:"channel_url,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	EmailFrom     *String   `xmlrpc:"email_from,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	PartnerIds    *Relation `xmlrpc:"partner_ids,omptempty"`
	Subject       *String   `xmlrpc:"subject,omptempty"`
	TemplateId    *Many2One `xmlrpc:"template_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelInvites represents array of slide.channel.invite model.
type SlideChannelInvites []SlideChannelInvite

// SlideChannelInviteModel is the odoo model name.
const SlideChannelInviteModel = "slide.channel.invite"

// Many2One convert SlideChannelInvite to *Many2One.
func (sci *SlideChannelInvite) Many2One() *Many2One {
	return NewMany2One(sci.Id.Get(), "")
}

// CreateSlideChannelInvite creates a new slide.channel.invite model and returns its id.
func (c *Client) CreateSlideChannelInvite(sci *SlideChannelInvite) (int64, error) {
	return c.Create(SlideChannelInviteModel, sci)
}

// UpdateSlideChannelInvite updates an existing slide.channel.invite record.
func (c *Client) UpdateSlideChannelInvite(sci *SlideChannelInvite) error {
	return c.UpdateSlideChannelInvites([]int64{sci.Id.Get()}, sci)
}

// UpdateSlideChannelInvites updates existing slide.channel.invite records.
// All records (represented by ids) will be updated by sci values.
func (c *Client) UpdateSlideChannelInvites(ids []int64, sci *SlideChannelInvite) error {
	return c.Update(SlideChannelInviteModel, ids, sci)
}

// DeleteSlideChannelInvite deletes an existing slide.channel.invite record.
func (c *Client) DeleteSlideChannelInvite(id int64) error {
	return c.DeleteSlideChannelInvites([]int64{id})
}

// DeleteSlideChannelInvites deletes existing slide.channel.invite records.
func (c *Client) DeleteSlideChannelInvites(ids []int64) error {
	return c.Delete(SlideChannelInviteModel, ids)
}

// GetSlideChannelInvite gets slide.channel.invite existing record.
func (c *Client) GetSlideChannelInvite(id int64) (*SlideChannelInvite, error) {
	scis, err := c.GetSlideChannelInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	if scis != nil && len(*scis) > 0 {
		return &((*scis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel.invite not found", id)
}

// GetSlideChannelInvites gets slide.channel.invite existing records.
func (c *Client) GetSlideChannelInvites(ids []int64) (*SlideChannelInvites, error) {
	scis := &SlideChannelInvites{}
	if err := c.Read(SlideChannelInviteModel, ids, nil, scis); err != nil {
		return nil, err
	}
	return scis, nil
}

// FindSlideChannelInvite finds slide.channel.invite record by querying it with criteria.
func (c *Client) FindSlideChannelInvite(criteria *Criteria) (*SlideChannelInvite, error) {
	scis := &SlideChannelInvites{}
	if err := c.SearchRead(SlideChannelInviteModel, criteria, NewOptions().Limit(1), scis); err != nil {
		return nil, err
	}
	if scis != nil && len(*scis) > 0 {
		return &((*scis)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel.invite was not found")
}

// FindSlideChannelInvites finds slide.channel.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelInvites(criteria *Criteria, options *Options) (*SlideChannelInvites, error) {
	scis := &SlideChannelInvites{}
	if err := c.SearchRead(SlideChannelInviteModel, criteria, options, scis); err != nil {
		return nil, err
	}
	return scis, nil
}

// FindSlideChannelInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelInviteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelInviteId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel.invite was not found")
}
