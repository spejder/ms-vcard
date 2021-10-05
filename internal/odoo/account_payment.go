package odoo

import (
	"fmt"
)

// AccountPayment represents account.payment model.
type AccountPayment struct {
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	Amount                      *Float     `xmlrpc:"amount,omptempty"`
	Communication               *String    `xmlrpc:"communication,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DestinationAccountId        *Many2One  `xmlrpc:"destination_account_id,omptempty"`
	DestinationJournalId        *Many2One  `xmlrpc:"destination_journal_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	HasInvoices                 *Bool      `xmlrpc:"has_invoices,omptempty"`
	HidePaymentMethod           *Bool      `xmlrpc:"hide_payment_method,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	InvoiceIds                  *Relation  `xmlrpc:"invoice_ids,omptempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
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
	MoveLineIds                 *Relation  `xmlrpc:"move_line_ids,omptempty"`
	MoveName                    *String    `xmlrpc:"move_name,omptempty"`
	MoveReconciled              *Bool      `xmlrpc:"move_reconciled,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PartnerBankAccountId        *Many2One  `xmlrpc:"partner_bank_account_id,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerType                 *Selection `xmlrpc:"partner_type,omptempty"`
	PaymentDate                 *Time      `xmlrpc:"payment_date,omptempty"`
	PaymentDifference           *Float     `xmlrpc:"payment_difference,omptempty"`
	PaymentDifferenceHandling   *Selection `xmlrpc:"payment_difference_handling,omptempty"`
	PaymentMethodCode           *String    `xmlrpc:"payment_method_code,omptempty"`
	PaymentMethodId             *Many2One  `xmlrpc:"payment_method_id,omptempty"`
	PaymentReference            *String    `xmlrpc:"payment_reference,omptempty"`
	PaymentTokenId              *Many2One  `xmlrpc:"payment_token_id,omptempty"`
	PaymentTransactionId        *Many2One  `xmlrpc:"payment_transaction_id,omptempty"`
	PaymentType                 *Selection `xmlrpc:"payment_type,omptempty"`
	ReconciledInvoiceIds        *Relation  `xmlrpc:"reconciled_invoice_ids,omptempty"`
	ReconciledInvoicesCount     *Int       `xmlrpc:"reconciled_invoices_count,omptempty"`
	RequirePartnerBankAccount   *Bool      `xmlrpc:"require_partner_bank_account,omptempty"`
	ShowPartnerBankAccount      *Bool      `xmlrpc:"show_partner_bank_account,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteoffAccountId           *Many2One  `xmlrpc:"writeoff_account_id,omptempty"`
	WriteoffLabel               *String    `xmlrpc:"writeoff_label,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountPayments represents array of account.payment model.
type AccountPayments []AccountPayment

// AccountPaymentModel is the odoo model name.
const AccountPaymentModel = "account.payment"

// Many2One convert AccountPayment to *Many2One.
func (ap *AccountPayment) Many2One() *Many2One {
	return NewMany2One(ap.Id.Get(), "")
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayment(ap *AccountPayment) (int64, error) {
	return c.Create(AccountPaymentModel, ap)
}

// UpdateAccountPayment updates an existing account.payment record.
func (c *Client) UpdateAccountPayment(ap *AccountPayment) error {
	return c.UpdateAccountPayments([]int64{ap.Id.Get()}, ap)
}

// UpdateAccountPayments updates existing account.payment records.
// All records (represented by ids) will be updated by ap values.
func (c *Client) UpdateAccountPayments(ids []int64, ap *AccountPayment) error {
	return c.Update(AccountPaymentModel, ids, ap)
}

// DeleteAccountPayment deletes an existing account.payment record.
func (c *Client) DeleteAccountPayment(id int64) error {
	return c.DeleteAccountPayments([]int64{id})
}

// DeleteAccountPayments deletes existing account.payment records.
func (c *Client) DeleteAccountPayments(ids []int64) error {
	return c.Delete(AccountPaymentModel, ids)
}

// GetAccountPayment gets account.payment existing record.
func (c *Client) GetAccountPayment(id int64) (*AccountPayment, error) {
	aps, err := c.GetAccountPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if aps != nil && len(*aps) > 0 {
		return &((*aps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment not found", id)
}

// GetAccountPayments gets account.payment existing records.
func (c *Client) GetAccountPayments(ids []int64) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.Read(AccountPaymentModel, ids, nil, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPayment finds account.payment record by querying it with criteria.
func (c *Client) FindAccountPayment(criteria *Criteria) (*AccountPayment, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, NewOptions().Limit(1), aps); err != nil {
		return nil, err
	}
	if aps != nil && len(*aps) > 0 {
		return &((*aps)[0]), nil
	}
	return nil, fmt.Errorf("account.payment was not found")
}

// FindAccountPayments finds account.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPayments(criteria *Criteria, options *Options) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, options, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.payment was not found")
}
