package goembedtest

import (
	"embed"
)

//go:embed content/*
var FS embed.FS
