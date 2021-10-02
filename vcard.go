package main

import (
	"strconv"

	"github.com/emersion/go-vcard"
	"github.com/spejder/ms-vcard/odoo"
)

func toCard(profile odoo.MemberProfile) vcard.Card {
	card := vcard.Card{}
	card.SetValue(vcard.FieldUID, uuid(profile.Id.Get()))

	//nolint:exhaustivestruct
	card.AddName(&vcard.Name{
		FamilyName: profile.Lastname.Get(),
		GivenName:  profile.Firstname.Get(),
	})

	if profile.ScoutName != nil {
		card.SetValue("NICKNAME", profile.ScoutName.Get())
	}

	if profile.OrganizationId != nil {
		card.SetValue("ORG", profile.OrganizationId.Name)
	}

	if profile.Email != nil {
		card.SetValue("EMAIL", profile.Email.Get())
	}

	if profile.Phone != nil {
		card.SetValue("TEL;TYPE=home", profile.Phone.Get())
	}

	if profile.MobileClean != nil {
		card.SetValue("TEL;TYPE=cell", profile.MobileClean.Get())
	}

	for _, id := range profile.RelationIds.Get() {
		card.AddValue("RELATED", uuid(id))
	}

	vcard.ToV4(card)

	return card
}

func uuid(id int64) string {
	//nolint:gomnd
	return "urn:uuid:ms-" + strconv.FormatInt(id, 10)
}
