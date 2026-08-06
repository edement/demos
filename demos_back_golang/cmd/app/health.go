package main

import "net/http"

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))

	//app.database.Posts.Create(r.Context(), nil)
}
