package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	host   = "localhost"
	port   = "5432"
	user   = "b21"
	dbname = ""
)

func (a *App) Init(password string) error {
	var psql = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", psql)
	if err != nil {
		return err
	}

	// a.Run()
	a.DB.SetMaxIdleConns(5)
	a.DB.SetMaxOpenConns(5)
	a.DB.SetConnMaxIdleTime(5 * time.Second)
	a.DB.SetConnMaxLifetime(20 * time.Second)

	a.Router = mux.NewRouter()
	a.InitRoutes()
	return nil

}

func (a *App) Run(address string) error {
	err := http.ListenAndServe(":"+address, a.Router)
	if err != nil {
		return err
	}
	return nil
}

func HttpErrorResponse(w http.ResponseWriter, Rcode int, message string) {
	JsonResponse(w, Rcode, map[string]string{"error": message})
}

func JsonResponse(w http.ResponseWriter, Rcode int, info interface{}) {
	response, _ := json.Marshal(info)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Rcode)
	w.Write(response)
}

// func (a *App) GetRateByZip

// func (a *App) InitRoutes() {
// 	a.Router.HandleFunc(path, f)
// }
