package odoo

import (
	"fmt"
)

// ImLivechatReportChannel represents im_livechat.report.channel model.
type ImLivechatReportChannel struct {
	ChannelId         *Many2One `xmlrpc:"channel_id,omptempty"`
	ChannelName       *String   `xmlrpc:"channel_name,omptempty"`
	CountryId         *Many2One `xmlrpc:"country_id,omptempty"`
	DayNumber         *String   `xmlrpc:"day_number,omptempty"`
	DaysOfActivity    *Int      `xmlrpc:"days_of_activity,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Duration          *Float    `xmlrpc:"duration,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	IsAnonymous       *Int      `xmlrpc:"is_anonymous,omptempty"`
	IsHappy           *Int      `xmlrpc:"is_happy,omptempty"`
	IsUnrated         *Int      `xmlrpc:"is_unrated,omptempty"`
	IsWithoutAnswer   *Int      `xmlrpc:"is_without_answer,omptempty"`
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	LivechatChannelId *Many2One `xmlrpc:"livechat_channel_id,omptempty"`
	NbrMessage        *Int      `xmlrpc:"nbr_message,omptempty"`
	NbrSpeaker        *Int      `xmlrpc:"nbr_speaker,omptempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omptempty"`
	Rating            *Int      `xmlrpc:"rating,omptempty"`
	RatingText        *String   `xmlrpc:"rating_text,omptempty"`
	StartDate         *Time     `xmlrpc:"start_date,omptempty"`
	StartDateHour     *String   `xmlrpc:"start_date_hour,omptempty"`
	StartHour         *String   `xmlrpc:"start_hour,omptempty"`
	TechnicalName     *String   `xmlrpc:"technical_name,omptempty"`
	TimeToAnswer      *Float    `xmlrpc:"time_to_answer,omptempty"`
	Uuid              *String   `xmlrpc:"uuid,omptempty"`
}

// ImLivechatReportChannels represents array of im_livechat.report.channel model.
type ImLivechatReportChannels []ImLivechatReportChannel

// ImLivechatReportChannelModel is the odoo model name.
const ImLivechatReportChannelModel = "im_livechat.report.channel"

// Many2One convert ImLivechatReportChannel to *Many2One.
func (irc *ImLivechatReportChannel) Many2One() *Many2One {
	return NewMany2One(irc.Id.Get(), "")
}

// CreateImLivechatReportChannel creates a new im_livechat.report.channel model and returns its id.
func (c *Client) CreateImLivechatReportChannel(irc *ImLivechatReportChannel) (int64, error) {
	return c.Create(ImLivechatReportChannelModel, irc)
}

// UpdateImLivechatReportChannel updates an existing im_livechat.report.channel record.
func (c *Client) UpdateImLivechatReportChannel(irc *ImLivechatReportChannel) error {
	return c.UpdateImLivechatReportChannels([]int64{irc.Id.Get()}, irc)
}

// UpdateImLivechatReportChannels updates existing im_livechat.report.channel records.
// All records (represented by ids) will be updated by irc values.
func (c *Client) UpdateImLivechatReportChannels(ids []int64, irc *ImLivechatReportChannel) error {
	return c.Update(ImLivechatReportChannelModel, ids, irc)
}

// DeleteImLivechatReportChannel deletes an existing im_livechat.report.channel record.
func (c *Client) DeleteImLivechatReportChannel(id int64) error {
	return c.DeleteImLivechatReportChannels([]int64{id})
}

// DeleteImLivechatReportChannels deletes existing im_livechat.report.channel records.
func (c *Client) DeleteImLivechatReportChannels(ids []int64) error {
	return c.Delete(ImLivechatReportChannelModel, ids)
}

// GetImLivechatReportChannel gets im_livechat.report.channel existing record.
func (c *Client) GetImLivechatReportChannel(id int64) (*ImLivechatReportChannel, error) {
	ircs, err := c.GetImLivechatReportChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if ircs != nil && len(*ircs) > 0 {
		return &((*ircs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of im_livechat.report.channel not found", id)
}

// GetImLivechatReportChannels gets im_livechat.report.channel existing records.
func (c *Client) GetImLivechatReportChannels(ids []int64) (*ImLivechatReportChannels, error) {
	ircs := &ImLivechatReportChannels{}
	if err := c.Read(ImLivechatReportChannelModel, ids, nil, ircs); err != nil {
		return nil, err
	}
	return ircs, nil
}

// FindImLivechatReportChannel finds im_livechat.report.channel record by querying it with criteria.
func (c *Client) FindImLivechatReportChannel(criteria *Criteria) (*ImLivechatReportChannel, error) {
	ircs := &ImLivechatReportChannels{}
	if err := c.SearchRead(ImLivechatReportChannelModel, criteria, NewOptions().Limit(1), ircs); err != nil {
		return nil, err
	}
	if ircs != nil && len(*ircs) > 0 {
		return &((*ircs)[0]), nil
	}
	return nil, fmt.Errorf("im_livechat.report.channel was not found")
}

// FindImLivechatReportChannels finds im_livechat.report.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatReportChannels(criteria *Criteria, options *Options) (*ImLivechatReportChannels, error) {
	ircs := &ImLivechatReportChannels{}
	if err := c.SearchRead(ImLivechatReportChannelModel, criteria, options, ircs); err != nil {
		return nil, err
	}
	return ircs, nil
}

// FindImLivechatReportChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatReportChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ImLivechatReportChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindImLivechatReportChannelId finds record id by querying it with criteria.
func (c *Client) FindImLivechatReportChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatReportChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("im_livechat.report.channel was not found")
}
