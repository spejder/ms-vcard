package odoo

import (
	"fmt"
)

// ProjectTask represents project.task model.
type ProjectTask struct {
	AccessToken                 *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omptempty"`
	ChildIds                    *Relation  `xmlrpc:"child_ids,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateAssign                  *Time      `xmlrpc:"date_assign,omptempty"`
	DateDeadline                *Time      `xmlrpc:"date_deadline,omptempty"`
	DateDeadlineFormatted       *String    `xmlrpc:"date_deadline_formatted,omptempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omptempty"`
	DateLastStageUpdate         *Time      `xmlrpc:"date_last_stage_update,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayedImageId            *Many2One  `xmlrpc:"displayed_image_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omptempty"`
	EmailFrom                   *String    `xmlrpc:"email_from,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omptempty"`
	KanbanStateLabel            *String    `xmlrpc:"kanban_state_label,omptempty"`
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	LegendBlocked               *String    `xmlrpc:"legend_blocked,omptempty"`
	LegendDone                  *String    `xmlrpc:"legend_done,omptempty"`
	LegendNormal                *String    `xmlrpc:"legend_normal,omptempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerCity                 *String    `xmlrpc:"partner_city,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PlannedHours                *Float     `xmlrpc:"planned_hours,omptempty"`
	Priority                    *Selection `xmlrpc:"priority,omptempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omptempty"`
	RatingAvg                   *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingCount                 *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback          *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage             *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastValue             *Float     `xmlrpc:"rating_last_value,omptempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omptempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omptempty"`
	SubtaskCount                *Int       `xmlrpc:"subtask_count,omptempty"`
	SubtaskPlannedHours         *Float     `xmlrpc:"subtask_planned_hours,omptempty"`
	SubtaskProjectId            *Many2One  `xmlrpc:"subtask_project_id,omptempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omptempty"`
	UserEmail                   *String    `xmlrpc:"user_email,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WorkingDaysClose            *Float     `xmlrpc:"working_days_close,omptempty"`
	WorkingDaysOpen             *Float     `xmlrpc:"working_days_open,omptempty"`
	WorkingHoursClose           *Float     `xmlrpc:"working_hours_close,omptempty"`
	WorkingHoursOpen            *Float     `xmlrpc:"working_hours_open,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProjectTasks represents array of project.task model.
type ProjectTasks []ProjectTask

// ProjectTaskModel is the odoo model name.
const ProjectTaskModel = "project.task"

// Many2One convert ProjectTask to *Many2One.
func (pt *ProjectTask) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTask(pt *ProjectTask) (int64, error) {
	return c.Create(ProjectTaskModel, pt)
}

// UpdateProjectTask updates an existing project.task record.
func (c *Client) UpdateProjectTask(pt *ProjectTask) error {
	return c.UpdateProjectTasks([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTasks updates existing project.task records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTasks(ids []int64, pt *ProjectTask) error {
	return c.Update(ProjectTaskModel, ids, pt)
}

// DeleteProjectTask deletes an existing project.task record.
func (c *Client) DeleteProjectTask(id int64) error {
	return c.DeleteProjectTasks([]int64{id})
}

// DeleteProjectTasks deletes existing project.task records.
func (c *Client) DeleteProjectTasks(ids []int64) error {
	return c.Delete(ProjectTaskModel, ids)
}

// GetProjectTask gets project.task existing record.
func (c *Client) GetProjectTask(id int64) (*ProjectTask, error) {
	pts, err := c.GetProjectTasks([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task not found", id)
}

// GetProjectTasks gets project.task existing records.
func (c *Client) GetProjectTasks(ids []int64) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.Read(ProjectTaskModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTask finds project.task record by querying it with criteria.
func (c *Client) FindProjectTask(criteria *Criteria) (*ProjectTask, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("project.task was not found")
}

// FindProjectTasks finds project.task records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTasks(criteria *Criteria, options *Options) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTaskIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task was not found")
}
