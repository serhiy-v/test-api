package pkg

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serhiy-v/test-api/pkg/db"
	"github.com/serhiy-v/test-api/pkg/handlers"
)

type Server struct {
	router *mux.Router
	db db.Database
}

func NewServer(db db.Database) *Server {
	return &Server{
		router: handlers.NewRouter(),
		db: db,
	}
}
func (s *Server) RunServer() error{
	return http.ListenAndServe(":8080", s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewConnection(ctx context.Context) *pgxpool.Pool{
	addr := "postgres://postgres:secret@localhost:5432/test-api?sslmode=disable"
	conn, err := pgxpool.Connect(ctx,addr)
	if err != nil{
		log.Fatalf("Error connecting to db")
	}
	if err = conn.Ping(ctx); err != nil {
		log.Fatalf("Error ping database")
	}
	return conn
}


