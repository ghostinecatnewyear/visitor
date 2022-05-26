package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"visitor/internal/entity"
	"visitor/internal/mysql"
)

type APIVisitor struct {
	db *mysql.Visitor
}

func NewAPIVisitor(db *mysql.Visitor) *APIVisitor {
	return &APIVisitor{
		db: db,
	}
}

func (v *APIVisitor) GetAll(w http.ResponseWriter, r *http.Request) {
	visitors, err := v.db.GetAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	respb, err := json.Marshal(visitors)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(respb))
}

type requestSave struct {
	Nickname string `json:"nickname"`
}

func (v *APIVisitor) Save(w http.ResponseWriter, r *http.Request) {
	reqb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reqVisitor := &requestSave{}

	err = json.Unmarshal(reqb, reqVisitor)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	visitor := &entity.Visitor{
		Nickname:  reqVisitor.Nickname,
		VisitTime: time.Now(),
	}

	if !visitor.IsValid() {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err = v.db.Save(visitor)

	if err != nil {
		code := http.StatusInternalServerError

		if errors.Is(err, mysql.ErrDuplicate) {
			code = http.StatusConflict
		}

		http.Error(w, http.StatusText(code), code)
		return
	}
}
