package odoo

import (
	"fmt"
)

// ThemeIrAttachment represents theme.ir.attachment model.
type ThemeIrAttachment struct {
	CopyIds     *Relation `xmlrpc:"copy_ids,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Key         *String   `xmlrpc:"key,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Url         *String   `xmlrpc:"url,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ThemeIrAttachments represents array of theme.ir.attachment model.
type ThemeIrAttachments []ThemeIrAttachment

// ThemeIrAttachmentModel is the odoo model name.
const ThemeIrAttachmentModel = "theme.ir.attachment"

// Many2One convert ThemeIrAttachment to *Many2One.
func (tia *ThemeIrAttachment) Many2One() *Many2One {
	return NewMany2One(tia.Id.Get(), "")
}

// CreateThemeIrAttachment creates a new theme.ir.attachment model and returns its id.
func (c *Client) CreateThemeIrAttachment(tia *ThemeIrAttachment) (int64, error) {
	return c.Create(ThemeIrAttachmentModel, tia)
}

// UpdateThemeIrAttachment updates an existing theme.ir.attachment record.
func (c *Client) UpdateThemeIrAttachment(tia *ThemeIrAttachment) error {
	return c.UpdateThemeIrAttachments([]int64{tia.Id.Get()}, tia)
}

// UpdateThemeIrAttachments updates existing theme.ir.attachment records.
// All records (represented by ids) will be updated by tia values.
func (c *Client) UpdateThemeIrAttachments(ids []int64, tia *ThemeIrAttachment) error {
	return c.Update(ThemeIrAttachmentModel, ids, tia)
}

// DeleteThemeIrAttachment deletes an existing theme.ir.attachment record.
func (c *Client) DeleteThemeIrAttachment(id int64) error {
	return c.DeleteThemeIrAttachments([]int64{id})
}

// DeleteThemeIrAttachments deletes existing theme.ir.attachment records.
func (c *Client) DeleteThemeIrAttachments(ids []int64) error {
	return c.Delete(ThemeIrAttachmentModel, ids)
}

// GetThemeIrAttachment gets theme.ir.attachment existing record.
func (c *Client) GetThemeIrAttachment(id int64) (*ThemeIrAttachment, error) {
	tias, err := c.GetThemeIrAttachments([]int64{id})
	if err != nil {
		return nil, err
	}
	if tias != nil && len(*tias) > 0 {
		return &((*tias)[0]), nil
	}
	return nil, fmt.Errorf("id %v of theme.ir.attachment not found", id)
}

// GetThemeIrAttachments gets theme.ir.attachment existing records.
func (c *Client) GetThemeIrAttachments(ids []int64) (*ThemeIrAttachments, error) {
	tias := &ThemeIrAttachments{}
	if err := c.Read(ThemeIrAttachmentModel, ids, nil, tias); err != nil {
		return nil, err
	}
	return tias, nil
}

// FindThemeIrAttachment finds theme.ir.attachment record by querying it with criteria.
func (c *Client) FindThemeIrAttachment(criteria *Criteria) (*ThemeIrAttachment, error) {
	tias := &ThemeIrAttachments{}
	if err := c.SearchRead(ThemeIrAttachmentModel, criteria, NewOptions().Limit(1), tias); err != nil {
		return nil, err
	}
	if tias != nil && len(*tias) > 0 {
		return &((*tias)[0]), nil
	}
	return nil, fmt.Errorf("theme.ir.attachment was not found")
}

// FindThemeIrAttachments finds theme.ir.attachment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrAttachments(criteria *Criteria, options *Options) (*ThemeIrAttachments, error) {
	tias := &ThemeIrAttachments{}
	if err := c.SearchRead(ThemeIrAttachmentModel, criteria, options, tias); err != nil {
		return nil, err
	}
	return tias, nil
}

// FindThemeIrAttachmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrAttachmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ThemeIrAttachmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindThemeIrAttachmentId finds record id by querying it with criteria.
func (c *Client) FindThemeIrAttachmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeIrAttachmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("theme.ir.attachment was not found")
}
