package odoo

import (
	"fmt"
)

// ThemeUtils represents theme.utils model.
type ThemeUtils struct {
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
}

// ThemeUtilss represents array of theme.utils model.
type ThemeUtilss []ThemeUtils

// ThemeUtilsModel is the odoo model name.
const ThemeUtilsModel = "theme.utils"

// Many2One convert ThemeUtils to *Many2One.
func (tu *ThemeUtils) Many2One() *Many2One {
	return NewMany2One(tu.Id.Get(), "")
}

// CreateThemeUtils creates a new theme.utils model and returns its id.
func (c *Client) CreateThemeUtils(tu *ThemeUtils) (int64, error) {
	return c.Create(ThemeUtilsModel, tu)
}

// UpdateThemeUtils updates an existing theme.utils record.
func (c *Client) UpdateThemeUtils(tu *ThemeUtils) error {
	return c.UpdateThemeUtilss([]int64{tu.Id.Get()}, tu)
}

// UpdateThemeUtilss updates existing theme.utils records.
// All records (represented by ids) will be updated by tu values.
func (c *Client) UpdateThemeUtilss(ids []int64, tu *ThemeUtils) error {
	return c.Update(ThemeUtilsModel, ids, tu)
}

// DeleteThemeUtils deletes an existing theme.utils record.
func (c *Client) DeleteThemeUtils(id int64) error {
	return c.DeleteThemeUtilss([]int64{id})
}

// DeleteThemeUtilss deletes existing theme.utils records.
func (c *Client) DeleteThemeUtilss(ids []int64) error {
	return c.Delete(ThemeUtilsModel, ids)
}

// GetThemeUtils gets theme.utils existing record.
func (c *Client) GetThemeUtils(id int64) (*ThemeUtils, error) {
	tus, err := c.GetThemeUtilss([]int64{id})
	if err != nil {
		return nil, err
	}
	if tus != nil && len(*tus) > 0 {
		return &((*tus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of theme.utils not found", id)
}

// GetThemeUtilss gets theme.utils existing records.
func (c *Client) GetThemeUtilss(ids []int64) (*ThemeUtilss, error) {
	tus := &ThemeUtilss{}
	if err := c.Read(ThemeUtilsModel, ids, nil, tus); err != nil {
		return nil, err
	}
	return tus, nil
}

// FindThemeUtils finds theme.utils record by querying it with criteria.
func (c *Client) FindThemeUtils(criteria *Criteria) (*ThemeUtils, error) {
	tus := &ThemeUtilss{}
	if err := c.SearchRead(ThemeUtilsModel, criteria, NewOptions().Limit(1), tus); err != nil {
		return nil, err
	}
	if tus != nil && len(*tus) > 0 {
		return &((*tus)[0]), nil
	}
	return nil, fmt.Errorf("theme.utils was not found")
}

// FindThemeUtilss finds theme.utils records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeUtilss(criteria *Criteria, options *Options) (*ThemeUtilss, error) {
	tus := &ThemeUtilss{}
	if err := c.SearchRead(ThemeUtilsModel, criteria, options, tus); err != nil {
		return nil, err
	}
	return tus, nil
}

// FindThemeUtilsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeUtilsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ThemeUtilsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindThemeUtilsId finds record id by querying it with criteria.
func (c *Client) FindThemeUtilsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeUtilsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("theme.utils was not found")
}
