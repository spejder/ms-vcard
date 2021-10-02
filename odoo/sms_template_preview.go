package odoo

import (
	"fmt"
)

// SmsTemplatePreview represents sms.template.preview model.
type SmsTemplatePreview struct {
	Body                *String    `xmlrpc:"body,omptempty"`
	Copyvalue           *String    `xmlrpc:"copyvalue,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	Lang                *Selection `xmlrpc:"lang,omptempty"`
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Model               *String    `xmlrpc:"model,omptempty"`
	ModelId             *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelObjectField    *Many2One  `xmlrpc:"model_object_field,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	NullValue           *String    `xmlrpc:"null_value,omptempty"`
	ResId               *Int       `xmlrpc:"res_id,omptempty"`
	ResourceRef         *String    `xmlrpc:"resource_ref,omptempty"`
	SidebarActionId     *Many2One  `xmlrpc:"sidebar_action_id,omptempty"`
	SmsTemplateId       *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	SubModelObjectField *Many2One  `xmlrpc:"sub_model_object_field,omptempty"`
	SubObject           *Many2One  `xmlrpc:"sub_object,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SmsTemplatePreviews represents array of sms.template.preview model.
type SmsTemplatePreviews []SmsTemplatePreview

// SmsTemplatePreviewModel is the odoo model name.
const SmsTemplatePreviewModel = "sms.template.preview"

// Many2One convert SmsTemplatePreview to *Many2One.
func (stp *SmsTemplatePreview) Many2One() *Many2One {
	return NewMany2One(stp.Id.Get(), "")
}

// CreateSmsTemplatePreview creates a new sms.template.preview model and returns its id.
func (c *Client) CreateSmsTemplatePreview(stp *SmsTemplatePreview) (int64, error) {
	return c.Create(SmsTemplatePreviewModel, stp)
}

// UpdateSmsTemplatePreview updates an existing sms.template.preview record.
func (c *Client) UpdateSmsTemplatePreview(stp *SmsTemplatePreview) error {
	return c.UpdateSmsTemplatePreviews([]int64{stp.Id.Get()}, stp)
}

// UpdateSmsTemplatePreviews updates existing sms.template.preview records.
// All records (represented by ids) will be updated by stp values.
func (c *Client) UpdateSmsTemplatePreviews(ids []int64, stp *SmsTemplatePreview) error {
	return c.Update(SmsTemplatePreviewModel, ids, stp)
}

// DeleteSmsTemplatePreview deletes an existing sms.template.preview record.
func (c *Client) DeleteSmsTemplatePreview(id int64) error {
	return c.DeleteSmsTemplatePreviews([]int64{id})
}

// DeleteSmsTemplatePreviews deletes existing sms.template.preview records.
func (c *Client) DeleteSmsTemplatePreviews(ids []int64) error {
	return c.Delete(SmsTemplatePreviewModel, ids)
}

// GetSmsTemplatePreview gets sms.template.preview existing record.
func (c *Client) GetSmsTemplatePreview(id int64) (*SmsTemplatePreview, error) {
	stps, err := c.GetSmsTemplatePreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if stps != nil && len(*stps) > 0 {
		return &((*stps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.template.preview not found", id)
}

// GetSmsTemplatePreviews gets sms.template.preview existing records.
func (c *Client) GetSmsTemplatePreviews(ids []int64) (*SmsTemplatePreviews, error) {
	stps := &SmsTemplatePreviews{}
	if err := c.Read(SmsTemplatePreviewModel, ids, nil, stps); err != nil {
		return nil, err
	}
	return stps, nil
}

// FindSmsTemplatePreview finds sms.template.preview record by querying it with criteria.
func (c *Client) FindSmsTemplatePreview(criteria *Criteria) (*SmsTemplatePreview, error) {
	stps := &SmsTemplatePreviews{}
	if err := c.SearchRead(SmsTemplatePreviewModel, criteria, NewOptions().Limit(1), stps); err != nil {
		return nil, err
	}
	if stps != nil && len(*stps) > 0 {
		return &((*stps)[0]), nil
	}
	return nil, fmt.Errorf("sms.template.preview was not found")
}

// FindSmsTemplatePreviews finds sms.template.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplatePreviews(criteria *Criteria, options *Options) (*SmsTemplatePreviews, error) {
	stps := &SmsTemplatePreviews{}
	if err := c.SearchRead(SmsTemplatePreviewModel, criteria, options, stps); err != nil {
		return nil, err
	}
	return stps, nil
}

// FindSmsTemplatePreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplatePreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsTemplatePreviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindSmsTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.template.preview was not found")
}
