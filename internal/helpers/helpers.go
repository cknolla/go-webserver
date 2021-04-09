package helpers

import (
	"fmt"
	"github.com/cknolla/go-webserver/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

func NewHelpers(config *config.AppConfig) {
	app = config
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}