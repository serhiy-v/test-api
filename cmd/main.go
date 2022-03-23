package main

import (
	"context"

	"github.com/serhiy-v/test-api/pkg"
	db2 "github.com/serhiy-v/test-api/pkg/db"
)

func main()  {
	ctx := context.Background()

	postgresClient := pkg.NewConnection(ctx)
	defer postgresClient.Close()

	db := db2.New(postgresClient)

	serv := pkg.NewServer(*db)

	serv.RunServer()


}
