package odoo

import (
	"fmt"
)

// SlideCourseType represents slide.course_type model.
type SlideCourseType struct {
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideCourseTypes represents array of slide.course_type model.
type SlideCourseTypes []SlideCourseType

// SlideCourseTypeModel is the odoo model name.
const SlideCourseTypeModel = "slide.course_type"

// Many2One convert SlideCourseType to *Many2One.
func (sc *SlideCourseType) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSlideCourseType creates a new slide.course_type model and returns its id.
func (c *Client) CreateSlideCourseType(sc *SlideCourseType) (int64, error) {
	return c.Create(SlideCourseTypeModel, sc)
}

// UpdateSlideCourseType updates an existing slide.course_type record.
func (c *Client) UpdateSlideCourseType(sc *SlideCourseType) error {
	return c.UpdateSlideCourseTypes([]int64{sc.Id.Get()}, sc)
}

// UpdateSlideCourseTypes updates existing slide.course_type records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSlideCourseTypes(ids []int64, sc *SlideCourseType) error {
	return c.Update(SlideCourseTypeModel, ids, sc)
}

// DeleteSlideCourseType deletes an existing slide.course_type record.
func (c *Client) DeleteSlideCourseType(id int64) error {
	return c.DeleteSlideCourseTypes([]int64{id})
}

// DeleteSlideCourseTypes deletes existing slide.course_type records.
func (c *Client) DeleteSlideCourseTypes(ids []int64) error {
	return c.Delete(SlideCourseTypeModel, ids)
}

// GetSlideCourseType gets slide.course_type existing record.
func (c *Client) GetSlideCourseType(id int64) (*SlideCourseType, error) {
	scs, err := c.GetSlideCourseTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.course_type not found", id)
}

// GetSlideCourseTypes gets slide.course_type existing records.
func (c *Client) GetSlideCourseTypes(ids []int64) (*SlideCourseTypes, error) {
	scs := &SlideCourseTypes{}
	if err := c.Read(SlideCourseTypeModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideCourseType finds slide.course_type record by querying it with criteria.
func (c *Client) FindSlideCourseType(criteria *Criteria) (*SlideCourseType, error) {
	scs := &SlideCourseTypes{}
	if err := c.SearchRead(SlideCourseTypeModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("slide.course_type was not found")
}

// FindSlideCourseTypes finds slide.course_type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideCourseTypes(criteria *Criteria, options *Options) (*SlideCourseTypes, error) {
	scs := &SlideCourseTypes{}
	if err := c.SearchRead(SlideCourseTypeModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideCourseTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideCourseTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideCourseTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideCourseTypeId finds record id by querying it with criteria.
func (c *Client) FindSlideCourseTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideCourseTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.course_type was not found")
}
