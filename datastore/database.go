package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/corvuscrypto/unity_bot/config"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

var connection *sql.DB

func initDBConnection() {
	connectionString := fmt.Sprintf(
		"%s:%s@%s:%d/%s",
		config.GlobalConfig.DB.User,
		config.GlobalConfig.DB.Password,
		config.GlobalConfig.DB.Address,
		config.GlobalConfig.DB.Port,
		config.GlobalConfig.DB.Name,
	)
	var err error
	connection, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Unable to initialize database! Exiting.")
	}
}

// GetDB is a convenience function to make code accessing the DB connection pool
// more comprehensible
func GetDB() *sql.DB {
	return connection
}

// WithTransaction is a convenience function wrapper that ensures transactions
// are carried through without error. If an error is returned from the passed
// function then the transaction is automatically rolled back. Otherwise it is
// automatically committed.
func WithTransaction(f func(*sql.Tx) error) {
	tx, _ := connection.Begin()

	err := f(tx)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}

func init() {
	initDBConnection()
}
