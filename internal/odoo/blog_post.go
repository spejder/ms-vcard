package odoo

import (
	"fmt"
)

// BlogPost represents blog.post model.
type BlogPost struct {
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	AuthorAvatar             *String   `xmlrpc:"author_avatar,omptempty"`
	AuthorId                 *Many2One `xmlrpc:"author_id,omptempty"`
	BlogId                   *Many2One `xmlrpc:"blog_id,omptempty"`
	CanPublish               *Bool     `xmlrpc:"can_publish,omptempty"`
	Content                  *String   `xmlrpc:"content,omptempty"`
	CoverProperties          *String   `xmlrpc:"cover_properties,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsPublished              *Bool     `xmlrpc:"is_published,omptempty"`
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
	PostDate                 *Time     `xmlrpc:"post_date,omptempty"`
	PublishedDate            *Time     `xmlrpc:"published_date,omptempty"`
	Subtitle                 *String   `xmlrpc:"subtitle,omptempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omptempty"`
	Teaser                   *String   `xmlrpc:"teaser,omptempty"`
	TeaserManual             *String   `xmlrpc:"teaser_manual,omptempty"`
	Visits                   *Int      `xmlrpc:"visits,omptempty"`
	WebsiteId                *Many2One `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription   *String   `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords      *String   `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg         *String   `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle         *String   `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished         *Bool     `xmlrpc:"website_published,omptempty"`
	WebsiteUrl               *String   `xmlrpc:"website_url,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BlogPosts represents array of blog.post model.
type BlogPosts []BlogPost

// BlogPostModel is the odoo model name.
const BlogPostModel = "blog.post"

// Many2One convert BlogPost to *Many2One.
func (bp *BlogPost) Many2One() *Many2One {
	return NewMany2One(bp.Id.Get(), "")
}

// CreateBlogPost creates a new blog.post model and returns its id.
func (c *Client) CreateBlogPost(bp *BlogPost) (int64, error) {
	return c.Create(BlogPostModel, bp)
}

// UpdateBlogPost updates an existing blog.post record.
func (c *Client) UpdateBlogPost(bp *BlogPost) error {
	return c.UpdateBlogPosts([]int64{bp.Id.Get()}, bp)
}

// UpdateBlogPosts updates existing blog.post records.
// All records (represented by ids) will be updated by bp values.
func (c *Client) UpdateBlogPosts(ids []int64, bp *BlogPost) error {
	return c.Update(BlogPostModel, ids, bp)
}

// DeleteBlogPost deletes an existing blog.post record.
func (c *Client) DeleteBlogPost(id int64) error {
	return c.DeleteBlogPosts([]int64{id})
}

// DeleteBlogPosts deletes existing blog.post records.
func (c *Client) DeleteBlogPosts(ids []int64) error {
	return c.Delete(BlogPostModel, ids)
}

// GetBlogPost gets blog.post existing record.
func (c *Client) GetBlogPost(id int64) (*BlogPost, error) {
	bps, err := c.GetBlogPosts([]int64{id})
	if err != nil {
		return nil, err
	}
	if bps != nil && len(*bps) > 0 {
		return &((*bps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of blog.post not found", id)
}

// GetBlogPosts gets blog.post existing records.
func (c *Client) GetBlogPosts(ids []int64) (*BlogPosts, error) {
	bps := &BlogPosts{}
	if err := c.Read(BlogPostModel, ids, nil, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBlogPost finds blog.post record by querying it with criteria.
func (c *Client) FindBlogPost(criteria *Criteria) (*BlogPost, error) {
	bps := &BlogPosts{}
	if err := c.SearchRead(BlogPostModel, criteria, NewOptions().Limit(1), bps); err != nil {
		return nil, err
	}
	if bps != nil && len(*bps) > 0 {
		return &((*bps)[0]), nil
	}
	return nil, fmt.Errorf("blog.post was not found")
}

// FindBlogPosts finds blog.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogPosts(criteria *Criteria, options *Options) (*BlogPosts, error) {
	bps := &BlogPosts{}
	if err := c.SearchRead(BlogPostModel, criteria, options, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBlogPostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBlogPostIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BlogPostModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBlogPostId finds record id by querying it with criteria.
func (c *Client) FindBlogPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BlogPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("blog.post was not found")
}
