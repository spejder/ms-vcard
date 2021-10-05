package odoo

import (
	"fmt"
)

// WebsiteMassMailingPopup represents website.mass_mailing.popup model.
type WebsiteMassMailingPopup struct {
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	MailingListId *Many2One `xmlrpc:"mailing_list_id,omptempty"`
	PopupContent  *String   `xmlrpc:"popup_content,omptempty"`
	WebsiteId     *Many2One `xmlrpc:"website_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WebsiteMassMailingPopups represents array of website.mass_mailing.popup model.
type WebsiteMassMailingPopups []WebsiteMassMailingPopup

// WebsiteMassMailingPopupModel is the odoo model name.
const WebsiteMassMailingPopupModel = "website.mass_mailing.popup"

// Many2One convert WebsiteMassMailingPopup to *Many2One.
func (wmp *WebsiteMassMailingPopup) Many2One() *Many2One {
	return NewMany2One(wmp.Id.Get(), "")
}

// CreateWebsiteMassMailingPopup creates a new website.mass_mailing.popup model and returns its id.
func (c *Client) CreateWebsiteMassMailingPopup(wmp *WebsiteMassMailingPopup) (int64, error) {
	return c.Create(WebsiteMassMailingPopupModel, wmp)
}

// UpdateWebsiteMassMailingPopup updates an existing website.mass_mailing.popup record.
func (c *Client) UpdateWebsiteMassMailingPopup(wmp *WebsiteMassMailingPopup) error {
	return c.UpdateWebsiteMassMailingPopups([]int64{wmp.Id.Get()}, wmp)
}

// UpdateWebsiteMassMailingPopups updates existing website.mass_mailing.popup records.
// All records (represented by ids) will be updated by wmp values.
func (c *Client) UpdateWebsiteMassMailingPopups(ids []int64, wmp *WebsiteMassMailingPopup) error {
	return c.Update(WebsiteMassMailingPopupModel, ids, wmp)
}

// DeleteWebsiteMassMailingPopup deletes an existing website.mass_mailing.popup record.
func (c *Client) DeleteWebsiteMassMailingPopup(id int64) error {
	return c.DeleteWebsiteMassMailingPopups([]int64{id})
}

// DeleteWebsiteMassMailingPopups deletes existing website.mass_mailing.popup records.
func (c *Client) DeleteWebsiteMassMailingPopups(ids []int64) error {
	return c.Delete(WebsiteMassMailingPopupModel, ids)
}

// GetWebsiteMassMailingPopup gets website.mass_mailing.popup existing record.
func (c *Client) GetWebsiteMassMailingPopup(id int64) (*WebsiteMassMailingPopup, error) {
	wmps, err := c.GetWebsiteMassMailingPopups([]int64{id})
	if err != nil {
		return nil, err
	}
	if wmps != nil && len(*wmps) > 0 {
		return &((*wmps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.mass_mailing.popup not found", id)
}

// GetWebsiteMassMailingPopups gets website.mass_mailing.popup existing records.
func (c *Client) GetWebsiteMassMailingPopups(ids []int64) (*WebsiteMassMailingPopups, error) {
	wmps := &WebsiteMassMailingPopups{}
	if err := c.Read(WebsiteMassMailingPopupModel, ids, nil, wmps); err != nil {
		return nil, err
	}
	return wmps, nil
}

// FindWebsiteMassMailingPopup finds website.mass_mailing.popup record by querying it with criteria.
func (c *Client) FindWebsiteMassMailingPopup(criteria *Criteria) (*WebsiteMassMailingPopup, error) {
	wmps := &WebsiteMassMailingPopups{}
	if err := c.SearchRead(WebsiteMassMailingPopupModel, criteria, NewOptions().Limit(1), wmps); err != nil {
		return nil, err
	}
	if wmps != nil && len(*wmps) > 0 {
		return &((*wmps)[0]), nil
	}
	return nil, fmt.Errorf("website.mass_mailing.popup was not found")
}

// FindWebsiteMassMailingPopups finds website.mass_mailing.popup records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMassMailingPopups(criteria *Criteria, options *Options) (*WebsiteMassMailingPopups, error) {
	wmps := &WebsiteMassMailingPopups{}
	if err := c.SearchRead(WebsiteMassMailingPopupModel, criteria, options, wmps); err != nil {
		return nil, err
	}
	return wmps, nil
}

// FindWebsiteMassMailingPopupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMassMailingPopupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteMassMailingPopupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteMassMailingPopupId finds record id by querying it with criteria.
func (c *Client) FindWebsiteMassMailingPopupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteMassMailingPopupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.mass_mailing.popup was not found")
}
