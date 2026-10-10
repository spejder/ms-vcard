package main

import (
	"fmt"

	"bitbucket.org/long174/go-odoo"
	"github.com/spejder/ms-vcard/internal/ms"
)

func profiles(c *ms.Client, withRelations bool) (*ms.MemberProfiles, error) {
	criteria := odoo.NewCriteria().Add("can_access_contact_info", "=", true)

	criteria.
		Add("state", "!=", "inactive").
		Add("state", "!=", "cancelled").
		Add("state", "!=", "draft")

	fields := []string{
		"birthdate",
		"city",
		"email",
		"firstname",
		"id",
		"lastname",
		"member_number",
		"mobile_clean",
		"organization_id",
		"phone",
		"partner_id",
		"scout_name",
		"street",
		"zip",
	}

	// Relations are expensive for Medlemsservice to compute, so only
	// fetch them when they are requested.
	if withRelations {
		fields = append(fields, "display_name", "relation_all_ids")
	}

	options := odoo.NewOptions().FetchFields(fields...)

	profiles, err := c.FindMemberProfiles(criteria, options)
	if err != nil {
		return &ms.MemberProfiles{}, fmt.Errorf("finding member profiles: %w", err)
	}

	return profiles, nil
}
