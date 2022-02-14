package app

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
	"github.com/kholidasad/products-API-go/app/handler"
	"github.com/kholidasad/products-API-go/app/model"
	"github.com/kholidasad/products-API-go/config"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}
	fmt.Println("Connected to DB")

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/product", a.handleRequest(handler.GetProducts))
	a.Post("/product", a.handleRequest(handler.PostProduct))
	a.Get("/product/{id}", a.handleRequest(handler.GetProduct))
	// a.Post("/product/{id}", a.handleRequest(handler.UpdateProducts))
	// a.Delete("/product/{id}", a.handleRequest(handler.DeleteProducts))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
// func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("DELETE")
// }

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}