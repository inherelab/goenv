package goenv

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

const (
	ModeGoEnv = "goenv"
	ModeBrew  = "brew"
)

var defaultFile = "~/.config/goenv/goenv.yml"
var Version = "1.0.1"
var ConfFile = envutil.Getenv("GOENV_CONF_FILE", defaultFile)

var CfgMgr *config.Config

// Cfg for the application
var Cfg = &appConf{}

// appConf struct
type appConf struct {
	// ConfFile path
	ConfFile string

	// Mode allow: goenv, brew
	Mode string `json:"mode" default:"goenv"`
	// BrewLibDir path
	BrewLibDir string `json:"brew_lib_dir" default:"/usr/local/opt"`

	// DlHost address, on use mode=goenv
	DlHost string `json:"dl_host" default:"https://golang.org/dl"`
	// InstallDir the go install dir.
	InstallDir string `json:"install_dir" default:"/usr/local/go"`
}

func (c *appConf) IsBrewMode() bool {
	return c.Mode == ModeBrew
}

func (c *appConf) IsGoEnvMode() bool {
	return c.Mode == ModeGoEnv
}

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

	cliutil.Infoln("load goenv config from", ConfFile)
	return CfgMgr.LoadFiles(realPath)
}
