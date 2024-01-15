package shared

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/swag"
	"gopkg.in/yaml.v3"
)

const (
	SWAGGER = iota
	UNKNOWN = -1
)

func DetectFileSpec(path string) int {
	_, filename := filepath.Split(path)
	if strings.HasSuffix(filename, ".yml") || strings.HasSuffix(filename, ".yaml") {
		yamlFile, err := os.ReadFile(path)
		if err != nil {
			log.Println(err)
			return UNKNOWN
		}
		c := map[string]interface{}{}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			log.Println(err)
			return UNKNOWN
		}
		if _, ok := c["swagger"]; ok {
			return SWAGGER
		}
		return UNKNOWN
	} else {
		return UNKNOWN
	}
}

func BaseImport(tgt string) (string, error) {
	absPath, err := filepath.Abs(tgt)
	if err != nil {
		return "", err
	}

	leafFolder := filepath.Base(absPath)
	return leafFolder, nil
}

func Pascalize(arg string) string {
	if len(arg) == 0 || arg[0] > '9' {
		return swag.ToGoName(arg)
	}
	if arg[0] == '+' {
		return swag.ToGoName("Plus " + arg[1:])
	}
	if arg[0] == '-' {
		return swag.ToGoName("Minus " + arg[1:])
	}

	return swag.ToGoName("Nr " + arg)
}
