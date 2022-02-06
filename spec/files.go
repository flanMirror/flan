package spec

import (
	"embed"
	"log"
)

//go:embed files
var files embed.FS

const (
	openAPISpecJSONPath = "files/api.json"
	openAPISpecYAMLPath = "files/openapi.yaml"
)

var (
	openAPISpecJSON []byte
	openAPISpecYAML []byte
)

func init() {
	if d, err := files.ReadFile(openAPISpecJSONPath); err != nil {
		log.Printf("error caching spec %s: %s", openAPISpecJSONPath, err)
	} else {
		openAPISpecJSON = d
	}

	if d, err := files.ReadFile(openAPISpecYAMLPath); err != nil {
		log.Printf("error caching spec %s: %s", openAPISpecYAMLPath, err)
	} else {
		openAPISpecYAML = d
	}
}

func JSON() []byte {
	return openAPISpecJSON
}

func YAML() []byte {
	return openAPISpecYAML
}
