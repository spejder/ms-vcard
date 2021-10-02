package odoo

import (
	"fmt"
)

// OpenacademyBundle represents openacademy.bundle model.
type OpenacademyBundle struct {
	Attachment  *String   `xmlrpc:"attachment,omptempty"`
	CourseIds   *Relation `xmlrpc:"course_ids,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Price       *Float    `xmlrpc:"price,omptempty"`
	Title       *String   `xmlrpc:"title,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OpenacademyBundles represents array of openacademy.bundle model.
type OpenacademyBundles []OpenacademyBundle

// OpenacademyBundleModel is the odoo model name.
const OpenacademyBundleModel = "openacademy.bundle"

// Many2One convert OpenacademyBundle to *Many2One.
func (ob *OpenacademyBundle) Many2One() *Many2One {
	return NewMany2One(ob.Id.Get(), "")
}

// CreateOpenacademyBundle creates a new openacademy.bundle model and returns its id.
func (c *Client) CreateOpenacademyBundle(ob *OpenacademyBundle) (int64, error) {
	return c.Create(OpenacademyBundleModel, ob)
}

// UpdateOpenacademyBundle updates an existing openacademy.bundle record.
func (c *Client) UpdateOpenacademyBundle(ob *OpenacademyBundle) error {
	return c.UpdateOpenacademyBundles([]int64{ob.Id.Get()}, ob)
}

// UpdateOpenacademyBundles updates existing openacademy.bundle records.
// All records (represented by ids) will be updated by ob values.
func (c *Client) UpdateOpenacademyBundles(ids []int64, ob *OpenacademyBundle) error {
	return c.Update(OpenacademyBundleModel, ids, ob)
}

// DeleteOpenacademyBundle deletes an existing openacademy.bundle record.
func (c *Client) DeleteOpenacademyBundle(id int64) error {
	return c.DeleteOpenacademyBundles([]int64{id})
}

// DeleteOpenacademyBundles deletes existing openacademy.bundle records.
func (c *Client) DeleteOpenacademyBundles(ids []int64) error {
	return c.Delete(OpenacademyBundleModel, ids)
}

// GetOpenacademyBundle gets openacademy.bundle existing record.
func (c *Client) GetOpenacademyBundle(id int64) (*OpenacademyBundle, error) {
	obs, err := c.GetOpenacademyBundles([]int64{id})
	if err != nil {
		return nil, err
	}
	if obs != nil && len(*obs) > 0 {
		return &((*obs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of openacademy.bundle not found", id)
}

// GetOpenacademyBundles gets openacademy.bundle existing records.
func (c *Client) GetOpenacademyBundles(ids []int64) (*OpenacademyBundles, error) {
	obs := &OpenacademyBundles{}
	if err := c.Read(OpenacademyBundleModel, ids, nil, obs); err != nil {
		return nil, err
	}
	return obs, nil
}

// FindOpenacademyBundle finds openacademy.bundle record by querying it with criteria.
func (c *Client) FindOpenacademyBundle(criteria *Criteria) (*OpenacademyBundle, error) {
	obs := &OpenacademyBundles{}
	if err := c.SearchRead(OpenacademyBundleModel, criteria, NewOptions().Limit(1), obs); err != nil {
		return nil, err
	}
	if obs != nil && len(*obs) > 0 {
		return &((*obs)[0]), nil
	}
	return nil, fmt.Errorf("openacademy.bundle was not found")
}

// FindOpenacademyBundles finds openacademy.bundle records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademyBundles(criteria *Criteria, options *Options) (*OpenacademyBundles, error) {
	obs := &OpenacademyBundles{}
	if err := c.SearchRead(OpenacademyBundleModel, criteria, options, obs); err != nil {
		return nil, err
	}
	return obs, nil
}

// FindOpenacademyBundleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademyBundleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OpenacademyBundleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOpenacademyBundleId finds record id by querying it with criteria.
func (c *Client) FindOpenacademyBundleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OpenacademyBundleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("openacademy.bundle was not found")
}
