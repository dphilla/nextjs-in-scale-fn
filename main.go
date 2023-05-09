package folder

import "embed"

//go:embed out/*
var Folder embed.FS
