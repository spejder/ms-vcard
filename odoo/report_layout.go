package odoo

import (
	"fmt"
)

// ReportLayout represents report.layout model.
type ReportLayout struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Image       *String   `xmlrpc:"image,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Pdf         *String   `xmlrpc:"pdf,omptempty"`
	ViewId      *Many2One `xmlrpc:"view_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ReportLayouts represents array of report.layout model.
type ReportLayouts []ReportLayout

// ReportLayoutModel is the odoo model name.
const ReportLayoutModel = "report.layout"

// Many2One convert ReportLayout to *Many2One.
func (rl *ReportLayout) Many2One() *Many2One {
	return NewMany2One(rl.Id.Get(), "")
}

// CreateReportLayout creates a new report.layout model and returns its id.
func (c *Client) CreateReportLayout(rl *ReportLayout) (int64, error) {
	return c.Create(ReportLayoutModel, rl)
}

// UpdateReportLayout updates an existing report.layout record.
func (c *Client) UpdateReportLayout(rl *ReportLayout) error {
	return c.UpdateReportLayouts([]int64{rl.Id.Get()}, rl)
}

// UpdateReportLayouts updates existing report.layout records.
// All records (represented by ids) will be updated by rl values.
func (c *Client) UpdateReportLayouts(ids []int64, rl *ReportLayout) error {
	return c.Update(ReportLayoutModel, ids, rl)
}

// DeleteReportLayout deletes an existing report.layout record.
func (c *Client) DeleteReportLayout(id int64) error {
	return c.DeleteReportLayouts([]int64{id})
}

// DeleteReportLayouts deletes existing report.layout records.
func (c *Client) DeleteReportLayouts(ids []int64) error {
	return c.Delete(ReportLayoutModel, ids)
}

// GetReportLayout gets report.layout existing record.
func (c *Client) GetReportLayout(id int64) (*ReportLayout, error) {
	rls, err := c.GetReportLayouts([]int64{id})
	if err != nil {
		return nil, err
	}
	if rls != nil && len(*rls) > 0 {
		return &((*rls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.layout not found", id)
}

// GetReportLayouts gets report.layout existing records.
func (c *Client) GetReportLayouts(ids []int64) (*ReportLayouts, error) {
	rls := &ReportLayouts{}
	if err := c.Read(ReportLayoutModel, ids, nil, rls); err != nil {
		return nil, err
	}
	return rls, nil
}

// FindReportLayout finds report.layout record by querying it with criteria.
func (c *Client) FindReportLayout(criteria *Criteria) (*ReportLayout, error) {
	rls := &ReportLayouts{}
	if err := c.SearchRead(ReportLayoutModel, criteria, NewOptions().Limit(1), rls); err != nil {
		return nil, err
	}
	if rls != nil && len(*rls) > 0 {
		return &((*rls)[0]), nil
	}
	return nil, fmt.Errorf("report.layout was not found")
}

// FindReportLayouts finds report.layout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportLayouts(criteria *Criteria, options *Options) (*ReportLayouts, error) {
	rls := &ReportLayouts{}
	if err := c.SearchRead(ReportLayoutModel, criteria, options, rls); err != nil {
		return nil, err
	}
	return rls, nil
}

// FindReportLayoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportLayoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportLayoutModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportLayoutId finds record id by querying it with criteria.
func (c *Client) FindReportLayoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportLayoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.layout was not found")
}
