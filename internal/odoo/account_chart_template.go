package odoo

import (
	"fmt"
)

// AccountChartTemplate represents account.chart.template model.
type AccountChartTemplate struct {
	AccountIds                            *Relation `xmlrpc:"account_ids,omptempty"`
	BankAccountCodePrefix                 *String   `xmlrpc:"bank_account_code_prefix,omptempty"`
	CashAccountCodePrefix                 *String   `xmlrpc:"cash_account_code_prefix,omptempty"`
	CodeDigits                            *Int      `xmlrpc:"code_digits,omptempty"`
	CompleteTaxSet                        *Bool     `xmlrpc:"complete_tax_set,omptempty"`
	CreateDate                            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId                            *Many2One `xmlrpc:"currency_id,omptempty"`
	DefaultCashDifferenceExpenseAccountId *Many2One `xmlrpc:"default_cash_difference_expense_account_id,omptempty"`
	DefaultCashDifferenceIncomeAccountId  *Many2One `xmlrpc:"default_cash_difference_income_account_id,omptempty"`
	DefaultPosReceivableAccountId         *Many2One `xmlrpc:"default_pos_receivable_account_id,omptempty"`
	DisplayName                           *String   `xmlrpc:"display_name,omptempty"`
	ExpenseCurrencyExchangeAccountId      *Many2One `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	Id                                    *Int      `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId       *Many2One `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	LastUpdate                            *Time     `xmlrpc:"__last_update,omptempty"`
	Name                                  *String   `xmlrpc:"name,omptempty"`
	ParentId                              *Many2One `xmlrpc:"parent_id,omptempty"`
	PropertyAccountExpenseCategId         *Many2One `xmlrpc:"property_account_expense_categ_id,omptempty"`
	PropertyAccountExpenseId              *Many2One `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeCategId          *Many2One `xmlrpc:"property_account_income_categ_id,omptempty"`
	PropertyAccountIncomeId               *Many2One `xmlrpc:"property_account_income_id,omptempty"`
	PropertyAccountPayableId              *Many2One `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountReceivableId           *Many2One `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyAdvanceTaxPaymentAccountId    *Many2One `xmlrpc:"property_advance_tax_payment_account_id,omptempty"`
	PropertyStockAccountInputCategId      *Many2One `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId     *Many2One `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockValuationAccountId       *Many2One `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	PropertyTaxPayableAccountId           *Many2One `xmlrpc:"property_tax_payable_account_id,omptempty"`
	PropertyTaxReceivableAccountId        *Many2One `xmlrpc:"property_tax_receivable_account_id,omptempty"`
	TaxTemplateIds                        *Relation `xmlrpc:"tax_template_ids,omptempty"`
	TransferAccountCodePrefix             *String   `xmlrpc:"transfer_account_code_prefix,omptempty"`
	UseAngloSaxon                         *Bool     `xmlrpc:"use_anglo_saxon,omptempty"`
	Visible                               *Bool     `xmlrpc:"visible,omptempty"`
	WriteDate                             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountChartTemplates represents array of account.chart.template model.
type AccountChartTemplates []AccountChartTemplate

// AccountChartTemplateModel is the odoo model name.
const AccountChartTemplateModel = "account.chart.template"

// Many2One convert AccountChartTemplate to *Many2One.
func (act *AccountChartTemplate) Many2One() *Many2One {
	return NewMany2One(act.Id.Get(), "")
}

// CreateAccountChartTemplate creates a new account.chart.template model and returns its id.
func (c *Client) CreateAccountChartTemplate(act *AccountChartTemplate) (int64, error) {
	return c.Create(AccountChartTemplateModel, act)
}

// UpdateAccountChartTemplate updates an existing account.chart.template record.
func (c *Client) UpdateAccountChartTemplate(act *AccountChartTemplate) error {
	return c.UpdateAccountChartTemplates([]int64{act.Id.Get()}, act)
}

// UpdateAccountChartTemplates updates existing account.chart.template records.
// All records (represented by ids) will be updated by act values.
func (c *Client) UpdateAccountChartTemplates(ids []int64, act *AccountChartTemplate) error {
	return c.Update(AccountChartTemplateModel, ids, act)
}

// DeleteAccountChartTemplate deletes an existing account.chart.template record.
func (c *Client) DeleteAccountChartTemplate(id int64) error {
	return c.DeleteAccountChartTemplates([]int64{id})
}

// DeleteAccountChartTemplates deletes existing account.chart.template records.
func (c *Client) DeleteAccountChartTemplates(ids []int64) error {
	return c.Delete(AccountChartTemplateModel, ids)
}

// GetAccountChartTemplate gets account.chart.template existing record.
func (c *Client) GetAccountChartTemplate(id int64) (*AccountChartTemplate, error) {
	acts, err := c.GetAccountChartTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if acts != nil && len(*acts) > 0 {
		return &((*acts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.chart.template not found", id)
}

// GetAccountChartTemplates gets account.chart.template existing records.
func (c *Client) GetAccountChartTemplates(ids []int64) (*AccountChartTemplates, error) {
	acts := &AccountChartTemplates{}
	if err := c.Read(AccountChartTemplateModel, ids, nil, acts); err != nil {
		return nil, err
	}
	return acts, nil
}

// FindAccountChartTemplate finds account.chart.template record by querying it with criteria.
func (c *Client) FindAccountChartTemplate(criteria *Criteria) (*AccountChartTemplate, error) {
	acts := &AccountChartTemplates{}
	if err := c.SearchRead(AccountChartTemplateModel, criteria, NewOptions().Limit(1), acts); err != nil {
		return nil, err
	}
	if acts != nil && len(*acts) > 0 {
		return &((*acts)[0]), nil
	}
	return nil, fmt.Errorf("account.chart.template was not found")
}

// FindAccountChartTemplates finds account.chart.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChartTemplates(criteria *Criteria, options *Options) (*AccountChartTemplates, error) {
	acts := &AccountChartTemplates{}
	if err := c.SearchRead(AccountChartTemplateModel, criteria, options, acts); err != nil {
		return nil, err
	}
	return acts, nil
}

// FindAccountChartTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChartTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountChartTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountChartTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountChartTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountChartTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.chart.template was not found")
}
