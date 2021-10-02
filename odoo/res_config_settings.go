package odoo

import (
	"fmt"
)

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	AccountBankReconciliationStart        *Time      `xmlrpc:"account_bank_reconciliation_start,omptempty"`
	ActiveUserCount                       *Int       `xmlrpc:"active_user_count,omptempty"`
	AliasDomain                           *String    `xmlrpc:"alias_domain,omptempty"`
	AuthSignupResetPassword               *Bool      `xmlrpc:"auth_signup_reset_password,omptempty"`
	AuthSignupTemplateUserId              *Many2One  `xmlrpc:"auth_signup_template_user_id,omptempty"`
	AuthSignupUninvited                   *Selection `xmlrpc:"auth_signup_uninvited,omptempty"`
	AutomaticInvoice                      *Bool      `xmlrpc:"automatic_invoice,omptempty"`
	CdnActivated                          *Bool      `xmlrpc:"cdn_activated,omptempty"`
	CdnFilters                            *String    `xmlrpc:"cdn_filters,omptempty"`
	CdnUrl                                *String    `xmlrpc:"cdn_url,omptempty"`
	ChannelId                             *Many2One  `xmlrpc:"channel_id,omptempty"`
	ChartTemplateId                       *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	CompanyCount                          *Int       `xmlrpc:"company_count,omptempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyInformations                   *String    `xmlrpc:"company_informations,omptempty"`
	CompanyName                           *String    `xmlrpc:"company_name,omptempty"`
	ConfirmationTemplateId                *Many2One  `xmlrpc:"confirmation_template_id,omptempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmAliasPrefix                        *String    `xmlrpc:"crm_alias_prefix,omptempty"`
	CrmDefaultTeamId                      *Many2One  `xmlrpc:"crm_default_team_id,omptempty"`
	CrmDefaultUserId                      *Many2One  `xmlrpc:"crm_default_user_id,omptempty"`
	CurrencyExchangeJournalId             *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultInvoicePolicy                  *Selection `xmlrpc:"default_invoice_policy,omptempty"`
	DefaultSaleOrderTemplateId            *Many2One  `xmlrpc:"default_sale_order_template_id,omptempty"`
	DepositDefaultProductId               *Many2One  `xmlrpc:"deposit_default_product_id,omptempty"`
	DigestEmails                          *Bool      `xmlrpc:"digest_emails,omptempty"`
	DigestId                              *Many2One  `xmlrpc:"digest_id,omptempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omptempty"`
	ExpenseAliasPrefix                    *String    `xmlrpc:"expense_alias_prefix,omptempty"`
	ExternalEmailServerDefault            *Bool      `xmlrpc:"external_email_server_default,omptempty"`
	ExternalReportLayoutId                *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	FailCounter                           *Int       `xmlrpc:"fail_counter,omptempty"`
	Favicon                               *String    `xmlrpc:"favicon,omptempty"`
	GenerateLeadFromAlias                 *Bool      `xmlrpc:"generate_lead_from_alias,omptempty"`
	GoogleAnalyticsKey                    *String    `xmlrpc:"google_analytics_key,omptempty"`
	GoogleManagementClientId              *String    `xmlrpc:"google_management_client_id,omptempty"`
	GoogleManagementClientSecret          *String    `xmlrpc:"google_management_client_secret,omptempty"`
	GoogleMapsApiKey                      *String    `xmlrpc:"google_maps_api_key,omptempty"`
	GroupAnalyticAccounting               *Bool      `xmlrpc:"group_analytic_accounting,omptempty"`
	GroupAnalyticTags                     *Bool      `xmlrpc:"group_analytic_tags,omptempty"`
	GroupAttendanceUsePin                 *Bool      `xmlrpc:"group_attendance_use_pin,omptempty"`
	GroupAutoDoneSetting                  *Bool      `xmlrpc:"group_auto_done_setting,omptempty"`
	GroupCashRounding                     *Bool      `xmlrpc:"group_cash_rounding,omptempty"`
	GroupDiscountPerSoLine                *Bool      `xmlrpc:"group_discount_per_so_line,omptempty"`
	GroupFiscalYear                       *Bool      `xmlrpc:"group_fiscal_year,omptempty"`
	GroupMassMailingCampaign              *Bool      `xmlrpc:"group_mass_mailing_campaign,omptempty"`
	GroupMultiCurrency                    *Bool      `xmlrpc:"group_multi_currency,omptempty"`
	GroupMultiWebsite                     *Bool      `xmlrpc:"group_multi_website,omptempty"`
	GroupProductPricelist                 *Bool      `xmlrpc:"group_product_pricelist,omptempty"`
	GroupProductVariant                   *Bool      `xmlrpc:"group_product_variant,omptempty"`
	GroupProformaSales                    *Bool      `xmlrpc:"group_proforma_sales,omptempty"`
	GroupProjectRating                    *Bool      `xmlrpc:"group_project_rating,omptempty"`
	GroupSaleDeliveryAddress              *Bool      `xmlrpc:"group_sale_delivery_address,omptempty"`
	GroupSaleOrderTemplate                *Bool      `xmlrpc:"group_sale_order_template,omptempty"`
	GroupSalePricelist                    *Bool      `xmlrpc:"group_sale_pricelist,omptempty"`
	GroupShowLineSubtotalsTaxExcluded     *Bool      `xmlrpc:"group_show_line_subtotals_tax_excluded,omptempty"`
	GroupShowLineSubtotalsTaxIncluded     *Bool      `xmlrpc:"group_show_line_subtotals_tax_included,omptempty"`
	GroupStockPackaging                   *Bool      `xmlrpc:"group_stock_packaging,omptempty"`
	GroupSubtaskProject                   *Bool      `xmlrpc:"group_subtask_project,omptempty"`
	GroupUom                              *Bool      `xmlrpc:"group_uom,omptempty"`
	GroupUseLead                          *Bool      `xmlrpc:"group_use_lead,omptempty"`
	GroupWarningAccount                   *Bool      `xmlrpc:"group_warning_account,omptempty"`
	GroupWarningSale                      *Bool      `xmlrpc:"group_warning_sale,omptempty"`
	HasAccountingEntries                  *Bool      `xmlrpc:"has_accounting_entries,omptempty"`
	HasChartOfAccounts                    *Bool      `xmlrpc:"has_chart_of_accounts,omptempty"`
	HasGoogleAnalytics                    *Bool      `xmlrpc:"has_google_analytics,omptempty"`
	HasGoogleAnalyticsDashboard           *Bool      `xmlrpc:"has_google_analytics_dashboard,omptempty"`
	HasGoogleMaps                         *Bool      `xmlrpc:"has_google_maps,omptempty"`
	HasSocialNetwork                      *Bool      `xmlrpc:"has_social_network,omptempty"`
	HrEmployeeSelfEdit                    *Bool      `xmlrpc:"hr_employee_self_edit,omptempty"`
	HrPresenceControlEmail                *Bool      `xmlrpc:"hr_presence_control_email,omptempty"`
	HrPresenceControlEmailAmount          *Int       `xmlrpc:"hr_presence_control_email_amount,omptempty"`
	HrPresenceControlIp                   *Bool      `xmlrpc:"hr_presence_control_ip,omptempty"`
	HrPresenceControlIpList               *String    `xmlrpc:"hr_presence_control_ip_list,omptempty"`
	HrPresenceControlLogin                *Bool      `xmlrpc:"hr_presence_control_login,omptempty"`
	Id                                    *Int       `xmlrpc:"id,omptempty"`
	IncotermId                            *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InvoiceIsEmail                        *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsPrint                        *Bool      `xmlrpc:"invoice_is_print,omptempty"`
	InvoiceIsSnailmail                    *Bool      `xmlrpc:"invoice_is_snailmail,omptempty"`
	InvoiceTerms                          *String    `xmlrpc:"invoice_terms,omptempty"`
	LanguageCount                         *Int       `xmlrpc:"language_count,omptempty"`
	LanguageIds                           *Relation  `xmlrpc:"language_ids,omptempty"`
	LastUpdate                            *Time      `xmlrpc:"__last_update,omptempty"`
	LeadEnrichAuto                        *Selection `xmlrpc:"lead_enrich_auto,omptempty"`
	LeadMiningInPipeline                  *Bool      `xmlrpc:"lead_mining_in_pipeline,omptempty"`
	MassMailingMailServerId               *Many2One  `xmlrpc:"mass_mailing_mail_server_id,omptempty"`
	MassMailingOutgoingMailServer         *Bool      `xmlrpc:"mass_mailing_outgoing_mail_server,omptempty"`
	ModuleAccountAccountant               *Bool      `xmlrpc:"module_account_accountant,omptempty"`
	ModuleAccountBankStatementImportCamt  *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omptempty"`
	ModuleAccountBankStatementImportCsv   *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omptempty"`
	ModuleAccountBankStatementImportOfx   *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omptempty"`
	ModuleAccountBankStatementImportQif   *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omptempty"`
	ModuleAccountBatchPayment             *Bool      `xmlrpc:"module_account_batch_payment,omptempty"`
	ModuleAccountBudget                   *Bool      `xmlrpc:"module_account_budget,omptempty"`
	ModuleAccountCheckPrinting            *Bool      `xmlrpc:"module_account_check_printing,omptempty"`
	ModuleAccountIntrastat                *Bool      `xmlrpc:"module_account_intrastat,omptempty"`
	ModuleAccountInvoiceExtract           *Bool      `xmlrpc:"module_account_invoice_extract,omptempty"`
	ModuleAccountPayment                  *Bool      `xmlrpc:"module_account_payment,omptempty"`
	ModuleAccountPlaid                    *Bool      `xmlrpc:"module_account_plaid,omptempty"`
	ModuleAccountReports                  *Bool      `xmlrpc:"module_account_reports,omptempty"`
	ModuleAccountSepa                     *Bool      `xmlrpc:"module_account_sepa,omptempty"`
	ModuleAccountSepaDirectDebit          *Bool      `xmlrpc:"module_account_sepa_direct_debit,omptempty"`
	ModuleAccountTaxcloud                 *Bool      `xmlrpc:"module_account_taxcloud,omptempty"`
	ModuleAccountYodlee                   *Bool      `xmlrpc:"module_account_yodlee,omptempty"`
	ModuleAuthLdap                        *Bool      `xmlrpc:"module_auth_ldap,omptempty"`
	ModuleAuthOauth                       *Bool      `xmlrpc:"module_auth_oauth,omptempty"`
	ModuleBaseGengo                       *Bool      `xmlrpc:"module_base_gengo,omptempty"`
	ModuleBaseGeolocalize                 *Bool      `xmlrpc:"module_base_geolocalize,omptempty"`
	ModuleBaseImport                      *Bool      `xmlrpc:"module_base_import,omptempty"`
	ModuleCrmIapLead                      *Bool      `xmlrpc:"module_crm_iap_lead,omptempty"`
	ModuleCrmIapLeadEnrich                *Bool      `xmlrpc:"module_crm_iap_lead_enrich,omptempty"`
	ModuleCrmIapLeadWebsite               *Bool      `xmlrpc:"module_crm_iap_lead_website,omptempty"`
	ModuleCurrencyRateLive                *Bool      `xmlrpc:"module_currency_rate_live,omptempty"`
	ModuleDelivery                        *Bool      `xmlrpc:"module_delivery,omptempty"`
	ModuleDeliveryBpost                   *Bool      `xmlrpc:"module_delivery_bpost,omptempty"`
	ModuleDeliveryDhl                     *Bool      `xmlrpc:"module_delivery_dhl,omptempty"`
	ModuleDeliveryEasypost                *Bool      `xmlrpc:"module_delivery_easypost,omptempty"`
	ModuleDeliveryFedex                   *Bool      `xmlrpc:"module_delivery_fedex,omptempty"`
	ModuleDeliveryUps                     *Bool      `xmlrpc:"module_delivery_ups,omptempty"`
	ModuleDeliveryUsps                    *Bool      `xmlrpc:"module_delivery_usps,omptempty"`
	ModuleEventBarcode                    *Bool      `xmlrpc:"module_event_barcode,omptempty"`
	ModuleEventSale                       *Bool      `xmlrpc:"module_event_sale,omptempty"`
	ModuleGoogleCalendar                  *Bool      `xmlrpc:"module_google_calendar,omptempty"`
	ModuleGoogleDrive                     *Bool      `xmlrpc:"module_google_drive,omptempty"`
	ModuleGoogleSpreadsheet               *Bool      `xmlrpc:"module_google_spreadsheet,omptempty"`
	ModuleHrAttendance                    *Bool      `xmlrpc:"module_hr_attendance,omptempty"`
	ModuleHrOrgChart                      *Bool      `xmlrpc:"module_hr_org_chart,omptempty"`
	ModuleHrPayrollExpense                *Bool      `xmlrpc:"module_hr_payroll_expense,omptempty"`
	ModuleHrPresence                      *Bool      `xmlrpc:"module_hr_presence,omptempty"`
	ModuleHrRecruitmentSurvey             *Bool      `xmlrpc:"module_hr_recruitment_survey,omptempty"`
	ModuleHrSkills                        *Bool      `xmlrpc:"module_hr_skills,omptempty"`
	ModuleHrTimesheet                     *Bool      `xmlrpc:"module_hr_timesheet,omptempty"`
	ModuleInterCompanyRules               *Bool      `xmlrpc:"module_inter_company_rules,omptempty"`
	ModuleL10NEuService                   *Bool      `xmlrpc:"module_l10n_eu_service,omptempty"`
	ModuleMassMailingSlides               *Bool      `xmlrpc:"module_mass_mailing_slides,omptempty"`
	ModulePad                             *Bool      `xmlrpc:"module_pad,omptempty"`
	ModulePartnerAutocomplete             *Bool      `xmlrpc:"module_partner_autocomplete,omptempty"`
	ModuleProductEmailTemplate            *Bool      `xmlrpc:"module_product_email_template,omptempty"`
	ModuleProductMargin                   *Bool      `xmlrpc:"module_product_margin,omptempty"`
	ModuleProjectForecast                 *Bool      `xmlrpc:"module_project_forecast,omptempty"`
	ModuleSaleAmazon                      *Bool      `xmlrpc:"module_sale_amazon,omptempty"`
	ModuleSaleCoupon                      *Bool      `xmlrpc:"module_sale_coupon,omptempty"`
	ModuleSaleMargin                      *Bool      `xmlrpc:"module_sale_margin,omptempty"`
	ModuleSaleProductConfigurator         *Bool      `xmlrpc:"module_sale_product_configurator,omptempty"`
	ModuleSaleProductMatrix               *Bool      `xmlrpc:"module_sale_product_matrix,omptempty"`
	ModuleSaleQuotationBuilder            *Bool      `xmlrpc:"module_sale_quotation_builder,omptempty"`
	ModuleSnailmailAccount                *Bool      `xmlrpc:"module_snailmail_account,omptempty"`
	ModuleVoip                            *Bool      `xmlrpc:"module_voip,omptempty"`
	ModuleWebsiteEventQuestions           *Bool      `xmlrpc:"module_website_event_questions,omptempty"`
	ModuleWebsiteEventSale                *Bool      `xmlrpc:"module_website_event_sale,omptempty"`
	ModuleWebsiteEventTrack               *Bool      `xmlrpc:"module_website_event_track,omptempty"`
	ModuleWebsiteHrRecruitment            *Bool      `xmlrpc:"module_website_hr_recruitment,omptempty"`
	ModuleWebsiteLinks                    *Bool      `xmlrpc:"module_website_links,omptempty"`
	ModuleWebsiteSaleDigital              *Bool      `xmlrpc:"module_website_sale_digital,omptempty"`
	ModuleWebsiteSaleSlides               *Bool      `xmlrpc:"module_website_sale_slides,omptempty"`
	ModuleWebsiteSlidesForum              *Bool      `xmlrpc:"module_website_slides_forum,omptempty"`
	ModuleWebsiteSlidesSurvey             *Bool      `xmlrpc:"module_website_slides_survey,omptempty"`
	ModuleWebsiteVersion                  *Bool      `xmlrpc:"module_website_version,omptempty"`
	ModuleWebUnsplash                     *Bool      `xmlrpc:"module_web_unsplash,omptempty"`
	PaperformatId                         *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	PartnerAutocompleteInsufficientCredit *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omptempty"`
	PortalConfirmationPay                 *Bool      `xmlrpc:"portal_confirmation_pay,omptempty"`
	PortalConfirmationSign                *Bool      `xmlrpc:"portal_confirmation_sign,omptempty"`
	PredictiveLeadScoringFields           *Relation  `xmlrpc:"predictive_lead_scoring_fields,omptempty"`
	PredictiveLeadScoringFieldsStr        *String    `xmlrpc:"predictive_lead_scoring_fields_str,omptempty"`
	PredictiveLeadScoringStartDate        *Time      `xmlrpc:"predictive_lead_scoring_start_date,omptempty"`
	PredictiveLeadScoringStartDateStr     *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omptempty"`
	ProductPricelistSetting               *Selection `xmlrpc:"product_pricelist_setting,omptempty"`
	ProductVolumeVolumeInCubicFeet        *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omptempty"`
	ProductWeightInLbs                    *Selection `xmlrpc:"product_weight_in_lbs,omptempty"`
	PurchaseTaxId                         *Many2One  `xmlrpc:"purchase_tax_id,omptempty"`
	QrCode                                *Bool      `xmlrpc:"qr_code,omptempty"`
	QuotationValidityDays                 *Int       `xmlrpc:"quotation_validity_days,omptempty"`
	ReportFooter                          *String    `xmlrpc:"report_footer,omptempty"`
	ResourceCalendarId                    *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	SaleTaxId                             *Many2One  `xmlrpc:"sale_tax_id,omptempty"`
	ShowBlacklistButtons                  *Bool      `xmlrpc:"show_blacklist_buttons,omptempty"`
	ShowEffect                            *Bool      `xmlrpc:"show_effect,omptempty"`
	ShowLineSubtotalsTaxSelection         *Selection `xmlrpc:"show_line_subtotals_tax_selection,omptempty"`
	SnailmailColor                        *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                        *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                       *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	SocialDefaultImage                    *String    `xmlrpc:"social_default_image,omptempty"`
	SocialFacebook                        *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                          *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram                       *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin                        *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                         *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                         *String    `xmlrpc:"social_youtube,omptempty"`
	SpecificUserAccount                   *Bool      `xmlrpc:"specific_user_account,omptempty"`
	TaxCalculationRoundingMethod          *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                 *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                        *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TemplateId                            *Many2One  `xmlrpc:"template_id,omptempty"`
	UnsplashAccessKey                     *String    `xmlrpc:"unsplash_access_key,omptempty"`
	UseInvoiceTerms                       *Bool      `xmlrpc:"use_invoice_terms,omptempty"`
	UseMailgateway                        *Bool      `xmlrpc:"use_mailgateway,omptempty"`
	UseQuotationValidityDays              *Bool      `xmlrpc:"use_quotation_validity_days,omptempty"`
	UserDefaultRights                     *Bool      `xmlrpc:"user_default_rights,omptempty"`
	WebsiteCompanyId                      *Many2One  `xmlrpc:"website_company_id,omptempty"`
	WebsiteCountryGroupIds                *Relation  `xmlrpc:"website_country_group_ids,omptempty"`
	WebsiteDefaultLangCode                *String    `xmlrpc:"website_default_lang_code,omptempty"`
	WebsiteDefaultLangId                  *Many2One  `xmlrpc:"website_default_lang_id,omptempty"`
	WebsiteDomain                         *String    `xmlrpc:"website_domain,omptempty"`
	WebsiteFormEnableMetadata             *Bool      `xmlrpc:"website_form_enable_metadata,omptempty"`
	WebsiteId                             *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteLanguageCount                  *Int       `xmlrpc:"website_language_count,omptempty"`
	WebsiteLogo                           *String    `xmlrpc:"website_logo,omptempty"`
	WebsiteName                           *String    `xmlrpc:"website_name,omptempty"`
	WebsiteSlideGoogleAppKey              *String    `xmlrpc:"website_slide_google_app_key,omptempty"`
	WebWindowTitle                        *String    `xmlrpc:"web_window_title,omptempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	return c.Create(ResConfigSettingsModel, rcs)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.config.settings not found", id)
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("res.config.settings was not found")
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.config.settings was not found")
}
