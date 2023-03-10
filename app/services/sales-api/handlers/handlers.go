package handlers

import (
	"expvar"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/ZweWT/backend-go/app/services/sales-api/handlers/debug/checkgrp"
	v1ProductGrp "github.com/ZweWT/backend-go/app/services/sales-api/handlers/v1/productgrp"
	v1TestGrp "github.com/ZweWT/backend-go/app/services/sales-api/handlers/v1/testgrp"
	v1UserGrp "github.com/ZweWT/backend-go/app/services/sales-api/handlers/v1/usergrp"
	productCore "github.com/ZweWT/backend-go/bussiness/core/product"
	userCore "github.com/ZweWT/backend-go/bussiness/core/user"
	"github.com/ZweWT/backend-go/bussiness/sys/auth"
	"github.com/ZweWT/backend-go/bussiness/web/mid"
	"github.com/ZweWT/backend-go/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}

// DebugMux registers all the debug standard library routes and then custom
// debug application routes for the service.
func DebugMux(build string, log *zap.SugaredLogger, db *sqlx.DB) http.Handler {
	mux := DebugStandardLibraryMux()

	// Register debug check endpoints.
	cgh := checkgrp.Handlers{
		Build: build,
		Log:   log,
		DB:    db,
	}

	mux.HandleFunc("/debug/readiness", cgh.Readiness)

	return mux
}

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Auth     *auth.Auth
	DB       *sqlx.DB
}

// APIMux constructs an http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {

	// Construct the web.App which holds all routes.
	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	)

	// Load the routes for the different versions of the API.
	v1(app, cfg)

	return app
}

// v1 binds all the version 1 routes.
func v1(app *web.App, cfg APIMuxConfig) {
	const version = "v1"

	tgh := v1TestGrp.Handlers{
		Log: cfg.Log,
	}
	app.Handle(http.MethodGet, version, "/test", tgh.Test)
	app.Handle(http.MethodGet, version, "/testauth", tgh.Test, mid.Authenticate(cfg.Auth), mid.Authorize("ADMIN"))

	// Register user and authenticaton endpoints
	ugh := v1UserGrp.Handlers{
		User: userCore.NewCore(cfg.Log, cfg.DB),
		Auth: cfg.Auth,
	}
	app.Handle(http.MethodGet, version, "/users/token", ugh.Token)
	app.Handle(http.MethodGet, version, "/users/:page/:rows", ugh.Query, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodGet, version, "/users/:id", ugh.QueryByID, mid.Authenticate(cfg.Auth))
	// app.Handle(http.MethodPost, version, "/users", ugh.Create, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))
	// app.Handle(http.MethodPut, version, "/users/:id", ugh.Update, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))
	// app.Handle(http.MethodDelete, version, "/users/:id", ugh.Delete, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))

	// Register product endpoints
	pgh := v1ProductGrp.Handlers{
		Product: productCore.NewCore(cfg.Log, cfg.DB),
		Auth:    cfg.Auth,
	}

	app.Handle(http.MethodGet, version, "/products/:page/:rows", pgh.Query, mid.Authenticate(cfg.Auth))
	app.Handle(http.MethodPost, version, "/products", pgh.Create, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodPut, version, "/products/:id", pgh.Update, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodDelete, version, "/products/:id", pgh.Delete, mid.Authenticate(cfg.Auth), mid.Authorize(auth.RoleAdmin))
}
