package devices

import "embed"

//go:embed database/migration/*.sql
var FS embed.FS
