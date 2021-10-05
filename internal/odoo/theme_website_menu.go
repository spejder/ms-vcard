package odoo

import (
	"fmt"
)

// ThemeWebsiteMenu represents theme.website.menu model.
type ThemeWebsiteMenu struct {
	CopyIds     *Relation `xmlrpc:"copy_ids,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	NewWindow   *Bool     `xmlrpc:"new_window,omptempty"`
	PageId      *Many2One `xmlrpc:"page_id,omptempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	Url         *String   `xmlrpc:"url,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ThemeWebsiteMenus represents array of theme.website.menu model.
type ThemeWebsiteMenus []ThemeWebsiteMenu

// ThemeWebsiteMenuModel is the odoo model name.
const ThemeWebsiteMenuModel = "theme.website.menu"

// Many2One convert ThemeWebsiteMenu to *Many2One.
func (twm *ThemeWebsiteMenu) Many2One() *Many2One {
	return NewMany2One(twm.Id.Get(), "")
}

// CreateThemeWebsiteMenu creates a new theme.website.menu model and returns its id.
func (c *Client) CreateThemeWebsiteMenu(twm *ThemeWebsiteMenu) (int64, error) {
	return c.Create(ThemeWebsiteMenuModel, twm)
}

// UpdateThemeWebsiteMenu updates an existing theme.website.menu record.
func (c *Client) UpdateThemeWebsiteMenu(twm *ThemeWebsiteMenu) error {
	return c.UpdateThemeWebsiteMenus([]int64{twm.Id.Get()}, twm)
}

// UpdateThemeWebsiteMenus updates existing theme.website.menu records.
// All records (represented by ids) will be updated by twm values.
func (c *Client) UpdateThemeWebsiteMenus(ids []int64, twm *ThemeWebsiteMenu) error {
	return c.Update(ThemeWebsiteMenuModel, ids, twm)
}

// DeleteThemeWebsiteMenu deletes an existing theme.website.menu record.
func (c *Client) DeleteThemeWebsiteMenu(id int64) error {
	return c.DeleteThemeWebsiteMenus([]int64{id})
}

// DeleteThemeWebsiteMenus deletes existing theme.website.menu records.
func (c *Client) DeleteThemeWebsiteMenus(ids []int64) error {
	return c.Delete(ThemeWebsiteMenuModel, ids)
}

// GetThemeWebsiteMenu gets theme.website.menu existing record.
func (c *Client) GetThemeWebsiteMenu(id int64) (*ThemeWebsiteMenu, error) {
	twms, err := c.GetThemeWebsiteMenus([]int64{id})
	if err != nil {
		return nil, err
	}
	if twms != nil && len(*twms) > 0 {
		return &((*twms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of theme.website.menu not found", id)
}

// GetThemeWebsiteMenus gets theme.website.menu existing records.
func (c *Client) GetThemeWebsiteMenus(ids []int64) (*ThemeWebsiteMenus, error) {
	twms := &ThemeWebsiteMenus{}
	if err := c.Read(ThemeWebsiteMenuModel, ids, nil, twms); err != nil {
		return nil, err
	}
	return twms, nil
}

// FindThemeWebsiteMenu finds theme.website.menu record by querying it with criteria.
func (c *Client) FindThemeWebsiteMenu(criteria *Criteria) (*ThemeWebsiteMenu, error) {
	twms := &ThemeWebsiteMenus{}
	if err := c.SearchRead(ThemeWebsiteMenuModel, criteria, NewOptions().Limit(1), twms); err != nil {
		return nil, err
	}
	if twms != nil && len(*twms) > 0 {
		return &((*twms)[0]), nil
	}
	return nil, fmt.Errorf("theme.website.menu was not found")
}

// FindThemeWebsiteMenus finds theme.website.menu records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeWebsiteMenus(criteria *Criteria, options *Options) (*ThemeWebsiteMenus, error) {
	twms := &ThemeWebsiteMenus{}
	if err := c.SearchRead(ThemeWebsiteMenuModel, criteria, options, twms); err != nil {
		return nil, err
	}
	return twms, nil
}

// FindThemeWebsiteMenuIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeWebsiteMenuIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ThemeWebsiteMenuModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindThemeWebsiteMenuId finds record id by querying it with criteria.
func (c *Client) FindThemeWebsiteMenuId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeWebsiteMenuModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("theme.website.menu was not found")
}
