package odoo

import (
	"fmt"
)

// ThemeIrUiView represents theme.ir.ui.view model.
type ThemeIrUiView struct {
	Active      *Bool      `xmlrpc:"active,omptempty"`
	Arch        *String    `xmlrpc:"arch,omptempty"`
	ArchFs      *String    `xmlrpc:"arch_fs,omptempty"`
	CopyIds     *Relation  `xmlrpc:"copy_ids,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	InheritId   *String    `xmlrpc:"inherit_id,omptempty"`
	Key         *String    `xmlrpc:"key,omptempty"`
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Mode        *Selection `xmlrpc:"mode,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Priority    *Int       `xmlrpc:"priority,omptempty"`
	Type        *String    `xmlrpc:"type,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ThemeIrUiViews represents array of theme.ir.ui.view model.
type ThemeIrUiViews []ThemeIrUiView

// ThemeIrUiViewModel is the odoo model name.
const ThemeIrUiViewModel = "theme.ir.ui.view"

// Many2One convert ThemeIrUiView to *Many2One.
func (tiuv *ThemeIrUiView) Many2One() *Many2One {
	return NewMany2One(tiuv.Id.Get(), "")
}

// CreateThemeIrUiView creates a new theme.ir.ui.view model and returns its id.
func (c *Client) CreateThemeIrUiView(tiuv *ThemeIrUiView) (int64, error) {
	return c.Create(ThemeIrUiViewModel, tiuv)
}

// UpdateThemeIrUiView updates an existing theme.ir.ui.view record.
func (c *Client) UpdateThemeIrUiView(tiuv *ThemeIrUiView) error {
	return c.UpdateThemeIrUiViews([]int64{tiuv.Id.Get()}, tiuv)
}

// UpdateThemeIrUiViews updates existing theme.ir.ui.view records.
// All records (represented by ids) will be updated by tiuv values.
func (c *Client) UpdateThemeIrUiViews(ids []int64, tiuv *ThemeIrUiView) error {
	return c.Update(ThemeIrUiViewModel, ids, tiuv)
}

// DeleteThemeIrUiView deletes an existing theme.ir.ui.view record.
func (c *Client) DeleteThemeIrUiView(id int64) error {
	return c.DeleteThemeIrUiViews([]int64{id})
}

// DeleteThemeIrUiViews deletes existing theme.ir.ui.view records.
func (c *Client) DeleteThemeIrUiViews(ids []int64) error {
	return c.Delete(ThemeIrUiViewModel, ids)
}

// GetThemeIrUiView gets theme.ir.ui.view existing record.
func (c *Client) GetThemeIrUiView(id int64) (*ThemeIrUiView, error) {
	tiuvs, err := c.GetThemeIrUiViews([]int64{id})
	if err != nil {
		return nil, err
	}
	if tiuvs != nil && len(*tiuvs) > 0 {
		return &((*tiuvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of theme.ir.ui.view not found", id)
}

// GetThemeIrUiViews gets theme.ir.ui.view existing records.
func (c *Client) GetThemeIrUiViews(ids []int64) (*ThemeIrUiViews, error) {
	tiuvs := &ThemeIrUiViews{}
	if err := c.Read(ThemeIrUiViewModel, ids, nil, tiuvs); err != nil {
		return nil, err
	}
	return tiuvs, nil
}

// FindThemeIrUiView finds theme.ir.ui.view record by querying it with criteria.
func (c *Client) FindThemeIrUiView(criteria *Criteria) (*ThemeIrUiView, error) {
	tiuvs := &ThemeIrUiViews{}
	if err := c.SearchRead(ThemeIrUiViewModel, criteria, NewOptions().Limit(1), tiuvs); err != nil {
		return nil, err
	}
	if tiuvs != nil && len(*tiuvs) > 0 {
		return &((*tiuvs)[0]), nil
	}
	return nil, fmt.Errorf("theme.ir.ui.view was not found")
}

// FindThemeIrUiViews finds theme.ir.ui.view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrUiViews(criteria *Criteria, options *Options) (*ThemeIrUiViews, error) {
	tiuvs := &ThemeIrUiViews{}
	if err := c.SearchRead(ThemeIrUiViewModel, criteria, options, tiuvs); err != nil {
		return nil, err
	}
	return tiuvs, nil
}

// FindThemeIrUiViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrUiViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ThemeIrUiViewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindThemeIrUiViewId finds record id by querying it with criteria.
func (c *Client) FindThemeIrUiViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeIrUiViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("theme.ir.ui.view was not found")
}
