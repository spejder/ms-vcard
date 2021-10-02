package odoo

import (
	"fmt"
)

// WebsitePublishedMixin represents website.published.mixin model.
type WebsitePublishedMixin struct {
	CanPublish       *Bool   `xmlrpc:"can_publish,omptempty"`
	DisplayName      *String `xmlrpc:"display_name,omptempty"`
	Id               *Int    `xmlrpc:"id,omptempty"`
	IsPublished      *Bool   `xmlrpc:"is_published,omptempty"`
	LastUpdate       *Time   `xmlrpc:"__last_update,omptempty"`
	WebsitePublished *Bool   `xmlrpc:"website_published,omptempty"`
	WebsiteUrl       *String `xmlrpc:"website_url,omptempty"`
}

// WebsitePublishedMixins represents array of website.published.mixin model.
type WebsitePublishedMixins []WebsitePublishedMixin

// WebsitePublishedMixinModel is the odoo model name.
const WebsitePublishedMixinModel = "website.published.mixin"

// Many2One convert WebsitePublishedMixin to *Many2One.
func (wpm *WebsitePublishedMixin) Many2One() *Many2One {
	return NewMany2One(wpm.Id.Get(), "")
}

// CreateWebsitePublishedMixin creates a new website.published.mixin model and returns its id.
func (c *Client) CreateWebsitePublishedMixin(wpm *WebsitePublishedMixin) (int64, error) {
	return c.Create(WebsitePublishedMixinModel, wpm)
}

// UpdateWebsitePublishedMixin updates an existing website.published.mixin record.
func (c *Client) UpdateWebsitePublishedMixin(wpm *WebsitePublishedMixin) error {
	return c.UpdateWebsitePublishedMixins([]int64{wpm.Id.Get()}, wpm)
}

// UpdateWebsitePublishedMixins updates existing website.published.mixin records.
// All records (represented by ids) will be updated by wpm values.
func (c *Client) UpdateWebsitePublishedMixins(ids []int64, wpm *WebsitePublishedMixin) error {
	return c.Update(WebsitePublishedMixinModel, ids, wpm)
}

// DeleteWebsitePublishedMixin deletes an existing website.published.mixin record.
func (c *Client) DeleteWebsitePublishedMixin(id int64) error {
	return c.DeleteWebsitePublishedMixins([]int64{id})
}

// DeleteWebsitePublishedMixins deletes existing website.published.mixin records.
func (c *Client) DeleteWebsitePublishedMixins(ids []int64) error {
	return c.Delete(WebsitePublishedMixinModel, ids)
}

// GetWebsitePublishedMixin gets website.published.mixin existing record.
func (c *Client) GetWebsitePublishedMixin(id int64) (*WebsitePublishedMixin, error) {
	wpms, err := c.GetWebsitePublishedMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if wpms != nil && len(*wpms) > 0 {
		return &((*wpms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.published.mixin not found", id)
}

// GetWebsitePublishedMixins gets website.published.mixin existing records.
func (c *Client) GetWebsitePublishedMixins(ids []int64) (*WebsitePublishedMixins, error) {
	wpms := &WebsitePublishedMixins{}
	if err := c.Read(WebsitePublishedMixinModel, ids, nil, wpms); err != nil {
		return nil, err
	}
	return wpms, nil
}

// FindWebsitePublishedMixin finds website.published.mixin record by querying it with criteria.
func (c *Client) FindWebsitePublishedMixin(criteria *Criteria) (*WebsitePublishedMixin, error) {
	wpms := &WebsitePublishedMixins{}
	if err := c.SearchRead(WebsitePublishedMixinModel, criteria, NewOptions().Limit(1), wpms); err != nil {
		return nil, err
	}
	if wpms != nil && len(*wpms) > 0 {
		return &((*wpms)[0]), nil
	}
	return nil, fmt.Errorf("website.published.mixin was not found")
}

// FindWebsitePublishedMixins finds website.published.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePublishedMixins(criteria *Criteria, options *Options) (*WebsitePublishedMixins, error) {
	wpms := &WebsitePublishedMixins{}
	if err := c.SearchRead(WebsitePublishedMixinModel, criteria, options, wpms); err != nil {
		return nil, err
	}
	return wpms, nil
}

// FindWebsitePublishedMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePublishedMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsitePublishedMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsitePublishedMixinId finds record id by querying it with criteria.
func (c *Client) FindWebsitePublishedMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePublishedMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.published.mixin was not found")
}
