package temp

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/lib/pq"
)

func PQ() (*sql.DB, error) {
	host := os.Getenv("PQ_HOST")
	port := os.Getenv("PQ_PORT")
	user := os.Getenv("PQ_USER")
	pass := os.Getenv("PQ_PASS")
	name := os.Getenv("PQ_NAME")

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=require",
		user, pass, host, port, name)
	conn, err := pq.NewConnector(dsn)
	if err != nil {
		return nil, err
	}
	return sql.OpenDB(conn), nil
}
