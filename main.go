package main

import (
	"embed"
	"io/fs"

	"github.com/aklinker1/url-shortener/server"
)

//go:embed ui/*
var embeddedUI embed.FS

//go:embed meta.json
var metaJSON string

func main() {
	uiFolder, err := fs.Sub(embeddedUI, "ui")
	if err != nil {
		panic("ui folder not found in embedded files");
	}

	server.Start(&uiFolder, metaJSON)
}
