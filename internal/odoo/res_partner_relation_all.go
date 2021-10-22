package odoo

import (
	"fmt"
)

// ResPartnerRelationAll represents res.partner.relation.all model.
type ResPartnerRelationAll struct {
	Active              *Bool      `xmlrpc:"active,omptempty"`
	ContactType         *Selection `xmlrpc:"contact_type,omptempty"`
	DateEnd             *Time      `xmlrpc:"date_end,omptempty"`
	DateStart           *Time      `xmlrpc:"date_start,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	GrantAccess         *Bool      `xmlrpc:"grant_access,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	OtherPartnerId      *Many2One  `xmlrpc:"other_partner_id,omptempty"`
	OtherPrimaryContact *Bool      `xmlrpc:"other_primary_contact,omptempty"`
	RecordType          *Selection `xmlrpc:"record_type,omptempty"`
	RelationId          *Many2One  `xmlrpc:"relation_id,omptempty"`
	ThisPartnerId       *Many2One  `xmlrpc:"this_partner_id,omptempty"`
	ThisPrimaryContact  *Bool      `xmlrpc:"this_primary_contact,omptempty"`
	TypeId              *Many2One  `xmlrpc:"type_id,omptempty"`
	TypeSelectionId     *Many2One  `xmlrpc:"type_selection_id,omptempty"`
}

// ResPartnerRelationAlls represents array of res.partner.relation.all model.
type ResPartnerRelationAlls []ResPartnerRelationAll

// ResPartnerRelationAllModel is the odoo model name.
const ResPartnerRelationAllModel = "res.partner.relation.all"

// Many2One convert ResPartnerRelationAll to *Many2One.
func (rpra *ResPartnerRelationAll) Many2One() *Many2One {
	return NewMany2One(rpra.Id.Get(), "")
}

// CreateResPartnerRelationAll creates a new res.partner.relation.all model and returns its id.
func (c *Client) CreateResPartnerRelationAll(rpra *ResPartnerRelationAll) (int64, error) {
	return c.Create(ResPartnerRelationAllModel, rpra)
}

// UpdateResPartnerRelationAll updates an existing res.partner.relation.all record.
func (c *Client) UpdateResPartnerRelationAll(rpra *ResPartnerRelationAll) error {
	return c.UpdateResPartnerRelationAlls([]int64{rpra.Id.Get()}, rpra)
}

// UpdateResPartnerRelationAlls updates existing res.partner.relation.all records.
// All records (represented by ids) will be updated by rpra values.
func (c *Client) UpdateResPartnerRelationAlls(ids []int64, rpra *ResPartnerRelationAll) error {
	return c.Update(ResPartnerRelationAllModel, ids, rpra)
}

// DeleteResPartnerRelationAll deletes an existing res.partner.relation.all record.
func (c *Client) DeleteResPartnerRelationAll(id int64) error {
	return c.DeleteResPartnerRelationAlls([]int64{id})
}

// DeleteResPartnerRelationAlls deletes existing res.partner.relation.all records.
func (c *Client) DeleteResPartnerRelationAlls(ids []int64) error {
	return c.Delete(ResPartnerRelationAllModel, ids)
}

// GetResPartnerRelationAll gets res.partner.relation.all existing record.
func (c *Client) GetResPartnerRelationAll(id int64) (*ResPartnerRelationAll, error) {
	rpras, err := c.GetResPartnerRelationAlls([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpras != nil && len(*rpras) > 0 {
		return &((*rpras)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.relation.all not found", id)
}

// GetResPartnerRelationAlls gets res.partner.relation.all existing records.
func (c *Client) GetResPartnerRelationAlls(ids []int64) (*ResPartnerRelationAlls, error) {
	rpras := &ResPartnerRelationAlls{}
	if err := c.Read(ResPartnerRelationAllModel, ids, nil, rpras); err != nil {
		return nil, err
	}
	return rpras, nil
}

// FindResPartnerRelationAll finds res.partner.relation.all record by querying it with criteria.
func (c *Client) FindResPartnerRelationAll(criteria *Criteria) (*ResPartnerRelationAll, error) {
	rpras := &ResPartnerRelationAlls{}
	if err := c.SearchRead(ResPartnerRelationAllModel, criteria, NewOptions().Limit(1), rpras); err != nil {
		return nil, err
	}
	if rpras != nil && len(*rpras) > 0 {
		return &((*rpras)[0]), nil
	}
	return nil, fmt.Errorf("res.partner.relation.all was not found")
}

// FindResPartnerRelationAlls finds res.partner.relation.all records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerRelationAlls(criteria *Criteria, options *Options) (*ResPartnerRelationAlls, error) {
	rpras := &ResPartnerRelationAlls{}
	if err := c.SearchRead(ResPartnerRelationAllModel, criteria, options, rpras); err != nil {
		return nil, err
	}
	return rpras, nil
}

// FindResPartnerRelationAllIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerRelationAllIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerRelationAllModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerRelationAllId finds record id by querying it with criteria.
func (c *Client) FindResPartnerRelationAllId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerRelationAllModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner.relation.all was not found")
}
