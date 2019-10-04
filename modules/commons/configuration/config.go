package configuration

import (
	"fmt"

	"github.com/tkanos/gonfig"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	ConfigDriver struct {
		filename []string
		Port     int
		Url      string
	}
)

func NewConfigDriver() ConfigDriver {
	// os.Setenv("Connection_String", "test")
	// os.Setenv("APP_PORT", "8081")

	config := ConfigDriver{}
	err := gonfig.GetConf(config.getFileName(), &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	fmt.Println("config.Port:", config.Port)
	fmt.Println("config.Url:", config.Url)
	return config
	// fmt.Println(config.Connection_String)
}
func (c *ConfigDriver) getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	c.filename = []string{"config/", "config-", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(c.filename, ""))

	return filePath
}
