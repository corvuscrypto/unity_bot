package datastore

import "database/sql"

// PreparedStatements is the map that holds all of the prepared SQL statements
var PreparedStatements map[string]*sql.Stmt

func initPreparedStatements() {
}

func init() {
	PreparedStatements = make(map[string]*sql.Stmt)
}
