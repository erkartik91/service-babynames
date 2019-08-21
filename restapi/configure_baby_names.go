// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	"github.com/erkartik91/service-babynames/controllers"
	"github.com/erkartik91/service-babynames/orm"
	"github.com/erkartik91/service-babynames/restapi/operations"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

//go:generate swagger generate server --target ../../service-babynames --name BabyNames --spec ../swagger/definition.yaml

func configureFlags(api *operations.BabyNamesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BabyNamesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	var db *gorm.DB
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// /	db, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}
	if db == nil {
		panic("unsupported database type")
	}
	o := orm.New(db)

	listController := controllers.List{
		ORM: o,
	}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetListIDHandler = operations.GetListIDHandlerFunc(listController.GetListIDHandlerFunc)
	api.PostListHandler = operations.PostListHandlerFunc(listController.PostListHandlerFunc)
	api.PutListIDAddBabyNameHandler = operations.PutListIDAddBabyNameHandlerFunc(listController.PutListIDAddBabyNameHandlerFunc)
	api.PutListIDRemoveBabyNameHandler = operations.PutListIDRemoveBabyNameHandlerFunc(listController.PutListIDRemoveBabyNameHandlerFunc)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         1000,
	})
	return corsHandler.Handler(handler)
}
