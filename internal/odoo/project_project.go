package odoo

import (
	"fmt"
)

// ProjectProject represents project.project model.
type ProjectProject struct {
	AccessToken                  *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                    *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning                *String    `xmlrpc:"access_warning,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	AliasContact                 *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults                *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain                  *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId           *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                      *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId                 *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                    *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId           *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId          *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId                  *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	AnalyticAccountId            *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	Color                        *Int       `xmlrpc:"color,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                         *Time      `xmlrpc:"date,omptempty"`
	DateStart                    *Time      `xmlrpc:"date_start,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	DocCount                     *Int       `xmlrpc:"doc_count,omptempty"`
	FavoriteUserIds              *Relation  `xmlrpc:"favorite_user_ids,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	IsFavorite                   *Bool      `xmlrpc:"is_favorite,omptempty"`
	LabelTasks                   *String    `xmlrpc:"label_tasks,omptempty"`
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds            *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId      *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter         *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omptempty"`
	PortalShowRating             *Bool      `xmlrpc:"portal_show_rating,omptempty"`
	PrivacyVisibility            *Selection `xmlrpc:"privacy_visibility,omptempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingPercentageSatisfaction *Int       `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RatingRequestDeadline        *Time      `xmlrpc:"rating_request_deadline,omptempty"`
	RatingStatus                 *Selection `xmlrpc:"rating_status,omptempty"`
	RatingStatusPeriod           *Selection `xmlrpc:"rating_status_period,omptempty"`
	ResourceCalendarId           *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omptempty"`
	SubtaskProjectId             *Many2One  `xmlrpc:"subtask_project_id,omptempty"`
	TaskCount                    *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                      *Relation  `xmlrpc:"task_ids,omptempty"`
	Tasks                        *Relation  `xmlrpc:"tasks,omptempty"`
	TypeIds                      *Relation  `xmlrpc:"type_ids,omptempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProjectProjects represents array of project.project model.
type ProjectProjects []ProjectProject

// ProjectProjectModel is the odoo model name.
const ProjectProjectModel = "project.project"

// Many2One convert ProjectProject to *Many2One.
func (pp *ProjectProject) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProject(pp *ProjectProject) (int64, error) {
	return c.Create(ProjectProjectModel, pp)
}

// UpdateProjectProject updates an existing project.project record.
func (c *Client) UpdateProjectProject(pp *ProjectProject) error {
	return c.UpdateProjectProjects([]int64{pp.Id.Get()}, pp)
}

// UpdateProjectProjects updates existing project.project records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProjectProjects(ids []int64, pp *ProjectProject) error {
	return c.Update(ProjectProjectModel, ids, pp)
}

// DeleteProjectProject deletes an existing project.project record.
func (c *Client) DeleteProjectProject(id int64) error {
	return c.DeleteProjectProjects([]int64{id})
}

// DeleteProjectProjects deletes existing project.project records.
func (c *Client) DeleteProjectProjects(ids []int64) error {
	return c.Delete(ProjectProjectModel, ids)
}

// GetProjectProject gets project.project existing record.
func (c *Client) GetProjectProject(id int64) (*ProjectProject, error) {
	pps, err := c.GetProjectProjects([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.project not found", id)
}

// GetProjectProjects gets project.project existing records.
func (c *Client) GetProjectProjects(ids []int64) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.Read(ProjectProjectModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProject finds project.project record by querying it with criteria.
func (c *Client) FindProjectProject(criteria *Criteria) (*ProjectProject, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("project.project was not found")
}

// FindProjectProjects finds project.project records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjects(criteria *Criteria, options *Options) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProjectIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectProjectId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.project was not found")
}
