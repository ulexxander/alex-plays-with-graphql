package main

import (
	"context"
	"database/sql"
	"fmt"
	"gqllasttry/graphql"
	"gqllasttry/graphql/generated"
	"gqllasttry/postgres"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DBWithLogging struct {
	db  *sql.DB
	log *logrus.Logger
}

func (withLog *DBWithLogging) stripQuery(query string) string {
	newLineIdx := strings.Index(query, "\n")
	return query[newLineIdx+1 : len(query)-1]
}

func (withLog *DBWithLogging) ExecContext(ctx context.Context, query string, params ...interface{}) (sql.Result, error) {
	withLog.log.WithField("query", withLog.stripQuery(query)).Trace("DB:ExecContext")
	return withLog.db.ExecContext(ctx, query, params...)
}

func (withLog *DBWithLogging) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	withLog.log.WithField("query", withLog.stripQuery(query)).Trace("DB:PrepareContext")
	return withLog.db.PrepareContext(ctx, query)
}

func (withLog *DBWithLogging) QueryContext(ctx context.Context, query string, params ...interface{}) (*sql.Rows, error) {
	withLog.log.WithField("query", withLog.stripQuery(query)).Trace("DB:QueryContext")
	return withLog.db.QueryContext(ctx, query, params...)
}

func (withLog *DBWithLogging) QueryRowContext(ctx context.Context, query string, params ...interface{}) *sql.Row {
	withLog.log.WithField("query", withLog.stripQuery(query)).Trace("DB:QueryRowContext")
	return withLog.db.QueryRowContext(ctx, query, params...)
}

func corsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	}
}

func run() error {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)

	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/gqlfun?sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to the postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error when pinging db: %v", err)
	}

	dbWithLog := &DBWithLogging{
		db:  db,
		log: log,
	}
	pg := postgres.New(dbWithLog)

	resolver := &graphql.Resolver{
		DB: pg,
	}
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: resolver})

	mux := http.NewServeMux()

	mux.Handle("/graphql", corsMiddleware(handler.NewDefaultServer(schema)))
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	server := &http.Server{
		Addr:    ":4015",
		Handler: mux,
	}

	return server.ListenAndServe()
}

func main() {
	err := run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
