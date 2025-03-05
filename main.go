package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/long174/go-odoo"
	"github.com/emersion/go-vcard"
	"github.com/spejder/ms-vcard/internal/ms"
)

func main() {
	addr := getAddr()

	http.HandleFunc("/", handler)

	//nolint:gosec
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func getAddr() string {
	addr, present := os.LookupEnv("ADDR")

	if !present {
		addr = ":80"
	}

	return addr
}

//nolint:funlen,cyclop
func handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	username, password, ok := r.BasicAuth()

	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	oc, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    username,
		Password: password,
		Database: os.Getenv("MS_DATABASE"),
		URL:      os.Getenv("MS_URL"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c := &ms.Client{Client: *oc}

	defer c.Close()

	companyID := r.URL.Query().Get("company_id")

	err = switchCompany(c, companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer func() {
		postCompanyID := r.URL.Query().Get("post_company_id")

		err = switchCompany(c, postCompanyID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	profiles, err := profiles(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	enc := vcard.NewEncoder(w)
	w.Header().Set("Content-Type", "text/vcard; charset=utf-8")

	for _, profile := range *profiles {
		relIDs := profile.RelationAllIds.Get()
		relations := &ms.ResPartnerRelationAlls{}

		if r.URL.Query().Has("relations") && len(relIDs) > 0 {
			relations, err = c.GetResPartnerRelationAlls(relIDs)
			if err != nil {
				relations = &ms.ResPartnerRelationAlls{}
			}
		}

		card := toCard(profile, relations)

		err = enc.Encode(card)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}

func switchCompany(c *ms.Client, companyIDParam string) error {
	if companyIDParam == "" {
		return nil
	}

	//nolint:gomnd
	companyID, err := strconv.ParseInt(companyIDParam, 10, 64)
	if err != nil {
		return fmt.Errorf("switching company parameters: %w", err)
	}

	err = c.Update(
		odoo.ResUsersModel,
		[]int64{c.UID()},
		//nolint:exhaustivestruct
		&odoo.ResUsers{
			CompanyId: &odoo.Many2One{ID: companyID},
		})
	if err != nil {
		return fmt.Errorf("switching company: %w", err)
	}

	return nil
}
