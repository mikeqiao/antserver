package f

import (
	"github.com/mikeqiao/antserver/function/mysql"
	"github.com/mikeqiao/antserver/function/redis"
	p "github.com/mikeqiao/common/processor"
)

type Function struct {
	Uid int64
}

func (f *Function) Init() {
	mysql.Init()
	redis.Init()
	p.SetP()
}
func (f *Function) Close() {
	mysql.Close()
	redis.Close()
}

var FF = new(Function)
