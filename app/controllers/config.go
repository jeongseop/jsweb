package controllers

import (
	"github.com/revel/config"
	"os"
)

func getString(c *config.Config, section, option string) string {
	v, err := c.String(section, option)
	if err != nil {
		return ""
	}
	return v
}

func getInt(c *config.Config, section, option string) int {
	v, err := c.Int(section, option)
	if err != nil {
		return 0
	}
	return v
}

func ReadConfig(fname string) (*config.Config, error) {
	defaultPath := os.Getenv("GOPATH")
	newName := fname
	c, err := config.ReadDefault(newName)
	if err != nil {
		return nil, err
	}
	return c, nil
}
