package odoo

import (
	"fmt"
)

// AccountInvoiceReport represents account.invoice.report model.
type AccountInvoiceReport struct {
	AccountId            *Many2One  `xmlrpc:"account_id,omptempty"`
	AmountTotal          *Float     `xmlrpc:"amount_total,omptempty"`
	AnalyticAccountId    *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CommercialPartnerId  *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId            *Many2One  `xmlrpc:"country_id,omptempty"`
	CurrencyId           *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionId     *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	InvoiceDate          *Time      `xmlrpc:"invoice_date,omptempty"`
	InvoiceDateDue       *Time      `xmlrpc:"invoice_date_due,omptempty"`
	InvoicePartnerBankId *Many2One  `xmlrpc:"invoice_partner_bank_id,omptempty"`
	InvoicePaymentState  *Selection `xmlrpc:"invoice_payment_state,omptempty"`
	InvoicePaymentTermId *Many2One  `xmlrpc:"invoice_payment_term_id,omptempty"`
	InvoiceUserId        *Many2One  `xmlrpc:"invoice_user_id,omptempty"`
	JournalId            *Many2One  `xmlrpc:"journal_id,omptempty"`
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	MoveId               *Many2One  `xmlrpc:"move_id,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	NbrLines             *Int       `xmlrpc:"nbr_lines,omptempty"`
	PartnerId            *Many2One  `xmlrpc:"partner_id,omptempty"`
	PriceAverage         *Float     `xmlrpc:"price_average,omptempty"`
	PriceSubtotal        *Float     `xmlrpc:"price_subtotal,omptempty"`
	ProductCategId       *Many2One  `xmlrpc:"product_categ_id,omptempty"`
	ProductId            *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductUomId         *Many2One  `xmlrpc:"product_uom_id,omptempty"`
	Quantity             *Float     `xmlrpc:"quantity,omptempty"`
	Residual             *Float     `xmlrpc:"residual,omptempty"`
	State                *Selection `xmlrpc:"state,omptempty"`
	TeamId               *Many2One  `xmlrpc:"team_id,omptempty"`
	Type                 *Selection `xmlrpc:"type,omptempty"`
}

// AccountInvoiceReports represents array of account.invoice.report model.
type AccountInvoiceReports []AccountInvoiceReport

// AccountInvoiceReportModel is the odoo model name.
const AccountInvoiceReportModel = "account.invoice.report"

// Many2One convert AccountInvoiceReport to *Many2One.
func (air *AccountInvoiceReport) Many2One() *Many2One {
	return NewMany2One(air.Id.Get(), "")
}

// CreateAccountInvoiceReport creates a new account.invoice.report model and returns its id.
func (c *Client) CreateAccountInvoiceReport(air *AccountInvoiceReport) (int64, error) {
	return c.Create(AccountInvoiceReportModel, air)
}

// UpdateAccountInvoiceReport updates an existing account.invoice.report record.
func (c *Client) UpdateAccountInvoiceReport(air *AccountInvoiceReport) error {
	return c.UpdateAccountInvoiceReports([]int64{air.Id.Get()}, air)
}

// UpdateAccountInvoiceReports updates existing account.invoice.report records.
// All records (represented by ids) will be updated by air values.
func (c *Client) UpdateAccountInvoiceReports(ids []int64, air *AccountInvoiceReport) error {
	return c.Update(AccountInvoiceReportModel, ids, air)
}

// DeleteAccountInvoiceReport deletes an existing account.invoice.report record.
func (c *Client) DeleteAccountInvoiceReport(id int64) error {
	return c.DeleteAccountInvoiceReports([]int64{id})
}

// DeleteAccountInvoiceReports deletes existing account.invoice.report records.
func (c *Client) DeleteAccountInvoiceReports(ids []int64) error {
	return c.Delete(AccountInvoiceReportModel, ids)
}

// GetAccountInvoiceReport gets account.invoice.report existing record.
func (c *Client) GetAccountInvoiceReport(id int64) (*AccountInvoiceReport, error) {
	airs, err := c.GetAccountInvoiceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if airs != nil && len(*airs) > 0 {
		return &((*airs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.invoice.report not found", id)
}

// GetAccountInvoiceReports gets account.invoice.report existing records.
func (c *Client) GetAccountInvoiceReports(ids []int64) (*AccountInvoiceReports, error) {
	airs := &AccountInvoiceReports{}
	if err := c.Read(AccountInvoiceReportModel, ids, nil, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceReport finds account.invoice.report record by querying it with criteria.
func (c *Client) FindAccountInvoiceReport(criteria *Criteria) (*AccountInvoiceReport, error) {
	airs := &AccountInvoiceReports{}
	if err := c.SearchRead(AccountInvoiceReportModel, criteria, NewOptions().Limit(1), airs); err != nil {
		return nil, err
	}
	if airs != nil && len(*airs) > 0 {
		return &((*airs)[0]), nil
	}
	return nil, fmt.Errorf("account.invoice.report was not found")
}

// FindAccountInvoiceReports finds account.invoice.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceReports(criteria *Criteria, options *Options) (*AccountInvoiceReports, error) {
	airs := &AccountInvoiceReports{}
	if err := c.SearchRead(AccountInvoiceReportModel, criteria, options, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountInvoiceReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountInvoiceReportId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.invoice.report was not found")
}
