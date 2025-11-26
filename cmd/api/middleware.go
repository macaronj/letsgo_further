package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer functions runs at end of panic event
		defer func() {
			// recover functions checks for panic
			if err := recover(); err != nil {
				// INFO: "Connection": "close" header closes connection automatically
				w.Header().Set("Connection", "close")
				// recover() returns any, Errorf normalizes to error, custom logger will log "ERROR" level
				// While we’re on the topic of errors, I’d like to mention that in certain scenarios Go’s
				// http.Server may still automatically generate and send plain-text HTTP responses. These
				// scenarios include when:
				// The HTTP request specifies an unsupported HTTP protocol version.
				// The HTTP request contains a missing or invalid Host header, or multiple Host headers.
				// The HTTP request contains a empty Content-Length header.
				// The HTTP request contains an unsupported Transfer-Encoding header.
				// The size of the HTTP request headers exceeds the server’s MaxHeaderBytes setting.
				// The client makes a HTTP request to a HTTPS serve
				app.serverErrorResponse(w, r, fmt.Errorf("s%", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
