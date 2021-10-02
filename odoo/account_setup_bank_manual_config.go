package odoo

import (
	"fmt"
)

// AccountSetupBankManualConfig represents account.setup.bank.manual.config model.
type AccountSetupBankManualConfig struct {
	AccHolderName             *String    `xmlrpc:"acc_holder_name,omptempty"`
	AccNumber                 *String    `xmlrpc:"acc_number,omptempty"`
	AccType                   *Selection `xmlrpc:"acc_type,omptempty"`
	BankBic                   *String    `xmlrpc:"bank_bic,omptempty"`
	BankId                    *Many2One  `xmlrpc:"bank_id,omptempty"`
	BankName                  *String    `xmlrpc:"bank_name,omptempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	JournalId                 *Relation  `xmlrpc:"journal_id,omptempty"`
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	LinkedJournalId           *Many2One  `xmlrpc:"linked_journal_id,omptempty"`
	NewJournalCode            *String    `xmlrpc:"new_journal_code,omptempty"`
	NewJournalName            *String    `xmlrpc:"new_journal_name,omptempty"`
	NumJournalsWithoutAccount *Int       `xmlrpc:"num_journals_without_account,omptempty"`
	PartnerId                 *Many2One  `xmlrpc:"partner_id,omptempty"`
	QrCodeValid               *Bool      `xmlrpc:"qr_code_valid,omptempty"`
	RelatedAccType            *Selection `xmlrpc:"related_acc_type,omptempty"`
	ResPartnerBankId          *Many2One  `xmlrpc:"res_partner_bank_id,omptempty"`
	SanitizedAccNumber        *String    `xmlrpc:"sanitized_acc_number,omptempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountSetupBankManualConfigs represents array of account.setup.bank.manual.config model.
type AccountSetupBankManualConfigs []AccountSetupBankManualConfig

// AccountSetupBankManualConfigModel is the odoo model name.
const AccountSetupBankManualConfigModel = "account.setup.bank.manual.config"

// Many2One convert AccountSetupBankManualConfig to *Many2One.
func (asbmc *AccountSetupBankManualConfig) Many2One() *Many2One {
	return NewMany2One(asbmc.Id.Get(), "")
}

// CreateAccountSetupBankManualConfig creates a new account.setup.bank.manual.config model and returns its id.
func (c *Client) CreateAccountSetupBankManualConfig(asbmc *AccountSetupBankManualConfig) (int64, error) {
	return c.Create(AccountSetupBankManualConfigModel, asbmc)
}

// UpdateAccountSetupBankManualConfig updates an existing account.setup.bank.manual.config record.
func (c *Client) UpdateAccountSetupBankManualConfig(asbmc *AccountSetupBankManualConfig) error {
	return c.UpdateAccountSetupBankManualConfigs([]int64{asbmc.Id.Get()}, asbmc)
}

// UpdateAccountSetupBankManualConfigs updates existing account.setup.bank.manual.config records.
// All records (represented by ids) will be updated by asbmc values.
func (c *Client) UpdateAccountSetupBankManualConfigs(ids []int64, asbmc *AccountSetupBankManualConfig) error {
	return c.Update(AccountSetupBankManualConfigModel, ids, asbmc)
}

// DeleteAccountSetupBankManualConfig deletes an existing account.setup.bank.manual.config record.
func (c *Client) DeleteAccountSetupBankManualConfig(id int64) error {
	return c.DeleteAccountSetupBankManualConfigs([]int64{id})
}

// DeleteAccountSetupBankManualConfigs deletes existing account.setup.bank.manual.config records.
func (c *Client) DeleteAccountSetupBankManualConfigs(ids []int64) error {
	return c.Delete(AccountSetupBankManualConfigModel, ids)
}

// GetAccountSetupBankManualConfig gets account.setup.bank.manual.config existing record.
func (c *Client) GetAccountSetupBankManualConfig(id int64) (*AccountSetupBankManualConfig, error) {
	asbmcs, err := c.GetAccountSetupBankManualConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if asbmcs != nil && len(*asbmcs) > 0 {
		return &((*asbmcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.setup.bank.manual.config not found", id)
}

// GetAccountSetupBankManualConfigs gets account.setup.bank.manual.config existing records.
func (c *Client) GetAccountSetupBankManualConfigs(ids []int64) (*AccountSetupBankManualConfigs, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.Read(AccountSetupBankManualConfigModel, ids, nil, asbmcs); err != nil {
		return nil, err
	}
	return asbmcs, nil
}

// FindAccountSetupBankManualConfig finds account.setup.bank.manual.config record by querying it with criteria.
func (c *Client) FindAccountSetupBankManualConfig(criteria *Criteria) (*AccountSetupBankManualConfig, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.SearchRead(AccountSetupBankManualConfigModel, criteria, NewOptions().Limit(1), asbmcs); err != nil {
		return nil, err
	}
	if asbmcs != nil && len(*asbmcs) > 0 {
		return &((*asbmcs)[0]), nil
	}
	return nil, fmt.Errorf("account.setup.bank.manual.config was not found")
}

// FindAccountSetupBankManualConfigs finds account.setup.bank.manual.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSetupBankManualConfigs(criteria *Criteria, options *Options) (*AccountSetupBankManualConfigs, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.SearchRead(AccountSetupBankManualConfigModel, criteria, options, asbmcs); err != nil {
		return nil, err
	}
	return asbmcs, nil
}

// FindAccountSetupBankManualConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSetupBankManualConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountSetupBankManualConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountSetupBankManualConfigId finds record id by querying it with criteria.
func (c *Client) FindAccountSetupBankManualConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSetupBankManualConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.setup.bank.manual.config was not found")
}
