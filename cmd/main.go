package main

import (
	"context"

	"github.com/serhiy-v/test-api/pkg"
	db2 "github.com/serhiy-v/test-api/pkg/db"
	"github.com/serhiy-v/test-api/pkg/handlers"
)

func main()  {
	ctx := context.Background()

	postgresClient := pkg.NewConnection(ctx)
	defer postgresClient.Close()

	db := db2.New(postgresClient)

	h := handlers.NewBaseHandler(db)

	serv := pkg.NewServer(*db, *h)

	serv.RunServer()


}
