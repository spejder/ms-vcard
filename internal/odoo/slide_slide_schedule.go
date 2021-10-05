package odoo

import (
	"fmt"
)

// SlideSlideSchedule represents slide.slide_schedule model.
type SlideSlideSchedule struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt   *Time     `xmlrpc:"deleted_at,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ScheduleId  *Many2One `xmlrpc:"schedule_id,omptempty"`
	SlideId     *Many2One `xmlrpc:"slide_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideSlideSchedules represents array of slide.slide_schedule model.
type SlideSlideSchedules []SlideSlideSchedule

// SlideSlideScheduleModel is the odoo model name.
const SlideSlideScheduleModel = "slide.slide_schedule"

// Many2One convert SlideSlideSchedule to *Many2One.
func (ss *SlideSlideSchedule) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSlideSlideSchedule creates a new slide.slide_schedule model and returns its id.
func (c *Client) CreateSlideSlideSchedule(ss *SlideSlideSchedule) (int64, error) {
	return c.Create(SlideSlideScheduleModel, ss)
}

// UpdateSlideSlideSchedule updates an existing slide.slide_schedule record.
func (c *Client) UpdateSlideSlideSchedule(ss *SlideSlideSchedule) error {
	return c.UpdateSlideSlideSchedules([]int64{ss.Id.Get()}, ss)
}

// UpdateSlideSlideSchedules updates existing slide.slide_schedule records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSlideSlideSchedules(ids []int64, ss *SlideSlideSchedule) error {
	return c.Update(SlideSlideScheduleModel, ids, ss)
}

// DeleteSlideSlideSchedule deletes an existing slide.slide_schedule record.
func (c *Client) DeleteSlideSlideSchedule(id int64) error {
	return c.DeleteSlideSlideSchedules([]int64{id})
}

// DeleteSlideSlideSchedules deletes existing slide.slide_schedule records.
func (c *Client) DeleteSlideSlideSchedules(ids []int64) error {
	return c.Delete(SlideSlideScheduleModel, ids)
}

// GetSlideSlideSchedule gets slide.slide_schedule existing record.
func (c *Client) GetSlideSlideSchedule(id int64) (*SlideSlideSchedule, error) {
	sss, err := c.GetSlideSlideSchedules([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide_schedule not found", id)
}

// GetSlideSlideSchedules gets slide.slide_schedule existing records.
func (c *Client) GetSlideSlideSchedules(ids []int64) (*SlideSlideSchedules, error) {
	sss := &SlideSlideSchedules{}
	if err := c.Read(SlideSlideScheduleModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlideSchedule finds slide.slide_schedule record by querying it with criteria.
func (c *Client) FindSlideSlideSchedule(criteria *Criteria) (*SlideSlideSchedule, error) {
	sss := &SlideSlideSchedules{}
	if err := c.SearchRead(SlideSlideScheduleModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide_schedule was not found")
}

// FindSlideSlideSchedules finds slide.slide_schedule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideSchedules(criteria *Criteria, options *Options) (*SlideSlideSchedules, error) {
	sss := &SlideSlideSchedules{}
	if err := c.SearchRead(SlideSlideScheduleModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlideScheduleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideScheduleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideScheduleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlideScheduleId finds record id by querying it with criteria.
func (c *Client) FindSlideSlideScheduleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideScheduleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide_schedule was not found")
}
