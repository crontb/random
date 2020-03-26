package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/crontb/random/pkg/http/handler"
	"github.com/crontb/random/pkg/middleware/auth"
	httpm "github.com/crontb/random/pkg/middleware/http"
	"github.com/crontb/random/pkg/middleware/logger"
)

func Run() {
	// @todo: Implement config autoloader
	// @todo: Implement SIGTERM/SIGKILL/SIGINT/os.Interrupt signals listener (graceful shutdown)
	runHttpService()
}

func runHttpService() {
	router := mux.NewRouter()
	// @todo: move router handlefunc definitions to pkg/endpoint
	router.HandleFunc("/card", handler.Card).Methods(http.MethodGet, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(httpm.MethodMiddleware)
	router.Use(logger.Middleware)
	router.Use(auth.Middleware)
	// These self-signed certificates are just for demo purpose.
	// @todo: implement secure certificates automount (e.g. from Vault by init container in Kubernetes Pod)
	// @todo: implement http to https redirect
	log.Fatal(http.ListenAndServeTLS(":8000", "./internal/crt/cert.pem", "./internal/crt/key.pem", router))
}
