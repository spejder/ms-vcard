package main

import (
	"slices"
	"testing"

	"bitbucket.org/long174/go-odoo"
	"github.com/spejder/ms-vcard/internal/ms"
)

func TestProfileRelations(t *testing.T) {
	t.Parallel()

	byID := map[int64]ms.ResPartnerRelationAll{
		1: relation(1, "one"),
		2: relation(2, "two"),
		3: relation(3, "three"),
	}

	tests := []struct {
		name string
		ids  []int64
		want []string
	}{
		{name: "no ids", ids: []int64{}, want: []string{}},
		{name: "keeps order of ids", ids: []int64{3, 1}, want: []string{"three", "one"}},
		{name: "skips missing ids", ids: []int64{4, 2, 5}, want: []string{"two"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := []string{}

			for _, relation := range *profileRelations(test.ids, byID) {
				got = append(got, relation.DisplayName.Get())
			}

			if !slices.Equal(got, test.want) {
				t.Errorf("profileRelations(%v) = %v, want %v", test.ids, got, test.want)
			}
		})
	}
}

func relation(id int64, displayName string) ms.ResPartnerRelationAll {
	var relation ms.ResPartnerRelationAll

	relation.Id = odoo.NewInt(id)
	relation.DisplayName = odoo.NewString(displayName)

	return relation
}
