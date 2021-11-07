package api

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"github.com/iamnande/ff8-api/internal/datastore"

	"github.com/iamnande/ff8-api/internal/api/server"
	"github.com/iamnande/ff8-api/internal/config"
)

// FF8API is the FF8 API application.
type FF8API struct {
	cfg *config.Config
	ds  datastore.Datastore
	log zerolog.Logger
	svr *server.Server
}

// NewFF8API initializes a fresh instance of the FF8 API.
func NewFF8API(cfg *config.Config, log zerolog.Logger, ds datastore.Datastore) *FF8API {

	//  api: initialize new application instance
	api := &FF8API{
		cfg: cfg,
		ds:  ds,
		log: log,
		svr: server.NewServer(cfg),
	}

	// api: initialize domain instances
	// TODO: move these into sub-packages and chain the APIs
	api.svr.POST("/calculate", api.HandleCalculate)

	// api: return initialized API
	return api

}

// Log will return the core logging instance handler.
func (api *FF8API) Log() *zerolog.Logger {
	logger := api.log.With().Logger()
	return &logger
}

// Address will return the current HTTP listener address.
func (api *FF8API) Address() string {
	return api.svr.Server.Addr
}

// Serve will start the HTTP server.
func (api *FF8API) Serve() error {
	return api.svr.StartServer(api.svr.Server)
}

// Shutdown will stop the HTTP server.
func (api *FF8API) Shutdown() error {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownCancel()
	return api.svr.Shutdown(shutdownCtx)
}
