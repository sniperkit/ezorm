package people

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/ezbuy/ezorm/db"
	"github.com/jmoiron/sqlx"
)

var (
	_sqlServer *db.SqlServer
	_db        *sql.DB
)

func MssqlSetUp(dataSourceName string) {
	conn, err := sqlx.Connect("mssql", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("[db.GetSqlServer] open sql fail:%s", err.Error()))
	}

	_sqlServer = &db.SqlServer{DB: conn}
	_db = conn.DB
}

func MssqlSetMaxOpenConns(maxOpenConns int) {
	_sqlServer.SetMaxOpenConns(maxOpenConns)
}

func MssqlSetMaxIdleConns(maxIdleConns int) {
	_sqlServer.SetMaxIdleConns(maxIdleConns)
}

func MssqlClose() error {
	return _sqlServer.Close()
}
