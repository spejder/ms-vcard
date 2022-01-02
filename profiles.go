package main

import (
	"fmt"

	"bitbucket.org/long174/go-odoo"
	"github.com/spejder/ms-vcard/internal/ms"
)

func profiles(c *ms.Client) (*ms.MemberProfiles, error) {
	criteria := odoo.NewCriteria().Add("can_access_contact_info", "=", true)

	criteria.
		Add("state", "!=", "inactive").
		Add("state", "!=", "cancelled").
		Add("state", "!=", "draft")

	options := odoo.NewOptions().FetchFields(
		"birthdate",
		"city",
		"display_name",
		"email",
		"firstname",
		"id",
		"lastname",
		"member_number",
		"mobile_clean",
		"organization_id",
		"phone",
		"partner_id",
		"relation_all_ids",
		"relation_ids",
		"scout_name",
		"street",
		"zip",
	)

	profiles, err := c.FindMemberProfiles(criteria, options)
	if err != nil {
		return &ms.MemberProfiles{}, fmt.Errorf("finding member profiles: %w", err)
	}

	return profiles, nil
}
