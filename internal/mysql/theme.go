package mysql

import (
	"database/sql"

	"visitor/internal/entity"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

type Theme struct {
	db *sql.DB
}

func NewTheme(db *sql.DB) *Theme {
	return &Theme{
		db: db,
	}
}

func (t *Theme) Get() (*entity.Theme, error) {
	q, _, err := goqu.
		Dialect("mysql").
		Select("color").
		From("theme").
		ToSQL()

	if err != nil {
		return nil, err
	}

	theme := &entity.Theme{}

	err = t.db.QueryRow(q).Scan(&theme.Color)

	return theme, err
}

func (t *Theme) Update(theme *entity.Theme) error {
	const id = 0

	q, _, err := goqu.
		Dialect("mysql").
		Update("theme").
		Set(goqu.Record{
			"color": theme.Color,
		}).
		Where(goqu.Ex{"id": id}).
		ToSQL()

	if err != nil {
		return err
	}

	_, err = t.db.Exec(q)

	return err
}
