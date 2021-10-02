package odoo

import (
	"fmt"
)

// OpenacademyWizard represents openacademy.wizard model.
type OpenacademyWizard struct {
	AttendeeIds *Relation `xmlrpc:"attendee_ids,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	SessionIds  *Relation `xmlrpc:"session_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OpenacademyWizards represents array of openacademy.wizard model.
type OpenacademyWizards []OpenacademyWizard

// OpenacademyWizardModel is the odoo model name.
const OpenacademyWizardModel = "openacademy.wizard"

// Many2One convert OpenacademyWizard to *Many2One.
func (ow *OpenacademyWizard) Many2One() *Many2One {
	return NewMany2One(ow.Id.Get(), "")
}

// CreateOpenacademyWizard creates a new openacademy.wizard model and returns its id.
func (c *Client) CreateOpenacademyWizard(ow *OpenacademyWizard) (int64, error) {
	return c.Create(OpenacademyWizardModel, ow)
}

// UpdateOpenacademyWizard updates an existing openacademy.wizard record.
func (c *Client) UpdateOpenacademyWizard(ow *OpenacademyWizard) error {
	return c.UpdateOpenacademyWizards([]int64{ow.Id.Get()}, ow)
}

// UpdateOpenacademyWizards updates existing openacademy.wizard records.
// All records (represented by ids) will be updated by ow values.
func (c *Client) UpdateOpenacademyWizards(ids []int64, ow *OpenacademyWizard) error {
	return c.Update(OpenacademyWizardModel, ids, ow)
}

// DeleteOpenacademyWizard deletes an existing openacademy.wizard record.
func (c *Client) DeleteOpenacademyWizard(id int64) error {
	return c.DeleteOpenacademyWizards([]int64{id})
}

// DeleteOpenacademyWizards deletes existing openacademy.wizard records.
func (c *Client) DeleteOpenacademyWizards(ids []int64) error {
	return c.Delete(OpenacademyWizardModel, ids)
}

// GetOpenacademyWizard gets openacademy.wizard existing record.
func (c *Client) GetOpenacademyWizard(id int64) (*OpenacademyWizard, error) {
	ows, err := c.GetOpenacademyWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ows != nil && len(*ows) > 0 {
		return &((*ows)[0]), nil
	}
	return nil, fmt.Errorf("id %v of openacademy.wizard not found", id)
}

// GetOpenacademyWizards gets openacademy.wizard existing records.
func (c *Client) GetOpenacademyWizards(ids []int64) (*OpenacademyWizards, error) {
	ows := &OpenacademyWizards{}
	if err := c.Read(OpenacademyWizardModel, ids, nil, ows); err != nil {
		return nil, err
	}
	return ows, nil
}

// FindOpenacademyWizard finds openacademy.wizard record by querying it with criteria.
func (c *Client) FindOpenacademyWizard(criteria *Criteria) (*OpenacademyWizard, error) {
	ows := &OpenacademyWizards{}
	if err := c.SearchRead(OpenacademyWizardModel, criteria, NewOptions().Limit(1), ows); err != nil {
		return nil, err
	}
	if ows != nil && len(*ows) > 0 {
		return &((*ows)[0]), nil
	}
	return nil, fmt.Errorf("openacademy.wizard was not found")
}

// FindOpenacademyWizards finds openacademy.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademyWizards(criteria *Criteria, options *Options) (*OpenacademyWizards, error) {
	ows := &OpenacademyWizards{}
	if err := c.SearchRead(OpenacademyWizardModel, criteria, options, ows); err != nil {
		return nil, err
	}
	return ows, nil
}

// FindOpenacademyWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOpenacademyWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OpenacademyWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOpenacademyWizardId finds record id by querying it with criteria.
func (c *Client) FindOpenacademyWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OpenacademyWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("openacademy.wizard was not found")
}
