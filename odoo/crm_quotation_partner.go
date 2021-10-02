package odoo

import (
	"fmt"
)

// CrmQuotationPartner represents crm.quotation.partner model.
type CrmQuotationPartner struct {
	Action      *Selection `xmlrpc:"action,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	LeadId      *Many2One  `xmlrpc:"lead_id,omptempty"`
	PartnerId   *Many2One  `xmlrpc:"partner_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CrmQuotationPartners represents array of crm.quotation.partner model.
type CrmQuotationPartners []CrmQuotationPartner

// CrmQuotationPartnerModel is the odoo model name.
const CrmQuotationPartnerModel = "crm.quotation.partner"

// Many2One convert CrmQuotationPartner to *Many2One.
func (cqp *CrmQuotationPartner) Many2One() *Many2One {
	return NewMany2One(cqp.Id.Get(), "")
}

// CreateCrmQuotationPartner creates a new crm.quotation.partner model and returns its id.
func (c *Client) CreateCrmQuotationPartner(cqp *CrmQuotationPartner) (int64, error) {
	return c.Create(CrmQuotationPartnerModel, cqp)
}

// UpdateCrmQuotationPartner updates an existing crm.quotation.partner record.
func (c *Client) UpdateCrmQuotationPartner(cqp *CrmQuotationPartner) error {
	return c.UpdateCrmQuotationPartners([]int64{cqp.Id.Get()}, cqp)
}

// UpdateCrmQuotationPartners updates existing crm.quotation.partner records.
// All records (represented by ids) will be updated by cqp values.
func (c *Client) UpdateCrmQuotationPartners(ids []int64, cqp *CrmQuotationPartner) error {
	return c.Update(CrmQuotationPartnerModel, ids, cqp)
}

// DeleteCrmQuotationPartner deletes an existing crm.quotation.partner record.
func (c *Client) DeleteCrmQuotationPartner(id int64) error {
	return c.DeleteCrmQuotationPartners([]int64{id})
}

// DeleteCrmQuotationPartners deletes existing crm.quotation.partner records.
func (c *Client) DeleteCrmQuotationPartners(ids []int64) error {
	return c.Delete(CrmQuotationPartnerModel, ids)
}

// GetCrmQuotationPartner gets crm.quotation.partner existing record.
func (c *Client) GetCrmQuotationPartner(id int64) (*CrmQuotationPartner, error) {
	cqps, err := c.GetCrmQuotationPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if cqps != nil && len(*cqps) > 0 {
		return &((*cqps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.quotation.partner not found", id)
}

// GetCrmQuotationPartners gets crm.quotation.partner existing records.
func (c *Client) GetCrmQuotationPartners(ids []int64) (*CrmQuotationPartners, error) {
	cqps := &CrmQuotationPartners{}
	if err := c.Read(CrmQuotationPartnerModel, ids, nil, cqps); err != nil {
		return nil, err
	}
	return cqps, nil
}

// FindCrmQuotationPartner finds crm.quotation.partner record by querying it with criteria.
func (c *Client) FindCrmQuotationPartner(criteria *Criteria) (*CrmQuotationPartner, error) {
	cqps := &CrmQuotationPartners{}
	if err := c.SearchRead(CrmQuotationPartnerModel, criteria, NewOptions().Limit(1), cqps); err != nil {
		return nil, err
	}
	if cqps != nil && len(*cqps) > 0 {
		return &((*cqps)[0]), nil
	}
	return nil, fmt.Errorf("crm.quotation.partner was not found")
}

// FindCrmQuotationPartners finds crm.quotation.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmQuotationPartners(criteria *Criteria, options *Options) (*CrmQuotationPartners, error) {
	cqps := &CrmQuotationPartners{}
	if err := c.SearchRead(CrmQuotationPartnerModel, criteria, options, cqps); err != nil {
		return nil, err
	}
	return cqps, nil
}

// FindCrmQuotationPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmQuotationPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmQuotationPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmQuotationPartnerId finds record id by querying it with criteria.
func (c *Client) FindCrmQuotationPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmQuotationPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.quotation.partner was not found")
}
