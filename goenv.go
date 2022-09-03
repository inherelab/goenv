package goenv

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

var defaultFile = "~/.config/goenv/goenv.yml"
var ConfFile = envutil.Getenv("GOENV_CONF_FILE", defaultFile)
var Version = "0.0.1"

var CfgMgr *config.Config

// Cfg for the application
var Cfg = &appConf{}

// Init config and more
func Init() error {
	if err := loadConfig(); err != nil {
		return err
	}

	if err := CfgMgr.Decode(Cfg); err != nil {
		return err
	}

	Cfg.ConfFile = ConfFile
	return nil
}

func loadConfig() error {
	realPath, err := InitConfigFile(false)
	if err != nil {
		return err
	}

	CfgMgr = config.NewWithOptions("goenv",
		config.ParseEnv,
		config.ParseDefault,
		config.WithTagName("json"),
	).WithDriver(yamlv3.Driver)

	// cliutil.Infoln("load goenv config from", realPath)
	return CfgMgr.LoadFiles(realPath)
}

func InitConfigFile(reinstall bool) (string, error) {
	realPath := sysutil.ExpandPath(ConfFile)
	if !reinstall && fsutil.IsFile(realPath) {
		return realPath, nil
	}

	if reinstall {
		cliutil.Magentaln("TIP: Reinstall the default config contents to config file")
	} else {
		cliutil.Infoln("TIP: config file not found, will init default contents")
	}

	_, err := fsutil.PutContents(realPath, defConf, fsutil.FsCWTFlags)
	if err != nil {
		return "", err
	}

	cliutil.Infoln("OK, init default config at:", ConfFile)
	return realPath, nil
}
