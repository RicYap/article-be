package article

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	Data struct {
		db   *sqlx.DB
		stmt *map[string]*sqlx.Stmt
	}

	statement struct {
		key   string
		query string
	}
)

// Query List to Prepare
const (
	getAllUser  = "GetAllUser"
	qGetAllUser = `SELECT UserID, u.Name, UserName, u.RoleID, r.Name AS RoleName, PhoneNumber 
					FROM user u
						JOIN role r ON u.RoleID = r.RoleID`
)

var (
	readStmt = []statement{
		{getAllUser, qGetAllUser},
	}
	insertStmt = []statement{}
	updateStmt = []statement{}
	deleteStmt = []statement{}
)

// New ...
func New(db *sqlx.DB) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)

	d := &Data{
		db:   db,
		stmt: &stmts,
	}

	d.InitStmt()

	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	*d.stmt = stmts
}
