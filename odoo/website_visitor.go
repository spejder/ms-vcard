package odoo

import (
	"fmt"
)

// WebsiteVisitor represents website.visitor model.
type WebsiteVisitor struct {
	AccessToken            *String    `xmlrpc:"access_token,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	CountryFlag            *String    `xmlrpc:"country_flag,omptempty"`
	CountryId              *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	Email                  *String    `xmlrpc:"email,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	IsConnected            *Bool      `xmlrpc:"is_connected,omptempty"`
	LangId                 *Many2One  `xmlrpc:"lang_id,omptempty"`
	LastConnectionDatetime *Time      `xmlrpc:"last_connection_datetime,omptempty"`
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	LastVisitedPageId      *Many2One  `xmlrpc:"last_visited_page_id,omptempty"`
	LeadCount              *Int       `xmlrpc:"lead_count,omptempty"`
	LeadIds                *Relation  `xmlrpc:"lead_ids,omptempty"`
	LivechatOperatorId     *Many2One  `xmlrpc:"livechat_operator_id,omptempty"`
	LivechatOperatorName   *String    `xmlrpc:"livechat_operator_name,omptempty"`
	MailChannelIds         *Relation  `xmlrpc:"mail_channel_ids,omptempty"`
	Mobile                 *String    `xmlrpc:"mobile,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	PageCount              *Int       `xmlrpc:"page_count,omptempty"`
	PageIds                *Relation  `xmlrpc:"page_ids,omptempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerImage           *String    `xmlrpc:"partner_image,omptempty"`
	SessionCount           *Int       `xmlrpc:"session_count,omptempty"`
	TimeSinceLastAction    *String    `xmlrpc:"time_since_last_action,omptempty"`
	Timezone               *Selection `xmlrpc:"timezone,omptempty"`
	VisitCount             *Int       `xmlrpc:"visit_count,omptempty"`
	VisitorPageCount       *Int       `xmlrpc:"visitor_page_count,omptempty"`
	WebsiteId              *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteTrackIds        *Relation  `xmlrpc:"website_track_ids,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WebsiteVisitors represents array of website.visitor model.
type WebsiteVisitors []WebsiteVisitor

// WebsiteVisitorModel is the odoo model name.
const WebsiteVisitorModel = "website.visitor"

// Many2One convert WebsiteVisitor to *Many2One.
func (wv *WebsiteVisitor) Many2One() *Many2One {
	return NewMany2One(wv.Id.Get(), "")
}

// CreateWebsiteVisitor creates a new website.visitor model and returns its id.
func (c *Client) CreateWebsiteVisitor(wv *WebsiteVisitor) (int64, error) {
	return c.Create(WebsiteVisitorModel, wv)
}

// UpdateWebsiteVisitor updates an existing website.visitor record.
func (c *Client) UpdateWebsiteVisitor(wv *WebsiteVisitor) error {
	return c.UpdateWebsiteVisitors([]int64{wv.Id.Get()}, wv)
}

// UpdateWebsiteVisitors updates existing website.visitor records.
// All records (represented by ids) will be updated by wv values.
func (c *Client) UpdateWebsiteVisitors(ids []int64, wv *WebsiteVisitor) error {
	return c.Update(WebsiteVisitorModel, ids, wv)
}

// DeleteWebsiteVisitor deletes an existing website.visitor record.
func (c *Client) DeleteWebsiteVisitor(id int64) error {
	return c.DeleteWebsiteVisitors([]int64{id})
}

// DeleteWebsiteVisitors deletes existing website.visitor records.
func (c *Client) DeleteWebsiteVisitors(ids []int64) error {
	return c.Delete(WebsiteVisitorModel, ids)
}

// GetWebsiteVisitor gets website.visitor existing record.
func (c *Client) GetWebsiteVisitor(id int64) (*WebsiteVisitor, error) {
	wvs, err := c.GetWebsiteVisitors([]int64{id})
	if err != nil {
		return nil, err
	}
	if wvs != nil && len(*wvs) > 0 {
		return &((*wvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.visitor not found", id)
}

// GetWebsiteVisitors gets website.visitor existing records.
func (c *Client) GetWebsiteVisitors(ids []int64) (*WebsiteVisitors, error) {
	wvs := &WebsiteVisitors{}
	if err := c.Read(WebsiteVisitorModel, ids, nil, wvs); err != nil {
		return nil, err
	}
	return wvs, nil
}

// FindWebsiteVisitor finds website.visitor record by querying it with criteria.
func (c *Client) FindWebsiteVisitor(criteria *Criteria) (*WebsiteVisitor, error) {
	wvs := &WebsiteVisitors{}
	if err := c.SearchRead(WebsiteVisitorModel, criteria, NewOptions().Limit(1), wvs); err != nil {
		return nil, err
	}
	if wvs != nil && len(*wvs) > 0 {
		return &((*wvs)[0]), nil
	}
	return nil, fmt.Errorf("website.visitor was not found")
}

// FindWebsiteVisitors finds website.visitor records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteVisitors(criteria *Criteria, options *Options) (*WebsiteVisitors, error) {
	wvs := &WebsiteVisitors{}
	if err := c.SearchRead(WebsiteVisitorModel, criteria, options, wvs); err != nil {
		return nil, err
	}
	return wvs, nil
}

// FindWebsiteVisitorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteVisitorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteVisitorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteVisitorId finds record id by querying it with criteria.
func (c *Client) FindWebsiteVisitorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteVisitorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.visitor was not found")
}
