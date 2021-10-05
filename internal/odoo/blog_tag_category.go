package odoo

import (
	"fmt"
)

// BlogTagCategory represents blog.tag.category model.
type BlogTagCategory struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	TagIds      *Relation `xmlrpc:"tag_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BlogTagCategorys represents array of blog.tag.category model.
type BlogTagCategorys []BlogTagCategory

// BlogTagCategoryModel is the odoo model name.
const BlogTagCategoryModel = "blog.tag.category"

// Many2One convert BlogTagCategory to *Many2One.
func (btc *BlogTagCategory) Many2One() *Many2One {
	return NewMany2One(btc.Id.Get(), "")
}

// CreateBlogTagCategory creates a new blog.tag.category model and returns its id.
func (c *Client) CreateBlogTagCategory(btc *BlogTagCategory) (int64, error) {
	return c.Create(BlogTagCategoryModel, btc)
}

// UpdateBlogTagCategory updates an existing blog.tag.category record.
func (c *Client) UpdateBlogTagCategory(btc *BlogTagCategory) error {
	return c.UpdateBlogTagCategorys([]int64{btc.Id.Get()}, btc)
}

// UpdateBlogTagCategorys updates existing blog.tag.category records.
// All records (represented by ids) will be updated by btc values.
func (c *Client) UpdateBlogTagCategorys(ids []int64, btc *BlogTagCategory) error {
	return c.Update(BlogTagCategoryModel, ids, btc)
}

// DeleteBlogTagCategory deletes an existing blog.tag.category record.
func (c *Client) DeleteBlogTagCategory(id int64) error {
	return c.DeleteBlogTagCategorys([]int64{id})
}

// DeleteBlogTagCategorys deletes existing blog.tag.category records.
func (c *Client) DeleteBlogTagCategorys(ids []int64) error {
	return c.Delete(BlogTagCategoryModel, ids)
}

// GetBlogTagCategory gets blog.tag.category existing record.
func (c *Client) GetBlogTagCategory(id int64) (*BlogTagCategory, error) {
	btcs, err := c.GetBlogTagCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if btcs != nil && len(*btcs) > 0 {
		return &((*btcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of blog.tag.category not found", id)
}

// GetBlogTagCategorys gets blog.tag.category existing records.
func (c *Client) GetBlogTagCategorys(ids []int64) (*BlogTagCategorys, error) {
	btcs := &BlogTagCategorys{}
	if err := c.Read(BlogTagCategoryModel, ids, nil, btcs); err != nil {
		return nil, err
	}
	return btcs, nil
}

// FindBlogTagCategory finds blog.tag.category record by querying it with criteria.
func (c *Client) FindBlogTagCategory(criteria *Criteria) (*BlogTagCategory, error) {
	btcs := &BlogTagCategorys{}
	if err := c.SearchRead(BlogTagCategoryModel, criteria, NewOptions().Limit(1), btcs); err != nil {
		return nil, err
	}
	if btcs != nil && len(*btcs) > 0 {
		return &((*btcs)[0]), nil
	}
	return nil, fmt.Errorf("blog.tag.category was not found")
}

// FindBlogTagCategorys finds blog.tag.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogTagCategorys(criteria *Criteria, options *Options) (*BlogTagCategorys, error) {
	btcs := &BlogTagCategorys{}
	if err := c.SearchRead(BlogTagCategoryModel, criteria, options, btcs); err != nil {
		return nil, err
	}
	return btcs, nil
}

// FindBlogTagCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogTagCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BlogTagCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBlogTagCategoryId finds record id by querying it with criteria.
func (c *Client) FindBlogTagCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BlogTagCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("blog.tag.category was not found")
}
