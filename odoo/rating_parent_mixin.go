package odoo

import (
	"fmt"
)

// RatingParentMixin represents rating.parent.mixin model.
type RatingParentMixin struct {
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	RatingIds                    *Relation `xmlrpc:"rating_ids,omptempty"`
	RatingPercentageSatisfaction *Int      `xmlrpc:"rating_percentage_satisfaction,omptempty"`
}

// RatingParentMixins represents array of rating.parent.mixin model.
type RatingParentMixins []RatingParentMixin

// RatingParentMixinModel is the odoo model name.
const RatingParentMixinModel = "rating.parent.mixin"

// Many2One convert RatingParentMixin to *Many2One.
func (rpm *RatingParentMixin) Many2One() *Many2One {
	return NewMany2One(rpm.Id.Get(), "")
}

// CreateRatingParentMixin creates a new rating.parent.mixin model and returns its id.
func (c *Client) CreateRatingParentMixin(rpm *RatingParentMixin) (int64, error) {
	return c.Create(RatingParentMixinModel, rpm)
}

// UpdateRatingParentMixin updates an existing rating.parent.mixin record.
func (c *Client) UpdateRatingParentMixin(rpm *RatingParentMixin) error {
	return c.UpdateRatingParentMixins([]int64{rpm.Id.Get()}, rpm)
}

// UpdateRatingParentMixins updates existing rating.parent.mixin records.
// All records (represented by ids) will be updated by rpm values.
func (c *Client) UpdateRatingParentMixins(ids []int64, rpm *RatingParentMixin) error {
	return c.Update(RatingParentMixinModel, ids, rpm)
}

// DeleteRatingParentMixin deletes an existing rating.parent.mixin record.
func (c *Client) DeleteRatingParentMixin(id int64) error {
	return c.DeleteRatingParentMixins([]int64{id})
}

// DeleteRatingParentMixins deletes existing rating.parent.mixin records.
func (c *Client) DeleteRatingParentMixins(ids []int64) error {
	return c.Delete(RatingParentMixinModel, ids)
}

// GetRatingParentMixin gets rating.parent.mixin existing record.
func (c *Client) GetRatingParentMixin(id int64) (*RatingParentMixin, error) {
	rpms, err := c.GetRatingParentMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpms != nil && len(*rpms) > 0 {
		return &((*rpms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of rating.parent.mixin not found", id)
}

// GetRatingParentMixins gets rating.parent.mixin existing records.
func (c *Client) GetRatingParentMixins(ids []int64) (*RatingParentMixins, error) {
	rpms := &RatingParentMixins{}
	if err := c.Read(RatingParentMixinModel, ids, nil, rpms); err != nil {
		return nil, err
	}
	return rpms, nil
}

// FindRatingParentMixin finds rating.parent.mixin record by querying it with criteria.
func (c *Client) FindRatingParentMixin(criteria *Criteria) (*RatingParentMixin, error) {
	rpms := &RatingParentMixins{}
	if err := c.SearchRead(RatingParentMixinModel, criteria, NewOptions().Limit(1), rpms); err != nil {
		return nil, err
	}
	if rpms != nil && len(*rpms) > 0 {
		return &((*rpms)[0]), nil
	}
	return nil, fmt.Errorf("rating.parent.mixin was not found")
}

// FindRatingParentMixins finds rating.parent.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingParentMixins(criteria *Criteria, options *Options) (*RatingParentMixins, error) {
	rpms := &RatingParentMixins{}
	if err := c.SearchRead(RatingParentMixinModel, criteria, options, rpms); err != nil {
		return nil, err
	}
	return rpms, nil
}

// FindRatingParentMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingParentMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RatingParentMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRatingParentMixinId finds record id by querying it with criteria.
func (c *Client) FindRatingParentMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RatingParentMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("rating.parent.mixin was not found")
}
