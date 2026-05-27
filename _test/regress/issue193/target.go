package main

import (
	"github.com/go-ruleguard/rg1"
	"go.uber.org/zap"
)

type Foo struct {
}

type Bar struct {
}

func (f *Foo) String() string { _ = "STUB: not implemented"; return "" }

func (f *Bar) Run() { _ = "STUB: not implemented"; return }

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	f := &Foo{}
	sugar.Infof("Stringer %+v", f)

	var err error
	sugar.Info(err)
	sugar.Info(err, "")
	sugar.Info(err.Error(), "")

	var w rg1.Worker = &Bar{}
	w.Run()

	sugar.Info("Test")
}
