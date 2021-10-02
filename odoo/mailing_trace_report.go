package odoo

import (
	"fmt"
)

// MailingTraceReport represents mailing.trace.report model.
type MailingTraceReport struct {
	Bounced       *Int       `xmlrpc:"bounced,omptempty"`
	Campaign      *String    `xmlrpc:"campaign,omptempty"`
	Clicked       *Int       `xmlrpc:"clicked,omptempty"`
	Delivered     *Int       `xmlrpc:"delivered,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom     *String    `xmlrpc:"email_from,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	MailingType   *Selection `xmlrpc:"mailing_type,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	Opened        *Int       `xmlrpc:"opened,omptempty"`
	Replied       *Int       `xmlrpc:"replied,omptempty"`
	ScheduledDate *Time      `xmlrpc:"scheduled_date,omptempty"`
	Sent          *Int       `xmlrpc:"sent,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
}

// MailingTraceReports represents array of mailing.trace.report model.
type MailingTraceReports []MailingTraceReport

// MailingTraceReportModel is the odoo model name.
const MailingTraceReportModel = "mailing.trace.report"

// Many2One convert MailingTraceReport to *Many2One.
func (mtr *MailingTraceReport) Many2One() *Many2One {
	return NewMany2One(mtr.Id.Get(), "")
}

// CreateMailingTraceReport creates a new mailing.trace.report model and returns its id.
func (c *Client) CreateMailingTraceReport(mtr *MailingTraceReport) (int64, error) {
	return c.Create(MailingTraceReportModel, mtr)
}

// UpdateMailingTraceReport updates an existing mailing.trace.report record.
func (c *Client) UpdateMailingTraceReport(mtr *MailingTraceReport) error {
	return c.UpdateMailingTraceReports([]int64{mtr.Id.Get()}, mtr)
}

// UpdateMailingTraceReports updates existing mailing.trace.report records.
// All records (represented by ids) will be updated by mtr values.
func (c *Client) UpdateMailingTraceReports(ids []int64, mtr *MailingTraceReport) error {
	return c.Update(MailingTraceReportModel, ids, mtr)
}

// DeleteMailingTraceReport deletes an existing mailing.trace.report record.
func (c *Client) DeleteMailingTraceReport(id int64) error {
	return c.DeleteMailingTraceReports([]int64{id})
}

// DeleteMailingTraceReports deletes existing mailing.trace.report records.
func (c *Client) DeleteMailingTraceReports(ids []int64) error {
	return c.Delete(MailingTraceReportModel, ids)
}

// GetMailingTraceReport gets mailing.trace.report existing record.
func (c *Client) GetMailingTraceReport(id int64) (*MailingTraceReport, error) {
	mtrs, err := c.GetMailingTraceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtrs != nil && len(*mtrs) > 0 {
		return &((*mtrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.trace.report not found", id)
}

// GetMailingTraceReports gets mailing.trace.report existing records.
func (c *Client) GetMailingTraceReports(ids []int64) (*MailingTraceReports, error) {
	mtrs := &MailingTraceReports{}
	if err := c.Read(MailingTraceReportModel, ids, nil, mtrs); err != nil {
		return nil, err
	}
	return mtrs, nil
}

// FindMailingTraceReport finds mailing.trace.report record by querying it with criteria.
func (c *Client) FindMailingTraceReport(criteria *Criteria) (*MailingTraceReport, error) {
	mtrs := &MailingTraceReports{}
	if err := c.SearchRead(MailingTraceReportModel, criteria, NewOptions().Limit(1), mtrs); err != nil {
		return nil, err
	}
	if mtrs != nil && len(*mtrs) > 0 {
		return &((*mtrs)[0]), nil
	}
	return nil, fmt.Errorf("mailing.trace.report was not found")
}

// FindMailingTraceReports finds mailing.trace.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraceReports(criteria *Criteria, options *Options) (*MailingTraceReports, error) {
	mtrs := &MailingTraceReports{}
	if err := c.SearchRead(MailingTraceReportModel, criteria, options, mtrs); err != nil {
		return nil, err
	}
	return mtrs, nil
}

// FindMailingTraceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingTraceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingTraceReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingTraceReportId finds record id by querying it with criteria.
func (c *Client) FindMailingTraceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingTraceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.trace.report was not found")
}
