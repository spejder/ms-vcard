package odoo

import (
	"fmt"
)

// UserProfile represents user.profile model.
type UserProfile struct {
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	NricBack    *String   `xmlrpc:"nric_back,omptempty"`
	NricFront   *String   `xmlrpc:"nric_front,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// UserProfiles represents array of user.profile model.
type UserProfiles []UserProfile

// UserProfileModel is the odoo model name.
const UserProfileModel = "user.profile"

// Many2One convert UserProfile to *Many2One.
func (up *UserProfile) Many2One() *Many2One {
	return NewMany2One(up.Id.Get(), "")
}

// CreateUserProfile creates a new user.profile model and returns its id.
func (c *Client) CreateUserProfile(up *UserProfile) (int64, error) {
	return c.Create(UserProfileModel, up)
}

// UpdateUserProfile updates an existing user.profile record.
func (c *Client) UpdateUserProfile(up *UserProfile) error {
	return c.UpdateUserProfiles([]int64{up.Id.Get()}, up)
}

// UpdateUserProfiles updates existing user.profile records.
// All records (represented by ids) will be updated by up values.
func (c *Client) UpdateUserProfiles(ids []int64, up *UserProfile) error {
	return c.Update(UserProfileModel, ids, up)
}

// DeleteUserProfile deletes an existing user.profile record.
func (c *Client) DeleteUserProfile(id int64) error {
	return c.DeleteUserProfiles([]int64{id})
}

// DeleteUserProfiles deletes existing user.profile records.
func (c *Client) DeleteUserProfiles(ids []int64) error {
	return c.Delete(UserProfileModel, ids)
}

// GetUserProfile gets user.profile existing record.
func (c *Client) GetUserProfile(id int64) (*UserProfile, error) {
	ups, err := c.GetUserProfiles([]int64{id})
	if err != nil {
		return nil, err
	}
	if ups != nil && len(*ups) > 0 {
		return &((*ups)[0]), nil
	}
	return nil, fmt.Errorf("id %v of user.profile not found", id)
}

// GetUserProfiles gets user.profile existing records.
func (c *Client) GetUserProfiles(ids []int64) (*UserProfiles, error) {
	ups := &UserProfiles{}
	if err := c.Read(UserProfileModel, ids, nil, ups); err != nil {
		return nil, err
	}
	return ups, nil
}

// FindUserProfile finds user.profile record by querying it with criteria.
func (c *Client) FindUserProfile(criteria *Criteria) (*UserProfile, error) {
	ups := &UserProfiles{}
	if err := c.SearchRead(UserProfileModel, criteria, NewOptions().Limit(1), ups); err != nil {
		return nil, err
	}
	if ups != nil && len(*ups) > 0 {
		return &((*ups)[0]), nil
	}
	return nil, fmt.Errorf("user.profile was not found")
}

// FindUserProfiles finds user.profile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUserProfiles(criteria *Criteria, options *Options) (*UserProfiles, error) {
	ups := &UserProfiles{}
	if err := c.SearchRead(UserProfileModel, criteria, options, ups); err != nil {
		return nil, err
	}
	return ups, nil
}

// FindUserProfileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUserProfileIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UserProfileModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUserProfileId finds record id by querying it with criteria.
func (c *Client) FindUserProfileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UserProfileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("user.profile was not found")
}
