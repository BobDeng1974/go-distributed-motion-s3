// Package dms3libs configuration structures and variables
//
package dms3libs

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// LibConfig contains dms3Libs configuration settings read from TOML file
var LibConfig *structConfig

type structConfig struct {
	SysCommands mapSysCommands
}

// mapSysCommands provides a location mapping of required system commands
type mapSysCommands map[string]string

// LoadLibConfig loads a TOML configuration file and parses entries into parameter values
func LoadLibConfig(configFile string) {

	if !IsFile(configFile) {
		log.Fatalln(configFile + " file not found")
	}

	if _, err := toml.DecodeFile(configFile, &LibConfig); err != nil {
		log.Fatalln(err.Error())
	}

}

// LoadComponentConfig loads a TOML configuration file and parses entries into parameter values
func LoadComponentConfig(structConfig interface{}, configFile string) {

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalln(configFile + " file not found")
	} else if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := toml.DecodeFile(configFile, structConfig); err != nil {
		log.Fatalln(err.Error())
	}

}

// SetLogFileLocation sets the location of the log file based on TOML configuration
func SetLogFileLocation(config *StructLogging) {

	projectDir := path.Dir(GetPackageDir())
	fail := false

	// if no config location set, attempt to set to development project folder
	if config.LogLocation == "" {
		if IsFile(projectDir) {
			config.LogLocation = projectDir
		} else {
			fail = true
		}
	} else if !IsFile(config.LogLocation) {
		fail = true
	}

	if fail {
		log.Fatalln("unable to set log location... check TOML configuration file")
	}

}
