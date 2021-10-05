package odoo

import (
	"fmt"
)

// SurveyInvite represents survey.invite model.
type SurveyInvite struct {
	AttachmentIds            *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorId                 *Many2One  `xmlrpc:"author_id,omptempty"`
	Body                     *String    `xmlrpc:"body,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Deadline                 *Time      `xmlrpc:"deadline,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	EmailFrom                *String    `xmlrpc:"email_from,omptempty"`
	Emails                   *String    `xmlrpc:"emails,omptempty"`
	ExistingEmails           *String    `xmlrpc:"existing_emails,omptempty"`
	ExistingMode             *Selection `xmlrpc:"existing_mode,omptempty"`
	ExistingPartnerIds       *Relation  `xmlrpc:"existing_partner_ids,omptempty"`
	ExistingText             *String    `xmlrpc:"existing_text,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	MailServerId             *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omptempty"`
	Subject                  *String    `xmlrpc:"subject,omptempty"`
	SurveyAccessMode         *Selection `xmlrpc:"survey_access_mode,omptempty"`
	SurveyId                 *Many2One  `xmlrpc:"survey_id,omptempty"`
	SurveyUrl                *String    `xmlrpc:"survey_url,omptempty"`
	SurveyUsersLoginRequired *Bool      `xmlrpc:"survey_users_login_required,omptempty"`
	TemplateId               *Many2One  `xmlrpc:"template_id,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyInvites represents array of survey.invite model.
type SurveyInvites []SurveyInvite

// SurveyInviteModel is the odoo model name.
const SurveyInviteModel = "survey.invite"

// Many2One convert SurveyInvite to *Many2One.
func (si *SurveyInvite) Many2One() *Many2One {
	return NewMany2One(si.Id.Get(), "")
}

// CreateSurveyInvite creates a new survey.invite model and returns its id.
func (c *Client) CreateSurveyInvite(si *SurveyInvite) (int64, error) {
	return c.Create(SurveyInviteModel, si)
}

// UpdateSurveyInvite updates an existing survey.invite record.
func (c *Client) UpdateSurveyInvite(si *SurveyInvite) error {
	return c.UpdateSurveyInvites([]int64{si.Id.Get()}, si)
}

// UpdateSurveyInvites updates existing survey.invite records.
// All records (represented by ids) will be updated by si values.
func (c *Client) UpdateSurveyInvites(ids []int64, si *SurveyInvite) error {
	return c.Update(SurveyInviteModel, ids, si)
}

// DeleteSurveyInvite deletes an existing survey.invite record.
func (c *Client) DeleteSurveyInvite(id int64) error {
	return c.DeleteSurveyInvites([]int64{id})
}

// DeleteSurveyInvites deletes existing survey.invite records.
func (c *Client) DeleteSurveyInvites(ids []int64) error {
	return c.Delete(SurveyInviteModel, ids)
}

// GetSurveyInvite gets survey.invite existing record.
func (c *Client) GetSurveyInvite(id int64) (*SurveyInvite, error) {
	sis, err := c.GetSurveyInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	if sis != nil && len(*sis) > 0 {
		return &((*sis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.invite not found", id)
}

// GetSurveyInvites gets survey.invite existing records.
func (c *Client) GetSurveyInvites(ids []int64) (*SurveyInvites, error) {
	sis := &SurveyInvites{}
	if err := c.Read(SurveyInviteModel, ids, nil, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindSurveyInvite finds survey.invite record by querying it with criteria.
func (c *Client) FindSurveyInvite(criteria *Criteria) (*SurveyInvite, error) {
	sis := &SurveyInvites{}
	if err := c.SearchRead(SurveyInviteModel, criteria, NewOptions().Limit(1), sis); err != nil {
		return nil, err
	}
	if sis != nil && len(*sis) > 0 {
		return &((*sis)[0]), nil
	}
	return nil, fmt.Errorf("survey.invite was not found")
}

// FindSurveyInvites finds survey.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyInvites(criteria *Criteria, options *Options) (*SurveyInvites, error) {
	sis := &SurveyInvites{}
	if err := c.SearchRead(SurveyInviteModel, criteria, options, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindSurveyInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyInviteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyInviteId finds record id by querying it with criteria.
func (c *Client) FindSurveyInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.invite was not found")
}
