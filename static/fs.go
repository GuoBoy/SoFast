package static

import "embed"

//go:embed dist/index.html
var IndexFile []byte

//go:embed dist/assets
var Assets embed.FS
