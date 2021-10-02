package odoo

import (
	"fmt"
)

// BaseDocumentLayout represents base.document.layout model.
type BaseDocumentLayout struct {
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomColors           *Bool      `xmlrpc:"custom_colors,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	ExternalReportLayoutId *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	Font                   *Selection `xmlrpc:"font,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Logo                   *String    `xmlrpc:"logo,omptempty"`
	LogoPrimaryColor       *String    `xmlrpc:"logo_primary_color,omptempty"`
	LogoSecondaryColor     *String    `xmlrpc:"logo_secondary_color,omptempty"`
	PaperformatId          *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	Preview                *String    `xmlrpc:"preview,omptempty"`
	PreviewLogo            *String    `xmlrpc:"preview_logo,omptempty"`
	PrimaryColor           *String    `xmlrpc:"primary_color,omptempty"`
	ReportFooter           *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader           *String    `xmlrpc:"report_header,omptempty"`
	ReportLayoutId         *Many2One  `xmlrpc:"report_layout_id,omptempty"`
	SecondaryColor         *String    `xmlrpc:"secondary_color,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BaseDocumentLayouts represents array of base.document.layout model.
type BaseDocumentLayouts []BaseDocumentLayout

// BaseDocumentLayoutModel is the odoo model name.
const BaseDocumentLayoutModel = "base.document.layout"

// Many2One convert BaseDocumentLayout to *Many2One.
func (bdl *BaseDocumentLayout) Many2One() *Many2One {
	return NewMany2One(bdl.Id.Get(), "")
}

// CreateBaseDocumentLayout creates a new base.document.layout model and returns its id.
func (c *Client) CreateBaseDocumentLayout(bdl *BaseDocumentLayout) (int64, error) {
	return c.Create(BaseDocumentLayoutModel, bdl)
}

// UpdateBaseDocumentLayout updates an existing base.document.layout record.
func (c *Client) UpdateBaseDocumentLayout(bdl *BaseDocumentLayout) error {
	return c.UpdateBaseDocumentLayouts([]int64{bdl.Id.Get()}, bdl)
}

// UpdateBaseDocumentLayouts updates existing base.document.layout records.
// All records (represented by ids) will be updated by bdl values.
func (c *Client) UpdateBaseDocumentLayouts(ids []int64, bdl *BaseDocumentLayout) error {
	return c.Update(BaseDocumentLayoutModel, ids, bdl)
}

// DeleteBaseDocumentLayout deletes an existing base.document.layout record.
func (c *Client) DeleteBaseDocumentLayout(id int64) error {
	return c.DeleteBaseDocumentLayouts([]int64{id})
}

// DeleteBaseDocumentLayouts deletes existing base.document.layout records.
func (c *Client) DeleteBaseDocumentLayouts(ids []int64) error {
	return c.Delete(BaseDocumentLayoutModel, ids)
}

// GetBaseDocumentLayout gets base.document.layout existing record.
func (c *Client) GetBaseDocumentLayout(id int64) (*BaseDocumentLayout, error) {
	bdls, err := c.GetBaseDocumentLayouts([]int64{id})
	if err != nil {
		return nil, err
	}
	if bdls != nil && len(*bdls) > 0 {
		return &((*bdls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.document.layout not found", id)
}

// GetBaseDocumentLayouts gets base.document.layout existing records.
func (c *Client) GetBaseDocumentLayouts(ids []int64) (*BaseDocumentLayouts, error) {
	bdls := &BaseDocumentLayouts{}
	if err := c.Read(BaseDocumentLayoutModel, ids, nil, bdls); err != nil {
		return nil, err
	}
	return bdls, nil
}

// FindBaseDocumentLayout finds base.document.layout record by querying it with criteria.
func (c *Client) FindBaseDocumentLayout(criteria *Criteria) (*BaseDocumentLayout, error) {
	bdls := &BaseDocumentLayouts{}
	if err := c.SearchRead(BaseDocumentLayoutModel, criteria, NewOptions().Limit(1), bdls); err != nil {
		return nil, err
	}
	if bdls != nil && len(*bdls) > 0 {
		return &((*bdls)[0]), nil
	}
	return nil, fmt.Errorf("base.document.layout was not found")
}

// FindBaseDocumentLayouts finds base.document.layout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseDocumentLayouts(criteria *Criteria, options *Options) (*BaseDocumentLayouts, error) {
	bdls := &BaseDocumentLayouts{}
	if err := c.SearchRead(BaseDocumentLayoutModel, criteria, options, bdls); err != nil {
		return nil, err
	}
	return bdls, nil
}

// FindBaseDocumentLayoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseDocumentLayoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseDocumentLayoutModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseDocumentLayoutId finds record id by querying it with criteria.
func (c *Client) FindBaseDocumentLayoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseDocumentLayoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.document.layout was not found")
}
