package odoo

import (
	"fmt"
)

// AccountReconcileModelTemplate represents account.reconcile.model.template model.
type AccountReconcileModelTemplate struct {
	AccountId                  *Many2One  `xmlrpc:"account_id,omptempty"`
	Amount                     *Float     `xmlrpc:"amount,omptempty"`
	AmountFromLabelRegex       *String    `xmlrpc:"amount_from_label_regex,omptempty"`
	AmountType                 *Selection `xmlrpc:"amount_type,omptempty"`
	AutoReconcile              *Bool      `xmlrpc:"auto_reconcile,omptempty"`
	ChartTemplateId            *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	DecimalSeparator           *String    `xmlrpc:"decimal_separator,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	ForceSecondTaxIncluded     *Bool      `xmlrpc:"force_second_tax_included,omptempty"`
	ForceTaxIncluded           *Bool      `xmlrpc:"force_tax_included,omptempty"`
	HasSecondLine              *Bool      `xmlrpc:"has_second_line,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	Label                      *String    `xmlrpc:"label,omptempty"`
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	MatchAmount                *Selection `xmlrpc:"match_amount,omptempty"`
	MatchAmountMax             *Float     `xmlrpc:"match_amount_max,omptempty"`
	MatchAmountMin             *Float     `xmlrpc:"match_amount_min,omptempty"`
	MatchJournalIds            *Relation  `xmlrpc:"match_journal_ids,omptempty"`
	MatchLabel                 *Selection `xmlrpc:"match_label,omptempty"`
	MatchLabelParam            *String    `xmlrpc:"match_label_param,omptempty"`
	MatchNature                *Selection `xmlrpc:"match_nature,omptempty"`
	MatchNote                  *Selection `xmlrpc:"match_note,omptempty"`
	MatchNoteParam             *String    `xmlrpc:"match_note_param,omptempty"`
	MatchPartner               *Bool      `xmlrpc:"match_partner,omptempty"`
	MatchPartnerCategoryIds    *Relation  `xmlrpc:"match_partner_category_ids,omptempty"`
	MatchPartnerIds            *Relation  `xmlrpc:"match_partner_ids,omptempty"`
	MatchSameCurrency          *Bool      `xmlrpc:"match_same_currency,omptempty"`
	MatchTotalAmount           *Bool      `xmlrpc:"match_total_amount,omptempty"`
	MatchTotalAmountParam      *Float     `xmlrpc:"match_total_amount_param,omptempty"`
	MatchTransactionType       *Selection `xmlrpc:"match_transaction_type,omptempty"`
	MatchTransactionTypeParam  *String    `xmlrpc:"match_transaction_type_param,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	RuleType                   *Selection `xmlrpc:"rule_type,omptempty"`
	SecondAccountId            *Many2One  `xmlrpc:"second_account_id,omptempty"`
	SecondAmount               *Float     `xmlrpc:"second_amount,omptempty"`
	SecondAmountFromLabelRegex *String    `xmlrpc:"second_amount_from_label_regex,omptempty"`
	SecondAmountType           *Selection `xmlrpc:"second_amount_type,omptempty"`
	SecondLabel                *String    `xmlrpc:"second_label,omptempty"`
	SecondTaxIds               *Relation  `xmlrpc:"second_tax_ids,omptempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omptempty"`
	TaxIds                     *Relation  `xmlrpc:"tax_ids,omptempty"`
	ToCheck                    *Bool      `xmlrpc:"to_check,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReconcileModelTemplates represents array of account.reconcile.model.template model.
type AccountReconcileModelTemplates []AccountReconcileModelTemplate

// AccountReconcileModelTemplateModel is the odoo model name.
const AccountReconcileModelTemplateModel = "account.reconcile.model.template"

// Many2One convert AccountReconcileModelTemplate to *Many2One.
func (armt *AccountReconcileModelTemplate) Many2One() *Many2One {
	return NewMany2One(armt.Id.Get(), "")
}

// CreateAccountReconcileModelTemplate creates a new account.reconcile.model.template model and returns its id.
func (c *Client) CreateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) (int64, error) {
	return c.Create(AccountReconcileModelTemplateModel, armt)
}

// UpdateAccountReconcileModelTemplate updates an existing account.reconcile.model.template record.
func (c *Client) UpdateAccountReconcileModelTemplate(armt *AccountReconcileModelTemplate) error {
	return c.UpdateAccountReconcileModelTemplates([]int64{armt.Id.Get()}, armt)
}

// UpdateAccountReconcileModelTemplates updates existing account.reconcile.model.template records.
// All records (represented by ids) will be updated by armt values.
func (c *Client) UpdateAccountReconcileModelTemplates(ids []int64, armt *AccountReconcileModelTemplate) error {
	return c.Update(AccountReconcileModelTemplateModel, ids, armt)
}

// DeleteAccountReconcileModelTemplate deletes an existing account.reconcile.model.template record.
func (c *Client) DeleteAccountReconcileModelTemplate(id int64) error {
	return c.DeleteAccountReconcileModelTemplates([]int64{id})
}

// DeleteAccountReconcileModelTemplates deletes existing account.reconcile.model.template records.
func (c *Client) DeleteAccountReconcileModelTemplates(ids []int64) error {
	return c.Delete(AccountReconcileModelTemplateModel, ids)
}

// GetAccountReconcileModelTemplate gets account.reconcile.model.template existing record.
func (c *Client) GetAccountReconcileModelTemplate(id int64) (*AccountReconcileModelTemplate, error) {
	armts, err := c.GetAccountReconcileModelTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if armts != nil && len(*armts) > 0 {
		return &((*armts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model.template not found", id)
}

// GetAccountReconcileModelTemplates gets account.reconcile.model.template existing records.
func (c *Client) GetAccountReconcileModelTemplates(ids []int64) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.Read(AccountReconcileModelTemplateModel, ids, nil, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplate finds account.reconcile.model.template record by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplate(criteria *Criteria) (*AccountReconcileModelTemplate, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, NewOptions().Limit(1), armts); err != nil {
		return nil, err
	}
	if armts != nil && len(*armts) > 0 {
		return &((*armts)[0]), nil
	}
	return nil, fmt.Errorf("account.reconcile.model.template was not found")
}

// FindAccountReconcileModelTemplates finds account.reconcile.model.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplates(criteria *Criteria, options *Options) (*AccountReconcileModelTemplates, error) {
	armts := &AccountReconcileModelTemplates{}
	if err := c.SearchRead(AccountReconcileModelTemplateModel, criteria, options, armts); err != nil {
		return nil, err
	}
	return armts, nil
}

// FindAccountReconcileModelTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.reconcile.model.template was not found")
}
