package main

import (
	"os"
	"testing"
)

var idx Index

func TestMain(m *testing.M) {
	idx = startServer(IN_MEM)
	exitVal := m.Run()
	os.Exit(exitVal)
}
