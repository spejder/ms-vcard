package odoo

import (
	"fmt"
)

// MailingListMerge represents mailing.list.merge model.
type MailingListMerge struct {
	ArchiveSrcLists *Bool      `xmlrpc:"archive_src_lists,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DestListId      *Many2One  `xmlrpc:"dest_list_id,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	MergeOptions    *Selection `xmlrpc:"merge_options,omptempty"`
	NewListName     *String    `xmlrpc:"new_list_name,omptempty"`
	SrcListIds      *Relation  `xmlrpc:"src_list_ids,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailingListMerges represents array of mailing.list.merge model.
type MailingListMerges []MailingListMerge

// MailingListMergeModel is the odoo model name.
const MailingListMergeModel = "mailing.list.merge"

// Many2One convert MailingListMerge to *Many2One.
func (mlm *MailingListMerge) Many2One() *Many2One {
	return NewMany2One(mlm.Id.Get(), "")
}

// CreateMailingListMerge creates a new mailing.list.merge model and returns its id.
func (c *Client) CreateMailingListMerge(mlm *MailingListMerge) (int64, error) {
	return c.Create(MailingListMergeModel, mlm)
}

// UpdateMailingListMerge updates an existing mailing.list.merge record.
func (c *Client) UpdateMailingListMerge(mlm *MailingListMerge) error {
	return c.UpdateMailingListMerges([]int64{mlm.Id.Get()}, mlm)
}

// UpdateMailingListMerges updates existing mailing.list.merge records.
// All records (represented by ids) will be updated by mlm values.
func (c *Client) UpdateMailingListMerges(ids []int64, mlm *MailingListMerge) error {
	return c.Update(MailingListMergeModel, ids, mlm)
}

// DeleteMailingListMerge deletes an existing mailing.list.merge record.
func (c *Client) DeleteMailingListMerge(id int64) error {
	return c.DeleteMailingListMerges([]int64{id})
}

// DeleteMailingListMerges deletes existing mailing.list.merge records.
func (c *Client) DeleteMailingListMerges(ids []int64) error {
	return c.Delete(MailingListMergeModel, ids)
}

// GetMailingListMerge gets mailing.list.merge existing record.
func (c *Client) GetMailingListMerge(id int64) (*MailingListMerge, error) {
	mlms, err := c.GetMailingListMerges([]int64{id})
	if err != nil {
		return nil, err
	}
	if mlms != nil && len(*mlms) > 0 {
		return &((*mlms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.list.merge not found", id)
}

// GetMailingListMerges gets mailing.list.merge existing records.
func (c *Client) GetMailingListMerges(ids []int64) (*MailingListMerges, error) {
	mlms := &MailingListMerges{}
	if err := c.Read(MailingListMergeModel, ids, nil, mlms); err != nil {
		return nil, err
	}
	return mlms, nil
}

// FindMailingListMerge finds mailing.list.merge record by querying it with criteria.
func (c *Client) FindMailingListMerge(criteria *Criteria) (*MailingListMerge, error) {
	mlms := &MailingListMerges{}
	if err := c.SearchRead(MailingListMergeModel, criteria, NewOptions().Limit(1), mlms); err != nil {
		return nil, err
	}
	if mlms != nil && len(*mlms) > 0 {
		return &((*mlms)[0]), nil
	}
	return nil, fmt.Errorf("mailing.list.merge was not found")
}

// FindMailingListMerges finds mailing.list.merge records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingListMerges(criteria *Criteria, options *Options) (*MailingListMerges, error) {
	mlms := &MailingListMerges{}
	if err := c.SearchRead(MailingListMergeModel, criteria, options, mlms); err != nil {
		return nil, err
	}
	return mlms, nil
}

// FindMailingListMergeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingListMergeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingListMergeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingListMergeId finds record id by querying it with criteria.
func (c *Client) FindMailingListMergeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingListMergeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.list.merge was not found")
}
