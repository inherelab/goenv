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
	realPath := sysutil.ExpandPath(ConfFile)
	if !fsutil.IsFile(realPath) {
		cliutil.Infoln("TIP: goenv config file not found, will init a default")

		realPath = sysutil.ExpandPath(defaultFile)
		_, err := fsutil.PutContents(realPath, defConf, fsutil.FsCWTFlags)
		if err != nil {
			return err
		}

		cliutil.Infoln("OK, init default config at: ", defaultFile)
	}

	CfgMgr = config.NewWithOptions("goenv",
		config.ParseEnv,
		config.ParseDefault,
		config.WithTagName("json"),
	).WithDriver(yamlv3.Driver)

	// cliutil.Infoln("load goenv config from", realPath)
	return CfgMgr.LoadFiles(realPath)
}
