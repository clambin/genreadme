package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/mod/modfile"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var input = flag.String("input", "go.mod", "path of the go.mod file to read")

func main() {
	flag.Parse()
	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	info, err := getModFile(f)
	if err != nil {
		panic(err)
	}
	writeREADME(os.Stdout, info)
}

type modInfo struct {
	fullPath     string
	strippedPath string
}

func getModFile(source io.Reader) (modInfo, error) {
	mod, err := io.ReadAll(source)
	if err != nil {
		return modInfo{}, err
	}
	file, err := modfile.Parse("go.mod", mod, nil)
	if err != nil {
		return modInfo{}, fmt.Errorf("parse: %w", err)
	}
	if !strings.HasPrefix(file.Module.Mod.Path, "github.com/") {
		return modInfo{}, errors.New("only supports github-hosted repos")
	}

	strippedPath := strings.TrimPrefix(file.Module.Mod.Path, "github.com/")
	return modInfo{fullPath: file.Module.Mod.Path, strippedPath: strippedPath}, err
}

func writeREADME(w io.Writer, info modInfo) {
	writeTitle(w, info)
	writeTag(w, info)
	writeCodeCov(w, info)
	writeBuild(w, info)
	writeGoReport(w, info)
	writeLicense(w, info)
}

func writeTitle(w io.Writer, info modInfo) {
	_, _ = w.Write([]byte("# " + filepath.Base(info.fullPath) + "\n"))
}

func writeTag(w io.Writer, info modInfo) {
	_, _ = w.Write([]byte("![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/" + info.strippedPath + "?color=green&label=Release&style=plastic)\n"))
}

func writeCodeCov(w io.Writer, info modInfo) {
	_, _ = w.Write([]byte("![Codecov](https://img.shields.io/codecov/c/gh/" + info.strippedPath + "?style=plastic)\n"))
}

func writeBuild(w io.Writer, info modInfo) {
	_, _ = w.Write([]byte("![Build](https://" + info.fullPath + "/workflows/Build/badge.svg)\n"))
}

func writeGoReport(w io.Writer, info modInfo) {
	_, _ = w.Write([]byte("![Go Report Card](https://goreportcard.com/badge/" + info.fullPath + ")\n"))
}

func writeLicense(w io.Writer, info modInfo) {
	_, _ = w.Write([]byte("![GitHub](https://img.shields.io/github/license/" + info.strippedPath + "?style=plastic)\n"))
}
