package main

import (
	"database/sql"
	"fmt"

	"visitor/internal/http"
	"visitor/internal/http/handler"
	"visitor/internal/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addr         = "<host:port>"
	templatesDir = "web/"
	privateToken = "<token>"

	dbUser     = "<username>"
	dbPassword = "<password>"
	dbHost     = "<host>"
	dbName     = "<dbname>"
)

func main() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName))

	if err != nil {
		fmt.Println(err)
		return
	}

	dbTheme := mysql.NewTheme(db)
	dbVisitor := mysql.NewVisitor(db)

	srv := http.NewServer(addr, &http.Handlers{
		Auth:       handler.NewAuth(privateToken),
		APITheme:   handler.NewAPITheme(dbTheme),
		APIVisitor: handler.NewAPIVisitor(dbVisitor),
		Page:       handler.NewPage(dbTheme, dbVisitor, templatesDir),
	})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
