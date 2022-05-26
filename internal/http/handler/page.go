package handler

import (
	"html/template"
	"net/http"

	"visitor/internal/entity"
	"visitor/internal/mysql"
)

type Page struct {
	templatesDir string

	dbTheme   *mysql.Theme
	dbVisitor *mysql.Visitor
}

func NewPage(dbTheme *mysql.Theme, dbVisitor *mysql.Visitor, templatesDir string) *Page {
	return &Page{
		templatesDir: templatesDir,
		dbTheme:      dbTheme,
		dbVisitor:    dbVisitor,
	}
}

func (p *Page) Load(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(p.templatesDir + "index.html")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	theme, err := p.dbTheme.Get()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	visitors, err := p.dbVisitor.GetAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, &struct {
		Theme         *entity.Theme
		Visitors      []*entity.Visitor
		VisitorsCount int
	}{
		Theme:         theme,
		Visitors:      visitors,
		VisitorsCount: len(visitors),
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
