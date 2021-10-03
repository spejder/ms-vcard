package main

import (
	"math/rand"

	"github.com/emersion/go-vcard"
	"github.com/google/uuid"
	"github.com/spejder/ms-vcard/odoo"
)

func toCard(profile odoo.MemberProfile) vcard.Card {
	card := vcard.Card{}
	card.SetValue(vcard.FieldUID, getUUID(profile.Id.Get()))

	//nolint:exhaustivestruct
	card.AddName(&vcard.Name{
		FamilyName: profile.Lastname.Get(),
		GivenName:  profile.Firstname.Get(),
	})

	if profile.MemberNumber != nil {
		card.SetValue("NOTE", "Medlemsnummer: "+profile.MemberNumber.Get())
	}

	if profile.Street != nil {
		//nolint:exhaustivestruct
		card.AddAddress(&vcard.Address{
			StreetAddress: profile.Street.Get(),
			Locality:      profile.City.Get(),
			PostalCode:    profile.Zip.Get(),
		})
	}

	if profile.ScoutName != nil {
		card.SetValue("NICKNAME;TYPE=Spejdernavn", profile.ScoutName.Get())
	}

	if profile.Birthdate != nil {
		card.SetValue("BDAY", profile.Birthdate.Get().Format("20060102"))
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
		card.AddValue("RELATED;TYPE=kin", getUUID(id))
	}

	vcard.ToV4(card)

	return card
}

// getUUID creates a deterministic UUID from Odoos int64 profile id.
func getUUID(id int64) string {
	//nolint:gosec // we don't use random for security
	rnd := rand.New(rand.NewSource(id))
	uuid.SetRand(rnd)
	u, _ := uuid.NewRandomFromReader(rnd)

	return u.URN()
}
