package odoo

import (
	"fmt"
)

// BlogBlog represents blog.blog model.
type BlogBlog struct {
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	Content                  *String   `xmlrpc:"content,omptempty"`
	CoverProperties          *String   `xmlrpc:"cover_properties,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsSeoOptimized           *Bool     `xmlrpc:"is_seo_optimized,omptempty"`
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	Subtitle                 *String   `xmlrpc:"subtitle,omptempty"`
	WebsiteId                *Many2One `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription   *String   `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords      *String   `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg         *String   `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle         *String   `xmlrpc:"website_meta_title,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BlogBlogs represents array of blog.blog model.
type BlogBlogs []BlogBlog

// BlogBlogModel is the odoo model name.
const BlogBlogModel = "blog.blog"

// Many2One convert BlogBlog to *Many2One.
func (bb *BlogBlog) Many2One() *Many2One {
	return NewMany2One(bb.Id.Get(), "")
}

// CreateBlogBlog creates a new blog.blog model and returns its id.
func (c *Client) CreateBlogBlog(bb *BlogBlog) (int64, error) {
	return c.Create(BlogBlogModel, bb)
}

// UpdateBlogBlog updates an existing blog.blog record.
func (c *Client) UpdateBlogBlog(bb *BlogBlog) error {
	return c.UpdateBlogBlogs([]int64{bb.Id.Get()}, bb)
}

// UpdateBlogBlogs updates existing blog.blog records.
// All records (represented by ids) will be updated by bb values.
func (c *Client) UpdateBlogBlogs(ids []int64, bb *BlogBlog) error {
	return c.Update(BlogBlogModel, ids, bb)
}

// DeleteBlogBlog deletes an existing blog.blog record.
func (c *Client) DeleteBlogBlog(id int64) error {
	return c.DeleteBlogBlogs([]int64{id})
}

// DeleteBlogBlogs deletes existing blog.blog records.
func (c *Client) DeleteBlogBlogs(ids []int64) error {
	return c.Delete(BlogBlogModel, ids)
}

// GetBlogBlog gets blog.blog existing record.
func (c *Client) GetBlogBlog(id int64) (*BlogBlog, error) {
	bbs, err := c.GetBlogBlogs([]int64{id})
	if err != nil {
		return nil, err
	}
	if bbs != nil && len(*bbs) > 0 {
		return &((*bbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of blog.blog not found", id)
}

// GetBlogBlogs gets blog.blog existing records.
func (c *Client) GetBlogBlogs(ids []int64) (*BlogBlogs, error) {
	bbs := &BlogBlogs{}
	if err := c.Read(BlogBlogModel, ids, nil, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBlogBlog finds blog.blog record by querying it with criteria.
func (c *Client) FindBlogBlog(criteria *Criteria) (*BlogBlog, error) {
	bbs := &BlogBlogs{}
	if err := c.SearchRead(BlogBlogModel, criteria, NewOptions().Limit(1), bbs); err != nil {
		return nil, err
	}
	if bbs != nil && len(*bbs) > 0 {
		return &((*bbs)[0]), nil
	}
	return nil, fmt.Errorf("blog.blog was not found")
}

// FindBlogBlogs finds blog.blog records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogBlogs(criteria *Criteria, options *Options) (*BlogBlogs, error) {
	bbs := &BlogBlogs{}
	if err := c.SearchRead(BlogBlogModel, criteria, options, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBlogBlogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogBlogIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BlogBlogModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBlogBlogId finds record id by querying it with criteria.
func (c *Client) FindBlogBlogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BlogBlogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("blog.blog was not found")
}
