package main

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

var pool * pgxpool.Pool

func closePool () {
    pool.Close()
}

func generatePool () {
    var databaseUrl = "postgres://postgres:kien25062000@localhost:5432/book"
    config, err := pgxpool.ParseConfig(databaseUrl)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to parse url to database: %v\n", err)
        os.Exit(1)
    }

    config.AfterConnect = func (ctx context.Context, conn *pgx.Conn) error {
        fmt.Println("Connected to datbase")
        return nil
    }

    pool, err = pgxpool.ConnectConfig(context.Background(), config)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
}

func QueryHandler (query func(conn *pgxpool.Conn) (pgx.Rows, error)) pgx.Rows {
    conn, err := pool.Acquire (context.Background())

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to acquire connection: %v\n", err)
        os.Exit(1)
    }

    rows,err := query(conn)

    if err != nil {
        fmt.Fprintf(os.Stderr, "There is error during querying: %v\n", err)
    }

    return rows
}

func RowsHandler (rows pgx.Rows) func (rowProcessor func(rows pgx.Rows) error) {
    return func (rowProcessor func(rows pgx.Rows) error) {
        for rows.Next() {
            err := rowProcessor(rows)
            if err != nil {
                fmt.Fprintf(os.Stderr, "There is error during handling rows: %v\n", err)
            }
        }
    }
}

func QueryPipeline (
    query func(conn *pgxpool.Conn) (pgx.Rows, error),
    rowProcessor func(rows pgx.Rows) error,
) {
    rows :=  QueryHandler(query)
    rowsHandler := RowsHandler(rows)
    rowsHandler(rowProcessor)
}
