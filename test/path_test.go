package test

import (
	"path"
	"testing"
)

func TestPathPrint(t *testing.T) {
	t.Log(path.Base("C:/WorkSpace/roPic_go/go.mod"))
}
