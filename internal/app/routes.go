package app

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()



	return mux
}