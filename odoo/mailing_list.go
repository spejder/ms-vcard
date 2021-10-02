package odoo

import (
	"fmt"
)

// MailingList represents mailing.list model.
type MailingList struct {
	Active          *Bool     `xmlrpc:"active,omptempty"`
	ContactIds      *Relation `xmlrpc:"contact_ids,omptempty"`
	ContactNbr      *Int      `xmlrpc:"contact_nbr,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	IsPublic        *Bool     `xmlrpc:"is_public,omptempty"`
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	SubscriptionIds *Relation `xmlrpc:"subscription_ids,omptempty"`
	ToastContent    *String   `xmlrpc:"toast_content,omptempty"`
	WebsitePopupIds *Relation `xmlrpc:"website_popup_ids,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailingLists represents array of mailing.list model.
type MailingLists []MailingList

// MailingListModel is the odoo model name.
const MailingListModel = "mailing.list"

// Many2One convert MailingList to *Many2One.
func (ml *MailingList) Many2One() *Many2One {
	return NewMany2One(ml.Id.Get(), "")
}

// CreateMailingList creates a new mailing.list model and returns its id.
func (c *Client) CreateMailingList(ml *MailingList) (int64, error) {
	return c.Create(MailingListModel, ml)
}

// UpdateMailingList updates an existing mailing.list record.
func (c *Client) UpdateMailingList(ml *MailingList) error {
	return c.UpdateMailingLists([]int64{ml.Id.Get()}, ml)
}

// UpdateMailingLists updates existing mailing.list records.
// All records (represented by ids) will be updated by ml values.
func (c *Client) UpdateMailingLists(ids []int64, ml *MailingList) error {
	return c.Update(MailingListModel, ids, ml)
}

// DeleteMailingList deletes an existing mailing.list record.
func (c *Client) DeleteMailingList(id int64) error {
	return c.DeleteMailingLists([]int64{id})
}

// DeleteMailingLists deletes existing mailing.list records.
func (c *Client) DeleteMailingLists(ids []int64) error {
	return c.Delete(MailingListModel, ids)
}

// GetMailingList gets mailing.list existing record.
func (c *Client) GetMailingList(id int64) (*MailingList, error) {
	mls, err := c.GetMailingLists([]int64{id})
	if err != nil {
		return nil, err
	}
	if mls != nil && len(*mls) > 0 {
		return &((*mls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.list not found", id)
}

// GetMailingLists gets mailing.list existing records.
func (c *Client) GetMailingLists(ids []int64) (*MailingLists, error) {
	mls := &MailingLists{}
	if err := c.Read(MailingListModel, ids, nil, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMailingList finds mailing.list record by querying it with criteria.
func (c *Client) FindMailingList(criteria *Criteria) (*MailingList, error) {
	mls := &MailingLists{}
	if err := c.SearchRead(MailingListModel, criteria, NewOptions().Limit(1), mls); err != nil {
		return nil, err
	}
	if mls != nil && len(*mls) > 0 {
		return &((*mls)[0]), nil
	}
	return nil, fmt.Errorf("mailing.list was not found")
}

// FindMailingLists finds mailing.list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingLists(criteria *Criteria, options *Options) (*MailingLists, error) {
	mls := &MailingLists{}
	if err := c.SearchRead(MailingListModel, criteria, options, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMailingListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingListIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingListModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingListId finds record id by querying it with criteria.
func (c *Client) FindMailingListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.list was not found")
}
