package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	json := `{"status": "available", "environment": %q, "version": %q}`
	json = fmt.Sprintf(json, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(json))
}
