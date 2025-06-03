package rpc

// Package rpc provides RPC functionality for the Leither project
type LApiStub struct {
	VarActStub
}

type VarActStub struct {
	GetVar          func(sid, name string, args ...string) (any, error)
	GetVarByContext func(sid, name string, mapOpt map[string]string) (any, error)
	Act             func(sid, name string, args ...string) error
	GetGobVar       func(sid, name string, args ...string) ([]byte, error)
}
