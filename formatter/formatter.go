package formatter

import (
	"io"
	"strings"

	"github.com/asciimoo/wuzz/config"
)

type ResponseFormatter interface {
	Format(writer io.Writer, data []byte) error
	Title() string
	Searchable() bool
}

func New(config *config.Config, contentType string) ResponseFormatter {
	if strings.Contains(contentType, "application/json") && config.General.FormatJSON {
		return &jsonFormatter{}
	} else if strings.Contains(contentType, "text/html") {
		return &htmlFormatter{}
	} else if strings.Index(contentType, "text") == -1 && strings.Index(contentType, "application") == -1 {
		return &binaryFormatter{}
	} else {
		return &textFormatter{}
	}
}
