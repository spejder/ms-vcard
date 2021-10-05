package odoo

import (
	"fmt"
)

// IrCron represents ir.cron model.
type IrCron struct {
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserFieldName         *String    `xmlrpc:"activity_user_field_name,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	ActivityUserType              *Selection `xmlrpc:"activity_user_type,omptempty"`
	BindingModelId                *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType                   *Selection `xmlrpc:"binding_type,omptempty"`
	BindingViewTypes              *String    `xmlrpc:"binding_view_types,omptempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	Code                          *String    `xmlrpc:"code,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CronName                      *String    `xmlrpc:"cron_name,omptempty"`
	CrudModelId                   *Many2One  `xmlrpc:"crud_model_id,omptempty"`
	CrudModelName                 *String    `xmlrpc:"crud_model_name,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	Doall                         *Bool      `xmlrpc:"doall,omptempty"`
	FieldsLines                   *Relation  `xmlrpc:"fields_lines,omptempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omptempty"`
	Help                          *String    `xmlrpc:"help,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	IntervalNumber                *Int       `xmlrpc:"interval_number,omptempty"`
	IntervalType                  *Selection `xmlrpc:"interval_type,omptempty"`
	IrActionsServerId             *Many2One  `xmlrpc:"ir_actions_server_id,omptempty"`
	Lastcall                      *Time      `xmlrpc:"lastcall,omptempty"`
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	LinkFieldId                   *Many2One  `xmlrpc:"link_field_id,omptempty"`
	ModelId                       *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName                     *String    `xmlrpc:"model_name,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	Nextcall                      *Time      `xmlrpc:"nextcall,omptempty"`
	Numbercall                    *Int       `xmlrpc:"numbercall,omptempty"`
	PartnerIds                    *Relation  `xmlrpc:"partner_ids,omptempty"`
	Priority                      *Int       `xmlrpc:"priority,omptempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omptempty"`
	SmsMassKeepLog                *Bool      `xmlrpc:"sms_mass_keep_log,omptempty"`
	SmsTemplateId                 *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	TemplateId                    *Many2One  `xmlrpc:"template_id,omptempty"`
	Type                          *String    `xmlrpc:"type,omptempty"`
	Usage                         *Selection `xmlrpc:"usage,omptempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsitePath                   *String    `xmlrpc:"website_path,omptempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId                         *String    `xmlrpc:"xml_id,omptempty"`
}

// IrCrons represents array of ir.cron model.
type IrCrons []IrCron

// IrCronModel is the odoo model name.
const IrCronModel = "ir.cron"

// Many2One convert IrCron to *Many2One.
func (ic *IrCron) Many2One() *Many2One {
	return NewMany2One(ic.Id.Get(), "")
}

// CreateIrCron creates a new ir.cron model and returns its id.
func (c *Client) CreateIrCron(ic *IrCron) (int64, error) {
	return c.Create(IrCronModel, ic)
}

// UpdateIrCron updates an existing ir.cron record.
func (c *Client) UpdateIrCron(ic *IrCron) error {
	return c.UpdateIrCrons([]int64{ic.Id.Get()}, ic)
}

// UpdateIrCrons updates existing ir.cron records.
// All records (represented by ids) will be updated by ic values.
func (c *Client) UpdateIrCrons(ids []int64, ic *IrCron) error {
	return c.Update(IrCronModel, ids, ic)
}

// DeleteIrCron deletes an existing ir.cron record.
func (c *Client) DeleteIrCron(id int64) error {
	return c.DeleteIrCrons([]int64{id})
}

// DeleteIrCrons deletes existing ir.cron records.
func (c *Client) DeleteIrCrons(ids []int64) error {
	return c.Delete(IrCronModel, ids)
}

// GetIrCron gets ir.cron existing record.
func (c *Client) GetIrCron(id int64) (*IrCron, error) {
	ics, err := c.GetIrCrons([]int64{id})
	if err != nil {
		return nil, err
	}
	if ics != nil && len(*ics) > 0 {
		return &((*ics)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.cron not found", id)
}

// GetIrCrons gets ir.cron existing records.
func (c *Client) GetIrCrons(ids []int64) (*IrCrons, error) {
	ics := &IrCrons{}
	if err := c.Read(IrCronModel, ids, nil, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindIrCron finds ir.cron record by querying it with criteria.
func (c *Client) FindIrCron(criteria *Criteria) (*IrCron, error) {
	ics := &IrCrons{}
	if err := c.SearchRead(IrCronModel, criteria, NewOptions().Limit(1), ics); err != nil {
		return nil, err
	}
	if ics != nil && len(*ics) > 0 {
		return &((*ics)[0]), nil
	}
	return nil, fmt.Errorf("ir.cron was not found")
}

// FindIrCrons finds ir.cron records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCrons(criteria *Criteria, options *Options) (*IrCrons, error) {
	ics := &IrCrons{}
	if err := c.SearchRead(IrCronModel, criteria, options, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindIrCronIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCronIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrCronModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrCronId finds record id by querying it with criteria.
func (c *Client) FindIrCronId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrCronModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.cron was not found")
}
