package main

import (
	"fmt"
	"slices"

	"bitbucket.org/long174/go-odoo"
	"github.com/spejder/ms-vcard/internal/ms"
)

// relationsByID reads the relations of all profiles and indexes them
// by relation ID.
//
// All relations are read in a single call to Medlemsservice. If that
// fails (e.g. because one of the relations is not accessible) we fall
// back to reading the relations one profile at a time, ignoring
// errors, so the other profiles still get their relations.
func relationsByID(c *ms.Client, profiles *ms.MemberProfiles) map[int64]ms.ResPartnerRelationAll {
	ids := []int64{}

	for _, profile := range *profiles {
		ids = append(ids, profile.RelationAllIds.Get()...)
	}

	slices.Sort(ids)
	ids = slices.Compact(ids)

	byID := make(map[int64]ms.ResPartnerRelationAll, len(ids))

	if len(ids) == 0 {
		return byID
	}

	relations, err := readRelations(c, ids)
	if err == nil {
		addRelations(byID, relations)

		return byID
	}

	for _, profile := range *profiles {
		relIDs := profile.RelationAllIds.Get()

		if len(relIDs) == 0 {
			continue
		}

		relations, err := readRelations(c, relIDs)
		if err != nil {
			continue
		}

		addRelations(byID, relations)
	}

	return byID
}

func readRelations(c *ms.Client, ids []int64) (*ms.ResPartnerRelationAlls, error) {
	relations := &ms.ResPartnerRelationAlls{}
	options := odoo.NewOptions().FetchFields("display_name")

	err := c.Read(ms.ResPartnerRelationAllModel, ids, options, relations)
	if err != nil {
		return nil, fmt.Errorf("reading partner relations: %w", err)
	}

	return relations, nil
}

func addRelations(byID map[int64]ms.ResPartnerRelationAll, relations *ms.ResPartnerRelationAlls) {
	for _, relation := range *relations {
		byID[relation.Id.Get()] = relation
	}
}

// profileRelations returns the relations matching ids in the order of
// ids. IDs not in byID are skipped.
func profileRelations(ids []int64, byID map[int64]ms.ResPartnerRelationAll) *ms.ResPartnerRelationAlls {
	relations := ms.ResPartnerRelationAlls{}

	for _, id := range ids {
		if relation, ok := byID[id]; ok {
			relations = append(relations, relation)
		}
	}

	return &relations
}
