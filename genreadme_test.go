package main

import (
	"bytes"
	"testing"
)

func TestWriteReadme(t *testing.T) {
	info := modInfo{
		fullPath:     "github.com/clambin/foo",
		strippedPath: "clambin/foo",
	}
	var out bytes.Buffer
	writeREADME(&out, info)

	if out.String() != `# foo
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/clambin/foo?color=green&label=Release&style=plastic)
![Codecov](https://img.shields.io/codecov/c/gh/clambin/foo?style=plastic)
![Build](https://github.com/clambin/foo/workflows/Build/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/clambin/foo)
![GitHub](https://img.shields.io/github/license/clambin/foo?style=plastic
` {
		t.Log(out.String())
		t.Fatal("unexpected output")
	}
}
