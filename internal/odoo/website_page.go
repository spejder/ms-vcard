package odoo

import (
	"fmt"
)

// WebsitePage represents website.page model.
type WebsitePage struct {
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	Arch                   *String    `xmlrpc:"arch,omptempty"`
	ArchBase               *String    `xmlrpc:"arch_base,omptempty"`
	ArchDb                 *String    `xmlrpc:"arch_db,omptempty"`
	ArchFs                 *String    `xmlrpc:"arch_fs,omptempty"`
	ArchPrev               *String    `xmlrpc:"arch_prev,omptempty"`
	ArchUpdated            *Bool      `xmlrpc:"arch_updated,omptempty"`
	CanPublish             *Bool      `xmlrpc:"can_publish,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomizeShow          *Bool      `xmlrpc:"customize_show,omptempty"`
	DatePublish            *Time      `xmlrpc:"date_publish,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	FieldParent            *String    `xmlrpc:"field_parent,omptempty"`
	FirstPageId            *Many2One  `xmlrpc:"first_page_id,omptempty"`
	GroupsId               *Relation  `xmlrpc:"groups_id,omptempty"`
	HeaderColor            *String    `xmlrpc:"header_color,omptempty"`
	HeaderOverlay          *Bool      `xmlrpc:"header_overlay,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	InheritChildrenIds     *Relation  `xmlrpc:"inherit_children_ids,omptempty"`
	InheritId              *Many2One  `xmlrpc:"inherit_id,omptempty"`
	IsHomepage             *Bool      `xmlrpc:"is_homepage,omptempty"`
	IsPublished            *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized         *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	IsVisible              *Bool      `xmlrpc:"is_visible,omptempty"`
	Key                    *String    `xmlrpc:"key,omptempty"`
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	MenuIds                *Relation  `xmlrpc:"menu_ids,omptempty"`
	Mode                   *Selection `xmlrpc:"mode,omptempty"`
	Model                  *String    `xmlrpc:"model,omptempty"`
	ModelDataId            *Many2One  `xmlrpc:"model_data_id,omptempty"`
	ModelIds               *Relation  `xmlrpc:"model_ids,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	PageIds                *Relation  `xmlrpc:"page_ids,omptempty"`
	Priority               *Int       `xmlrpc:"priority,omptempty"`
	ThemeTemplateId        *Many2One  `xmlrpc:"theme_template_id,omptempty"`
	Track                  *Bool      `xmlrpc:"track,omptempty"`
	Type                   *Selection `xmlrpc:"type,omptempty"`
	Url                    *String    `xmlrpc:"url,omptempty"`
	ViewId                 *Many2One  `xmlrpc:"view_id,omptempty"`
	WebsiteId              *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteIndexed         *Bool      `xmlrpc:"website_indexed,omptempty"`
	WebsiteMetaDescription *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords    *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg       *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle       *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished       *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl             *String    `xmlrpc:"website_url,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId                  *String    `xmlrpc:"xml_id,omptempty"`
}

// WebsitePages represents array of website.page model.
type WebsitePages []WebsitePage

// WebsitePageModel is the odoo model name.
const WebsitePageModel = "website.page"

// Many2One convert WebsitePage to *Many2One.
func (wp *WebsitePage) Many2One() *Many2One {
	return NewMany2One(wp.Id.Get(), "")
}

// CreateWebsitePage creates a new website.page model and returns its id.
func (c *Client) CreateWebsitePage(wp *WebsitePage) (int64, error) {
	return c.Create(WebsitePageModel, wp)
}

// UpdateWebsitePage updates an existing website.page record.
func (c *Client) UpdateWebsitePage(wp *WebsitePage) error {
	return c.UpdateWebsitePages([]int64{wp.Id.Get()}, wp)
}

// UpdateWebsitePages updates existing website.page records.
// All records (represented by ids) will be updated by wp values.
func (c *Client) UpdateWebsitePages(ids []int64, wp *WebsitePage) error {
	return c.Update(WebsitePageModel, ids, wp)
}

// DeleteWebsitePage deletes an existing website.page record.
func (c *Client) DeleteWebsitePage(id int64) error {
	return c.DeleteWebsitePages([]int64{id})
}

// DeleteWebsitePages deletes existing website.page records.
func (c *Client) DeleteWebsitePages(ids []int64) error {
	return c.Delete(WebsitePageModel, ids)
}

// GetWebsitePage gets website.page existing record.
func (c *Client) GetWebsitePage(id int64) (*WebsitePage, error) {
	wps, err := c.GetWebsitePages([]int64{id})
	if err != nil {
		return nil, err
	}
	if wps != nil && len(*wps) > 0 {
		return &((*wps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.page not found", id)
}

// GetWebsitePages gets website.page existing records.
func (c *Client) GetWebsitePages(ids []int64) (*WebsitePages, error) {
	wps := &WebsitePages{}
	if err := c.Read(WebsitePageModel, ids, nil, wps); err != nil {
		return nil, err
	}
	return wps, nil
}

// FindWebsitePage finds website.page record by querying it with criteria.
func (c *Client) FindWebsitePage(criteria *Criteria) (*WebsitePage, error) {
	wps := &WebsitePages{}
	if err := c.SearchRead(WebsitePageModel, criteria, NewOptions().Limit(1), wps); err != nil {
		return nil, err
	}
	if wps != nil && len(*wps) > 0 {
		return &((*wps)[0]), nil
	}
	return nil, fmt.Errorf("website.page was not found")
}

// FindWebsitePages finds website.page records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePages(criteria *Criteria, options *Options) (*WebsitePages, error) {
	wps := &WebsitePages{}
	if err := c.SearchRead(WebsitePageModel, criteria, options, wps); err != nil {
		return nil, err
	}
	return wps, nil
}

// FindWebsitePageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsitePageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsitePageId finds record id by querying it with criteria.
func (c *Client) FindWebsitePageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.page was not found")
}
