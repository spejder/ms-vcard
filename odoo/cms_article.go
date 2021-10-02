package odoo

import (
	"fmt"
)

// CmsArticle represents cms.article model.
type CmsArticle struct {
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DeletedAt        *Time     `xmlrpc:"deleted_at,omptempty"`
	Description      *String   `xmlrpc:"description,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Image            *String   `xmlrpc:"image,omptempty"`
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	ShortDescription *String   `xmlrpc:"short_description,omptempty"`
	Status           *Bool     `xmlrpc:"status,omptempty"`
	Title            *String   `xmlrpc:"title,omptempty"`
	Slug             *String   `xmlrpc:"slug,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CmsArticles represents array of cms.article model.
type CmsArticles []CmsArticle

// CmsArticleModel is the odoo model name.
const CmsArticleModel = "cms.article"

// Many2One convert CmsArticle to *Many2One.
func (ca *CmsArticle) Many2One() *Many2One {
	return NewMany2One(ca.Id.Get(), "")
}

// CreateCmsArticle creates a new cms.article model and returns its id.
func (c *Client) CreateCmsArticle(ca *CmsArticle) (int64, error) {
	return c.Create(CmsArticleModel, ca)
}

// UpdateCmsArticle updates an existing cms.article record.
func (c *Client) UpdateCmsArticle(ca *CmsArticle) error {
	return c.UpdateCmsArticles([]int64{ca.Id.Get()}, ca)
}

// UpdateCmsArticles updates existing cms.article records.
// All records (represented by ids) will be updated by ca values.
func (c *Client) UpdateCmsArticles(ids []int64, ca *CmsArticle) error {
	return c.Update(CmsArticleModel, ids, ca)
}

// DeleteCmsArticle deletes an existing cms.article record.
func (c *Client) DeleteCmsArticle(id int64) error {
	return c.DeleteCmsArticles([]int64{id})
}

// DeleteCmsArticles deletes existing cms.article records.
func (c *Client) DeleteCmsArticles(ids []int64) error {
	return c.Delete(CmsArticleModel, ids)
}

// GetCmsArticle gets cms.article existing record.
func (c *Client) GetCmsArticle(id int64) (*CmsArticle, error) {
	cas, err := c.GetCmsArticles([]int64{id})
	if err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cms.article not found", id)
}

// GetCmsArticles gets cms.article existing records.
func (c *Client) GetCmsArticles(ids []int64) (*CmsArticles, error) {
	cas := &CmsArticles{}
	if err := c.Read(CmsArticleModel, ids, nil, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCmsArticle finds cms.article record by querying it with criteria.
func (c *Client) FindCmsArticle(criteria *Criteria) (*CmsArticle, error) {
	cas := &CmsArticles{}
	if err := c.SearchRead(CmsArticleModel, criteria, NewOptions().Limit(1), cas); err != nil {
		return nil, err
	}
	if cas != nil && len(*cas) > 0 {
		return &((*cas)[0]), nil
	}
	return nil, fmt.Errorf("cms.article was not found")
}

// FindCmsArticles finds cms.article records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCmsArticles(criteria *Criteria, options *Options) (*CmsArticles, error) {
	cas := &CmsArticles{}
	if err := c.SearchRead(CmsArticleModel, criteria, options, cas); err != nil {
		return nil, err
	}
	return cas, nil
}

// FindCmsArticleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCmsArticleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CmsArticleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCmsArticleId finds record id by querying it with criteria.
func (c *Client) FindCmsArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CmsArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cms.article was not found")
}
