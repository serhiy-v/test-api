package pkg

import (
	"context"
	"net/http"

	"github.com/serhiy-v/test-api/pkg/db"
	"github.com/serhiy-v/test-api/pkg/handlers"
)

func RunServer() error{
	s := &http.Server{
		Addr: ":8080",
		Handler: handlers.NewRouter(),
	}

	ctx := context.Background()

	postgresClient := db.NewConnection(ctx)
	defer postgresClient.Close()

	return s.ListenAndServe()
}

