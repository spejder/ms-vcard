package odoo

import (
	"fmt"
)

// AccountFiscalYear represents account.fiscal.year model.
type AccountFiscalYear struct {
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DateFrom    *Time     `xmlrpc:"date_from,omptempty"`
	DateTo      *Time     `xmlrpc:"date_to,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountFiscalYears represents array of account.fiscal.year model.
type AccountFiscalYears []AccountFiscalYear

// AccountFiscalYearModel is the odoo model name.
const AccountFiscalYearModel = "account.fiscal.year"

// Many2One convert AccountFiscalYear to *Many2One.
func (afy *AccountFiscalYear) Many2One() *Many2One {
	return NewMany2One(afy.Id.Get(), "")
}

// CreateAccountFiscalYear creates a new account.fiscal.year model and returns its id.
func (c *Client) CreateAccountFiscalYear(afy *AccountFiscalYear) (int64, error) {
	return c.Create(AccountFiscalYearModel, afy)
}

// UpdateAccountFiscalYear updates an existing account.fiscal.year record.
func (c *Client) UpdateAccountFiscalYear(afy *AccountFiscalYear) error {
	return c.UpdateAccountFiscalYears([]int64{afy.Id.Get()}, afy)
}

// UpdateAccountFiscalYears updates existing account.fiscal.year records.
// All records (represented by ids) will be updated by afy values.
func (c *Client) UpdateAccountFiscalYears(ids []int64, afy *AccountFiscalYear) error {
	return c.Update(AccountFiscalYearModel, ids, afy)
}

// DeleteAccountFiscalYear deletes an existing account.fiscal.year record.
func (c *Client) DeleteAccountFiscalYear(id int64) error {
	return c.DeleteAccountFiscalYears([]int64{id})
}

// DeleteAccountFiscalYears deletes existing account.fiscal.year records.
func (c *Client) DeleteAccountFiscalYears(ids []int64) error {
	return c.Delete(AccountFiscalYearModel, ids)
}

// GetAccountFiscalYear gets account.fiscal.year existing record.
func (c *Client) GetAccountFiscalYear(id int64) (*AccountFiscalYear, error) {
	afys, err := c.GetAccountFiscalYears([]int64{id})
	if err != nil {
		return nil, err
	}
	if afys != nil && len(*afys) > 0 {
		return &((*afys)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.fiscal.year not found", id)
}

// GetAccountFiscalYears gets account.fiscal.year existing records.
func (c *Client) GetAccountFiscalYears(ids []int64) (*AccountFiscalYears, error) {
	afys := &AccountFiscalYears{}
	if err := c.Read(AccountFiscalYearModel, ids, nil, afys); err != nil {
		return nil, err
	}
	return afys, nil
}

// FindAccountFiscalYear finds account.fiscal.year record by querying it with criteria.
func (c *Client) FindAccountFiscalYear(criteria *Criteria) (*AccountFiscalYear, error) {
	afys := &AccountFiscalYears{}
	if err := c.SearchRead(AccountFiscalYearModel, criteria, NewOptions().Limit(1), afys); err != nil {
		return nil, err
	}
	if afys != nil && len(*afys) > 0 {
		return &((*afys)[0]), nil
	}
	return nil, fmt.Errorf("account.fiscal.year was not found")
}

// FindAccountFiscalYears finds account.fiscal.year records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalYears(criteria *Criteria, options *Options) (*AccountFiscalYears, error) {
	afys := &AccountFiscalYears{}
	if err := c.SearchRead(AccountFiscalYearModel, criteria, options, afys); err != nil {
		return nil, err
	}
	return afys, nil
}

// FindAccountFiscalYearIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalYearIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFiscalYearModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFiscalYearId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalYearId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalYearModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.fiscal.year was not found")
}
