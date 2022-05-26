package mysql

import (
	"database/sql"
	"errors"

	"visitor/internal/entity"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/go-sql-driver/mysql"
)

var (
	ErrDuplicate = errors.New("duplicated")
)

type Visitor struct {
	db *sql.DB
}

func NewVisitor(db *sql.DB) *Visitor {
	return &Visitor{
		db: db,
	}
}

func (v *Visitor) GetAll() ([]*entity.Visitor, error) {
	q, _, err := goqu.
		Dialect("mysql").
		Select("nickname", "visit_time").
		From("visitor").
		ToSQL()

	if err != nil {
		return nil, err
	}

	rows, err := v.db.Query(q)

	if err != nil {
		return nil, err
	}

	visitors := []*entity.Visitor{}

	for rows.Next() {
		visitor := &entity.Visitor{}

		err := rows.Scan(&visitor.Nickname, &visitor.VisitTime)

		if err != nil {
			return nil, err
		}

		visitors = append(visitors, visitor)
	}

	return visitors, nil
}

func (v *Visitor) Save(visitor *entity.Visitor) error {
	q, _, err := goqu.
		Dialect("mysql").
		Insert("visitor").
		Cols("nickname").
		Vals(goqu.Vals{visitor.Nickname}).
		ToSQL()

	if err != nil {
		return err
	}

	_, err = v.db.Exec(q)

	mysqlErr := &mysql.MySQLError{}

	const duplCode = 1062

	if errors.As(err, &mysqlErr) && mysqlErr.Number == duplCode {
		return ErrDuplicate
	}

	return err
}
