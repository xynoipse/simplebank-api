package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank-api/utils"
	"testing"

	_ "github.com/lib/pq"
)

var (
	config utils.Config
)

func init() {
	var err error

	config, err = utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
