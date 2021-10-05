package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/emersion/go-vcard"
	"github.com/spejder/ms-vcard/internal/odoo"
)

func main() {
	addr := getAddr()

	http.HandleFunc("/", handler)

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

//nolint:funlen
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

	c, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    username,
		Password: password,
		Database: os.Getenv("MS_DATABASE"),
		URL:      os.Getenv("MS_URL"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

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
		card := toCard(profile)

		err = enc.Encode(card)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}

func switchCompany(c *odoo.Client, companyIDParam string) error {
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
