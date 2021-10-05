package odoo

import (
	"fmt"
)

// WebsiteSeoMetadata represents website.seo.metadata model.
type WebsiteSeoMetadata struct {
	DisplayName            *String `xmlrpc:"display_name,omptempty"`
	Id                     *Int    `xmlrpc:"id,omptempty"`
	IsSeoOptimized         *Bool   `xmlrpc:"is_seo_optimized,omptempty"`
	LastUpdate             *Time   `xmlrpc:"__last_update,omptempty"`
	WebsiteMetaDescription *String `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords    *String `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg       *String `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle       *String `xmlrpc:"website_meta_title,omptempty"`
}

// WebsiteSeoMetadatas represents array of website.seo.metadata model.
type WebsiteSeoMetadatas []WebsiteSeoMetadata

// WebsiteSeoMetadataModel is the odoo model name.
const WebsiteSeoMetadataModel = "website.seo.metadata"

// Many2One convert WebsiteSeoMetadata to *Many2One.
func (wsm *WebsiteSeoMetadata) Many2One() *Many2One {
	return NewMany2One(wsm.Id.Get(), "")
}

// CreateWebsiteSeoMetadata creates a new website.seo.metadata model and returns its id.
func (c *Client) CreateWebsiteSeoMetadata(wsm *WebsiteSeoMetadata) (int64, error) {
	return c.Create(WebsiteSeoMetadataModel, wsm)
}

// UpdateWebsiteSeoMetadata updates an existing website.seo.metadata record.
func (c *Client) UpdateWebsiteSeoMetadata(wsm *WebsiteSeoMetadata) error {
	return c.UpdateWebsiteSeoMetadatas([]int64{wsm.Id.Get()}, wsm)
}

// UpdateWebsiteSeoMetadatas updates existing website.seo.metadata records.
// All records (represented by ids) will be updated by wsm values.
func (c *Client) UpdateWebsiteSeoMetadatas(ids []int64, wsm *WebsiteSeoMetadata) error {
	return c.Update(WebsiteSeoMetadataModel, ids, wsm)
}

// DeleteWebsiteSeoMetadata deletes an existing website.seo.metadata record.
func (c *Client) DeleteWebsiteSeoMetadata(id int64) error {
	return c.DeleteWebsiteSeoMetadatas([]int64{id})
}

// DeleteWebsiteSeoMetadatas deletes existing website.seo.metadata records.
func (c *Client) DeleteWebsiteSeoMetadatas(ids []int64) error {
	return c.Delete(WebsiteSeoMetadataModel, ids)
}

// GetWebsiteSeoMetadata gets website.seo.metadata existing record.
func (c *Client) GetWebsiteSeoMetadata(id int64) (*WebsiteSeoMetadata, error) {
	wsms, err := c.GetWebsiteSeoMetadatas([]int64{id})
	if err != nil {
		return nil, err
	}
	if wsms != nil && len(*wsms) > 0 {
		return &((*wsms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.seo.metadata not found", id)
}

// GetWebsiteSeoMetadatas gets website.seo.metadata existing records.
func (c *Client) GetWebsiteSeoMetadatas(ids []int64) (*WebsiteSeoMetadatas, error) {
	wsms := &WebsiteSeoMetadatas{}
	if err := c.Read(WebsiteSeoMetadataModel, ids, nil, wsms); err != nil {
		return nil, err
	}
	return wsms, nil
}

// FindWebsiteSeoMetadata finds website.seo.metadata record by querying it with criteria.
func (c *Client) FindWebsiteSeoMetadata(criteria *Criteria) (*WebsiteSeoMetadata, error) {
	wsms := &WebsiteSeoMetadatas{}
	if err := c.SearchRead(WebsiteSeoMetadataModel, criteria, NewOptions().Limit(1), wsms); err != nil {
		return nil, err
	}
	if wsms != nil && len(*wsms) > 0 {
		return &((*wsms)[0]), nil
	}
	return nil, fmt.Errorf("website.seo.metadata was not found")
}

// FindWebsiteSeoMetadatas finds website.seo.metadata records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSeoMetadatas(criteria *Criteria, options *Options) (*WebsiteSeoMetadatas, error) {
	wsms := &WebsiteSeoMetadatas{}
	if err := c.SearchRead(WebsiteSeoMetadataModel, criteria, options, wsms); err != nil {
		return nil, err
	}
	return wsms, nil
}

// FindWebsiteSeoMetadataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSeoMetadataIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteSeoMetadataModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteSeoMetadataId finds record id by querying it with criteria.
func (c *Client) FindWebsiteSeoMetadataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteSeoMetadataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.seo.metadata was not found")
}
