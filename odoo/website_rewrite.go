package odoo

import (
	"fmt"
)

// WebsiteRewrite represents website.rewrite model.
type WebsiteRewrite struct {
	Active       *Bool      `xmlrpc:"active,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	Name         *String    `xmlrpc:"name,omptempty"`
	RedirectType *Selection `xmlrpc:"redirect_type,omptempty"`
	RouteId      *Many2One  `xmlrpc:"route_id,omptempty"`
	Sequence     *Int       `xmlrpc:"sequence,omptempty"`
	UrlFrom      *String    `xmlrpc:"url_from,omptempty"`
	UrlTo        *String    `xmlrpc:"url_to,omptempty"`
	WebsiteId    *Many2One  `xmlrpc:"website_id,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WebsiteRewrites represents array of website.rewrite model.
type WebsiteRewrites []WebsiteRewrite

// WebsiteRewriteModel is the odoo model name.
const WebsiteRewriteModel = "website.rewrite"

// Many2One convert WebsiteRewrite to *Many2One.
func (wr *WebsiteRewrite) Many2One() *Many2One {
	return NewMany2One(wr.Id.Get(), "")
}

// CreateWebsiteRewrite creates a new website.rewrite model and returns its id.
func (c *Client) CreateWebsiteRewrite(wr *WebsiteRewrite) (int64, error) {
	return c.Create(WebsiteRewriteModel, wr)
}

// UpdateWebsiteRewrite updates an existing website.rewrite record.
func (c *Client) UpdateWebsiteRewrite(wr *WebsiteRewrite) error {
	return c.UpdateWebsiteRewrites([]int64{wr.Id.Get()}, wr)
}

// UpdateWebsiteRewrites updates existing website.rewrite records.
// All records (represented by ids) will be updated by wr values.
func (c *Client) UpdateWebsiteRewrites(ids []int64, wr *WebsiteRewrite) error {
	return c.Update(WebsiteRewriteModel, ids, wr)
}

// DeleteWebsiteRewrite deletes an existing website.rewrite record.
func (c *Client) DeleteWebsiteRewrite(id int64) error {
	return c.DeleteWebsiteRewrites([]int64{id})
}

// DeleteWebsiteRewrites deletes existing website.rewrite records.
func (c *Client) DeleteWebsiteRewrites(ids []int64) error {
	return c.Delete(WebsiteRewriteModel, ids)
}

// GetWebsiteRewrite gets website.rewrite existing record.
func (c *Client) GetWebsiteRewrite(id int64) (*WebsiteRewrite, error) {
	wrs, err := c.GetWebsiteRewrites([]int64{id})
	if err != nil {
		return nil, err
	}
	if wrs != nil && len(*wrs) > 0 {
		return &((*wrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.rewrite not found", id)
}

// GetWebsiteRewrites gets website.rewrite existing records.
func (c *Client) GetWebsiteRewrites(ids []int64) (*WebsiteRewrites, error) {
	wrs := &WebsiteRewrites{}
	if err := c.Read(WebsiteRewriteModel, ids, nil, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRewrite finds website.rewrite record by querying it with criteria.
func (c *Client) FindWebsiteRewrite(criteria *Criteria) (*WebsiteRewrite, error) {
	wrs := &WebsiteRewrites{}
	if err := c.SearchRead(WebsiteRewriteModel, criteria, NewOptions().Limit(1), wrs); err != nil {
		return nil, err
	}
	if wrs != nil && len(*wrs) > 0 {
		return &((*wrs)[0]), nil
	}
	return nil, fmt.Errorf("website.rewrite was not found")
}

// FindWebsiteRewrites finds website.rewrite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRewrites(criteria *Criteria, options *Options) (*WebsiteRewrites, error) {
	wrs := &WebsiteRewrites{}
	if err := c.SearchRead(WebsiteRewriteModel, criteria, options, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRewriteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRewriteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteRewriteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteRewriteId finds record id by querying it with criteria.
func (c *Client) FindWebsiteRewriteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteRewriteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.rewrite was not found")
}
