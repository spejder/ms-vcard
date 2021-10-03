package main

import (
	"fmt"

	"github.com/spejder/ms-vcard/odoo"
)

func profiles(c *odoo.Client) (*odoo.MemberProfiles, error) {
	criteria := odoo.NewCriteria().Add("can_access_contact_info", "=", true)

	criteria.
		Add("state", "!=", "inactive").
		Add("state", "!=", "cancelled").
		Add("state", "!=", "draft").
		Add("state", "!=", "waiting")

	options := odoo.NewOptions().FetchFields(
		"birthdate",
		"city",
		"country_id",
		"email",
		"firstname",
		"id",
		"lastname",
		"member_number",
		"mobile_clean",
		"organization_id",
		"phone",
		"relation_ids",
		"scout_name",
		"street",
		"zip",
	)

	profiles, err := c.FindMemberProfiles(criteria, options)
	if err != nil {
		return &odoo.MemberProfiles{}, fmt.Errorf("finding member profiles: %w", err)
	}

	return profiles, nil
}
