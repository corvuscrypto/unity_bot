package ui

import "net/http"

// All things related to the user interface's web server go here. This excludes routes, which should be
// placed into routes.go.

//CreateServer uses the global configuration to create a server object that will be used for the Admin UI
func CreateServer() (server *http.Server) {
	server = new(http.Server)
	return
}
