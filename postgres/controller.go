package postgres

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"time"
	"fmt"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	host   = "localhost"
	port   = "5432"
	user   = "b21"
	dbname = string
)

func (a *App) Initialize(password string) error {
	var psql = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	a.DB, err := sql.Open("postgres", psql)
	if err != nil {
		return err
	}

	a.DB.SetMaxIdleConns(5)
	a.DB.SetMaxOpenConns(5)
	a.DB.SetConnMaxIdleTime(5 * time.Second)
	a.DB.SetConnMaxLifeTime(20 * time.Second)
}

func http(w http.ResponseWriter, *r http.Request) {

}
