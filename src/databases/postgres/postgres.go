package dbpostgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getMaxConnections(db *sql.DB) (int, error) {
	var maxConnections int
	row := db.QueryRowContext(context.Background(), "SHOW max_connections")
	err := row.Scan(&maxConnections)
	if err != nil {
		return 0, err
	}
	return maxConnections, nil
}

func NewDB() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbParams := os.Getenv("DB_PARAMS")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbname, dbParams)
	fmt.Println(connection)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(20)

	err = db.PingContext(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	maxConnections, err := getMaxConnections(db)
	if err != nil {
		db.Close()
		return nil, err
	}

	maxOpenConns := int(float64(maxConnections) * 0.9)
	fmt.Println(maxConnections)
	fmt.Println(maxOpenConns)
	db.SetMaxOpenConns(maxOpenConns)

	return db, nil

}
