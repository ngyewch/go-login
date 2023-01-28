package resources

import "embed"

var (
	//go:embed templates
	TemplatesFS embed.FS
)
