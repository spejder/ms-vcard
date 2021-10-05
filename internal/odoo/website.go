package odoo

import (
	"fmt"
)

// Website represents website model.
type Website struct {
	AuthSignupUninvited          *Selection `xmlrpc:"auth_signup_uninvited,omptempty"`
	AutoRedirectLang             *Bool      `xmlrpc:"auto_redirect_lang,omptempty"`
	CdnActivated                 *Bool      `xmlrpc:"cdn_activated,omptempty"`
	CdnFilters                   *String    `xmlrpc:"cdn_filters,omptempty"`
	CdnUrl                       *String    `xmlrpc:"cdn_url,omptempty"`
	ChannelId                    *Many2One  `xmlrpc:"channel_id,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryGroupIds              *Relation  `xmlrpc:"country_group_ids,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmDefaultTeamId             *Many2One  `xmlrpc:"crm_default_team_id,omptempty"`
	CrmDefaultUserId             *Many2One  `xmlrpc:"crm_default_user_id,omptempty"`
	DefaultLangId                *Many2One  `xmlrpc:"default_lang_id,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	Domain                       *String    `xmlrpc:"domain,omptempty"`
	Favicon                      *String    `xmlrpc:"favicon,omptempty"`
	GoogleAnalyticsKey           *String    `xmlrpc:"google_analytics_key,omptempty"`
	GoogleManagementClientId     *String    `xmlrpc:"google_management_client_id,omptempty"`
	GoogleManagementClientSecret *String    `xmlrpc:"google_management_client_secret,omptempty"`
	GoogleMapsApiKey             *String    `xmlrpc:"google_maps_api_key,omptempty"`
	HomepageId                   *Many2One  `xmlrpc:"homepage_id,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	KarmaProfileMin              *Int       `xmlrpc:"karma_profile_min,omptempty"`
	LanguageIds                  *Relation  `xmlrpc:"language_ids,omptempty"`
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	Logo                         *String    `xmlrpc:"logo,omptempty"`
	MenuId                       *Many2One  `xmlrpc:"menu_id,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omptempty"`
	SocialDefaultImage           *String    `xmlrpc:"social_default_image,omptempty"`
	SocialFacebook               *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                 *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram              *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin               *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                *String    `xmlrpc:"social_youtube,omptempty"`
	SpecificUserAccount          *Bool      `xmlrpc:"specific_user_account,omptempty"`
	ThemeId                      *Many2One  `xmlrpc:"theme_id,omptempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteFormEnableMetadata    *Bool      `xmlrpc:"website_form_enable_metadata,omptempty"`
	WebsiteSlideGoogleAppKey     *String    `xmlrpc:"website_slide_google_app_key,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
	ZoomEmail                    *String    `xmlrpc:"zoom_email,omptempty"`
}

// Websites represents array of website model.
type Websites []Website

// WebsiteModel is the odoo model name.
const WebsiteModel = "website"

// Many2One convert Website to *Many2One.
func (w *Website) Many2One() *Many2One {
	return NewMany2One(w.Id.Get(), "")
}

// CreateWebsite creates a new website model and returns its id.
func (c *Client) CreateWebsite(w *Website) (int64, error) {
	return c.Create(WebsiteModel, w)
}

// UpdateWebsite updates an existing website record.
func (c *Client) UpdateWebsite(w *Website) error {
	return c.UpdateWebsites([]int64{w.Id.Get()}, w)
}

// UpdateWebsites updates existing website records.
// All records (represented by ids) will be updated by w values.
func (c *Client) UpdateWebsites(ids []int64, w *Website) error {
	return c.Update(WebsiteModel, ids, w)
}

// DeleteWebsite deletes an existing website record.
func (c *Client) DeleteWebsite(id int64) error {
	return c.DeleteWebsites([]int64{id})
}

// DeleteWebsites deletes existing website records.
func (c *Client) DeleteWebsites(ids []int64) error {
	return c.Delete(WebsiteModel, ids)
}

// GetWebsite gets website existing record.
func (c *Client) GetWebsite(id int64) (*Website, error) {
	ws, err := c.GetWebsites([]int64{id})
	if err != nil {
		return nil, err
	}
	if ws != nil && len(*ws) > 0 {
		return &((*ws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website not found", id)
}

// GetWebsites gets website existing records.
func (c *Client) GetWebsites(ids []int64) (*Websites, error) {
	ws := &Websites{}
	if err := c.Read(WebsiteModel, ids, nil, ws); err != nil {
		return nil, err
	}
	return ws, nil
}

// FindWebsite finds website record by querying it with criteria.
func (c *Client) FindWebsite(criteria *Criteria) (*Website, error) {
	ws := &Websites{}
	if err := c.SearchRead(WebsiteModel, criteria, NewOptions().Limit(1), ws); err != nil {
		return nil, err
	}
	if ws != nil && len(*ws) > 0 {
		return &((*ws)[0]), nil
	}
	return nil, fmt.Errorf("website was not found")
}

// FindWebsites finds website records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsites(criteria *Criteria, options *Options) (*Websites, error) {
	ws := &Websites{}
	if err := c.SearchRead(WebsiteModel, criteria, options, ws); err != nil {
		return nil, err
	}
	return ws, nil
}

// FindWebsiteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteId finds record id by querying it with criteria.
func (c *Client) FindWebsiteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website was not found")
}
