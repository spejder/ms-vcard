package odoo

import (
	"fmt"
)

// UserPayment represents user.payment model.
type UserPayment struct {
	Attachment     *String   `xmlrpc:"attachment,omptempty"`
	AttachmentName *String   `xmlrpc:"attachment_name,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// UserPayments represents array of user.payment model.
type UserPayments []UserPayment

// UserPaymentModel is the odoo model name.
const UserPaymentModel = "user.payment"

// Many2One convert UserPayment to *Many2One.
func (up *UserPayment) Many2One() *Many2One {
	return NewMany2One(up.Id.Get(), "")
}

// CreateUserPayment creates a new user.payment model and returns its id.
func (c *Client) CreateUserPayment(up *UserPayment) (int64, error) {
	return c.Create(UserPaymentModel, up)
}

// UpdateUserPayment updates an existing user.payment record.
func (c *Client) UpdateUserPayment(up *UserPayment) error {
	return c.UpdateUserPayments([]int64{up.Id.Get()}, up)
}

// UpdateUserPayments updates existing user.payment records.
// All records (represented by ids) will be updated by up values.
func (c *Client) UpdateUserPayments(ids []int64, up *UserPayment) error {
	return c.Update(UserPaymentModel, ids, up)
}

// DeleteUserPayment deletes an existing user.payment record.
func (c *Client) DeleteUserPayment(id int64) error {
	return c.DeleteUserPayments([]int64{id})
}

// DeleteUserPayments deletes existing user.payment records.
func (c *Client) DeleteUserPayments(ids []int64) error {
	return c.Delete(UserPaymentModel, ids)
}

// GetUserPayment gets user.payment existing record.
func (c *Client) GetUserPayment(id int64) (*UserPayment, error) {
	ups, err := c.GetUserPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if ups != nil && len(*ups) > 0 {
		return &((*ups)[0]), nil
	}
	return nil, fmt.Errorf("id %v of user.payment not found", id)
}

// GetUserPayments gets user.payment existing records.
func (c *Client) GetUserPayments(ids []int64) (*UserPayments, error) {
	ups := &UserPayments{}
	if err := c.Read(UserPaymentModel, ids, nil, ups); err != nil {
		return nil, err
	}
	return ups, nil
}

// FindUserPayment finds user.payment record by querying it with criteria.
func (c *Client) FindUserPayment(criteria *Criteria) (*UserPayment, error) {
	ups := &UserPayments{}
	if err := c.SearchRead(UserPaymentModel, criteria, NewOptions().Limit(1), ups); err != nil {
		return nil, err
	}
	if ups != nil && len(*ups) > 0 {
		return &((*ups)[0]), nil
	}
	return nil, fmt.Errorf("user.payment was not found")
}

// FindUserPayments finds user.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUserPayments(criteria *Criteria, options *Options) (*UserPayments, error) {
	ups := &UserPayments{}
	if err := c.SearchRead(UserPaymentModel, criteria, options, ups); err != nil {
		return nil, err
	}
	return ups, nil
}

// FindUserPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUserPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UserPaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUserPaymentId finds record id by querying it with criteria.
func (c *Client) FindUserPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UserPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("user.payment was not found")
}
