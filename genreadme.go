package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/mod/modfile"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filename := "go.mod"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	file, err := getModFile(filename)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	writeREADME(&out, file)
	fmt.Println(out.String())
}

type modInfo struct {
	fullPath     string
	strippedPath string
}

func getModFile(path string) (modInfo, error) {
	mod, err := os.ReadFile(path)
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

	strippedPath, _ := strings.CutPrefix(file.Module.Mod.Path, "github.com/")
	return modInfo{fullPath: file.Module.Mod.Path, strippedPath: strippedPath}, err
}

func writeREADME(readme io.Writer, info modInfo) {
	writeTitle(readme, info)
	writeTag(readme, info)
	writeCodeCov(readme, info)
	writeBuild(readme, info)
	writeGoReport(readme, info)
	writeLicense(readme, info)
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
	_, _ = w.Write([]byte("![GitHub](https://img.shields.io/github/license/" + info.strippedPath + "?style=plastic\n"))
}
