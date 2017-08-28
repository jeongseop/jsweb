package controllers

import (
	"github.com/revel/config"
	"path/filepath"
	"log"
	"os"
)

type CustomConfig struct {
	Config *config.Context
	FilePaths []string
	CurrentSection string
}
const (
	//jsweb path
	jsweb = "github.com/jeongseop/jsweb"
)
var (
	currentSection string
)
func NewCustomConfig() *CustomConfig {
	c := CustomConfig{}
	jswebConf := filepath.Join(jsweb, "conf")
	c.AddFilePath(jswebConf)
	c.CurrentSection = config.DefaultSection
	return &c
}
func (c *CustomConfig) AddFilePath(path string) {
	sourcePath := filepath.Join(os.Getenv("GOPATH"), "src")
	c.FilePaths = append(
		[]string{filepath.Join(sourcePath, path)},
		c.FilePaths...
	)
}
func (c *CustomConfig) LoadConfig(fname string) error {
	conf, err := config.LoadContext(fname, c.FilePaths)
	if err != nil {
		return err
	}
	c.Config = conf
	log.Println(conf.Raw().Options("DEFAULT"))
	return nil
}
func (c *CustomConfig) SetSection(section string) bool {
	if c.Config == nil {
		log.Fatal("Not Opened Config File!!")
		return false
	}
	if !c.Config.HasSection(section) {
		log.Fatal("Hasn't Section Name!!")
		return false
	}
	c.CurrentSection = section
	c.Config.SetSection(c.CurrentSection)
	return true
}
func (c *CustomConfig) GetStringDefault(option, defaultValue string) string {
	return c.Config.StringDefault(option, defaultValue)
}
func (c *CustomConfig) GetIntDefault(option string, defaultValue int) int {
	return c.Config.IntDefault(option, defaultValue)
}
func (c *CustomConfig) GetBoolDefault(option string, defaultValue bool) bool {
	return c.Config.BoolDefault(option, defaultValue)
}