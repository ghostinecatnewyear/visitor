package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"visitor/internal/entity"
	"visitor/internal/mysql"
)

type APITheme struct {
	db *mysql.Theme
}

func NewAPITheme(db *mysql.Theme) *APITheme {
	return &APITheme{
		db: db,
	}
}

func (t *APITheme) Get(w http.ResponseWriter, r *http.Request) {
	theme, err := t.db.Get()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	respb, err := json.Marshal(theme)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(respb))
}

func (t *APITheme) Update(w http.ResponseWriter, r *http.Request) {
	reqb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utheme := &entity.UpdateTheme{}

	err = json.Unmarshal(reqb, utheme)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	theme, err := t.db.Get()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ctheme := utheme.Apply(*theme)

	if !ctheme.IsValid() {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err = t.db.Update(ctheme)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
