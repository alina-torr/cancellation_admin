package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getArrayQuery[K any](dbpool *pgxpool.Pool, sql string, args ...any) ([]K, error) {
	rows, err := dbpool.Query(context.Background(), sql, args...)
	if err != nil {
		fmt.Println("query ", err.Error())
		return nil, err
	}
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[K])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return products, nil
}

// func getQuery[K any](dbpool *pgxpool.Pool, sql string, args ...any) (K, error) {
// 	rows, err := dbpool.Query(context.Background(), sql, args...)
// 	if err != nil {
// 		fmt.Println("query ", err.Error())
// 		return nil, err
// 	}
// 	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[K])
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	return products, nil
// }
