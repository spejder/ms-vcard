package odoo

import (
	"fmt"
)

// OpenacademyCourse represents openacademy.course model.
type OpenacademyCourse struct {
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	Description   *String   `xmlrpc:"description,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	ResponsibleId *Many2One `xmlrpc:"responsible_id,omptempty"`
	SessionIds    *Relation `xmlrpc:"session_ids,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OpenacademyCourses represents array of openacademy.course model.
type OpenacademyCourses []OpenacademyCourse

// OpenacademyCourseModel is the odoo model name.
const OpenacademyCourseModel = "openacademy.course"

// Many2One convert OpenacademyCourse to *Many2One.
func (oc *OpenacademyCourse) Many2One() *Many2One {
	return NewMany2One(oc.Id.Get(), "")
}

// CreateOpenacademyCourse creates a new openacademy.course model and returns its id.
func (c *Client) CreateOpenacademyCourse(oc *OpenacademyCourse) (int64, error) {
	return c.Create(OpenacademyCourseModel, oc)
}

// UpdateOpenacademyCourse updates an existing openacademy.course record.
func (c *Client) UpdateOpenacademyCourse(oc *OpenacademyCourse) error {
	return c.UpdateOpenacademyCourses([]int64{oc.Id.Get()}, oc)
}

// UpdateOpenacademyCourses updates existing openacademy.course records.
// All records (represented by ids) will be updated by oc values.
func (c *Client) UpdateOpenacademyCourses(ids []int64, oc *OpenacademyCourse) error {
	return c.Update(OpenacademyCourseModel, ids, oc)
}

// DeleteOpenacademyCourse deletes an existing openacademy.course record.
func (c *Client) DeleteOpenacademyCourse(id int64) error {
	return c.DeleteOpenacademyCourses([]int64{id})
}

// DeleteOpenacademyCourses deletes existing openacademy.course records.
func (c *Client) DeleteOpenacademyCourses(ids []int64) error {
	return c.Delete(OpenacademyCourseModel, ids)
}

// GetOpenacademyCourse gets openacademy.course existing record.
func (c *Client) GetOpenacademyCourse(id int64) (*OpenacademyCourse, error) {
	ocs, err := c.GetOpenacademyCourses([]int64{id})
	if err != nil {
		return nil, err
	}
	if ocs != nil && len(*ocs) > 0 {
		return &((*ocs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of openacademy.course not found", id)
}

// GetOpenacademyCourses gets openacademy.course existing records.
func (c *Client) GetOpenacademyCourses(ids []int64) (*OpenacademyCourses, error) {
	ocs := &OpenacademyCourses{}
	if err := c.Read(OpenacademyCourseModel, ids, nil, ocs); err != nil {
		return nil, err
	}
	return ocs, nil
}

// FindOpenacademyCourse finds openacademy.course record by querying it with criteria.
func (c *Client) FindOpenacademyCourse(criteria *Criteria) (*OpenacademyCourse, error) {
	ocs := &OpenacademyCourses{}
	if err := c.SearchRead(OpenacademyCourseModel, criteria, NewOptions().Limit(1), ocs); err != nil {
		return nil, err
	}
	if ocs != nil && len(*ocs) > 0 {
		return &((*ocs)[0]), nil
	}
	return nil, fmt.Errorf("openacademy.course was not found")
}

// FindOpenacademyCourses finds openacademy.course records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademyCourses(criteria *Criteria, options *Options) (*OpenacademyCourses, error) {
	ocs := &OpenacademyCourses{}
	if err := c.SearchRead(OpenacademyCourseModel, criteria, options, ocs); err != nil {
		return nil, err
	}
	return ocs, nil
}

// FindOpenacademyCourseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademyCourseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OpenacademyCourseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOpenacademyCourseId finds record id by querying it with criteria.
func (c *Client) FindOpenacademyCourseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OpenacademyCourseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("openacademy.course was not found")
}
