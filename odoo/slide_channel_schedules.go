package odoo

import (
	"fmt"
)

// SlideChannelSchedules represents slide.channel_schedules model.
type SlideChannelSchedules struct {
	Active       *Bool     `xmlrpc:"active,omptempty"`
	AutoReview   *Bool     `xmlrpc:"auto_review,omptempty"`
	ChannelId    *Many2One `xmlrpc:"channel_id,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt    *Time     `xmlrpc:"deleted_at,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	EndDate      *String   `xmlrpc:"end_date,omptempty"`
	EndTime      *String   `xmlrpc:"end_time,omptempty"`
	HourBreak    *Int      `xmlrpc:"hour_break,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	MaxSize      *Int      `xmlrpc:"max_size,omptempty"`
	ScheduleType *String   `xmlrpc:"schedule_type,omptempty"`
	StartDate    *String   `xmlrpc:"start_date,omptempty"`
	StartTime    *String   `xmlrpc:"start_time,omptempty"`
	TotalDays    *Int      `xmlrpc:"total_days,omptempty"`
	Venue        *String   `xmlrpc:"venue,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlideChannelScheduless represents array of slide.channel_schedules model.
type SlideChannelScheduless []SlideChannelSchedules

// SlideChannelSchedulesModel is the odoo model name.
const SlideChannelSchedulesModel = "slide.channel_schedules"

// Many2One convert SlideChannelSchedules to *Many2One.
func (sc *SlideChannelSchedules) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSlideChannelSchedules creates a new slide.channel_schedules model and returns its id.
func (c *Client) CreateSlideChannelSchedules(sc *SlideChannelSchedules) (int64, error) {
	return c.Create(SlideChannelSchedulesModel, sc)
}

// UpdateSlideChannelSchedules updates an existing slide.channel_schedules record.
func (c *Client) UpdateSlideChannelSchedules(sc *SlideChannelSchedules) error {
	return c.UpdateSlideChannelScheduless([]int64{sc.Id.Get()}, sc)
}

// UpdateSlideChannelScheduless updates existing slide.channel_schedules records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSlideChannelScheduless(ids []int64, sc *SlideChannelSchedules) error {
	return c.Update(SlideChannelSchedulesModel, ids, sc)
}

// DeleteSlideChannelSchedules deletes an existing slide.channel_schedules record.
func (c *Client) DeleteSlideChannelSchedules(id int64) error {
	return c.DeleteSlideChannelScheduless([]int64{id})
}

// DeleteSlideChannelScheduless deletes existing slide.channel_schedules records.
func (c *Client) DeleteSlideChannelScheduless(ids []int64) error {
	return c.Delete(SlideChannelSchedulesModel, ids)
}

// GetSlideChannelSchedules gets slide.channel_schedules existing record.
func (c *Client) GetSlideChannelSchedules(id int64) (*SlideChannelSchedules, error) {
	scs, err := c.GetSlideChannelScheduless([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel_schedules not found", id)
}

// GetSlideChannelScheduless gets slide.channel_schedules existing records.
func (c *Client) GetSlideChannelScheduless(ids []int64) (*SlideChannelScheduless, error) {
	scs := &SlideChannelScheduless{}
	if err := c.Read(SlideChannelSchedulesModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelSchedules finds slide.channel_schedules record by querying it with criteria.
func (c *Client) FindSlideChannelSchedules(criteria *Criteria) (*SlideChannelSchedules, error) {
	scs := &SlideChannelScheduless{}
	if err := c.SearchRead(SlideChannelSchedulesModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel_schedules was not found")
}

// FindSlideChannelScheduless finds slide.channel_schedules records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelScheduless(criteria *Criteria, options *Options) (*SlideChannelScheduless, error) {
	scs := &SlideChannelScheduless{}
	if err := c.SearchRead(SlideChannelSchedulesModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelSchedulesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelSchedulesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelSchedulesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelSchedulesId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelSchedulesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelSchedulesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel_schedules was not found")
}
