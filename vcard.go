package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/emersion/go-vcard"
	"github.com/google/uuid"
	"github.com/spejder/ms-vcard/internal/ms"
)

type group int

const (
	ScoutName group = iota + 1
	MemberNumber
)

func (g group) Group() string {
	return fmt.Sprintf("item%d", g)
}

//nolint:funlen
func toCard(profile ms.MemberProfile, relations *ms.ResPartnerRelationAlls) vcard.Card {
	card := vcard.Card{}
	card.SetValue(vcard.FieldUID, getUUID(profile.PartnerId.ID))

	card.AddName(&vcard.Name{
		FamilyName: profile.Lastname.Get(),
		GivenName:  profile.Firstname.Get(),
	})

	if profile.MemberNumber != nil {
		card.Add("NOTE", &vcard.Field{
			Value: "Medlemsnummer: " + profile.MemberNumber.Get(),
			Group: MemberNumber.Group(),
		})
		card.Add("X-ABLabel", &vcard.Field{
			Value: "_$!<Medlemsnummer>!$_",
			Group: MemberNumber.Group(),
		})
	}

	if profile.Street != nil {
		card.AddAddress(&vcard.Address{
			StreetAddress: profile.Street.Get(),
			Locality:      profile.City.Get(),
			PostalCode:    profile.Zip.Get(),
		})
	}

	if profile.ScoutName != nil {
		card.Set("NICKNAME", &vcard.Field{
			Value: profile.ScoutName.Get(),
			Group: ScoutName.Group(),
		})
		card.Add("X-ABLabel", &vcard.Field{
			Value: "_$!<Spejdernavn>!$_",
			Group: ScoutName.Group(),
		})
	}

	if profile.Birthdate != nil {
		card.SetValue("BDAY", profile.Birthdate.Get().Format("2006-01-02"))
	}

	if profile.OrganizationId != nil {
		card.SetValue("ORG", profile.OrganizationId.Name)
	}

	if profile.Email != nil {
		card.SetValue("EMAIL", profile.Email.Get())
	}

	if profile.Phone != nil {
		home := vcard.Params{}
		home.Set("TYPE", "home")

		card.Set("TEL", &vcard.Field{
			Value:  profile.Phone.Get(),
			Params: home,
		})
	}

	if profile.MobileClean != nil {
		cell := vcard.Params{}
		cell.Set("TYPE", "cell")

		card.Set("TEL", &vcard.Field{
			Value:  profile.MobileClean.Get(),
			Params: cell,
		})
	}

	for _, relation := range *relations {
		related := strings.TrimSpace(strings.Replace(relation.DisplayName.Get(), profile.DisplayName.Get(), "", 1))

		card.Add("NOTE", &vcard.Field{
			Value: related,
		})
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
