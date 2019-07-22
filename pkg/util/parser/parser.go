package parser

import (
	"path/filepath"

	"github.com/instrumenta/conftest/pkg/util/parser/terraform"
	"github.com/instrumenta/conftest/pkg/util/parser/toml"
	"github.com/instrumenta/conftest/pkg/util/parser/yaml"
)

// Parser is the interface implemented by objects that can unmarshal
// bytes into a golang interface
type Parser interface {
	Unmarshal(p []byte, v interface{}) error
}

// GetParser returns a Parser for the given file type. Defaults to returning the YAML parser.
func GetParser(fileName string) Parser {
	suffix := filepath.Ext(fileName)

	switch suffix {
	case ".toml":
		return &toml.Parser{
			FileName: fileName,
		}
	case ".tf":
		return &terraform.Parser{
			FileName: fileName,
		}
	default:
		return &yaml.Parser{
			FileName: fileName,
		}
	}
}
