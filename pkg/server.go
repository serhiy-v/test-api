package pkg

import (
	"net/http"

	"github.com/serhiy-v/test-api/pkg/handlers"
)

func RunServer() error{
	s := &http.Server{
		Addr: ":8080",
		Handler: handlers.NewRouter(),
	}


	return s.ListenAndServe()
}

