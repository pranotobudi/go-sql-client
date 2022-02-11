package sqlclient

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

const (
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var (
	IsMocked bool
)

type Client struct {
	db *sql.DB
}

func (c *Client) Query(query string, args ...interface{}) (*SqlRows, error) {
	returnedRows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	sqlRows := &SqlRows{
		rows: returnedRows,
	}
	return sqlRows, nil
}

type SqlClient interface {
	Query(query string, args ...interface{}) (*SqlRows, error)
}

func StartMockupServer() {
	IsMocked = true
}
func StopMockupServer() {
	IsMocked = false
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
func Open(driverName, dataSourceName string) (*Client, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	client := &Client{
		db: db,
	}

	return client, nil
}
