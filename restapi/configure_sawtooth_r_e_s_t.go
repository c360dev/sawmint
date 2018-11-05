// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/c360dev/sawmint/restapi/operations"
)

//go:generate swagger generate server --target .. --name SawtoothREST --spec ../openapi.yaml

func configureFlags(api *operations.SawtoothRESTAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SawtoothRESTAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinConsumer = runtime.ByteStreamConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetBatchStatusesHandler = operations.GetBatchStatusesHandlerFunc(func(params operations.GetBatchStatusesParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBatchStatuses has not yet been implemented")
	})
	api.GetBatchesHandler = operations.GetBatchesHandlerFunc(func(params operations.GetBatchesParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBatches has not yet been implemented")
	})
	api.GetBatchesBatchIDHandler = operations.GetBatchesBatchIDHandlerFunc(func(params operations.GetBatchesBatchIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBatchesBatchID has not yet been implemented")
	})
	api.GetBlocksHandler = operations.GetBlocksHandlerFunc(func(params operations.GetBlocksParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBlocks has not yet been implemented")
	})
	api.GetBlocksBlockIDHandler = operations.GetBlocksBlockIDHandlerFunc(func(params operations.GetBlocksBlockIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBlocksBlockID has not yet been implemented")
	})
	api.GetPeersHandler = operations.GetPeersHandlerFunc(func(params operations.GetPeersParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetPeers has not yet been implemented")
	})
	api.GetReceiptsHandler = operations.GetReceiptsHandlerFunc(func(params operations.GetReceiptsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetReceipts has not yet been implemented")
	})
	api.GetStateHandler = operations.GetStateHandlerFunc(func(params operations.GetStateParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetState has not yet been implemented")
	})
	api.GetStateAddressHandler = operations.GetStateAddressHandlerFunc(func(params operations.GetStateAddressParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetStateAddress has not yet been implemented")
	})
	api.GetStatusHandler = operations.GetStatusHandlerFunc(func(params operations.GetStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetStatus has not yet been implemented")
	})
	api.GetTransactionsHandler = operations.GetTransactionsHandlerFunc(func(params operations.GetTransactionsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetTransactions has not yet been implemented")
	})
	api.GetTransactionsTransactionIDHandler = operations.GetTransactionsTransactionIDHandlerFunc(func(params operations.GetTransactionsTransactionIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetTransactionsTransactionID has not yet been implemented")
	})
	api.PostBatchStatusesHandler = operations.PostBatchStatusesHandlerFunc(func(params operations.PostBatchStatusesParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostBatchStatuses has not yet been implemented")
	})
	api.PostBatchesHandler = operations.PostBatchesHandlerFunc(func(params operations.PostBatchesParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostBatches has not yet been implemented")
	})
	api.PostReceiptsHandler = operations.PostReceiptsHandlerFunc(func(params operations.PostReceiptsParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostReceipts has not yet been implemented")
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
