package odoo

import (
	"fmt"
)

// MailingSmsTest represents mailing.sms.test model.
type MailingSmsTest struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	MailingId   *Many2One `xmlrpc:"mailing_id,omptempty"`
	Numbers     *String   `xmlrpc:"numbers,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailingSmsTests represents array of mailing.sms.test model.
type MailingSmsTests []MailingSmsTest

// MailingSmsTestModel is the odoo model name.
const MailingSmsTestModel = "mailing.sms.test"

// Many2One convert MailingSmsTest to *Many2One.
func (mst *MailingSmsTest) Many2One() *Many2One {
	return NewMany2One(mst.Id.Get(), "")
}

// CreateMailingSmsTest creates a new mailing.sms.test model and returns its id.
func (c *Client) CreateMailingSmsTest(mst *MailingSmsTest) (int64, error) {
	return c.Create(MailingSmsTestModel, mst)
}

// UpdateMailingSmsTest updates an existing mailing.sms.test record.
func (c *Client) UpdateMailingSmsTest(mst *MailingSmsTest) error {
	return c.UpdateMailingSmsTests([]int64{mst.Id.Get()}, mst)
}

// UpdateMailingSmsTests updates existing mailing.sms.test records.
// All records (represented by ids) will be updated by mst values.
func (c *Client) UpdateMailingSmsTests(ids []int64, mst *MailingSmsTest) error {
	return c.Update(MailingSmsTestModel, ids, mst)
}

// DeleteMailingSmsTest deletes an existing mailing.sms.test record.
func (c *Client) DeleteMailingSmsTest(id int64) error {
	return c.DeleteMailingSmsTests([]int64{id})
}

// DeleteMailingSmsTests deletes existing mailing.sms.test records.
func (c *Client) DeleteMailingSmsTests(ids []int64) error {
	return c.Delete(MailingSmsTestModel, ids)
}

// GetMailingSmsTest gets mailing.sms.test existing record.
func (c *Client) GetMailingSmsTest(id int64) (*MailingSmsTest, error) {
	msts, err := c.GetMailingSmsTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if msts != nil && len(*msts) > 0 {
		return &((*msts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.sms.test not found", id)
}

// GetMailingSmsTests gets mailing.sms.test existing records.
func (c *Client) GetMailingSmsTests(ids []int64) (*MailingSmsTests, error) {
	msts := &MailingSmsTests{}
	if err := c.Read(MailingSmsTestModel, ids, nil, msts); err != nil {
		return nil, err
	}
	return msts, nil
}

// FindMailingSmsTest finds mailing.sms.test record by querying it with criteria.
func (c *Client) FindMailingSmsTest(criteria *Criteria) (*MailingSmsTest, error) {
	msts := &MailingSmsTests{}
	if err := c.SearchRead(MailingSmsTestModel, criteria, NewOptions().Limit(1), msts); err != nil {
		return nil, err
	}
	if msts != nil && len(*msts) > 0 {
		return &((*msts)[0]), nil
	}
	return nil, fmt.Errorf("mailing.sms.test was not found")
}

// FindMailingSmsTests finds mailing.sms.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingSmsTests(criteria *Criteria, options *Options) (*MailingSmsTests, error) {
	msts := &MailingSmsTests{}
	if err := c.SearchRead(MailingSmsTestModel, criteria, options, msts); err != nil {
		return nil, err
	}
	return msts, nil
}

// FindMailingSmsTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingSmsTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingSmsTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingSmsTestId finds record id by querying it with criteria.
func (c *Client) FindMailingSmsTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingSmsTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.sms.test was not found")
}
