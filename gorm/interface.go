package gorm

import "database/sql"


// Dialector GORM database dialector
type Dialector interface {
	Name() string
	Initialize(*DB) error
	Explain(sql string, vars ...interface{}) string
}

// SavePointerDialectorInterface save pointer interface
type SavePointerDialectorInterface interface {
	SavePoint(tx *DB, name string) error
	RollbackTo(tx *DB, name string) error
}
 
// SQLCommon is the minimal database connection functionality gorm requires.  Implemented by *sql.DB.
type SQLCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type sqlDb interface {
	Begin() (*sql.Tx, error)
}

type sqlTx interface {
	Commit() error
	Rollback() error
}
