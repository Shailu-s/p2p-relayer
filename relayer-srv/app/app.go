package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	relayer_srv "github.com/volmexfinance/relayers/relayer-srv"
)

const shutdownTimeout = 3 * time.Second

// App ...
type App struct {
	ctx     context.Context
	logger  *logrus.Logger
	router  *mux.Router
	server  *http.Server
	relayer *relayer_srv.RelayerSrv
}

// NewApp is initializes the app
func NewApp(logger *logrus.Logger, addr string, relayerSrv *relayer_srv.RelayerSrv) *App {
	// create new app
	inst := &App{
		logger:  logger,
		router:  mux.NewRouter(),
		server:  &http.Server{Addr: addr},
		relayer: relayerSrv,
	}
	// set router
	inst.router = mux.NewRouter()
	inst.setRouters()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})

	// Bypass all origins: https://pkg.go.dev/github.com/gorilla/handlers#CORSOption
	// Passing in a []string{"*"} will allow any domain.
	origins := handlers.AllowedOrigins([]string{"*"})

	inst.server.Handler = handlers.CORS(headers, methods, origins)(inst.router)

	return inst
}

// get wraps the router for GET method
func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}

// post wraps the router for Post method
func (a *App) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("POST")
}

func (a *App) setRouters() {
	a.get("/", a.endpointsGetHandler)
	a.post("/insertOrder/{chain}", a.ordersPostHandler)
	a.get("/getOrders/{trader}/{chain}", a.ordersGetHandler)
	a.get("/getOrderQueue/{chain}", a.ordersQueueGetHandler)
	a.get("/depth/{token}/{chain}", a.depthGetHandler)
	// a.Get("/signature/{amount}/{recipientAddress}/{destinationChainID}", a.SignatureHandler)
	// a.Get("/resend_tx/{id}", a.ResendTxHandler)
	// a.Get("/set_mode/{mode}", a.SetModeHandler)
}

// Run the app on it's router
func (a *App) Run(ctx context.Context) {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal(err)
		}
	}()

	a.logger.Infof("Perp Relayer Service has started on %s\nPress ctrl + C to exit.", a.server.Addr)
	a.relayer.Run()

	<-ctx.Done()

	a.logger.Infoln("Perp Relayer Service has stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	a.relayer.Stop()
	if err := a.server.Shutdown(ctxShutDown); err != nil {
		a.logger.Fatalf("Shutdown: %v\n", err)
	}
}
