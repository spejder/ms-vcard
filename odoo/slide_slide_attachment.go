package odoo

import (
	"fmt"
)

// SlideSlideAttachment represents slide.slide_attachment model.
type SlideSlideAttachment struct {
	Attachment  *String   `xmlrpc:"attachment,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt   *Time     `xmlrpc:"deleted_at,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	SlideId     *Many2One `xmlrpc:"slide_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideSlideAttachments represents array of slide.slide_attachment model.
type SlideSlideAttachments []SlideSlideAttachment

// SlideSlideAttachmentModel is the odoo model name.
const SlideSlideAttachmentModel = "slide.slide_attachment"

// Many2One convert SlideSlideAttachment to *Many2One.
func (ss *SlideSlideAttachment) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSlideSlideAttachment creates a new slide.slide_attachment model and returns its id.
func (c *Client) CreateSlideSlideAttachment(ss *SlideSlideAttachment) (int64, error) {
	return c.Create(SlideSlideAttachmentModel, ss)
}

// UpdateSlideSlideAttachment updates an existing slide.slide_attachment record.
func (c *Client) UpdateSlideSlideAttachment(ss *SlideSlideAttachment) error {
	return c.UpdateSlideSlideAttachments([]int64{ss.Id.Get()}, ss)
}

// UpdateSlideSlideAttachments updates existing slide.slide_attachment records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSlideSlideAttachments(ids []int64, ss *SlideSlideAttachment) error {
	return c.Update(SlideSlideAttachmentModel, ids, ss)
}

// DeleteSlideSlideAttachment deletes an existing slide.slide_attachment record.
func (c *Client) DeleteSlideSlideAttachment(id int64) error {
	return c.DeleteSlideSlideAttachments([]int64{id})
}

// DeleteSlideSlideAttachments deletes existing slide.slide_attachment records.
func (c *Client) DeleteSlideSlideAttachments(ids []int64) error {
	return c.Delete(SlideSlideAttachmentModel, ids)
}

// GetSlideSlideAttachment gets slide.slide_attachment existing record.
func (c *Client) GetSlideSlideAttachment(id int64) (*SlideSlideAttachment, error) {
	sss, err := c.GetSlideSlideAttachments([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide_attachment not found", id)
}

// GetSlideSlideAttachments gets slide.slide_attachment existing records.
func (c *Client) GetSlideSlideAttachments(ids []int64) (*SlideSlideAttachments, error) {
	sss := &SlideSlideAttachments{}
	if err := c.Read(SlideSlideAttachmentModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlideAttachment finds slide.slide_attachment record by querying it with criteria.
func (c *Client) FindSlideSlideAttachment(criteria *Criteria) (*SlideSlideAttachment, error) {
	sss := &SlideSlideAttachments{}
	if err := c.SearchRead(SlideSlideAttachmentModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide_attachment was not found")
}

// FindSlideSlideAttachments finds slide.slide_attachment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideAttachments(criteria *Criteria, options *Options) (*SlideSlideAttachments, error) {
	sss := &SlideSlideAttachments{}
	if err := c.SearchRead(SlideSlideAttachmentModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlideAttachmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideAttachmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideAttachmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlideAttachmentId finds record id by querying it with criteria.
func (c *Client) FindSlideSlideAttachmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideAttachmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide_attachment was not found")
}
