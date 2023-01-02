package server

import "github.com/krable55/go-server/pkg/server/routes"

func Create() {
	var api = router.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/customers/getAll", routes.GetAll)
}
