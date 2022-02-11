package sqlclient

import "database/sql"

type SqlRows struct {
	rows *sql.Rows
}

type Rows interface {
	HasNext() bool
	Close() error
	Scan(destinations ...interface{}) error
}

func (r *SqlRows) HasNext() bool {
	return r.rows.Next()
}

func (r *SqlRows) Close() error {
	return r.rows.Close()
}

func (r *SqlRows) Scan(destinations ...interface{}) error {
	return r.rows.Scan(destinations...)
}
