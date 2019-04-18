package mysql

import (
	"github.com/mikeqiao/common/db"
)

func Init() {
	db.InitDB()

}

func Close() {
	db.OnClose()

}
