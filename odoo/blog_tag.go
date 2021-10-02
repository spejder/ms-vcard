package odoo

import (
	"fmt"
)

// BlogTag represents blog.tag model.
type BlogTag struct {
	CategoryId             *Many2One `xmlrpc:"category_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	IsSeoOptimized         *Bool     `xmlrpc:"is_seo_optimized,omptempty"`
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	Name                   *String   `xmlrpc:"name,omptempty"`
	PostIds                *Relation `xmlrpc:"post_ids,omptempty"`
	WebsiteMetaDescription *String   `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords    *String   `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg       *String   `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle       *String   `xmlrpc:"website_meta_title,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BlogTags represents array of blog.tag model.
type BlogTags []BlogTag

// BlogTagModel is the odoo model name.
const BlogTagModel = "blog.tag"

// Many2One convert BlogTag to *Many2One.
func (bt *BlogTag) Many2One() *Many2One {
	return NewMany2One(bt.Id.Get(), "")
}

// CreateBlogTag creates a new blog.tag model and returns its id.
func (c *Client) CreateBlogTag(bt *BlogTag) (int64, error) {
	return c.Create(BlogTagModel, bt)
}

// UpdateBlogTag updates an existing blog.tag record.
func (c *Client) UpdateBlogTag(bt *BlogTag) error {
	return c.UpdateBlogTags([]int64{bt.Id.Get()}, bt)
}

// UpdateBlogTags updates existing blog.tag records.
// All records (represented by ids) will be updated by bt values.
func (c *Client) UpdateBlogTags(ids []int64, bt *BlogTag) error {
	return c.Update(BlogTagModel, ids, bt)
}

// DeleteBlogTag deletes an existing blog.tag record.
func (c *Client) DeleteBlogTag(id int64) error {
	return c.DeleteBlogTags([]int64{id})
}

// DeleteBlogTags deletes existing blog.tag records.
func (c *Client) DeleteBlogTags(ids []int64) error {
	return c.Delete(BlogTagModel, ids)
}

// GetBlogTag gets blog.tag existing record.
func (c *Client) GetBlogTag(id int64) (*BlogTag, error) {
	bts, err := c.GetBlogTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if bts != nil && len(*bts) > 0 {
		return &((*bts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of blog.tag not found", id)
}

// GetBlogTags gets blog.tag existing records.
func (c *Client) GetBlogTags(ids []int64) (*BlogTags, error) {
	bts := &BlogTags{}
	if err := c.Read(BlogTagModel, ids, nil, bts); err != nil {
		return nil, err
	}
	return bts, nil
}

// FindBlogTag finds blog.tag record by querying it with criteria.
func (c *Client) FindBlogTag(criteria *Criteria) (*BlogTag, error) {
	bts := &BlogTags{}
	if err := c.SearchRead(BlogTagModel, criteria, NewOptions().Limit(1), bts); err != nil {
		return nil, err
	}
	if bts != nil && len(*bts) > 0 {
		return &((*bts)[0]), nil
	}
	return nil, fmt.Errorf("blog.tag was not found")
}

// FindBlogTags finds blog.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogTags(criteria *Criteria, options *Options) (*BlogTags, error) {
	bts := &BlogTags{}
	if err := c.SearchRead(BlogTagModel, criteria, options, bts); err != nil {
		return nil, err
	}
	return bts, nil
}

// FindBlogTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BlogTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBlogTagId finds record id by querying it with criteria.
func (c *Client) FindBlogTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BlogTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("blog.tag was not found")
}
