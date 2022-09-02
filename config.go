package goenv

var defConf = `# 
# config for https://github.com/inherelab/goenv
# author: https://github.com/inhere
#

# mode allow: auto, goenv, brew
mode: brew

# on use brew mode
brew_lib_dir: /usr/local/opt

# use goenv mode.
# https://golang.org/dl
dl_host: https://golang.google.cn/dl
# install go dir.
install_dir: /usr/local/go

# custom add env map
env_map:
  APP_ENV: dev

current:
  version: 1.16
  path: /usr/local/opt/go@1.16

`

// mode consts
const (
	ModeAuto  = "auto"
	ModeGoEnv = "goenv"
	ModeBrew  = "brew"
	ModeScoop = "scoop"
)

// appConf struct
type appConf struct {
	// ConfFile path
	ConfFile string

	// Mode allow: auto, goenv, brew
	Mode string `json:"mode" default:"auto"`
	// BrewLibDir path
	BrewLibDir string `json:"brew_lib_dir" default:"/usr/local/opt"`

	// DlHost address, on use mode=goenv
	DlHost string `json:"dl_host" default:"https://golang.org/dl"`
	// InstallDir the go install dir.
	InstallDir string `json:"install_dir" default:"/usr/local/go"`
}

// IsBrewMode check
func (c *appConf) IsBrewMode() bool {
	return c.Mode == ModeBrew
}

// IsGoEnvMode check
func (c *appConf) IsGoEnvMode() bool {
	return c.Mode == ModeGoEnv
}
