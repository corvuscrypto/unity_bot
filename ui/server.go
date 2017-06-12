package ui

import (
	"net/http"
	"strconv"

	"github.com/corvuscrypto/unity_bot/config"
)

// All things related to the user interface's web server go here. This excludes routes, which should be
// placed into routes.go.

//CreateServer uses the global configuration to create a server object that will be used for the Admin UI
func CreateServer() (server *http.Server) {
	server = new(http.Server)
	server.Addr = ":" + strconv.Itoa(config.GlobalConfig.UIPort)
	return
}
