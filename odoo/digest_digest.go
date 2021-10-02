package odoo

import (
	"fmt"
)

// DigestDigest represents digest.digest model.
type DigestDigest struct {
	AvailableFields                    *String    `xmlrpc:"available_fields,omptempty"`
	CompanyId                          *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                          *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                         *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                        *String    `xmlrpc:"display_name,omptempty"`
	Id                                 *Int       `xmlrpc:"id,omptempty"`
	IsSubscribed                       *Bool      `xmlrpc:"is_subscribed,omptempty"`
	KpiAccountTotalRevenue             *Bool      `xmlrpc:"kpi_account_total_revenue,omptempty"`
	KpiAccountTotalRevenueValue        *Float     `xmlrpc:"kpi_account_total_revenue_value,omptempty"`
	KpiAllSaleTotal                    *Bool      `xmlrpc:"kpi_all_sale_total,omptempty"`
	KpiAllSaleTotalValue               *Float     `xmlrpc:"kpi_all_sale_total_value,omptempty"`
	KpiCrmLeadCreated                  *Bool      `xmlrpc:"kpi_crm_lead_created,omptempty"`
	KpiCrmLeadCreatedValue             *Int       `xmlrpc:"kpi_crm_lead_created_value,omptempty"`
	KpiCrmOpportunitiesWon             *Bool      `xmlrpc:"kpi_crm_opportunities_won,omptempty"`
	KpiCrmOpportunitiesWonValue        *Int       `xmlrpc:"kpi_crm_opportunities_won_value,omptempty"`
	KpiHrRecruitmentNewColleagues      *Bool      `xmlrpc:"kpi_hr_recruitment_new_colleagues,omptempty"`
	KpiHrRecruitmentNewColleaguesValue *Int       `xmlrpc:"kpi_hr_recruitment_new_colleagues_value,omptempty"`
	KpiMailMessageTotal                *Bool      `xmlrpc:"kpi_mail_message_total,omptempty"`
	KpiMailMessageTotalValue           *Int       `xmlrpc:"kpi_mail_message_total_value,omptempty"`
	KpiProjectTaskOpened               *Bool      `xmlrpc:"kpi_project_task_opened,omptempty"`
	KpiProjectTaskOpenedValue          *Int       `xmlrpc:"kpi_project_task_opened_value,omptempty"`
	KpiResUsersConnected               *Bool      `xmlrpc:"kpi_res_users_connected,omptempty"`
	KpiResUsersConnectedValue          *Int       `xmlrpc:"kpi_res_users_connected_value,omptempty"`
	LastUpdate                         *Time      `xmlrpc:"__last_update,omptempty"`
	Name                               *String    `xmlrpc:"name,omptempty"`
	NextRunDate                        *Time      `xmlrpc:"next_run_date,omptempty"`
	Periodicity                        *Selection `xmlrpc:"periodicity,omptempty"`
	State                              *Selection `xmlrpc:"state,omptempty"`
	TemplateId                         *Many2One  `xmlrpc:"template_id,omptempty"`
	UserIds                            *Relation  `xmlrpc:"user_ids,omptempty"`
	WriteDate                          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DigestDigests represents array of digest.digest model.
type DigestDigests []DigestDigest

// DigestDigestModel is the odoo model name.
const DigestDigestModel = "digest.digest"

// Many2One convert DigestDigest to *Many2One.
func (dd *DigestDigest) Many2One() *Many2One {
	return NewMany2One(dd.Id.Get(), "")
}

// CreateDigestDigest creates a new digest.digest model and returns its id.
func (c *Client) CreateDigestDigest(dd *DigestDigest) (int64, error) {
	return c.Create(DigestDigestModel, dd)
}

// UpdateDigestDigest updates an existing digest.digest record.
func (c *Client) UpdateDigestDigest(dd *DigestDigest) error {
	return c.UpdateDigestDigests([]int64{dd.Id.Get()}, dd)
}

// UpdateDigestDigests updates existing digest.digest records.
// All records (represented by ids) will be updated by dd values.
func (c *Client) UpdateDigestDigests(ids []int64, dd *DigestDigest) error {
	return c.Update(DigestDigestModel, ids, dd)
}

// DeleteDigestDigest deletes an existing digest.digest record.
func (c *Client) DeleteDigestDigest(id int64) error {
	return c.DeleteDigestDigests([]int64{id})
}

// DeleteDigestDigests deletes existing digest.digest records.
func (c *Client) DeleteDigestDigests(ids []int64) error {
	return c.Delete(DigestDigestModel, ids)
}

// GetDigestDigest gets digest.digest existing record.
func (c *Client) GetDigestDigest(id int64) (*DigestDigest, error) {
	dds, err := c.GetDigestDigests([]int64{id})
	if err != nil {
		return nil, err
	}
	if dds != nil && len(*dds) > 0 {
		return &((*dds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of digest.digest not found", id)
}

// GetDigestDigests gets digest.digest existing records.
func (c *Client) GetDigestDigests(ids []int64) (*DigestDigests, error) {
	dds := &DigestDigests{}
	if err := c.Read(DigestDigestModel, ids, nil, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDigestDigest finds digest.digest record by querying it with criteria.
func (c *Client) FindDigestDigest(criteria *Criteria) (*DigestDigest, error) {
	dds := &DigestDigests{}
	if err := c.SearchRead(DigestDigestModel, criteria, NewOptions().Limit(1), dds); err != nil {
		return nil, err
	}
	if dds != nil && len(*dds) > 0 {
		return &((*dds)[0]), nil
	}
	return nil, fmt.Errorf("digest.digest was not found")
}

// FindDigestDigests finds digest.digest records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestDigests(criteria *Criteria, options *Options) (*DigestDigests, error) {
	dds := &DigestDigests{}
	if err := c.SearchRead(DigestDigestModel, criteria, options, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDigestDigestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestDigestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DigestDigestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDigestDigestId finds record id by querying it with criteria.
func (c *Client) FindDigestDigestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DigestDigestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("digest.digest was not found")
}
