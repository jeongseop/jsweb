package controllers

import (
	"github.com/revel/config"
)

func getString(c *config.Config, section, option string) string {
	v, err := c.String(section, option)
	if err != nil {
		return nil
	}
	return v
}

func getInt(c *config.Config, section, option string) (int, error) {
	v, err := c.Int(section, option)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ReadConfig(fname string) (*config.Config, error) {
	c, err := config.ReadDefault(fname)
	if err != nil {
		return nil, err
	}
	return c, nil
}