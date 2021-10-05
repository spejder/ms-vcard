package odoo

import (
	"fmt"
)

// WebsiteMultiMixin represents website.multi.mixin model.
type WebsiteMultiMixin struct {
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	WebsiteId   *Many2One `xmlrpc:"website_id,omptempty"`
}

// WebsiteMultiMixins represents array of website.multi.mixin model.
type WebsiteMultiMixins []WebsiteMultiMixin

// WebsiteMultiMixinModel is the odoo model name.
const WebsiteMultiMixinModel = "website.multi.mixin"

// Many2One convert WebsiteMultiMixin to *Many2One.
func (wmm *WebsiteMultiMixin) Many2One() *Many2One {
	return NewMany2One(wmm.Id.Get(), "")
}

// CreateWebsiteMultiMixin creates a new website.multi.mixin model and returns its id.
func (c *Client) CreateWebsiteMultiMixin(wmm *WebsiteMultiMixin) (int64, error) {
	return c.Create(WebsiteMultiMixinModel, wmm)
}

// UpdateWebsiteMultiMixin updates an existing website.multi.mixin record.
func (c *Client) UpdateWebsiteMultiMixin(wmm *WebsiteMultiMixin) error {
	return c.UpdateWebsiteMultiMixins([]int64{wmm.Id.Get()}, wmm)
}

// UpdateWebsiteMultiMixins updates existing website.multi.mixin records.
// All records (represented by ids) will be updated by wmm values.
func (c *Client) UpdateWebsiteMultiMixins(ids []int64, wmm *WebsiteMultiMixin) error {
	return c.Update(WebsiteMultiMixinModel, ids, wmm)
}

// DeleteWebsiteMultiMixin deletes an existing website.multi.mixin record.
func (c *Client) DeleteWebsiteMultiMixin(id int64) error {
	return c.DeleteWebsiteMultiMixins([]int64{id})
}

// DeleteWebsiteMultiMixins deletes existing website.multi.mixin records.
func (c *Client) DeleteWebsiteMultiMixins(ids []int64) error {
	return c.Delete(WebsiteMultiMixinModel, ids)
}

// GetWebsiteMultiMixin gets website.multi.mixin existing record.
func (c *Client) GetWebsiteMultiMixin(id int64) (*WebsiteMultiMixin, error) {
	wmms, err := c.GetWebsiteMultiMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if wmms != nil && len(*wmms) > 0 {
		return &((*wmms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.multi.mixin not found", id)
}

// GetWebsiteMultiMixins gets website.multi.mixin existing records.
func (c *Client) GetWebsiteMultiMixins(ids []int64) (*WebsiteMultiMixins, error) {
	wmms := &WebsiteMultiMixins{}
	if err := c.Read(WebsiteMultiMixinModel, ids, nil, wmms); err != nil {
		return nil, err
	}
	return wmms, nil
}

// FindWebsiteMultiMixin finds website.multi.mixin record by querying it with criteria.
func (c *Client) FindWebsiteMultiMixin(criteria *Criteria) (*WebsiteMultiMixin, error) {
	wmms := &WebsiteMultiMixins{}
	if err := c.SearchRead(WebsiteMultiMixinModel, criteria, NewOptions().Limit(1), wmms); err != nil {
		return nil, err
	}
	if wmms != nil && len(*wmms) > 0 {
		return &((*wmms)[0]), nil
	}
	return nil, fmt.Errorf("website.multi.mixin was not found")
}

// FindWebsiteMultiMixins finds website.multi.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMultiMixins(criteria *Criteria, options *Options) (*WebsiteMultiMixins, error) {
	wmms := &WebsiteMultiMixins{}
	if err := c.SearchRead(WebsiteMultiMixinModel, criteria, options, wmms); err != nil {
		return nil, err
	}
	return wmms, nil
}

// FindWebsiteMultiMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMultiMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteMultiMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteMultiMixinId finds record id by querying it with criteria.
func (c *Client) FindWebsiteMultiMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteMultiMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.multi.mixin was not found")
}
