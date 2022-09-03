package goenv

var defConf = `# 
# config for https://github.com/inherelab/goenv
# author: https://github.com/inhere
#

# select adaptor name.
# allow: auto, goenv, brew
adaptor: auto

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

// appConf struct
type appConf struct {
	// Adaptor name. allow: auto, goenv, brew, scoop
	Adaptor string `json:"adaptor" default:"auto"`
	// ConfFile path
	ConfFile string `json:"conf_file"`

	// BrewLibDir path, on use adaptor=brew
	BrewLibDir string `json:"brew_lib_dir" default:"/usr/local/opt"`

	// DlHost address, on use adaptor=goenv
	DlHost string `json:"dl_host" default:"https://golang.org/dl"`
	// InstallDir the go install dir.
	InstallDir string `json:"install_dir" default:"/usr/local/go"`
}
