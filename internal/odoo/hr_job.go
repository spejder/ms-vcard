package odoo

import (
	"fmt"
)

// HrJob represents hr.job model.
type HrJob struct {
	AddressId                *Many2One  `xmlrpc:"address_id,omptempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId              *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	ApplicationCount         *Int       `xmlrpc:"application_count,omptempty"`
	ApplicationIds           *Relation  `xmlrpc:"application_ids,omptempty"`
	CanPublish               *Bool      `xmlrpc:"can_publish,omptempty"`
	Color                    *Int       `xmlrpc:"color,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	DocumentIds              *Relation  `xmlrpc:"document_ids,omptempty"`
	DocumentsCount           *Int       `xmlrpc:"documents_count,omptempty"`
	EmployeeIds              *Relation  `xmlrpc:"employee_ids,omptempty"`
	ExpectedEmployees        *Int       `xmlrpc:"expected_employees,omptempty"`
	FavoriteUserIds          *Relation  `xmlrpc:"favorite_user_ids,omptempty"`
	HrResponsibleId          *Many2One  `xmlrpc:"hr_responsible_id,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IsFavorite               *Bool      `xmlrpc:"is_favorite,omptempty"`
	IsPublished              *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized           *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	NewApplicationCount      *Int       `xmlrpc:"new_application_count,omptempty"`
	NoOfEmployee             *Int       `xmlrpc:"no_of_employee,omptempty"`
	NoOfHiredEmployee        *Int       `xmlrpc:"no_of_hired_employee,omptempty"`
	NoOfRecruitment          *Int       `xmlrpc:"no_of_recruitment,omptempty"`
	Requirements             *String    `xmlrpc:"requirements,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteDescription       *String    `xmlrpc:"website_description,omptempty"`
	WebsiteId                *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription   *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords      *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg         *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle         *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished         *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl               *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrJobs represents array of hr.job model.
type HrJobs []HrJob

// HrJobModel is the odoo model name.
const HrJobModel = "hr.job"

// Many2One convert HrJob to *Many2One.
func (hj *HrJob) Many2One() *Many2One {
	return NewMany2One(hj.Id.Get(), "")
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJob(hj *HrJob) (int64, error) {
	return c.Create(HrJobModel, hj)
}

// UpdateHrJob updates an existing hr.job record.
func (c *Client) UpdateHrJob(hj *HrJob) error {
	return c.UpdateHrJobs([]int64{hj.Id.Get()}, hj)
}

// UpdateHrJobs updates existing hr.job records.
// All records (represented by ids) will be updated by hj values.
func (c *Client) UpdateHrJobs(ids []int64, hj *HrJob) error {
	return c.Update(HrJobModel, ids, hj)
}

// DeleteHrJob deletes an existing hr.job record.
func (c *Client) DeleteHrJob(id int64) error {
	return c.DeleteHrJobs([]int64{id})
}

// DeleteHrJobs deletes existing hr.job records.
func (c *Client) DeleteHrJobs(ids []int64) error {
	return c.Delete(HrJobModel, ids)
}

// GetHrJob gets hr.job existing record.
func (c *Client) GetHrJob(id int64) (*HrJob, error) {
	hjs, err := c.GetHrJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	if hjs != nil && len(*hjs) > 0 {
		return &((*hjs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.job not found", id)
}

// GetHrJobs gets hr.job existing records.
func (c *Client) GetHrJobs(ids []int64) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.Read(HrJobModel, ids, nil, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJob finds hr.job record by querying it with criteria.
func (c *Client) FindHrJob(criteria *Criteria) (*HrJob, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, NewOptions().Limit(1), hjs); err != nil {
		return nil, err
	}
	if hjs != nil && len(*hjs) > 0 {
		return &((*hjs)[0]), nil
	}
	return nil, fmt.Errorf("hr.job was not found")
}

// FindHrJobs finds hr.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobs(criteria *Criteria, options *Options) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, options, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrJobId finds record id by querying it with criteria.
func (c *Client) FindHrJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.job was not found")
}
