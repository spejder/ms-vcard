package ms

import (
	"fmt"

	"bitbucket.org/long174/go-odoo"
)

// ResPartnerRelationAll represents res.partner.relation.all model.
type ResPartnerRelationAll struct {
	Active              *odoo.Bool      `xmlrpc:"active,omptempty"`
	ContactType         *odoo.Selection `xmlrpc:"contact_type,omptempty"`
	DateEnd             *odoo.Time      `xmlrpc:"date_end,omptempty"`
	DateStart           *odoo.Time      `xmlrpc:"date_start,omptempty"`
	DisplayName         *odoo.String    `xmlrpc:"display_name,omptempty"`
	GrantAccess         *odoo.Bool      `xmlrpc:"grant_access,omptempty"`
	Id                  *odoo.Int       `xmlrpc:"id,omptempty"`
	LastUpdate          *odoo.Time      `xmlrpc:"__last_update,omptempty"`
	OtherPartnerId      *odoo.Many2One  `xmlrpc:"other_partner_id,omptempty"`
	OtherPrimaryContact *odoo.Bool      `xmlrpc:"other_primary_contact,omptempty"`
	RecordType          *odoo.Selection `xmlrpc:"record_type,omptempty"`
	RelationId          *odoo.Many2One  `xmlrpc:"relation_id,omptempty"`
	ThisPartnerId       *odoo.Many2One  `xmlrpc:"this_partner_id,omptempty"`
	ThisPrimaryContact  *odoo.Bool      `xmlrpc:"this_primary_contact,omptempty"`
	TypeId              *odoo.Many2One  `xmlrpc:"type_id,omptempty"`
	TypeSelectionId     *odoo.Many2One  `xmlrpc:"type_selection_id,omptempty"`
}

// ResPartnerRelationAlls represents array of res.partner.relation.all model.
type ResPartnerRelationAlls []ResPartnerRelationAll

// ResPartnerRelationAllModel is the odoo model name.
const ResPartnerRelationAllModel = "res.partner.relation.all"

// Many2One convert ResPartnerRelationAll to *Many2One.
func (rpra *ResPartnerRelationAll) Many2One() *odoo.Many2One {
	return odoo.NewMany2One(rpra.Id.Get(), "")
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
func (c *Client) FindResPartnerRelationAll(criteria *odoo.Criteria) (*ResPartnerRelationAll, error) {
	rpras := &ResPartnerRelationAlls{}
	if err := c.SearchRead(ResPartnerRelationAllModel, criteria, odoo.NewOptions().Limit(1), rpras); err != nil {
		return nil, err
	}
	if rpras != nil && len(*rpras) > 0 {
		return &((*rpras)[0]), nil
	}
	return nil, fmt.Errorf("res.partner.relation.all was not found")
}

// FindResPartnerRelationAlls finds res.partner.relation.all records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerRelationAlls(criteria *odoo.Criteria, options *odoo.Options) (*ResPartnerRelationAlls, error) {
	rpras := &ResPartnerRelationAlls{}
	if err := c.SearchRead(ResPartnerRelationAllModel, criteria, options, rpras); err != nil {
		return nil, err
	}
	return rpras, nil
}

// FindResPartnerRelationAllIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerRelationAllIds(criteria *odoo.Criteria, options *odoo.Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerRelationAllModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerRelationAllId finds record id by querying it with criteria.
func (c *Client) FindResPartnerRelationAllId(criteria *odoo.Criteria, options *odoo.Options) (int64, error) {
	ids, err := c.Search(ResPartnerRelationAllModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner.relation.all was not found")
}
