package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"github.com/anmolbabu/product_service/pkg/db"
)

var (
	router    *mux.Router
	context   = ApplicationContext{}
	once      sync.Once
	dbCleanUp func() error
)

func Cleanup() {
	defer dbCleanUp()
}

func RegisterAPI(route string, method string, handlerName string) {
	once.Do(func() {
		router = mux.NewRouter()

		client, cleanupFunc, err := db.GetPSQLClientPool(
			os.Getenv("PSQL_HOST"),
			os.Getenv("PSQL_PORT"),
			os.Getenv("PSQL_USER"),
			os.Getenv("PSQL_PASS"),
			"product_service",
		)

		fmt.Printf("Client is : %v\n", client)

		if err != nil {
			fmt.Errorf("error creating psql client")
			os.Exit(1)
		}

		context.client = client

		dbCleanUp = cleanupFunc
	})

	handlerIface := reflect.ValueOf(context).MethodByName(handlerName).Interface()

	handler := handlerIface.(func(w http.ResponseWriter, r *http.Request))

	router.HandleFunc(route, handler).Methods(method)
}

func StartRestServer() {
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
