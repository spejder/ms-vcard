package odoo

import (
	"fmt"
)

// ImLivechatChannel represents im_livechat.channel model.
type ImLivechatChannel struct {
	AreYouInside                 *Bool     `xmlrpc:"are_you_inside,omptempty"`
	ButtonText                   *String   `xmlrpc:"button_text,omptempty"`
	CanPublish                   *Bool     `xmlrpc:"can_publish,omptempty"`
	ChannelIds                   *Relation `xmlrpc:"channel_ids,omptempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omptempty"`
	DefaultMessage               *String   `xmlrpc:"default_message,omptempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	Image128                     *String   `xmlrpc:"image_128,omptempty"`
	InputPlaceholder             *String   `xmlrpc:"input_placeholder,omptempty"`
	IsPublished                  *Bool     `xmlrpc:"is_published,omptempty"`
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	Name                         *String   `xmlrpc:"name,omptempty"`
	NbrChannel                   *Int      `xmlrpc:"nbr_channel,omptempty"`
	RatingIds                    *Relation `xmlrpc:"rating_ids,omptempty"`
	RatingPercentageSatisfaction *Int      `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RuleIds                      *Relation `xmlrpc:"rule_ids,omptempty"`
	ScriptExternal               *String   `xmlrpc:"script_external,omptempty"`
	UserIds                      *Relation `xmlrpc:"user_ids,omptempty"`
	WebPage                      *String   `xmlrpc:"web_page,omptempty"`
	WebsiteDescription           *String   `xmlrpc:"website_description,omptempty"`
	WebsitePublished             *Bool     `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                   *String   `xmlrpc:"website_url,omptempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ImLivechatChannels represents array of im_livechat.channel model.
type ImLivechatChannels []ImLivechatChannel

// ImLivechatChannelModel is the odoo model name.
const ImLivechatChannelModel = "im_livechat.channel"

// Many2One convert ImLivechatChannel to *Many2One.
func (ic *ImLivechatChannel) Many2One() *Many2One {
	return NewMany2One(ic.Id.Get(), "")
}

// CreateImLivechatChannel creates a new im_livechat.channel model and returns its id.
func (c *Client) CreateImLivechatChannel(ic *ImLivechatChannel) (int64, error) {
	return c.Create(ImLivechatChannelModel, ic)
}

// UpdateImLivechatChannel updates an existing im_livechat.channel record.
func (c *Client) UpdateImLivechatChannel(ic *ImLivechatChannel) error {
	return c.UpdateImLivechatChannels([]int64{ic.Id.Get()}, ic)
}

// UpdateImLivechatChannels updates existing im_livechat.channel records.
// All records (represented by ids) will be updated by ic values.
func (c *Client) UpdateImLivechatChannels(ids []int64, ic *ImLivechatChannel) error {
	return c.Update(ImLivechatChannelModel, ids, ic)
}

// DeleteImLivechatChannel deletes an existing im_livechat.channel record.
func (c *Client) DeleteImLivechatChannel(id int64) error {
	return c.DeleteImLivechatChannels([]int64{id})
}

// DeleteImLivechatChannels deletes existing im_livechat.channel records.
func (c *Client) DeleteImLivechatChannels(ids []int64) error {
	return c.Delete(ImLivechatChannelModel, ids)
}

// GetImLivechatChannel gets im_livechat.channel existing record.
func (c *Client) GetImLivechatChannel(id int64) (*ImLivechatChannel, error) {
	ics, err := c.GetImLivechatChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if ics != nil && len(*ics) > 0 {
		return &((*ics)[0]), nil
	}
	return nil, fmt.Errorf("id %v of im_livechat.channel not found", id)
}

// GetImLivechatChannels gets im_livechat.channel existing records.
func (c *Client) GetImLivechatChannels(ids []int64) (*ImLivechatChannels, error) {
	ics := &ImLivechatChannels{}
	if err := c.Read(ImLivechatChannelModel, ids, nil, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindImLivechatChannel finds im_livechat.channel record by querying it with criteria.
func (c *Client) FindImLivechatChannel(criteria *Criteria) (*ImLivechatChannel, error) {
	ics := &ImLivechatChannels{}
	if err := c.SearchRead(ImLivechatChannelModel, criteria, NewOptions().Limit(1), ics); err != nil {
		return nil, err
	}
	if ics != nil && len(*ics) > 0 {
		return &((*ics)[0]), nil
	}
	return nil, fmt.Errorf("im_livechat.channel was not found")
}

// FindImLivechatChannels finds im_livechat.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannels(criteria *Criteria, options *Options) (*ImLivechatChannels, error) {
	ics := &ImLivechatChannels{}
	if err := c.SearchRead(ImLivechatChannelModel, criteria, options, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindImLivechatChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ImLivechatChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindImLivechatChannelId finds record id by querying it with criteria.
func (c *Client) FindImLivechatChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("im_livechat.channel was not found")
}
