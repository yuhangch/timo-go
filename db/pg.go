package db

import (
	"context"
	"fmt"

	pg "github.com/jackc/pgx/v4/pgxpool"
)

var Conn *pg.Pool

// Pool init pqx pool.
func Pool() (*pg.Pool, error) {
	cfg, err := pg.ParseConfig("postgresql://postgres:chen2code@134.175.83.238/timvt")
	if err != nil {
		fmt.Println(err)
	}
	DB, err := pg.ConnectConfig(context.Background(), cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(DB)
	return DB, nil
}

func init() {
	Conn, _ = Pool()
}
