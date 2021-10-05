package odoo

import (
	"fmt"
)

// WebsiteTrack represents website.track model.
type WebsiteTrack struct {
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	PageId        *Many2One `xmlrpc:"page_id,omptempty"`
	Url           *String   `xmlrpc:"url,omptempty"`
	VisitDatetime *Time     `xmlrpc:"visit_datetime,omptempty"`
	VisitorId     *Many2One `xmlrpc:"visitor_id,omptempty"`
}

// WebsiteTracks represents array of website.track model.
type WebsiteTracks []WebsiteTrack

// WebsiteTrackModel is the odoo model name.
const WebsiteTrackModel = "website.track"

// Many2One convert WebsiteTrack to *Many2One.
func (wt *WebsiteTrack) Many2One() *Many2One {
	return NewMany2One(wt.Id.Get(), "")
}

// CreateWebsiteTrack creates a new website.track model and returns its id.
func (c *Client) CreateWebsiteTrack(wt *WebsiteTrack) (int64, error) {
	return c.Create(WebsiteTrackModel, wt)
}

// UpdateWebsiteTrack updates an existing website.track record.
func (c *Client) UpdateWebsiteTrack(wt *WebsiteTrack) error {
	return c.UpdateWebsiteTracks([]int64{wt.Id.Get()}, wt)
}

// UpdateWebsiteTracks updates existing website.track records.
// All records (represented by ids) will be updated by wt values.
func (c *Client) UpdateWebsiteTracks(ids []int64, wt *WebsiteTrack) error {
	return c.Update(WebsiteTrackModel, ids, wt)
}

// DeleteWebsiteTrack deletes an existing website.track record.
func (c *Client) DeleteWebsiteTrack(id int64) error {
	return c.DeleteWebsiteTracks([]int64{id})
}

// DeleteWebsiteTracks deletes existing website.track records.
func (c *Client) DeleteWebsiteTracks(ids []int64) error {
	return c.Delete(WebsiteTrackModel, ids)
}

// GetWebsiteTrack gets website.track existing record.
func (c *Client) GetWebsiteTrack(id int64) (*WebsiteTrack, error) {
	wts, err := c.GetWebsiteTracks([]int64{id})
	if err != nil {
		return nil, err
	}
	if wts != nil && len(*wts) > 0 {
		return &((*wts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.track not found", id)
}

// GetWebsiteTracks gets website.track existing records.
func (c *Client) GetWebsiteTracks(ids []int64) (*WebsiteTracks, error) {
	wts := &WebsiteTracks{}
	if err := c.Read(WebsiteTrackModel, ids, nil, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWebsiteTrack finds website.track record by querying it with criteria.
func (c *Client) FindWebsiteTrack(criteria *Criteria) (*WebsiteTrack, error) {
	wts := &WebsiteTracks{}
	if err := c.SearchRead(WebsiteTrackModel, criteria, NewOptions().Limit(1), wts); err != nil {
		return nil, err
	}
	if wts != nil && len(*wts) > 0 {
		return &((*wts)[0]), nil
	}
	return nil, fmt.Errorf("website.track was not found")
}

// FindWebsiteTracks finds website.track records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteTracks(criteria *Criteria, options *Options) (*WebsiteTracks, error) {
	wts := &WebsiteTracks{}
	if err := c.SearchRead(WebsiteTrackModel, criteria, options, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWebsiteTrackIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteTrackIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteTrackModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteTrackId finds record id by querying it with criteria.
func (c *Client) FindWebsiteTrackId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteTrackModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.track was not found")
}
