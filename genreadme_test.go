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
![GitHub](https://img.shields.io/github/license/clambin/foo?style=plastic)
` {
		t.Fatalf("unexpected output:\n%s", out.String())
	}
}

func Test_getModFile(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  modInfo
		pass  bool
	}{
		{
			name: "pass",
			input: `
module github.com/clambin/foo

go 1.21
`,
			want: modInfo{
				fullPath:     "github.com/clambin/foo",
				strippedPath: "clambin/foo",
			},
			pass: true,
		},
		{
			name: "non-github",
			input: `
module foo

go 1.21
`,
			pass: false,
		},
		{
			name:  "invalid",
			input: `invalid`,
			pass:  false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mod, err := getModFile(bytes.NewBufferString(tt.input))

			if tt.pass && err != nil {
				t.Fatalf("expected to pass, but got error: %v", err)
			}
			if !tt.pass && err == nil {
				t.Fatal("expected to fail, but passed")
			}

			if mod != tt.want {
				t.Fatalf("unexpected result: %v", mod)
			}
		})
	}
}
