package app

import (
	"net/http"
)

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.reverseProxyHandler)

	return mux
}
