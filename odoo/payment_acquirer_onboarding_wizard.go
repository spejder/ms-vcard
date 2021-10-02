package odoo

import (
	"fmt"
)

// PaymentAcquirerOnboardingWizard represents payment.acquirer.onboarding.wizard model.
type PaymentAcquirerOnboardingWizard struct {
	AccNumber            *String    `xmlrpc:"acc_number,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	JournalName          *String    `xmlrpc:"journal_name,omptempty"`
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	ManualName           *String    `xmlrpc:"manual_name,omptempty"`
	ManualPostMsg        *String    `xmlrpc:"manual_post_msg,omptempty"`
	PaymentMethod        *Selection `xmlrpc:"payment_method,omptempty"`
	PaypalEmailAccount   *String    `xmlrpc:"paypal_email_account,omptempty"`
	PaypalPdtToken       *String    `xmlrpc:"paypal_pdt_token,omptempty"`
	PaypalSellerAccount  *String    `xmlrpc:"paypal_seller_account,omptempty"`
	PaypalUserType       *Selection `xmlrpc:"paypal_user_type,omptempty"`
	StripePublishableKey *String    `xmlrpc:"stripe_publishable_key,omptempty"`
	StripeSecretKey      *String    `xmlrpc:"stripe_secret_key,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentAcquirerOnboardingWizards represents array of payment.acquirer.onboarding.wizard model.
type PaymentAcquirerOnboardingWizards []PaymentAcquirerOnboardingWizard

// PaymentAcquirerOnboardingWizardModel is the odoo model name.
const PaymentAcquirerOnboardingWizardModel = "payment.acquirer.onboarding.wizard"

// Many2One convert PaymentAcquirerOnboardingWizard to *Many2One.
func (paow *PaymentAcquirerOnboardingWizard) Many2One() *Many2One {
	return NewMany2One(paow.Id.Get(), "")
}

// CreatePaymentAcquirerOnboardingWizard creates a new payment.acquirer.onboarding.wizard model and returns its id.
func (c *Client) CreatePaymentAcquirerOnboardingWizard(paow *PaymentAcquirerOnboardingWizard) (int64, error) {
	return c.Create(PaymentAcquirerOnboardingWizardModel, paow)
}

// UpdatePaymentAcquirerOnboardingWizard updates an existing payment.acquirer.onboarding.wizard record.
func (c *Client) UpdatePaymentAcquirerOnboardingWizard(paow *PaymentAcquirerOnboardingWizard) error {
	return c.UpdatePaymentAcquirerOnboardingWizards([]int64{paow.Id.Get()}, paow)
}

// UpdatePaymentAcquirerOnboardingWizards updates existing payment.acquirer.onboarding.wizard records.
// All records (represented by ids) will be updated by paow values.
func (c *Client) UpdatePaymentAcquirerOnboardingWizards(ids []int64, paow *PaymentAcquirerOnboardingWizard) error {
	return c.Update(PaymentAcquirerOnboardingWizardModel, ids, paow)
}

// DeletePaymentAcquirerOnboardingWizard deletes an existing payment.acquirer.onboarding.wizard record.
func (c *Client) DeletePaymentAcquirerOnboardingWizard(id int64) error {
	return c.DeletePaymentAcquirerOnboardingWizards([]int64{id})
}

// DeletePaymentAcquirerOnboardingWizards deletes existing payment.acquirer.onboarding.wizard records.
func (c *Client) DeletePaymentAcquirerOnboardingWizards(ids []int64) error {
	return c.Delete(PaymentAcquirerOnboardingWizardModel, ids)
}

// GetPaymentAcquirerOnboardingWizard gets payment.acquirer.onboarding.wizard existing record.
func (c *Client) GetPaymentAcquirerOnboardingWizard(id int64) (*PaymentAcquirerOnboardingWizard, error) {
	paows, err := c.GetPaymentAcquirerOnboardingWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if paows != nil && len(*paows) > 0 {
		return &((*paows)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.acquirer.onboarding.wizard not found", id)
}

// GetPaymentAcquirerOnboardingWizards gets payment.acquirer.onboarding.wizard existing records.
func (c *Client) GetPaymentAcquirerOnboardingWizards(ids []int64) (*PaymentAcquirerOnboardingWizards, error) {
	paows := &PaymentAcquirerOnboardingWizards{}
	if err := c.Read(PaymentAcquirerOnboardingWizardModel, ids, nil, paows); err != nil {
		return nil, err
	}
	return paows, nil
}

// FindPaymentAcquirerOnboardingWizard finds payment.acquirer.onboarding.wizard record by querying it with criteria.
func (c *Client) FindPaymentAcquirerOnboardingWizard(criteria *Criteria) (*PaymentAcquirerOnboardingWizard, error) {
	paows := &PaymentAcquirerOnboardingWizards{}
	if err := c.SearchRead(PaymentAcquirerOnboardingWizardModel, criteria, NewOptions().Limit(1), paows); err != nil {
		return nil, err
	}
	if paows != nil && len(*paows) > 0 {
		return &((*paows)[0]), nil
	}
	return nil, fmt.Errorf("payment.acquirer.onboarding.wizard was not found")
}

// FindPaymentAcquirerOnboardingWizards finds payment.acquirer.onboarding.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentAcquirerOnboardingWizards(criteria *Criteria, options *Options) (*PaymentAcquirerOnboardingWizards, error) {
	paows := &PaymentAcquirerOnboardingWizards{}
	if err := c.SearchRead(PaymentAcquirerOnboardingWizardModel, criteria, options, paows); err != nil {
		return nil, err
	}
	return paows, nil
}

// FindPaymentAcquirerOnboardingWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentAcquirerOnboardingWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentAcquirerOnboardingWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentAcquirerOnboardingWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentAcquirerOnboardingWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentAcquirerOnboardingWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.acquirer.onboarding.wizard was not found")
}
