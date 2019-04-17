package f

import (
	"github.com/mikeqiao/antserver/function/db"
	p "github.com/mikeqiao/common/processor"
)

type Function struct {
	Uid int64
}

func (f *Function) Init() {
	db.InitDB()
	p.SetP()
}
func (f *Function) Close() {
	db.Close()
}

var FF = new(Function)
