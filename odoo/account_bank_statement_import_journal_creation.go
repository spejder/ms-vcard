package odoo

import (
	"fmt"
)

// AccountBankStatementImportJournalCreation represents account.bank.statement.import.journal.creation model.
type AccountBankStatementImportJournalCreation struct {
	AccountControlIds           *Relation  `xmlrpc:"account_control_ids,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AliasDomain                 *String    `xmlrpc:"alias_domain,omptempty"`
	AliasId                     *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasName                   *String    `xmlrpc:"alias_name,omptempty"`
	AtLeastOneInbound           *Bool      `xmlrpc:"at_least_one_inbound,omptempty"`
	AtLeastOneOutbound          *Bool      `xmlrpc:"at_least_one_outbound,omptempty"`
	BankAccNumber               *String    `xmlrpc:"bank_acc_number,omptempty"`
	BankAccountId               *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	BankId                      *Many2One  `xmlrpc:"bank_id,omptempty"`
	BankStatementsSource        *Selection `xmlrpc:"bank_statements_source,omptempty"`
	Code                        *String    `xmlrpc:"code,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyPartnerId            *Many2One  `xmlrpc:"company_partner_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCreditAccountId      *Many2One  `xmlrpc:"default_credit_account_id,omptempty"`
	DefaultDebitAccountId       *Many2One  `xmlrpc:"default_debit_account_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	InboundPaymentMethodIds     *Relation  `xmlrpc:"inbound_payment_method_ids,omptempty"`
	InvoiceReferenceModel       *Selection `xmlrpc:"invoice_reference_model,omptempty"`
	InvoiceReferenceType        *Selection `xmlrpc:"invoice_reference_type,omptempty"`
	JournalGroupIds             *Relation  `xmlrpc:"journal_group_ids,omptempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omptempty"`
	JsonActivityData            *String    `xmlrpc:"json_activity_data,omptempty"`
	KanbanDashboard             *String    `xmlrpc:"kanban_dashboard,omptempty"`
	KanbanDashboardGraph        *String    `xmlrpc:"kanban_dashboard_graph,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	LossAccountId               *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	OutboundPaymentMethodIds    *Relation  `xmlrpc:"outbound_payment_method_ids,omptempty"`
	PostAt                      *Selection `xmlrpc:"post_at,omptempty"`
	ProfitAccountId             *Many2One  `xmlrpc:"profit_account_id,omptempty"`
	RefundSequence              *Bool      `xmlrpc:"refund_sequence,omptempty"`
	RefundSequenceId            *Many2One  `xmlrpc:"refund_sequence_id,omptempty"`
	RefundSequenceNumberNext    *Int       `xmlrpc:"refund_sequence_number_next,omptempty"`
	RestrictModeHashTable       *Bool      `xmlrpc:"restrict_mode_hash_table,omptempty"`
	SecureSequenceId            *Many2One  `xmlrpc:"secure_sequence_id,omptempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omptempty"`
	SequenceId                  *Many2One  `xmlrpc:"sequence_id,omptempty"`
	SequenceNumberNext          *Int       `xmlrpc:"sequence_number_next,omptempty"`
	ShowOnDashboard             *Bool      `xmlrpc:"show_on_dashboard,omptempty"`
	Type                        *Selection `xmlrpc:"type,omptempty"`
	TypeControlIds              *Relation  `xmlrpc:"type_control_ids,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatementImportJournalCreations represents array of account.bank.statement.import.journal.creation model.
type AccountBankStatementImportJournalCreations []AccountBankStatementImportJournalCreation

// AccountBankStatementImportJournalCreationModel is the odoo model name.
const AccountBankStatementImportJournalCreationModel = "account.bank.statement.import.journal.creation"

// Many2One convert AccountBankStatementImportJournalCreation to *Many2One.
func (absijc *AccountBankStatementImportJournalCreation) Many2One() *Many2One {
	return NewMany2One(absijc.Id.Get(), "")
}

// CreateAccountBankStatementImportJournalCreation creates a new account.bank.statement.import.journal.creation model and returns its id.
func (c *Client) CreateAccountBankStatementImportJournalCreation(absijc *AccountBankStatementImportJournalCreation) (int64, error) {
	return c.Create(AccountBankStatementImportJournalCreationModel, absijc)
}

// UpdateAccountBankStatementImportJournalCreation updates an existing account.bank.statement.import.journal.creation record.
func (c *Client) UpdateAccountBankStatementImportJournalCreation(absijc *AccountBankStatementImportJournalCreation) error {
	return c.UpdateAccountBankStatementImportJournalCreations([]int64{absijc.Id.Get()}, absijc)
}

// UpdateAccountBankStatementImportJournalCreations updates existing account.bank.statement.import.journal.creation records.
// All records (represented by ids) will be updated by absijc values.
func (c *Client) UpdateAccountBankStatementImportJournalCreations(ids []int64, absijc *AccountBankStatementImportJournalCreation) error {
	return c.Update(AccountBankStatementImportJournalCreationModel, ids, absijc)
}

// DeleteAccountBankStatementImportJournalCreation deletes an existing account.bank.statement.import.journal.creation record.
func (c *Client) DeleteAccountBankStatementImportJournalCreation(id int64) error {
	return c.DeleteAccountBankStatementImportJournalCreations([]int64{id})
}

// DeleteAccountBankStatementImportJournalCreations deletes existing account.bank.statement.import.journal.creation records.
func (c *Client) DeleteAccountBankStatementImportJournalCreations(ids []int64) error {
	return c.Delete(AccountBankStatementImportJournalCreationModel, ids)
}

// GetAccountBankStatementImportJournalCreation gets account.bank.statement.import.journal.creation existing record.
func (c *Client) GetAccountBankStatementImportJournalCreation(id int64) (*AccountBankStatementImportJournalCreation, error) {
	absijcs, err := c.GetAccountBankStatementImportJournalCreations([]int64{id})
	if err != nil {
		return nil, err
	}
	if absijcs != nil && len(*absijcs) > 0 {
		return &((*absijcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.import.journal.creation not found", id)
}

// GetAccountBankStatementImportJournalCreations gets account.bank.statement.import.journal.creation existing records.
func (c *Client) GetAccountBankStatementImportJournalCreations(ids []int64) (*AccountBankStatementImportJournalCreations, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.Read(AccountBankStatementImportJournalCreationModel, ids, nil, absijcs); err != nil {
		return nil, err
	}
	return absijcs, nil
}

// FindAccountBankStatementImportJournalCreation finds account.bank.statement.import.journal.creation record by querying it with criteria.
func (c *Client) FindAccountBankStatementImportJournalCreation(criteria *Criteria) (*AccountBankStatementImportJournalCreation, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.SearchRead(AccountBankStatementImportJournalCreationModel, criteria, NewOptions().Limit(1), absijcs); err != nil {
		return nil, err
	}
	if absijcs != nil && len(*absijcs) > 0 {
		return &((*absijcs)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.statement.import.journal.creation was not found")
}

// FindAccountBankStatementImportJournalCreations finds account.bank.statement.import.journal.creation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportJournalCreations(criteria *Criteria, options *Options) (*AccountBankStatementImportJournalCreations, error) {
	absijcs := &AccountBankStatementImportJournalCreations{}
	if err := c.SearchRead(AccountBankStatementImportJournalCreationModel, criteria, options, absijcs); err != nil {
		return nil, err
	}
	return absijcs, nil
}

// FindAccountBankStatementImportJournalCreationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementImportJournalCreationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementImportJournalCreationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementImportJournalCreationId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementImportJournalCreationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementImportJournalCreationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.statement.import.journal.creation was not found")
}
