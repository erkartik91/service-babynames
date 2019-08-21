// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/erkartik91/service-babynames/restapi/operations"
)

//go:generate swagger generate server --target ../../service-babynames --name BabyNames --spec ../swagger/definition.yaml --exclude-main

func configureFlags(api *operations.BabyNamesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BabyNamesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetListIDHandler = operations.GetListIDHandlerFunc(func(params operations.GetListIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetListID has not yet been implemented")
	})
	api.PostListHandler = operations.PostListHandlerFunc(func(params operations.PostListParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostList has not yet been implemented")
	})
	api.PutListIDAddBabyNameHandler = operations.PutListIDAddBabyNameHandlerFunc(func(params operations.PutListIDAddBabyNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutListIDAddBabyName has not yet been implemented")
	})
	api.PutListIDRemoveBabyNameHandler = operations.PutListIDRemoveBabyNameHandlerFunc(func(params operations.PutListIDRemoveBabyNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutListIDRemoveBabyName has not yet been implemented")
	})

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
	return handler
}
