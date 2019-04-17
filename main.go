package main

import (
	"runtime"

	"github.com/mikeqiao/ant"
	"github.com/mikeqiao/antserver/conf"
	f "github.com/mikeqiao/antserver/function"
)

func main() {
	runtime.GOMAXPROCS(4)
	conf.InitConfig()
	ant.Start(f.FF)
}
