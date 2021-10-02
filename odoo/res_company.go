package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	AccountBankReconciliationStart        *Time      `xmlrpc:"account_bank_reconciliation_start,omptempty"`
	AccountDashboardOnboardingState       *Selection `xmlrpc:"account_dashboard_onboarding_state,omptempty"`
	AccountDefaultPosReceivableAccountId  *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omptempty"`
	AccountId                             *String    `xmlrpc:"account_id,omptempty"`
	AccountInvoiceOnboardingState         *Selection `xmlrpc:"account_invoice_onboarding_state,omptempty"`
	AccountLinkUrl                        *String    `xmlrpc:"account_link_url,omptempty"`
	AccountNo                             *String    `xmlrpc:"account_no,omptempty"`
	AccountOnboardingInvoiceLayoutState   *Selection `xmlrpc:"account_onboarding_invoice_layout_state,omptempty"`
	AccountOnboardingSaleTaxState         *Selection `xmlrpc:"account_onboarding_sale_tax_state,omptempty"`
	AccountOnboardingSampleInvoiceState   *Selection `xmlrpc:"account_onboarding_sample_invoice_state,omptempty"`
	AccountOpeningDate                    *Time      `xmlrpc:"account_opening_date,omptempty"`
	AccountOpeningJournalId               *Many2One  `xmlrpc:"account_opening_journal_id,omptempty"`
	AccountOpeningMoveId                  *Many2One  `xmlrpc:"account_opening_move_id,omptempty"`
	AccountPurchaseTaxId                  *Many2One  `xmlrpc:"account_purchase_tax_id,omptempty"`
	AccountSaleTaxId                      *Many2One  `xmlrpc:"account_sale_tax_id,omptempty"`
	AccountSetupBankDataState             *Selection `xmlrpc:"account_setup_bank_data_state,omptempty"`
	AccountSetupCoaState                  *Selection `xmlrpc:"account_setup_coa_state,omptempty"`
	AccountSetupFyDataState               *Selection `xmlrpc:"account_setup_fy_data_state,omptempty"`
	AccrualDefaultJournalId               *Many2One  `xmlrpc:"accrual_default_journal_id,omptempty"`
	AngloSaxonAccounting                  *Bool      `xmlrpc:"anglo_saxon_accounting,omptempty"`
	BankAccountCodePrefix                 *String    `xmlrpc:"bank_account_code_prefix,omptempty"`
	BankIds                               *Relation  `xmlrpc:"bank_ids,omptempty"`
	BankJournalIds                        *Relation  `xmlrpc:"bank_journal_ids,omptempty"`
	BaseOnboardingCompanyState            *Selection `xmlrpc:"base_onboarding_company_state,omptempty"`
	CashAccountCodePrefix                 *String    `xmlrpc:"cash_account_code_prefix,omptempty"`
	Catchall                              *String    `xmlrpc:"catchall,omptempty"`
	ChargesEnabled                        *Bool      `xmlrpc:"charges_enabled,omptempty"`
	ChartTemplateId                       *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	ChildIds                              *Relation  `xmlrpc:"child_ids,omptempty"`
	City                                  *String    `xmlrpc:"city,omptempty"`
	CompanyRegistry                       *String    `xmlrpc:"company_registry,omptempty"`
	CountryId                             *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyExchangeJournalId             *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCashDifferenceExpenseAccountId *Many2One  `xmlrpc:"default_cash_difference_expense_account_id,omptempty"`
	DefaultCashDifferenceIncomeAccountId  *Many2One  `xmlrpc:"default_cash_difference_income_account_id,omptempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omptempty"`
	Email                                 *String    `xmlrpc:"email,omptempty"`
	ExpectsChartOfAccounts                *Bool      `xmlrpc:"expects_chart_of_accounts,omptempty"`
	ExpenseAccrualAccountId               *Many2One  `xmlrpc:"expense_accrual_account_id,omptempty"`
	ExpenseCurrencyExchangeAccountId      *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExternalReportLayoutId                *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	Favicon                               *String    `xmlrpc:"favicon,omptempty"`
	FiscalyearLastDay                     *Int       `xmlrpc:"fiscalyear_last_day,omptempty"`
	FiscalyearLastMonth                   *Selection `xmlrpc:"fiscalyear_last_month,omptempty"`
	FiscalyearLockDate                    *Time      `xmlrpc:"fiscalyear_lock_date,omptempty"`
	Font                                  *Selection `xmlrpc:"font,omptempty"`
	HrPresenceControlEmailAmount          *Int       `xmlrpc:"hr_presence_control_email_amount,omptempty"`
	HrPresenceControlIpList               *String    `xmlrpc:"hr_presence_control_ip_list,omptempty"`
	Id                                    *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId       *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	IncotermId                            *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InvoiceIsEmail                        *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsPrint                        *Bool      `xmlrpc:"invoice_is_print,omptempty"`
	InvoiceIsSnailmail                    *Bool      `xmlrpc:"invoice_is_snailmail,omptempty"`
	InvoiceTerms                          *String    `xmlrpc:"invoice_terms,omptempty"`
	L10NSgUniqueEntityNumber              *String    `xmlrpc:"l10n_sg_unique_entity_number,omptempty"`
	LastUpdate                            *Time      `xmlrpc:"__last_update,omptempty"`
	Logo                                  *String    `xmlrpc:"logo,omptempty"`
	LogoWeb                               *String    `xmlrpc:"logo_web,omptempty"`
	Name                                  *String    `xmlrpc:"name,omptempty"`
	NomenclatureId                        *Many2One  `xmlrpc:"nomenclature_id,omptempty"`
	PaperformatId                         *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	ParentId                              *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerGid                            *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentAcquirerOnboardingState        *Selection `xmlrpc:"payment_acquirer_onboarding_state,omptempty"`
	PaymentOnboardingPaymentMethod        *Selection `xmlrpc:"payment_onboarding_payment_method,omptempty"`
	PeriodLockDate                        *Time      `xmlrpc:"period_lock_date,omptempty"`
	Phone                                 *String    `xmlrpc:"phone,omptempty"`
	PortalConfirmationPay                 *Bool      `xmlrpc:"portal_confirmation_pay,omptempty"`
	PortalConfirmationSign                *Bool      `xmlrpc:"portal_confirmation_sign,omptempty"`
	PrimaryColor                          *String    `xmlrpc:"primary_color,omptempty"`
	PropertyStockAccountInputCategId      *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId     *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockValuationAccountId       *Many2One  `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	QrCode                                *Bool      `xmlrpc:"qr_code,omptempty"`
	QuotationValidityDays                 *Int       `xmlrpc:"quotation_validity_days,omptempty"`
	ReportFooter                          *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader                          *String    `xmlrpc:"report_header,omptempty"`
	ResourceCalendarId                    *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceCalendarIds                   *Relation  `xmlrpc:"resource_calendar_ids,omptempty"`
	RevenueAccrualAccountId               *Many2One  `xmlrpc:"revenue_accrual_account_id,omptempty"`
	SaleOnboardingOrderConfirmationState  *Selection `xmlrpc:"sale_onboarding_order_confirmation_state,omptempty"`
	SaleOnboardingPaymentMethod           *Selection `xmlrpc:"sale_onboarding_payment_method,omptempty"`
	SaleOnboardingSampleQuotationState    *Selection `xmlrpc:"sale_onboarding_sample_quotation_state,omptempty"`
	SaleQuotationOnboardingState          *Selection `xmlrpc:"sale_quotation_onboarding_state,omptempty"`
	SecondaryColor                        *String    `xmlrpc:"secondary_color,omptempty"`
	Sequence                              *Int       `xmlrpc:"sequence,omptempty"`
	SnailmailColor                        *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                        *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                       *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	SocialFacebook                        *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                          *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram                       *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin                        *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                         *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                         *String    `xmlrpc:"social_youtube,omptempty"`
	StateId                               *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                                *String    `xmlrpc:"street,omptempty"`
	Street2                               *String    `xmlrpc:"street2,omptempty"`
	TaxCalculationRoundingMethod          *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                 *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                        *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TaxLockDate                           *Time      `xmlrpc:"tax_lock_date,omptempty"`
	TransferAccountCodePrefix             *String    `xmlrpc:"transfer_account_code_prefix,omptempty"`
	TransferAccountId                     *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	UserIds                               *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                                   *String    `xmlrpc:"vat,omptempty"`
	Website                               *String    `xmlrpc:"website,omptempty"`
	WebsiteThemeOnboardingDone            *Bool      `xmlrpc:"website_theme_onboarding_done,omptempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                                   *String    `xmlrpc:"zip,omptempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	return c.Create(ResCompanyModel, rc)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.company not found", id)
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.company was not found")
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.company was not found")
}
