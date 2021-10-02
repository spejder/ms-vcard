package odoo

import (
	"fmt"
)

// SlideSlide represents slide.slide model.
type SlideSlide struct {
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	CanPublish               *Bool      `xmlrpc:"can_publish,omptempty"`
	CategoryId               *Many2One  `xmlrpc:"category_id,omptempty"`
	ChannelAllowComment      *Bool      `xmlrpc:"channel_allow_comment,omptempty"`
	ChannelId                *Many2One  `xmlrpc:"channel_id,omptempty"`
	ChannelType              *Selection `xmlrpc:"channel_type,omptempty"`
	CommentsCount            *Int       `xmlrpc:"comments_count,omptempty"`
	CompletionTime           *Float     `xmlrpc:"completion_time,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Datas                    *String    `xmlrpc:"datas,omptempty"`
	DatePublished            *Time      `xmlrpc:"date_published,omptempty"`
	DeletedAt                *Time      `xmlrpc:"deleted_at,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	Dislikes                 *Int       `xmlrpc:"dislikes,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	DocumentId               *String    `xmlrpc:"document_id,omptempty"`
	EmbedCode                *String    `xmlrpc:"embed_code,omptempty"`
	EmbedcountIds            *Relation  `xmlrpc:"embedcount_ids,omptempty"`
	HtmlContent              *String    `xmlrpc:"html_content,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Image1024                *String    `xmlrpc:"image_1024,omptempty"`
	Image128                 *String    `xmlrpc:"image_128,omptempty"`
	Image1920                *String    `xmlrpc:"image_1920,omptempty"`
	Image256                 *String    `xmlrpc:"image_256,omptempty"`
	Image512                 *String    `xmlrpc:"image_512,omptempty"`
	IsCategory               *Bool      `xmlrpc:"is_category,omptempty"`
	IsPreview                *Bool      `xmlrpc:"is_preview,omptempty"`
	IsPublished              *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized           *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Likes                    *Int       `xmlrpc:"likes,omptempty"`
	LinkIds                  *Relation  `xmlrpc:"link_ids,omptempty"`
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
	MimeType                 *String    `xmlrpc:"mime_type,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	NbrDocument              *Int       `xmlrpc:"nbr_document,omptempty"`
	NbrInfographic           *Int       `xmlrpc:"nbr_infographic,omptempty"`
	NbrPresentation          *Int       `xmlrpc:"nbr_presentation,omptempty"`
	NbrQuiz                  *Int       `xmlrpc:"nbr_quiz,omptempty"`
	NbrVideo                 *Int       `xmlrpc:"nbr_video,omptempty"`
	NbrWebpage               *Int       `xmlrpc:"nbr_webpage,omptempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omptempty"`
	PublicViews              *Int       `xmlrpc:"public_views,omptempty"`
	QuestionIds              *Relation  `xmlrpc:"question_ids,omptempty"`
	QuestionsCount           *Int       `xmlrpc:"questions_count,omptempty"`
	QuizFirstAttemptReward   *Int       `xmlrpc:"quiz_first_attempt_reward,omptempty"`
	QuizFourthAttemptReward  *Int       `xmlrpc:"quiz_fourth_attempt_reward,omptempty"`
	QuizSecondAttemptReward  *Int       `xmlrpc:"quiz_second_attempt_reward,omptempty"`
	QuizThirdAttemptReward   *Int       `xmlrpc:"quiz_third_attempt_reward,omptempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omptempty"`
	SubSequence              *Int       `xmlrpc:"sub_sequence,omptempty"`
	SlideIds                 *Relation  `xmlrpc:"slide_ids,omptempty"`
	SlidePartnerIds          *Relation  `xmlrpc:"slide_partner_ids,omptempty"`
	SlideType                *Selection `xmlrpc:"slide_type,omptempty"`
	SlideViews               *Int       `xmlrpc:"slide_views,omptempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omptempty"`
	TotalSlides              *Int       `xmlrpc:"total_slides,omptempty"`
	TotalViews               *Int       `xmlrpc:"total_views,omptempty"`
	QuizPassRate             *Int       `xmlrpc:"quiz_pass_rate,omptempty"`
	Url                      *String    `xmlrpc:"url,omptempty"`
	ImageUrl                 *String    `xmlrpc:"image_url,omptempty"`
	DocumentUrl              *String    `xmlrpc:"document_url,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	UserMembershipId         *Many2One  `xmlrpc:"user_membership_id,omptempty"`
	UserVote                 *Int       `xmlrpc:"user_vote,omptempty"`
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

// SlideSlides represents array of slide.slide model.
type SlideSlides []SlideSlide

// SlideSlideModel is the odoo model name.
const SlideSlideModel = "slide.slide"

// Many2One convert SlideSlide to *Many2One.
func (ss *SlideSlide) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSlideSlide creates a new slide.slide model and returns its id.
func (c *Client) CreateSlideSlide(ss *SlideSlide) (int64, error) {
	return c.Create(SlideSlideModel, ss)
}

// UpdateSlideSlide updates an existing slide.slide record.
func (c *Client) UpdateSlideSlide(ss *SlideSlide) error {
	return c.UpdateSlideSlides([]int64{ss.Id.Get()}, ss)
}

// UpdateSlideSlides updates existing slide.slide records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSlideSlides(ids []int64, ss *SlideSlide) error {
	return c.Update(SlideSlideModel, ids, ss)
}

// DeleteSlideSlide deletes an existing slide.slide record.
func (c *Client) DeleteSlideSlide(id int64) error {
	return c.DeleteSlideSlides([]int64{id})
}

// DeleteSlideSlides deletes existing slide.slide records.
func (c *Client) DeleteSlideSlides(ids []int64) error {
	return c.Delete(SlideSlideModel, ids)
}

// GetSlideSlide gets slide.slide existing record.
func (c *Client) GetSlideSlide(id int64) (*SlideSlide, error) {
	sss, err := c.GetSlideSlides([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide not found", id)
}

// GetSlideSlides gets slide.slide existing records.
func (c *Client) GetSlideSlides(ids []int64) (*SlideSlides, error) {
	sss := &SlideSlides{}
	if err := c.Read(SlideSlideModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlide finds slide.slide record by querying it with criteria.
func (c *Client) FindSlideSlide(criteria *Criteria) (*SlideSlide, error) {
	sss := &SlideSlides{}
	if err := c.SearchRead(SlideSlideModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide was not found")
}

// FindSlideSlides finds slide.slide records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlides(criteria *Criteria, options *Options) (*SlideSlides, error) {
	sss := &SlideSlides{}
	if err := c.SearchRead(SlideSlideModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlideIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlideId finds record id by querying it with criteria.
func (c *Client) FindSlideSlideId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide was not found")
}
