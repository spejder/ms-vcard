package odoo

import (
	"fmt"
)

// AccountAnalyticAccount represents account.analytic.account model.
type AccountAnalyticAccount struct {
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	Balance                  *Float    `xmlrpc:"balance,omptempty"`
	Code                     *String   `xmlrpc:"code,omptempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	Credit                   *Float    `xmlrpc:"credit,omptempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omptempty"`
	Debit                    *Float    `xmlrpc:"debit,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	GroupId                  *Many2One `xmlrpc:"group_id,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	LineIds                  *Relation `xmlrpc:"line_ids,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	PartnerId                *Many2One `xmlrpc:"partner_id,omptempty"`
	ProjectCount             *Int      `xmlrpc:"project_count,omptempty"`
	ProjectIds               *Relation `xmlrpc:"project_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountAnalyticAccounts represents array of account.analytic.account model.
type AccountAnalyticAccounts []AccountAnalyticAccount

// AccountAnalyticAccountModel is the odoo model name.
const AccountAnalyticAccountModel = "account.analytic.account"

// Many2One convert AccountAnalyticAccount to *Many2One.
func (aaa *AccountAnalyticAccount) Many2One() *Many2One {
	return NewMany2One(aaa.Id.Get(), "")
}

// CreateAccountAnalyticAccount creates a new account.analytic.account model and returns its id.
func (c *Client) CreateAccountAnalyticAccount(aaa *AccountAnalyticAccount) (int64, error) {
	return c.Create(AccountAnalyticAccountModel, aaa)
}

// UpdateAccountAnalyticAccount updates an existing account.analytic.account record.
func (c *Client) UpdateAccountAnalyticAccount(aaa *AccountAnalyticAccount) error {
	return c.UpdateAccountAnalyticAccounts([]int64{aaa.Id.Get()}, aaa)
}

// UpdateAccountAnalyticAccounts updates existing account.analytic.account records.
// All records (represented by ids) will be updated by aaa values.
func (c *Client) UpdateAccountAnalyticAccounts(ids []int64, aaa *AccountAnalyticAccount) error {
	return c.Update(AccountAnalyticAccountModel, ids, aaa)
}

// DeleteAccountAnalyticAccount deletes an existing account.analytic.account record.
func (c *Client) DeleteAccountAnalyticAccount(id int64) error {
	return c.DeleteAccountAnalyticAccounts([]int64{id})
}

// DeleteAccountAnalyticAccounts deletes existing account.analytic.account records.
func (c *Client) DeleteAccountAnalyticAccounts(ids []int64) error {
	return c.Delete(AccountAnalyticAccountModel, ids)
}

// GetAccountAnalyticAccount gets account.analytic.account existing record.
func (c *Client) GetAccountAnalyticAccount(id int64) (*AccountAnalyticAccount, error) {
	aaas, err := c.GetAccountAnalyticAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaas != nil && len(*aaas) > 0 {
		return &((*aaas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.account not found", id)
}

// GetAccountAnalyticAccounts gets account.analytic.account existing records.
func (c *Client) GetAccountAnalyticAccounts(ids []int64) (*AccountAnalyticAccounts, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.Read(AccountAnalyticAccountModel, ids, nil, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
}

// FindAccountAnalyticAccount finds account.analytic.account record by querying it with criteria.
func (c *Client) FindAccountAnalyticAccount(criteria *Criteria) (*AccountAnalyticAccount, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.SearchRead(AccountAnalyticAccountModel, criteria, NewOptions().Limit(1), aaas); err != nil {
		return nil, err
	}
	if aaas != nil && len(*aaas) > 0 {
		return &((*aaas)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.account was not found")
}

// FindAccountAnalyticAccounts finds account.analytic.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticAccounts(criteria *Criteria, options *Options) (*AccountAnalyticAccounts, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.SearchRead(AccountAnalyticAccountModel, criteria, options, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
}

// FindAccountAnalyticAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.account was not found")
}
