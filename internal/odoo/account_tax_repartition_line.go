package odoo

import (
	"fmt"
)

// AccountTaxRepartitionLine represents account.tax.repartition.line model.
type AccountTaxRepartitionLine struct {
	AccountId       *Many2One  `xmlrpc:"account_id,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId       *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Factor          *Float     `xmlrpc:"factor,omptempty"`
	FactorPercent   *Float     `xmlrpc:"factor_percent,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	InvoiceTaxId    *Many2One  `xmlrpc:"invoice_tax_id,omptempty"`
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	RefundTaxId     *Many2One  `xmlrpc:"refund_tax_id,omptempty"`
	RepartitionType *Selection `xmlrpc:"repartition_type,omptempty"`
	Sequence        *Int       `xmlrpc:"sequence,omptempty"`
	TagIds          *Relation  `xmlrpc:"tag_ids,omptempty"`
	TaxId           *Many2One  `xmlrpc:"tax_id,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountTaxRepartitionLines represents array of account.tax.repartition.line model.
type AccountTaxRepartitionLines []AccountTaxRepartitionLine

// AccountTaxRepartitionLineModel is the odoo model name.
const AccountTaxRepartitionLineModel = "account.tax.repartition.line"

// Many2One convert AccountTaxRepartitionLine to *Many2One.
func (atrl *AccountTaxRepartitionLine) Many2One() *Many2One {
	return NewMany2One(atrl.Id.Get(), "")
}

// CreateAccountTaxRepartitionLine creates a new account.tax.repartition.line model and returns its id.
func (c *Client) CreateAccountTaxRepartitionLine(atrl *AccountTaxRepartitionLine) (int64, error) {
	return c.Create(AccountTaxRepartitionLineModel, atrl)
}

// UpdateAccountTaxRepartitionLine updates an existing account.tax.repartition.line record.
func (c *Client) UpdateAccountTaxRepartitionLine(atrl *AccountTaxRepartitionLine) error {
	return c.UpdateAccountTaxRepartitionLines([]int64{atrl.Id.Get()}, atrl)
}

// UpdateAccountTaxRepartitionLines updates existing account.tax.repartition.line records.
// All records (represented by ids) will be updated by atrl values.
func (c *Client) UpdateAccountTaxRepartitionLines(ids []int64, atrl *AccountTaxRepartitionLine) error {
	return c.Update(AccountTaxRepartitionLineModel, ids, atrl)
}

// DeleteAccountTaxRepartitionLine deletes an existing account.tax.repartition.line record.
func (c *Client) DeleteAccountTaxRepartitionLine(id int64) error {
	return c.DeleteAccountTaxRepartitionLines([]int64{id})
}

// DeleteAccountTaxRepartitionLines deletes existing account.tax.repartition.line records.
func (c *Client) DeleteAccountTaxRepartitionLines(ids []int64) error {
	return c.Delete(AccountTaxRepartitionLineModel, ids)
}

// GetAccountTaxRepartitionLine gets account.tax.repartition.line existing record.
func (c *Client) GetAccountTaxRepartitionLine(id int64) (*AccountTaxRepartitionLine, error) {
	atrls, err := c.GetAccountTaxRepartitionLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if atrls != nil && len(*atrls) > 0 {
		return &((*atrls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax.repartition.line not found", id)
}

// GetAccountTaxRepartitionLines gets account.tax.repartition.line existing records.
func (c *Client) GetAccountTaxRepartitionLines(ids []int64) (*AccountTaxRepartitionLines, error) {
	atrls := &AccountTaxRepartitionLines{}
	if err := c.Read(AccountTaxRepartitionLineModel, ids, nil, atrls); err != nil {
		return nil, err
	}
	return atrls, nil
}

// FindAccountTaxRepartitionLine finds account.tax.repartition.line record by querying it with criteria.
func (c *Client) FindAccountTaxRepartitionLine(criteria *Criteria) (*AccountTaxRepartitionLine, error) {
	atrls := &AccountTaxRepartitionLines{}
	if err := c.SearchRead(AccountTaxRepartitionLineModel, criteria, NewOptions().Limit(1), atrls); err != nil {
		return nil, err
	}
	if atrls != nil && len(*atrls) > 0 {
		return &((*atrls)[0]), nil
	}
	return nil, fmt.Errorf("account.tax.repartition.line was not found")
}

// FindAccountTaxRepartitionLines finds account.tax.repartition.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxRepartitionLines(criteria *Criteria, options *Options) (*AccountTaxRepartitionLines, error) {
	atrls := &AccountTaxRepartitionLines{}
	if err := c.SearchRead(AccountTaxRepartitionLineModel, criteria, options, atrls); err != nil {
		return nil, err
	}
	return atrls, nil
}

// FindAccountTaxRepartitionLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxRepartitionLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxRepartitionLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxRepartitionLineId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxRepartitionLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxRepartitionLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tax.repartition.line was not found")
}
