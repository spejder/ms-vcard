package main

import (
	"net/http"
	"os"

	"github.com/emersion/go-vcard"
	"github.com/spejder/ms-vcard/odoo"
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

	profiles, err := profiles(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	enc := vcard.NewEncoder(w)
	w.Header().Set("Content-Type", "text/vcard")

	for _, profile := range *profiles {
		card := toCard(profile)

		err = enc.Encode(card)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}
